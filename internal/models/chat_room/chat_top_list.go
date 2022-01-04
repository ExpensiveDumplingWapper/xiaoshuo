package chat_room

import (
	"context"

	"gorm.io/gorm"
)

type ChatTopList struct {
	Id          int   `json:"id"`
	UID         int   `json:"uid"`
	CreateTime  int64 `json:"create_time"`
	MemberCount int   `json:"member_count"`
	IsTop       int   `json:"is_top"`
}

func (ChatTopList) TableName() string {
	return "chat_top_list"
}

func DeleteByUid(ctx context.Context, tx *gorm.DB, acceptData *ChatTopList) (err error) {
	return tx.Where(" uid = ? ", acceptData.UID).Delete(&acceptData).Error
}

func CreatTopList(ctx context.Context, tx *gorm.DB, acceptData *ChatTopList) (err error) {
	return tx.Create(acceptData).Error
}
