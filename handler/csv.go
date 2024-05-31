package handler

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
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

	for i := range chunks {
		roleParams, _, _, err := parseRoleChunk(chunks[i])
		if err != nil {
			return util.InternalServerErrorJson(c, err.Error())
		}
		log.Println(roleParams)
	}

	if len(chunks) < 1 {
		return util.BadRequestJson(c, "No records found")
	}

	return c.JSON(200, "Success")
}

func parseRoleChunk(chunk [][]string) (models.CreateRoleParams, []models.CreateAbilityDetailParams, []models.CreatePassiveDetailParams, error) {
	roleParams := models.CreateRoleParams{}
	roleAbilityDetailParams := []models.CreateAbilityDetailParams{}
	roleStatusDetailParams := []models.CreatePassiveDetailParams{}
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
		return roleParams, roleAbilityDetailParams, roleStatusDetailParams, errors.New("Invalid alignment")
	}

	return roleParams, roleAbilityDetailParams, roleStatusDetailParams, nil
}

func parseAbility(row []string) (models.CreateAbilityDetailParams, error) {
	abilityDetail := models.CreateAbilityDetailParams{}
	abilityDetail.Name = row[1]
	charge, err := strconv.Atoi(row[2])
	if err != nil {
		return abilityDetail, err
	}
	abCharge := pgtype.Int4{}
	abCharge.Int32 = int32(charge)
	abilityDetail.DefaultCharges = abCharge
	switch row[3] {
	case "*":
		abTrue := pgtype.Bool{}
		abTrue.Bool = true
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
		abTrue := pgtype.Bool{}
		abTrue.Bool = true
		abilityDetail.AnyAbility = abTrue
		// abilityDetail.RoleSpecific = roleName
		abilityDetail.Rarity = models.RarityROLESPECIFIC
	case "":
		abFalse := pgtype.Bool{}
		abFalse.Bool = true
		abilityDetail.AnyAbility = abFalse
		// abilityDetail.RoleSpecific = roleName
		abilityDetail.Rarity = models.RarityROLESPECIFIC
	}
	return abilityDetail, nil
}

// func parseStatusChunk(chunk [][]string) (models.StatusDetail, error) {}
