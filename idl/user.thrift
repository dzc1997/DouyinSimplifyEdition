namespace go user

//用户注册接口
struct UserRegisterRequest {
  1: string username // 注册用户名，最长32个字符
  2: string password // 密码，最长32个字符
}

struct UserRegisterResponse {
  1: i32 StatusCode // 状态码，0-成功，其他值-失败
  2: optional string StatusMsg// 返回状态描述
  3: i64 UserId // 用户id
  4: string Token // 用户鉴权token
}

//用户登录接口
struct UserLoginRequest {
  1: string username // 登录用户名
  2: string password // 登录密码
}

struct UserLoginResponse {
  1: i32 StatusCode // 状态码，0-成功，其他值-失败
  2: optional string StatusMsg // 返回状态描述
  3: required i64 UserId // 用户id
  4: required string Token // 用户鉴权token
}

//用户信息
struct UserRequest {
  1: i64 UserId // 用户id
  2: string Token // 用户鉴权token
}

struct UserResponse {
  1: i32 StatusCode // 状态码，0-成功，其他值-失败
  2: optional string StatusMsg// 返回状态描述
  3: User user// 用户信息
}

struct User {
  1: i64 Id// 用户id
  2: string Name // 用户名称
  3: optional i64 FollowCount // 关注总数
  4: optional i64 FollowerCount // 粉丝总数
  5: bool IsFollow // true-已关注，false-未关注
}

service UserService {
    UserRegisterResponse UserRegister (1: UserRegisterRequest Req) //注册
    UserLoginResponse UserLogin (1: UserLoginRequest Req)   //登录
    UserResponse UserInfo (1: UserRequest Req)  //获取用户信息
}