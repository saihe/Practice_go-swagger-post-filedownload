package controller

import (
	"app/restapi/operations"
	"app/service"
	"log"

	"github.com/go-openapi/runtime/middleware"
)

type FileController interface {
	Generate(request operations.PostFileParams) middleware.Responder
}

type FileControllerImpl struct {
	service service.FileService
}

func NewFileController(s service.FileService) FileController {
	return &FileControllerImpl{
		service: s,
	}
}

func (c *FileControllerImpl) Generate(request operations.PostFileParams) middleware.Responder {
	log.Printf("%v\n", request.HTTPRequest)
	body := request.File
	fn, b := c.service.Generate(*body.ID, *body.Name)
	log.Printf("%v\n", b)

	return operations.NewPostFileOK().WithPayload(&operations.PostFileOKBody{
		Name: fn,
		Data: b,
	})
}
