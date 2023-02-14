namespace go favorite

//赞操作
struct FavoriteActionRequest {
  1: string Token // 用户鉴权token
  2: i64 VideoId // 视频id
  3: i32 ActionType // 1-点赞，2-取消点赞
}

struct FavoriteActionResponse {
  1: i32 StatusCode // 状态码，0-成功，其他值-失败
  2: optional string StatusMsg // 返回状态描述
}

//喜欢列表
struct FavoriteListRequest {
  1: i64 UserId // 用户id
  2: string Token // 用户鉴权token
}

struct FavoriteListResponse {
  1: i32 StatusCode // 状态码，0-成功，其他值-失败
  2: optional string StatusMsg // 返回状态描述
  3: list<Video> VideoList // 用户点赞视频列表
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

service FavoriteService {
    FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest Req)  //用户点赞
    FavoriteListResponse FavoriteList(1:FavoriteListRequest Req)    //用户点赞列表
}