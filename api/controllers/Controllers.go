package controllers

import (
	"CustomFileStorage/repositories/database/MongoDB"
	"CustomFileStorage/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

type StorageControllersInterface interface {
	UploadPhoto(ctx *gin.Context)
	GetPhoto(ctx *gin.Context)
}

type StorageControllers struct {
	Service service.ServiceInterface
}

func NewStorageControllers(serviceInterface service.ServiceInterface) StorageControllersInterface {
	return &StorageControllers{
		Service: serviceInterface,
	}
}

func (s *StorageControllers) UploadPhoto(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("img")
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	defer file.Close()
	fileHeaderType := header.Header.Get("Content-Type")
	data := ctx.Request.FormValue("metadata")
	var metadata *MongoDB.PhotoMetadata
	if err = json.Unmarshal([]byte(data), &metadata); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	metadata.Type = fileHeaderType

	if err = s.Service.SavePhoto(file, metadata); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, metadata)
}
func (s *StorageControllers) GetPhoto(ctx *gin.Context) {
	id := ctx.Query("id")
	photo, err := s.Service.FindPhotoByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}
	writer := multipart.NewWriter(ctx.Writer)
	ctx.Header("Content-Type", writer.FormDataContentType())

	defer writer.Close()
	metadataPart, err := writer.CreateFormField("metadata")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	metadataJSON, _ := json.Marshal(photo)
	metadataPart.Write(metadataJSON)
	filePart, err := writer.CreateFormFile("file", photo.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if err := s.Service.DownloadPhotoByID(id, filePart); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, photo)

}
