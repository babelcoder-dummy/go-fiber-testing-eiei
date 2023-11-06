package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2/utils"
)

type Storage struct{}

func (s *Storage) Save(fh *multipart.FileHeader) (string, error) {
	if fh == nil {
		return "", nil
	}

	file, err := fh.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	id := utils.UUIDv4()
	path := fmt.Sprintf("uploads/%s", id)
	os.MkdirAll(path, os.ModePerm)
	filename := path + "/" + fh.Filename

	dst, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	return filename, err
}

func (s *Storage) Remove(path string) error {
	dir := filepath.Dir(path)
	err := os.RemoveAll(dir)

	return err
}
