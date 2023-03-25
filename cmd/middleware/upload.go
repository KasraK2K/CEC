package middleware

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"

	"app/pkg/helper"
)

func HandleMultipart(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")
	fmt.Println(contentType)
	if strings.Contains(contentType, "multipart/form-data") {
		err := upload(c)
		if err != nil {
			helper.JSON(c, err.Error(), http.StatusInternalServerError)
		}
		return c.Next()
	}
	return c.Next()
}

func upload(c *fiber.Ctx) error {
	// Parse the multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	// Get the field name for the files
	mainFolder := form.Value["m"][0]
	subFolder := form.Value["s"][0]

	// Create a new directory based on the field name
	dirPath := filepath.Join("uploads", mainFolder, subFolder)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	// Loop through all the files in the form data
	for _, fileHeader := range form.File {
		// Loop through all the files for this field name
		for _, file := range fileHeader {
			fileExt := filepath.Ext(file.Filename)
			fileName := file.Filename[0 : len(file.Filename)-len(fileExt)]

			// Open the file
			fileContent, err := file.Open()
			if err != nil {
				return err
			}
			defer fileContent.Close()

			// Create a new file on the server
			newFileName := filepath.Join(dirPath, fileName+".jpg")
			newFile, err := os.Create(newFileName)
			if err != nil {
				return err
			}
			defer newFile.Close()

			// Write the file to the server
			if _, err := io.Copy(newFile, fileContent); err != nil {
				return err
			}
		}
	}

	return nil
}
