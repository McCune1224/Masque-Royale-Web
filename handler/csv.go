package handler

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
)

func (h *Handler) SyncRolesCsv(c echo.Context) error {
	form, err := c.FormFile("file")
	if err != nil {
		return util.BadRequestJson(c, err.Error())
	}

	if form == nil {
		return util.BadRequestJson(c, "No file provided")
	}

	rawFile, err := form.Open()
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	defer rawFile.Close()

	reader := csv.NewReader(rawFile)
	chunks := [][][]string{}
	currChunk := [][]string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			chunks = append(chunks, currChunk)
			break
		}
		if err != nil {
			return util.InternalServerErrorJson(c, err.Error())
		}
		if record[1] == "" {
			chunks = append(chunks, currChunk)
			currChunk = [][]string{}
		} else {
			currChunk = append(currChunk, record)
		}
	}
	chunks = chunks[1:]

	if len(chunks) < 1 {
		return util.BadRequestJson(c, "No records found")
	}

	type bulkRoleCreate struct {
		R models.CreateRoleParams
		A []models.CreateAbilityDetailParams
		P []models.CreatePassiveDetailParams
	}

	bulkRoleCreateList := []bulkRoleCreate{}

	for i := range chunks {
		roleParams, roleAbilityDetailParams, rolePassiveDetailParams, err := parseRoleChunk(chunks[i])
		if err != nil {
			log.Println("Error Parsing Roles CSV into chunks", err)
			return util.InternalServerErrorJson(c, err.Error())
		}
		bulkEntry := bulkRoleCreate{
			R: roleParams,
			A: roleAbilityDetailParams,
			P: rolePassiveDetailParams,
		}
		bulkRoleCreateList = append(bulkRoleCreateList, bulkEntry)
	}

	q := models.New(h.Db)

	err = q.NukeRoles(c.Request().Context())
	if err != nil {
		log.Println("Error Nuking Roles", err)
		return util.InternalServerErrorJson(c, err.Error())
	}

	// NOTE: Need to create the role first before creating the ability/passive, otherwise the ability/passive will be created with the wrong role_id
	// hence why this is in its own loop
	roleIds := pq.Int32Array{}
	for _, roleParams := range bulkRoleCreateList {
		r, err := q.CreateRole(c.Request().Context(), roleParams.R)
		if err != nil {
			log.Println("Error Creating Role", err)
			return util.InternalServerErrorJson(c, err.Error())
		}
		roleIds = append(roleIds, r.ID)
	}
	for i, roleParams := range bulkRoleCreateList {
		for _, a := range roleParams.A {
			roleID := roleIds[i]
			dbAbility, err := q.CreateAbilityDetail(c.Request().Context(), a)

			if err != nil {
				if util.ErrorContains(err, pgerrcode.UniqueViolation) {
					log.Println(a.Name, "already exists")
				} else {
					log.Println(err, roleParams.R.Name, a.Name)
					return util.InternalServerErrorJson(c, err.Error())
				}
			}
			_, err = q.CreateRoleAbilityJoin(c.Request().Context(), models.CreateRoleAbilityJoinParams{RoleID: roleID, AbilityID: dbAbility.ID})
			if err != nil {
				log.Println(err, roleParams.R.Name, a.Name)
				return util.InternalServerErrorJson(c, err.Error())
			}
		}

		for _, p := range roleParams.P {
			rId := roleIds[i]
			log.Println("\t", p.Name)
			dbPassive, err := q.CreatePassiveDetail(c.Request().Context(), p)
			// FIXME: For some god forsaken reason, this error is not being processed right, for right now im ignoring it
			err = nil
			if err != nil {
				if util.ErrorContains(err, "23505") {
					log.Println(p.Name, "already exists")
				} else {
					log.Println(err, roleParams.R.Name, p.Name)
					return util.InternalServerErrorJson(c, err.Error())
				}
				log.Println(err, roleParams.R.Name, p.Name)
				return util.InternalServerErrorJson(c, err.Error())
			}
			// insert entry into role_passives_join
			_, err = q.CreateRolePassiveJoin(c.Request().Context(), models.CreateRolePassiveJoinParams{RoleID: rId, PassiveID: dbPassive.ID})
			if err != nil {
				log.Println(err, roleParams.R.Name, p.Name)
				return util.InternalServerErrorJson(c, err.Error())
			}
		}

	}

	return c.JSON(200, "Success")

}

