package util

import (
	"io"
	"mime/multipart"

	"github.com/szyhf/go-excel"
)

func ReadImportExcel[T any](file *multipart.FileHeader) (data T, err error) {
	multipartFile, err := file.Open()
	if err != nil {
		return
	}
	byteFile, err := io.ReadAll(multipartFile)
	if err != nil {
		return
	}

	conn := excel.NewConnecter()
	if errs := conn.OpenBinary(byteFile); errs != nil {
		return
	}
	defer conn.Close()

	rd, err := conn.NewReaderByConfig(&excel.Config{
		Sheet: "Data",
	})
	if err != nil {
		return
	}
	defer rd.Close()

	if errs := rd.ReadAll(&data); errs != nil {
		return
	}
	return
}
