package constants

import "time"

const (
	MySQLDefaultDSN    = "root:19971002@tcp(127.0.0.1:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local" //MySQL DSN
	EtcdAddress        = "127.0.0.1:2379"                                                                        //Etcd 地址
	ApiAddress         = ":8080"                                                                                 //Api层 地址
	FeedAddress        = "127.0.0.1:8081"                                                                        //Feed 服务地址
	PublishAddress     = "127.0.0.1:8082"                                                                        //Publish 服务地址
	UserAddress        = "127.0.0.1:8083"                                                                        //User服务地址
	FavoriteAddress    = "127.0.0.1:8084"                                                                        //Favorite服务地址
	CommentAddress     = "127.0.0.1:8085"                                                                        //Comment服务地址
	RelationAddress    = "127.0.0.1:8086"                                                                        //Relation服务地址
	OssEndPoint        = "oss-cn-hangzhou.aliyuncs.com"                                                          //Oss
	OssAccessKeyId     = ""
	OssAccessKeySecret = ""
	OssBucket          = "dynn"
	VideoTableName    = "video"
	UserTableName     = "user"
	FavoriteTableName = "favorite"
	CommentTableName  = "comment"
	RelationTableName = "relation"
	SecretKey           = "secret key"
	IdentiryKey         = "id"
	TimeFormat = "2006-01-02 15:04:05"
	Like   = 1
	Unlike = 2
	AddComment = 1
	DelComment = 2
	Follow   = 1
	UnFollow = 2
	ApiServiceName      = "api"
	FeedServiceName     = "feed"
	PublishServiceName  = "publish"
	UserServiceName     = "user"
	FavoriteServiceName = "favorite"
	CommentServiceName  = "comment"
	RelationServiceName = "relation"
	NeedCPURateLimit = false
	CPURateLimit     = 80.0
	MySQLMaxIdleConns    = 20
	MySQLMaxOpenConns    = 100
	MySQLConnMaxLifetime = 2 * time.Hour

)
