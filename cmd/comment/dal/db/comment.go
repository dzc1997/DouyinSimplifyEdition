package db

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/gorm"
	"time"
)

type CommentRaw struct {
	gorm.Model
	UserId   int64  `gorm:"column:user_id;not null;index:idx_userid"`
	VideoId  int64  `gorm:"column:video_id;not null;index:idx_videoid"`
	Contents string `gorm:"column:contents;type:varchar(255);not null"`
}

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
	Name          string `gorm:"column:name;index:idx_username,unique;type:varchar(32);not null"`
	Password      string `gorm:"column:password;type:varchar(32);not null"`
	FollowCount   int64  `gorm:"column:follow_count;default:0"`
	FollowerCount int64  `gorm:"column:follower_count;default:0"`
}

type RelationRaw struct {
	gorm.Model
	UserId   int64 `gorm:"column:user_id;not null;index:idx_userid"`
	ToUserId int64 `gorm:"column:to_user_id;not null;index:idx_touserid"`
}


func (CommentRaw) TableName() string {
	return "comment"
}


func CreateComment(ctx context.Context, comment *CommentRaw) error {
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("comment").Create(&comment).Error
		if err != nil {
			klog.Error("create comment fail " + err.Error())
			return err
		}
		err = tx.Table("video").Where("id = ?", comment.VideoId).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			klog.Error("AddCommentCount error " + err.Error())
			return err
		}
		err = tx.Table("comment").First(&comment).Error
		if err != nil {
			klog.Errorf("find comment %v fail, %v", comment, err.Error())
			return err
		}
		return nil
	})
	return nil
}

func DeleteComment(ctx context.Context, commentId int64) (*CommentRaw, error) {
	var commentRaw *CommentRaw
	DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Table("comment").Where("id = ?", commentId).First(&commentRaw).Error
		if err == gorm.ErrRecordNotFound {
			klog.Errorf("not find comment %v, %v", commentRaw, err.Error())
			return err
		}
		if err != nil {
			klog.Errorf("find comment %v fail, %v", commentRaw, err.Error())
			return err
		}
		err = tx.Table("comment").Where("id = ?", commentId).Delete(&CommentRaw{}).Error
		if err != nil {
			klog.Error("delete comment fail " + err.Error())
			return err
		}
		err = tx.Table("video").Where("id = ?", commentRaw.VideoId).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error
		if err != nil {
			klog.Error("AddCommentCount error " + err.Error())
			return err
		}
		return nil
	})
	return commentRaw, nil
}

//通过评论id查询一组评论信息
func QueryCommentByCommentIds(ctx context.Context, commentIds []int64) ([]*CommentRaw, error) {
	var comments []*CommentRaw
	err := DB.WithContext(ctx).Table("comment").Where("id In ?", commentIds).Find(&comments).Error
	if err != nil {
		klog.Error("query comment by comment id fail " + err.Error())
		return nil, err
	}
	return comments, nil
}

//通过视频id号倒序返回一组评论信息
func QueryCommentByVideoId(ctx context.Context, videoId int64) ([]*CommentRaw, error) {
	var comments []*CommentRaw
	err := DB.WithContext(ctx).Table("comment").Order("updated_at desc").Where("video_id = ?", videoId).Find(&comments).Error
	if err != nil {
		klog.Error("query comment by video id fail " + err.Error())
		return nil, err
	}
	return comments, nil
}

func QueryVideoByVideoIds(ctx context.Context, videoIds []int64) ([]*VideoRaw, error) {
	var videos []*VideoRaw
	err := DB.WithContext(ctx).Where("id in (?)", videoIds).Find(&videos).Error
	if err != nil {
		klog.Error("QueryVideoByVideoIds error " + err.Error())
		return nil, err
	}
	return videos, nil
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
