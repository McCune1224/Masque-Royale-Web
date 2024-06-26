package handler

import (
	"encoding/csv"
	"errors"
	"fmt"
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
		A []TempCreateAbilityDetailParams
		P []models.CreatePassiveDetailParams
	}

	bulkRoleCreateList := []bulkRoleCreate{}

	// TODO: Remove this hardcoded limit after testing
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

	realAbility := models.CreateAbilityDetailParams{}
	for i, roleParams := range bulkRoleCreateList {
		for _, a := range roleParams.A {
			roleID := roleIds[i]

			realAbility.Name = a.Name
			realAbility.Description = a.Description
			realAbility.DefaultCharges = a.DefaultCharges
			realAbility.Rarity = a.Rarity
			realAbility.AnyAbility = a.AnyAbility

			dbAbility, err := q.CreateAbilityDetail(c.Request().Context(), realAbility)

			if err != nil {
				if util.ErrorContains(err, pgerrcode.UniqueViolation) {
					log.Println(a.Name, "already exists")
				} else {
					log.Println(err, roleParams.R.Name, a.Name)
					return util.InternalServerErrorJson(c, err.Error())
				}
			}

			for _, categoryName := range a.CategoryNames {
				dbCategory, err := q.GetCategoryByName(c.Request().Context(), pgtype.Text{String: strings.ToUpper(categoryName), Valid: true})
				if err != nil {
					log.Println("Error Getting Category ID", categoryName, err)
					return util.InternalServerErrorJson(c, err.Error())
				}
				q.CreateAbilityDetailsCategoriesJoin(c.Request().Context(), models.CreateAbilityDetailsCategoriesJoinParams{
					AbilityDetailsID: dbAbility.ID,
					CategoriesID:     dbCategory.ID,
				})
			}

			_, err = q.CreateRoleAbilityJoin(c.Request().Context(), models.CreateRoleAbilityJoinParams{RoleID: roleID, AbilityID: dbAbility.ID})
			if err != nil {
				log.Println(err, roleParams.R.Name, a.Name)
				return util.InternalServerErrorJson(c, err.Error())
			}
		}

		for _, p := range roleParams.P {
			rId := roleIds[i]
			dbPassive, err := q.CreatePassiveDetail(c.Request().Context(), p)
			if err != nil {
				if !util.ErrorContains(err, "23505") {
					log.Println(err, roleParams.R.Name, p.Name)
					return util.InternalServerErrorJson(c, err.Error())
				}
				// Passive already exists, so just grab it here before proceeding
				dbPassive, err = q.GetPassiveDetailsByName(c.Request().Context(), p.Name)
				if err != nil {
					log.Println(err, roleParams.R.Name, p.Name)
					return util.InternalServerErrorJson(c, err.Error())
				}
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

// FIXME: This is a temporary solution to not making abilities within its own function.
// Right now I dont have a db context here so I can't immediatly get the assocaited cateogries
type TempCreateAbilityDetailParams struct {
	models.CreateAbilityDetailParams
	CategoryNames []string
}

func parseRoleChunk(chunk [][]string) (models.CreateRoleParams, []TempCreateAbilityDetailParams, []models.CreatePassiveDetailParams, error) {
	roleParams := models.CreateRoleParams{}
	tempRoleAbilityDetailParams := []TempCreateAbilityDetailParams{}
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
		return roleParams, tempRoleAbilityDetailParams, rolePassiveDetailParams, errors.New("Invalid alignment")
	}

	abParseIndex := 5
	for chunk[abParseIndex][1] != "Passives:" {
		ab, err := parseAbility(chunk[abParseIndex])
		if err != nil {
			return roleParams, tempRoleAbilityDetailParams, rolePassiveDetailParams, err
		}

		tempRoleAbilityDetailParams = append(tempRoleAbilityDetailParams, ab)
		abParseIndex++
	}
	for _, p := range chunk[abParseIndex+1:] {
		createPassive := models.CreatePassiveDetailParams{Name: p[1], Description: p[2]}
		rolePassiveDetailParams = append(rolePassiveDetailParams, createPassive)
	}

	return roleParams, tempRoleAbilityDetailParams, rolePassiveDetailParams, nil
}

func parseAbility(row []string) (TempCreateAbilityDetailParams, error) {
	abilityDetail := TempCreateAbilityDetailParams{}
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

	abilityDetail.CategoryNames = strings.Split(row[5], "/")
	return abilityDetail, nil
}

func (h *Handler) SyncStatusDetailsCSV(c echo.Context) error {
	form, err := c.FormFile("file")
	if err != nil {
		log.Println("Error getting form file", err)
		return util.BadRequestJson(c, err.Error())
	}

	if form == nil {
		log.Println("No file provided")
		return util.BadRequestJson(c, "No file provided")
	}

	rawFile, err := form.Open()
	if err != nil {
		log.Println("Error opening file", err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	defer rawFile.Close()

	reader := csv.NewReader(rawFile)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("Error reading file", err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	// Drop first and second row as they are headers / useless
	records = records[2:]

	q := models.New(h.Db)

	csvAnyAbilityDetails := []TempCreateAbilityDetailParams{}
	for i, r := range records {
		// WARNING: For some god forsaken reason there's 80 empty lines in the CSV's normally so need to manually check
		// for the actual "end of file" line
		if r[1] == "" {
			break
		}

		csvAnyAbilityLine := TempCreateAbilityDetailParams{}
		switch strings.TrimSpace(r[1]) {
		case "Common":
			csvAnyAbilityLine.Rarity = models.RarityCOMMON
		case "Uncommon":
			csvAnyAbilityLine.Rarity = models.RarityUNCOMMON
		case "Rare":
			csvAnyAbilityLine.Rarity = models.RarityRARE
		case "Epic":
			csvAnyAbilityLine.Rarity = models.RarityEPIC
		case "Legendary":
			csvAnyAbilityLine.Rarity = models.RarityLEGENDARY
		case "Mythical":
			csvAnyAbilityLine.Rarity = models.RarityMYTHICAL
		case "Role Specific Ability (Non AA)":
			csvAnyAbilityLine.Rarity = models.RarityROLESPECIFIC
		default:
			log.Printf("Invalid Rarity: %s\tName: %s\t LINE: %d", r[1], r[2], i)
			return util.BadRequestJson(c, "Invalid Rarity")
		}
		csvAnyAbilityLine.Name = r[2]
		csvAnyAbilityLine.Description = r[6]
		csvAnyAbilityLine.AnyAbility = pgtype.Bool{Bool: true, Valid: true}
		csvAnyAbilityLine.CategoryNames = strings.Split(r[5], "/")
		csvAnyAbilityDetails = append(csvAnyAbilityDetails, csvAnyAbilityLine)
	}

	for _, a := range csvAnyAbilityDetails {
		dbAbility, err := q.CreateAbilityDetail(c.Request().Context(), models.CreateAbilityDetailParams{
			Name:        a.Name,
			Description: a.Description,
			AnyAbility:  a.AnyAbility,
			Rarity:      a.Rarity,
		})
		if err != nil {
			// WARNING: we can just skip this as the CSV can have duplicate entries from the roles csv...that should be changed but w/e
			if util.ErrorContains(err, "23505") {
				continue
			} else {
				log.Println("Unhandled error on upload", err, a.Name)
				return util.InternalServerErrorJson(c, fmt.Sprintf("Error creating CSV Any Ability details for %s: %s", a.Name, err.Error()))
			}
		}
		for _, cat := range a.CategoryNames {
			dbCat, err := q.GetCategoryByName(c.Request().Context(), pgtype.Text{String: strings.ToUpper(cat), Valid: true})
			if err != nil {
				log.Println("Error getting category", cat, err)
				return util.InternalServerErrorJson(c, err.Error())
			}
			q.CreateAbilityDetailsCategoriesJoin(c.Request().Context(), models.CreateAbilityDetailsCategoriesJoinParams{
				AbilityDetailsID: dbAbility.ID,
				CategoriesID:     dbCat.ID,
			})
		}
	}

	return c.JSON(200, "Success")
}
