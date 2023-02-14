namespace go feed

//视频流接口
struct FeedRequest {
    1: optional i64 LatestTime// 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string Token// 可选参数，登录用户设置
}

struct FeedResponse {
    1: i32 StatusCode// 状态码，0-成功，其他值-失败
    2: optional string StatusMsg// 返回状态描述
    3: Video VideoList// 视频列表
    4: optional i64 NextTime// 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

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

service FeedService {
    FeedResponse Feed(1: FeedRequest Req)   //获取视频流
}