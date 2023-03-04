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

func GetMessage(ctx context.Context, myID int64, toID int64, t *int64) (res []*MessageRaw, err error) {
	curTime := time.Unix(*t, 0)
	conn := DB.WithContext(ctx).Where("created_at <= ?", curTime).Find(&res)
	if err = conn.Error; err != nil {
		return res, err
	}
	return res, nil
}

func PostMessage(ctx context.Context, message MessageRaw) (err error) {
	err = DB.WithContext(ctx).Table("message").Create(&message).Error
	if err != nil {
		klog.Error("create message fail " + err.Error())
		return err
	}
	return nil
}
