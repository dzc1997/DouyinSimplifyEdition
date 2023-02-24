package oss

import (
	"bytes"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PublishVideoToPublic(video []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		klog.Errorf("create %v fail, %v", filePath, err.Error())
		return err
	}
	defer file.Close()
	_, err = file.Write(video)
	if err != nil {
		klog.Errorf("write file fail, %v", err.Error())
		return err
	}
	return nil
}

func PublishVideoToOss(objectKey string, filePath string) error {
	err := Bucket.UploadFile(objectKey, filePath, 1024*1024, oss.Routines(3))
	if err != nil {
		klog.Errorf("publish %v to Oss fail, %v ", filePath, err.Error())
		return err
	}
	return nil
}

func QueryOssVideoURL(objectKey string) (string, error) {
	signedURL, err := Bucket.SignURL(objectKey, oss.HTTPPut, 60)
	if err != nil {
		klog.Errorf("Query %v Video URL fail, %v", objectKey, err.Error())
		return "", err
	}
	return signedURL, nil
}

func PublishCoverToOss(objectKey string, coverReader *bytes.Reader) error {
	err := Bucket.PutObject(objectKey, coverReader)
	if err != nil {
		klog.Errorf("publish %v to Oss fail, %v ", objectKey, err.Error())
		return err
	}
	return nil
}

func QueryOssCoverURL(objectKey string) (string, error) {
	signedURL, err := Bucket.SignURL(objectKey, oss.HTTPPut, 60)
	if err != nil {
		klog.Errorf("Query %v Cover URL fail, %v", objectKey, err.Error())
		return "", err
	}

	klog.Infof("QueryOssCoverURL objectKey %v signedURL %v\n", objectKey, signedURL)
	return signedURL, nil
}
