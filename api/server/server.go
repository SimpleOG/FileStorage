package server

import (
	"CustomFileStorage/api/controllers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	Storage controllers.StorageControllersInterface
}
type ServerInterface interface {
	Start(address string) error
}

func NewServer(gin *gin.Engine, Storage controllers.StorageControllersInterface) ServerInterface {
	return &Server{
		router:  gin,
		Storage: Storage}
}
func (s *Server) Start(address string) error {
	s.AddRoutes()
	return s.router.Run(address)
}

func (s *Server) AddRoutes() {
	s.router.POST("/upload_image", s.Storage.UploadPhoto)
	s.router.GET("/get_image", s.Storage.GetPhoto)
}
