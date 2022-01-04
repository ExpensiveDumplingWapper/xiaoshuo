package chat_room

import (
	"context"

	"gorm.io/gorm"
)

type ChatLive struct {
	RoomId              int64  `json:"room_id"`
	RoomName            string `json:"room_name"`
	CreateTime          int    `json:"create_time"`
	RoomMembers         string `json:"room_members"`
	IsPublic            int    `json:"is_public"`
	RoomOwner           int    `json:"room_owner"`
	RoomLanguage        string `json:"room_language"`
	RoomMemberTotal     int    `json:"room_member_total"`
	RoomMemberLockCount int    `json:"room_member_lock_count"`
	Region              string `json:"region"`
	RoomMemberCount     int    `json:"room_member_count"`
	RoomAllCount        int    `json:"room_all_count"`
	RoomTop             int    `json:"room_top"`
	Tags                string `json:"tags"`
}

func (ChatLive) TableName() string {
	return "chat_live"
}

// NewAcceptLog
func NewAcceptLog(ctx context.Context, tx *gorm.DB, acceptData *ChatLive) (err error) {
	// span, _ := apm.StartSpan(ctx, "mantch.join", "mysql")
	// defer span.End()
	err = tx.Create(&acceptData).Error
	return
}

func CheckByRoomId(ctx context.Context, tx *gorm.DB, acceptData *ChatLive) (err error, res ChatLive) {
	tx.Model(&acceptData).Where(" room_id = ? ", acceptData.RoomId).First(&res)
	return nil, res
}
