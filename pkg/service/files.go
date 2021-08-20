package service

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/labstack/echo"
)

func (s *Service) Upload(c echo.Context) error {
	_file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := _file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create("file.json")
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.NoContent(200)
}

func (s *Service) Load(c echo.Context) error {
	file, err := os.Open("file.json")
	if err != nil {
		return err
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return c.HTML(200, string(content))
}
