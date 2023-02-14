namespace go relation

//关系操作
struct RelationActionRequest {
    1: string Token //用户鉴权token
    2: i64 ToUserId //对方用户id
    3: i32 ActionType // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值-失败
    2: optional string StatusMsg //返回状态描述
}

//用户关注列表
struct RelationFollowListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}

struct RelationFollowListResponse {
    1: i32 StatusCode //状态码，0成功，其他值-失败
    2: optional string StatusMsg //返回状态描述
    3: list<User> UserList //用户信息列表
}

struct User {
    1: i64 Id //用户id
    2: string Name //用户名称
    3: optional i64 FollowCount //关注总数
    4: optional i64 FollowerCount //粉丝总数
    5: bool IsFollow // true- 已关注，false-未关注
}

//用户粉丝列表
struct RelationFollowerListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}

struct RelationFollowerListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: optional string StatusMsg //返回状态描述
    3: list<User> UserList //用户列表
}

//用户好友列表
struct RelationFriendListRequest {
    1: i64 UserId //用户id
    2: string Token //用户鉴权token
}

struct RelationFriendListResponse {
    1: i32 StatusCode //状态码，0- 成功，其他值失败
    2: optional string StatusMsg //返回状态描述
    3: list<FriendUser> UserList //用户列表
}

struct FriendUser {
    1: i64 Id //用户id
    2: string Name //用户名称
    3: optional i64 FollowCount //关注总数
    4: optional i64 FollowerCount //粉丝总数
    5: bool IsFollow // true- 已关注，false-未关注
    6: string Avatar    //用户头像url
    7: optional string Message //和好友的最新聊天记录
    8: i64 MsgType  //message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

service RelationService {
    RelationActionResponse RelationAction(1:RelationActionRequest Req)// 关系操作
    RelationFollowListResponse RelationFollowList(1:RelationFollowListRequest Req)// 关注列表
    RelationFollowerListResponse RelationFollowerList(1:RelationFollowerListRequest Req)// 粉丝列表
    RelationFriendListResponse RelationFriendList(1:RelationFriendListRequest Req)// 好友列表
}
