package constants

import "time"

const (
	MySQLDefaultDSN    = "root:19971002@tcp(127.0.0.1:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddress        = "127.0.0.1:2379"
	ApiAddress         = ":8080"
	FeedAddress        = "127.0.0.1:8081"
	PublishAddress     = "127.0.0.1:8082"
	UserAddress        = "127.0.0.1:8083"
	FavoriteAddress    = "127.0.0.1:8084"
	CommentAddress     = "127.0.0.1:8085"
	RelationAddress    = "127.0.0.1:8086"
	MessageAddress	= "127.0.0.1:8087"

	OssEndPoint        = "oss-cn-hangzhou.aliyuncs.com"
	OssAccessKeyId     = ""
	OssAccessKeySecret = ""
	OssBucket          = "dynn"

	VideoTableName    = "video"
	UserTableName     = "user"
	FavoriteTableName = "favorite"
	CommentTableName  = "comment"
	RelationTableName = "relation"
	FriendTableName   = "friend"
	MessageTableName = "message"

	SecretKey   = "dzc1997"
	IdentiryKey = "id"

	TimeFormat = "2006-01-02 15:04:05"

	Like       = 1
	Unlike     = 2
	AddComment = 1
	DelComment = 2
	Follow     = 1
	UnFollow   = 2
	SendMessage = 1

	ApiServiceName      = "api"
	FeedServiceName     = "feed"
	PublishServiceName  = "publish"
	UserServiceName     = "user"
	FavoriteServiceName = "favorite"
	CommentServiceName  = "comment"
	RelationServiceName = "relation"
	MessageServiceName = "message"

	MySQLMaxIdleConns    = 10
	MySQLMaxOpenConns    = 100
	MySQLConnMaxLifetime = time.Hour
)
