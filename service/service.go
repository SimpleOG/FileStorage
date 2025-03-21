package service

import (
	"CustomFileStorage/repositories/database/MongoDB"
	"fmt"
	"io"
	"time"
)

type ServiceInterface interface {
	SavePhoto(file io.Reader, metadata *MongoDB.PhotoMetadata) error
	FindPhotoByID(id string) (*MongoDB.PhotoMetadata, error)
	DownloadPhotoByID(id string, writer io.Writer) error
}

type Service struct {
	db MongoDB.Queries
}

func NewService(db MongoDB.Queries) ServiceInterface {
	return &Service{db: db}
}
func (s *Service) SavePhoto(file io.Reader, metadata *MongoDB.PhotoMetadata) error {
	fileID := fmt.Sprintf("%d", time.Now().UnixNano())
	metadata.ID = fileID
	err := s.db.SavePhoto(file, metadata)
	if err != nil {
		return err
	}
	return err
}
func (s *Service) FindPhotoByID(id string) (*MongoDB.PhotoMetadata, error) {
	meta, err := s.db.FindPhotoByID(id)
	if err != nil {
		return nil, err
	}
	return meta, nil

}
func (s *Service) DownloadPhotoByID(id string, writer io.Writer) error {
	err := s.db.DownloadPhotoByID(id, writer)
	if err != nil {
		return err
	}
	return nil
}
