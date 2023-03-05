package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"gorm.io/gorm"
	"time"
)

type MessageRaw struct {
	gorm.Model
	FromUserID int64     `gorm:"column:from_user_id;NOT NULL"`
	ToUserID   int64     `gorm:"column:to_user_id;NOT NULL"`
	Content    string    `gorm:"column:content;NOT NULL"`
	CreateTime time.Time `gorm:"column:create_time;not null;index:idx_create"`
}

func (MessageRaw) TableName() string {
	return constants.MessageTableName
}

func PostMessage(ctx context.Context, message MessageRaw) (err error) {
	err = DB.WithContext(ctx).Table("message").Create(&message).Error
	if err != nil {
		klog.Error("create message fail " + err.Error())
		return err
	}
	return nil
}

func QueryMessageById(ctx context.Context, userId int64, toUserId int64) ([]*MessageRaw, error) {
	var messages []*MessageRaw
	err := DB.WithContext(ctx).Table("message").Where("from_user_id = ? and to_user_id = ?) or (to_user_id = ? and from_user_id = ?", userId, toUserId, userId, toUserId).Find(&messages).Error
	if err != nil {
		klog.Error("query message by id fail " + err.Error())
		return nil, err
	}
	return messages, nil
}
