namespace go publish

struct Video {
    1: i64 Id// 视频唯一标识
    2: User Author // 视频作者信息
    3: string PlayUrl // 视频播放地址
    4: string CoverUrl// 视频封面地址
    5: i64 FavoriteCount// 视频的点赞总数
    6: i64 CommentCount // 视频的评论总数
    7: bool IsFavorite// true-已点赞，false-未点赞
    8: string title// 视频标题
}

struct User {
    1: i64 Id// 用户id
    2: string Name // 用户名称
    3: optional i64 FollowCount // 关注总数
    4: optional i64 FollowerCount// 粉丝总数
    5: required bool IsFollow// true-已关注，false-未关注
}

//视频投稿
struct PublishActionRequest {
    1: string Token// 用户鉴权token
    2: binary Data
    3: string Title// 视频标题
}

struct PublishActionResponse {
    1: i32 StatusCode// 状态码，0-成功，其他值-失败
    2: optional string StatusMsg// 返回状态描述
}

//发布列表
struct PublishListRequest {
    1: i64 UserId// 用户id
    2: string Token// 用户鉴权token
}

struct PublishListResponse {
    1: i32 StatusCode// 状态码，0-成功，其他值-失败
    2: optional string StatusMsg// 返回状态描述
    3: list<Video> VideoList// 用户发布的视频列表
}

service PublishService {
    PublishListResponse PublishList (1: PublishListRequest Req) //获取用户发布作品
    PublishActionResponse PublishAction (1: PublishActionRequest Req)   //视频投稿
}