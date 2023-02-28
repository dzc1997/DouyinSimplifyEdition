namespace go relation

//关系操作
struct RelationActionRequest {
    1: string token //用户鉴权token
    2: i64 to_user_id //对方用户id
    3: i32 action_type // 1-关注，2-取消关注
}

struct RelationActionResponse {
    1: i32 status_code //状态码，0- 成功，其他值-失败
    2: optional string status_msg //返回状态描述
}

//用户关注列表
struct RelationFollowListRequest {
    1: i64 user_id //用户id
    2: string token //用户鉴权token
}

struct RelationFollowListResponse {
    1: i32 status_code //状态码，0- 成功，其他值-失败
    2: optional string status_msg //返回状态描述
    3: list<User> user_list //用户信息列表
}

struct User {
  1: i64 id// 用户id
  2: string name // 用户名称
  3: optional i64 follow_count // 关注总数
  4: optional i64 follower_count // 粉丝总数
  5: bool is_follow // true-已关注，false-未关注
  6: optional string avatar // 用户头像
  7: optional string background_image // 用户个人页顶部大图
  8: optional string signature // 个人简介
  9: optional i64 total_favorited // 获赞数量
  10: optional i64 work_count  // 作品数量
  11: optional i64 favorite_count //点赞数量
}

//用户粉丝列表
struct RelationFollowerListRequest {
    1: i64 user_id //用户id
    2: string token //用户鉴权token
}

struct RelationFollowerListResponse {
    1: i32 status_code //状态码，0- 成功，其他值-失败
    2: optional string status_msg //返回状态描述
    3: list<User> user_list //用户列表
}

//用户好友列表
struct RelationFriendListRequest {
    1: i64 user_id //用户id
    2: string token //用户鉴权token
}

struct RelationFriendListResponse {
    1: i32 status_code //状态码，0- 成功，其他值-失败
    2: optional string status_msg //返回状态描述1: i32 StatusCode //状态码，0- 成功，其他值失败
    3: list<FriendUser> user_list //用户列表
}

struct FriendUser {
    1: i64 id// 用户id
    2: string name // 用户名称
    3: optional i64 follow_count // 关注总数
    4: optional i64 follower_count // 粉丝总数
    5: bool is_follow // true-已关注，false-未关注
    6: optional string avatar // 用户头像
    7: optional string background_image // 用户个人页顶部大图
    8: optional string signature // 个人简介
    9: optional i64 total_favorited // 获赞数量
    10: optional i64 work_count  // 作品数量
    11: optional i64 favorite_count //点赞数量
    12: optional string message //和好友的最新聊天记录
    13: optional i64 msg_type  //message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

service RelationService {
    RelationActionResponse RelationAction(1:RelationActionRequest Req)// 关系操作
    RelationFollowListResponse RelationFollowList(1:RelationFollowListRequest Req)// 关注列表
    RelationFollowerListResponse RelationFollowerList(1:RelationFollowerListRequest Req)// 粉丝列表
    RelationFriendListResponse RelationFriendList(1:RelationFriendListRequest Req)// 好友列表
}
