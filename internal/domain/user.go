package domain

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:varchar(100)"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null;size:50;"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null;size:100;"`
	AvatarURL string    `json:"avatar_url" gorm:"size:255;"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	CreateOrUpdate(user *User) error
	GetByID(id string) (*User, error)
}

type UserUseCase interface {
	SyncUserFromAuth(user *User) error
}
