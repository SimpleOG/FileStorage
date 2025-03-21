package MongoDB

import "time"

type PhotoMetadata struct {
	ID         string    `bson:"_id"`         // Уникальный идентификатор файла
	UserID     int32     `bson:"user_id"`     // ID пользователя
	Name       string    `bson:"name"`        // Имя файла
	Size       int64     `bson:"size"`        // Размер файла
	Type       string    `bson:"type"`        // MIME-тип файла
	UploadedAt time.Time `bson:"uploaded_at"` // Время загрузки
	Tags       []string  `bson:"tags"`        // Кастомные теги
}
