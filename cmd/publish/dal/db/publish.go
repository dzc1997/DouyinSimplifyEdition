package db

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"gorm.io/gorm"
)

type VideoRaw struct {
	gorm.Model
	UserId        int64     `gorm:"column:user_id;not null;index:idx_userid"`
	Title         string    `gorm:"column:title;type:varchar(128);not null"`
	PlayUrl       string    `gorm:"column:play_url;varchar(128);not null"`
	CoverUrl      string    `gorm:"column:cover_url;varchar(128);not null"`
	FavoriteCount int64     `gorm:"column:favorite_count;default:0"`
	CommentCount  int64     `gorm:"column:comment_count;default:0"`
	UpdatedAt     time.Time `gorm:"column:update_time;not null;index:idx_update"`
}

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

type FavoriteRaw struct {
	gorm.Model
	UserId  int64 `gorm:"column:user_id;not null;index:idx_userid"`
	VideoId int64 `gorm:"column:video_id;not null;index:idx_videoid"`
}

type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

func (FavoriteRaw) TableName() string {
	return constants.FavoriteTableName
}

func (VideoRaw) TableName() string {
	return constants.VideoTableName
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

func PublishVideoInfo(ctx context.Context, videoData *VideoRaw) error {
	if err := DB.WithContext(ctx).Create(&videoData).Error; err != nil {
		klog.Error("PublishVideoData error " + err.Error())
		return err
	}
	return nil
}

func QueryVideoByUserId(ctx context.Context, userId int64) ([]*VideoRaw, error) {
	var videos []*VideoRaw
	err := DB.Table("video").WithContext(ctx).Order("update_time desc").Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByUserId find video error " + err.Error())
		return nil, err
	}
	return videos, nil
}

func QueryFavoriteByIds(ctx context.Context, currentId int64, videoIds []int64) (map[int64]*FavoriteRaw, error) {
	var favorites []*FavoriteRaw
	err := DB.WithContext(ctx).Where("user_id = ? AND video_id IN ?", currentId, videoIds).Find(&favorites).Error
	if err != nil {
		klog.Error("query favorite record fail " + err.Error())
		return nil, err
	}
	favoriteMap := make(map[int64]*FavoriteRaw)
	for _, favorite := range favorites {
		favoriteMap[favorite.VideoId] = favorite
	}
	return favoriteMap, nil
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
