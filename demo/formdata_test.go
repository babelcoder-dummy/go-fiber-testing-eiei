package demo_test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestFormData(t *testing.T) {
	app := fiber.New()

	setupRoutes := func(createPath func(*multipart.FileHeader) string) {
		app.Post("/files", func(c *fiber.Ctx) error {
			file, _ := c.FormFile("file")

			c.SaveFile(file, createPath(file))

			return c.Status(fiber.StatusCreated).JSON(fiber.Map{"file": file.Filename})
		})
	}

	tempFile, _ := os.CreateTemp(os.TempDir(), "test-")
	defer func(file *os.File) {
		err := file.Close()
		assert.NoError(t, err)
		err = os.Remove(file.Name())
		assert.NoError(t, err)
	}(tempFile)

	createPath := func(file *multipart.FileHeader) string {
		return tempFile.Name()
	}

	setupRoutes(createPath)

	formData := &bytes.Buffer{}
	writer := multipart.NewWriter(formData)
	tempFilePath := tempFile.Name()
	_, tempFileName := filepath.Split(tempFilePath)
	part, _ := writer.CreateFormFile("file", tempFileName)
	part.Write([]byte("hello world"))
	writer.Close()

	req := httptest.NewRequest("POST", "/files", formData)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", tempFilePath))

	res, err := app.Test(req)
	body, _ := io.ReadAll(res.Body)
	savedFile, _ := os.ReadFile(tempFilePath)

	assert.NoError(t, err)
	assert.Equal(t, 201, res.StatusCode)
	assert.JSONEq(t, fmt.Sprintf(`{"file":"%s"}`, tempFileName), string(body))
	assert.Equal(t, "hello world", string(savedFile))
}
