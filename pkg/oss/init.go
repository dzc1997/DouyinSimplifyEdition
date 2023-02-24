package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"os"
	"strings"
)

var Bucket *oss.Bucket
var Path string

func Init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Path = strings.Split(dir, "/cmd")[0]
	endpoint := constants.OssEndPoint
	accesskeyid := constants.OssAccessKeyId
	accessKeySecret := constants.OssAccessKeySecret
	bucket := constants.OssBucket
	client, err := oss.New(endpoint, accesskeyid, accessKeySecret)
	if err != nil {
		panic(err)
	}
	Bucket, err = client.Bucket(bucket)
	if err != nil {
		panic(err)
	}
}
