package db

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"gorm.io/gorm"
)

type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}

type FriendRaw struct {
	gorm.Model
	UserId     int64 `gorm:"column:user_id;not null;index:idx_userid"`
	FriendList int64 `gorm:"column:friend_list;index:idx_friendlist"`
}

func (FriendRaw) TableName() string {
	return constants.FriendTableName
}

func (RelationRaw) TableName() string {
	return constants.RelationTableName
}

func (UserRaw) TableName() string {
	return constants.UserTableName
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

func Create(ctx context.Context, currentId int64, toUserId int64) error {
	relationRaw := &RelationRaw{
		UserId:   currentId,
		ToUserId: toUserId,
	}
	_ = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error
		if err != nil {
			klog.Error("add user follow_count fail " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error
		if err != nil {
			klog.Error("add user follower_count fail " + err.Error())
			return err
		}

		err = tx.Table("relation").Create(&relationRaw).Error
		if err != nil {
			klog.Error("create relation record fail " + err.Error())
			return err
		}

		return nil
	})
	return nil
}

func Delete(ctx context.Context, currentId int64, toUserId int64) error {
	var relationRaw *RelationRaw
	_ = DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("user").Where("id = ?", currentId).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user follow_count fail " + err.Error())
			return err
		}

		err = tx.Table("user").Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error
		if err != nil {
			klog.Error("sub user follower_count fail " + err.Error())
			return err
		}

		err = tx.Table("relation").Where("user_id = ? AND to_user_id = ?", currentId, toUserId).Delete(&relationRaw).Error
		if err != nil {
			klog.Error("delete relation record fali " + err.Error())
			return err
		}
		return nil
	})
	return nil
}

func QueryFollowById(ctx context.Context, userId int64) ([]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follow by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
}

func QueryFollowerById(ctx context.Context, userId int64) ([]*RelationRaw, error) {
	var relations []*RelationRaw
	err := DB.WithContext(ctx).Table("relation").Where("to_user_id = ?", userId).Find(&relations).Error
	if err != nil {
		klog.Error("query follower by id fail " + err.Error())
		return nil, err
	}
	return relations, nil
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
	TotalFavorited  int64 `gorm:"column:total_favorited;default:0"` // 获赞数量
	WorkCount       int64  `gorm:"column:work_count;default:0"`                        // 作品数量
	FavoriteCount   int64  `gorm:"column:favorite_count;default:0"`                    // 点赞数量
}

type FriendUser struct {
	Name            string
	Password        string
	FollowCount     int64
	FollowerCount   int64
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  string
	WorkCount       int64
	FavoriteCount   int64
	Message         string
	MsgType         int64
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

