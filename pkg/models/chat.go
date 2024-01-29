package models

import "gorm.io/gorm"

type Contact struct {
	UserId    int64 `json:"userId" gorm:"primaryKey"`
	ContactId int64 `json:"contactId" gorm:"primaryKey"`
	Approved  bool  `json:"approved" gorm:"not null,default:false"`
}

type Chat struct {
	Id             int64 `json:"id" gorm:"primaryKey"`
	User1ID        int64 `json:"user1Id"`
	User2ID        int64 `json:"user2Id"`
	gorm.DeletedAt `json:"deletedAt"`
}

type Message struct {
	ChatId      int64  `json:"chat_id" gorm:"references:Chat,constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User1Id     int64  `json:"user1Id"`
	User2Id     int64  `json:"user2Id"`
	Original    string `json:"original"`
	Translation string `json:"translation"`
}
