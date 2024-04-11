package controlles

import (
	"fmt"
	"reflect"

	"e-monitoring/app/models"
	"e-monitoring/pkg/database"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)


func toExcelColumnName(index int) string {
	columnName := ""
	for index > 0 {
		index-- // Adjust index to be 0-based
		remainder := index % 26
		columnName = string(rune('A'+remainder)) + columnName
		index = index / 26
	}
	return columnName
}

func Download(c *fiber.Ctx) error {
	var komitmen []models.Komitmen
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			c.Status(fiber.StatusInternalServerError).SendString(err.Error())
			return
		}
	}()

	// Create a new sheet
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Set cell value in Sheet1
	f.SetCellValue("Sheet1", "B2", 100)

	// Retrieve data from database
	if err := database.DB.Find(&komitmen).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Add headers based on Struct, if komitmen is not empty
	
	
	// Usage in your code
	if len(komitmen) > 0 {
		t := reflect.TypeOf(komitmen[0])
		for i := 0; i < t.NumField(); i++ {
			columnName := toExcelColumnName(i+1) + "1" // "+1" to start from 1 instead of 0
			f.SetCellValue("Sheet2", columnName, t.Field(i).Name)
		}
	}
	// Populate the sheet with data
	for i, d := range komitmen {
		val := reflect.ValueOf(d)
		for j := 0; j < val.NumField(); j++ {
			cell := toExcelColumnName(j+1)  + fmt.Sprintf("%d", i+2)
			f.SetCellValue("Sheet2", cell, val.Field(j).Interface())
		}
	}
	f.SetActiveSheet(index)

	// Save the file and send as response
	fileName := "komitmen.xlsx"
	if err := f.SaveAs(fileName); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Download(fileName)
}
