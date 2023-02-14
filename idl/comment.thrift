namespace go comment

//评论操作
struct CommentActionRequest {
    1: string Token //用户鉴权token
    2: i64 VideoId //视频id
    3: i32 ActionType // 1- 发布评论，2- 删除评论
    4: optional string CommentText //用户填写的评论内容，在action_type=1 的时候使用
    5: optional i64 CommentId //要删除的评论id,在action_type=2的时候使用
}

struct CommentActionResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: optional string StatusMsg //返回状态描述
    3: optional Comment Comment //评论成功返回评论内容，不需要重新拉取整个列表
}

struct Comment {
    1: i64 Id //视频评论id
    2: User User //评论用户信息
    3: string Content //评论内容
    4: string CreateDate //评论发布日期，格式mm-dd
}

//视频评论列表
struct CommentListRequest {
    1: string Token //用户鉴权token
    2: i64 VideoId //视频id
}

struct CommentListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: optional string StatusMsg //返回状态描述
    3: list<Comment> CommentList //评论列表
}

struct User {
    1: i64 Id //用户id
    2: string Name //用户名称
    3: optional i64 FollowCount //关注总数
    4: optional i64 FollowerCount //粉丝总数
    5: bool IsFollow // true-已关注，false-未关注
}

service CommentService {
    // 用户评论
    CommentActionResponse CommentAction(1:CommentActionRequest Req)// 用户评论
    CommentListResponse CommentList(1:CommentListRequest Req)// 用户评论列表
}