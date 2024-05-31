package handler

import (
	"encoding/csv"
	"io"
	"log"

	"github.com/labstack/echo/v4"
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
		log.Println(chunks[i][1][1])
	}

	if len(chunks) < 1 {
		return util.BadRequestJson(c, "No records found")
	}

	return c.JSON(200, "Success")

}
