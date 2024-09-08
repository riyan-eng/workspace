package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"server/env"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type fileStruct struct {
	c *gin.Context
}

func NewFile(c *gin.Context) *fileStruct {
	return &fileStruct{c: c}
}

type fileMeta struct {
	Id          *string
	Name        *string
	Type        *string
	Size        *int
	ContentType *string
	Path        *string
	Url         *string
}

func (m *fileStruct) SaveLocal(file *multipart.FileHeader) (*fileMeta, *Error) {
	data := new(fileMeta)

	fileName := new(string)
	fileName = &file.Filename
	multipartFile, errT := file.Open()
	if errT != nil {
		return data, &Error{Errors: errT.Error()}
	}

	byteFile, errT := io.ReadAll(multipartFile)
	if errT != nil {
		return data, &Error{Errors: errT.Error()}
	}

	contentType := http.DetectContentType(byteFile)
	ext := filepath.Ext(file.Filename)
	fileType := ext[1:]
	fileSize := int(file.Size / 1000)

	fileId := uuid.NewString()

	path := fmt.Sprintf(`./media/%s.%s`, fileId, fileType)
	url := fmt.Sprintf(`%s/object/%s/%s`, env.NewEnv().SERVER_HOST_BE, fileId, *fileName)

	go func(file *multipart.FileHeader, path string) {
		m.c.SaveUploadedFile(file, path)
	}(file, path)

	return &fileMeta{
		Name:        fileName,
		Type:        &fileType,
		Size:        &fileSize,
		ContentType: &contentType,
		Path:        &path,
		Id:          &fileId,
		Url:         &url,
	}, &Error{}
}

func (f *fileStruct) GetFileSizeString(size int) (stringSize string) {
	switch len(strconv.Itoa(size)) {
	case 4:
		stringSize = fmt.Sprintf(`%.2f MB`, float64(size)/1000)
		return
	default:
		stringSize = fmt.Sprintf(`%v KB`, size)
		return
	}
}
