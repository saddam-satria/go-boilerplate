package app

import (
	"time"

	"github.com/google/uuid"
)

type UserCredential struct {
	UserCredentialId uuid.UUID  `json:"userCredentialId" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email            string     `json:"email" gorm:"index:email,unique;null;type:varchar(255)"`
	Password         string     `json:"password" gorm:"type:varchar(255)"`
	CreatedAt        *time.Time `json:"createdAt" gorm:"default_now()"`
	UpdatedAt        *time.Time `json:"updatedAt" gorm:"autoUpdateTime:milli"`
	IsDeleted        bool       `json:"isDeleted" gorm:"default:false"`
}
