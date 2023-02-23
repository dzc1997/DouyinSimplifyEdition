package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"gorm.io/gorm"
)

type UserRaw struct {
	gorm.Model
	Name            string `gorm:"column:name;index:idx_username,unique;type:varchar(32);not null"`
	Password        string `gorm:"column:password;type:varchar(32);not null"`
	FollowCount     int64  `gorm:"column:follow_count;default:0"`
	FollowerCount   int64  `gorm:"column:follower_count;default:0"`
	Avatar          string `gorm:"column:avatar,type:varchar(100);not null"`           // 用户头像
	BackgroundImage string `gorm:"column:background_image,type:varchar(100);not null"` // 用户个人页顶部大图
	Signature       string `gorm:"column:signature,type:varchar(1000);not null"`       // 个人简介
	TotalFavorited  string `gorm:"column:total_favorited;type:varchar(1000);not null"` // 获赞数量
	WorkCount       int64  `gorm:"column:work_count;default:0"`                        // 作品数量
	FavoriteCount   int64  `gorm:"column:favorite_count;default:0"`                    // 点赞数量
}

type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

func (UserRaw) TableName() string {
	return constants.UserTableName
}

func (RelationRaw) TableName() string {
	return constants.RelationTableName
}

func QueryUserByIds(ctx context.Context, userIds []int64) ([]*UserRaw, error) {
	var users []*UserRaw
	err := DB.WithContext(ctx).Where("id in (?)", userIds).Find(&users).Error
	if err != nil {
		klog.Error("query user by ids fail " + err.Error())
		return nil, err
	}
	return users, nil
}

func QueryUserByName(ctx context.Context, userName string) ([]*UserRaw, error) {
	var users []*UserRaw
	err := DB.WithContext(ctx).Where("name = ?", userName).Find(&users).Error
	if err != nil {
		klog.Error("query user by name fail " + err.Error())
		return nil, err
	}
	return users, nil
}

func UploadUserInfo(ctx context.Context, username string, password string) (int64, error) {
	user := &UserRaw{
		Name:          username,
		Password:      password,
		FollowCount:   0,
		FollowerCount: 0,
	}
	err := DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		klog.Error("upload user data fail " + err.Error())
		return 0, err
	}
	return int64(user.ID), nil
}

func QueryRelationByIds(ctx context.Context, currentId int64, userIds []int64) (map[int64]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Where("user_id = ? AND to_user_id IN ?", currentId, userIds).Find(&relations).Error
	if err != nil {
		klog.Error("query relation by ids " + err.Error())
		return nil, err
	}
	relationMap := make(map[int64]*RelationRaw)
	for _, relation := range relations {
		relationMap[relation.ToUserId] = relation
	}
	return relationMap, nil
}
