package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/feed/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/feed"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
	"os"
	"testing"
	"time"
)

var Token string

func TestMain(m *testing.M) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(1),
	})
	if err != nil {
		klog.Errorf("create token fail, %v", err.Error())
		panic(err)
	}
	Token = token
	db.Init()
	oss.Init()
	path := oss.Path + "/public/test.mp4"
	file, err := os.Open(path)
	if err != nil {
		klog.Errorf("open local file %v fail", path)
	}
	defer file.Close()
	m.Run()
}

func TestFeed(t *testing.T) {
	type reqArgs struct {
		latestTime int64
		token      string
	}
	feedTest := []struct {
		name   string
		args   reqArgs
		result bool
	}{
		{
			name: "feed流",
			args: reqArgs{
				latestTime: time.Now().Unix(),
				token:      Token,
			},
			result: false,
		},
	}
	for _, tt := range feedTest {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := NewFeedService(context.Background()).Feed(&feed.FeedRequest{
				LatestTime: &tt.args.latestTime,
				Token:      &tt.args.token,
			})
			if (err != nil) != tt.result {
				t.Errorf("Feed() error = %v, result %v", err, tt.result)
				return
			}
			klog.Info(tt.name + "success")
		})
	}
}