func parseRoleChunk(chunk [][]string) (models.CreateRoleParams, []models.CreateAbilityDetailParams, []models.CreatePassiveDetailParams, error) {
	roleParams := models.CreateRoleParams{}
	roleAbilityDetailParams := []models.CreateAbilityDetailParams{}
	rolePassiveDetailParams := []models.CreatePassiveDetailParams{}
	roleParams.Name = chunk[1][1]
	switch strings.ToUpper(chunk[3][1]) {
	case string(models.AlignmentCHAOTIC):
		roleParams.Alignment = models.AlignmentCHAOTIC
	case string(models.AlignmentLAWFUL):
		roleParams.Alignment = models.AlignmentLAWFUL
	case string(models.AlignmentOUTLANDER):
		roleParams.Alignment = models.AlignmentOUTLANDER
	default:
		log.Println(chunk[3][1])
		return roleParams, roleAbilityDetailParams, rolePassiveDetailParams, errors.New("Invalid alignment")
	}

	abParseIndex := 5
	for chunk[abParseIndex][1] != "Passives:" {
		ab, err := parseAbility(chunk[abParseIndex])
		if err != nil {
			return roleParams, roleAbilityDetailParams, rolePassiveDetailParams, err
		}
		roleAbilityDetailParams = append(roleAbilityDetailParams, ab)
		abParseIndex++
	}
	for _, p := range chunk[abParseIndex+1:] {
		createPassive := models.CreatePassiveDetailParams{Name: p[1], Description: p[2]}
		rolePassiveDetailParams = append(rolePassiveDetailParams, createPassive)
	}

	return roleParams, roleAbilityDetailParams, rolePassiveDetailParams, nil
}

func parseAbility(row []string) (models.CreateAbilityDetailParams, error) {
	abilityDetail := models.CreateAbilityDetailParams{}
	abilityDetail.Name = row[1]
	abilityDetail.Description = row[4]

	iCharge := int32(999999)
	if row[2] != "∞" {
		charge, err := strconv.Atoi(row[2])
		if err != nil {
			log.Println("ERR ON", abilityDetail.Name)
			return abilityDetail, err
		}
		iCharge = int32(charge)
	}

	abCharge := pgtype.Int4{
		Int32: iCharge,
		Valid: true,
	}
	abilityDetail.DefaultCharges = abCharge
	switch row[3] {
	case "*":
		abTrue := pgtype.Bool{
			Bool:  true,
			Valid: true,
		}
		abilityDetail.AnyAbility = abTrue
		// abilityDetail.RoleSpecific = roleName
		switch models.Rarity(strings.ToUpper(row[6])) {
		case models.RarityCOMMON:
			abilityDetail.Rarity = models.RarityCOMMON
		case models.RarityUNCOMMON:
			abilityDetail.Rarity = models.RarityUNCOMMON
		case models.RarityRARE:
			abilityDetail.Rarity = models.RarityRARE
		case models.RarityEPIC:
			abilityDetail.Rarity = models.RarityEPIC
		case models.RarityLEGENDARY:
			abilityDetail.Rarity = models.RarityLEGENDARY
		case models.RarityMYTHICAL:
			abilityDetail.Rarity = models.RarityMYTHICAL
		}
	case "^":
		abTrue := pgtype.Bool{
			Bool:  true,
			Valid: true,
		}
		abilityDetail.AnyAbility = abTrue
		// abilityDetail.RoleSpecific = roleName
		abilityDetail.Rarity = models.RarityROLESPECIFIC
	case "":
		abFalse := pgtype.Bool{
			Bool:  true,
			Valid: true,
		}
		abilityDetail.AnyAbility = abFalse
		// abilityDetail.RoleSpecific = roleName
		abilityDetail.Rarity = models.RarityROLESPECIFIC
	}
	return abilityDetail, nil
}

func (h *Handler) SyncStatusDetailsCSV(c echo.Context) error {
	form, err := c.FormFile("file")
	if err != nil {
		return util.BadRequestJson(c, err.Error())
	}

	if form == nil {
		return util.BadRequestJson(c, "No file provided")
	}

	rawFile, err := form.Open()
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	defer rawFile.Close()

	//read all csv func
	reader := csv.NewReader(rawFile)
	records, err := reader.ReadAll()
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}

	if len(records) < 1 {
		return util.BadRequestJson(c, "No records found")
	}

	q := models.New(h.Db)
	for i := 0; i < len(records); i++ {
		desc := pgtype.Text{
			String: records[i][2],
			Valid:  true,
		}

		status := models.CreateStatusDetailParams{
			Name:        records[i][1],
			Description: desc,
		}
		_, err = q.CreateStatusDetail(c.Request().Context(), status)
		if err != nil {
			return util.InternalServerErrorJson(c, err.Error())
		}
	}

	return c.JSON(200, "Success")
}
