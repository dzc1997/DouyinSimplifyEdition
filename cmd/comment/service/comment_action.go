package service

import (
	"context"
	"errors"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/comment/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/comment/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/comment"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"sync"
)

type CommentActionService struct {
	ctx context.Context
}

func NewCommentActionService(ctx context.Context) *CommentActionService {
	return &CommentActionService{ctx: ctx}
}

func (s *CommentActionService) CommentAction(req *comment.CommentActionRequest) (*comment.Comment, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, err := Jwt.CheckToken(req.Token)
	if err != nil {
		return nil, err
	}
	videos, err := db.QueryVideoByVideoIds(s.ctx, []int64{req.VideoId})
	if err != nil {
		return nil, err
	}
	if len(videos) == 0 {
		return nil, errors.New("video not exist")
	}
	if req.ActionType == 1 {
		return CreateComment(currentId, req, s.ctx)
	} //评论
	if req.ActionType == 2 {
		return DeleteComment(currentId, req, s.ctx)
	}
	return nil, errors.New("create or delete fail")
}

func CreateComment(currentId int64, req *comment.CommentActionRequest, ctx context.Context) (*comment.Comment, error) {
	commentRaw := &db.CommentRaw{
		UserId:   currentId,
		VideoId:  req.VideoId,
		Contents: *req.CommentText,
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var user *db.UserRaw
	var commentErr, userErr error
	//创建评论记录并增加视频评论数
	go func() {
		defer wg.Done()
		err := db.CreateComment(ctx, commentRaw)
		if err != nil {
			commentErr = err
			return
		}
	}()
	//获取当前用户信息
	go func() {
		defer wg.Done()
		users, err := db.QueryUserByIds(ctx, []int64{currentId})
		if err != nil {
			userErr = err
			return
		}
		user = users[0]
	}()
	wg.Wait()
	if commentErr != nil {
		return nil, commentErr
	}
	if userErr != nil {
		return nil, userErr
	}

	comment_ := pack.CommentInfo(commentRaw, user)
	return comment_, nil
}

func DeleteComment(currentId int64, req *comment.CommentActionRequest, ctx context.Context) (*comment.Comment, error) {
	comments, err := db.QueryCommentByCommentIds(ctx, []int64{*req.CommentId})
	if err != nil {
		return nil, err
	}
	if len(comments) == 0 {
		return nil, errors.New("commentId not exist")
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var commentRaw *db.CommentRaw
	var userRaw *db.UserRaw
	var commentErr, userErr error
	//删除评论记录并减少视频评论数
	go func() {
		defer wg.Done()
		commentRaw, err = db.DeleteComment(ctx, *req.CommentId)
		if err != nil {
			commentErr = err
			return
		}
	}()
	//获取用户信息
	go func() {
		defer wg.Done()
		users, err := db.QueryUserByIds(ctx, []int64{currentId})
		if err != nil {
			userErr = err
			return
		}
		userRaw = users[0]
	}()
	wg.Wait()
	if commentErr != nil {
		return nil, commentErr
	}
	if userErr != nil {
		return nil, userErr
	}

	comment_ := pack.CommentInfo(commentRaw, userRaw)
	return comment_, nil
}
