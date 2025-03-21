package MongoDB

import "io"

type Queries interface {
	SavePhoto(file io.Reader, metaData *PhotoMetadata) error
	FindPhotoByID(id string) (*PhotoMetadata, error)
	DownloadPhotoByID(id string, writer io.Writer) error
}
