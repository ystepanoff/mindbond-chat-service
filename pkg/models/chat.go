package models

import "gorm.io/gorm"

type Chat struct {
	Id      int64 `json:"id" gorm:"primaryKey"`
	User1ID int64 `json:"user1Id"`
	User2ID int64 `json:"user2Id"`
	gorm.DeletedAt
}
