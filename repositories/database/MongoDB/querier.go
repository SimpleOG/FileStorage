package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"io"
	"time"
)

type Querier struct {
	db         *mongo.Database
	bucket     *gridfs.Bucket
	collection *mongo.Collection
}

func NewQueries(db *mongo.Database) Queries {
	bucket, _ := gridfs.NewBucket(db)
	collection := db.Collection("photos_metadata")
	return &Querier{
		db:         db,
		bucket:     bucket,
		collection: collection,
	}
}
func (q *Querier) SavePhoto(file io.Reader, metaData *PhotoMetadata) error {
	//создать уникальный id фотографии

	stream, err := q.bucket.OpenUploadStream(metaData.ID)
	if err != nil {
		return err
	}
	defer stream.Close()
	fileData, err := io.Copy(stream, file)
	if err != nil {
		return err
	}
	metaData.Size = fileData
	metaData.UploadedAt = time.Now()
	_, err = q.collection.InsertOne(context.Background(), metaData)
	if err != nil {
		return err
	}
	return nil
}
func (q *Querier) FindPhotoByID(id string) (*PhotoMetadata, error) {
	var photo PhotoMetadata
	err := q.collection.FindOne(context.Background(), bson.M{"_user_id": id}).Decode(&photo)
	if err != nil {
		return nil, err
	}
	return &photo, nil
}
func (q *Querier) DownloadPhotoByID(id string, writer io.Writer) error {
	_, err := q.bucket.DownloadToStreamByName(id, writer)
	return err
}

var _ Queries = (*Querier)(nil)
