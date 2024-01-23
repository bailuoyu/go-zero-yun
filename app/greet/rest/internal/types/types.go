// Code generated by goctl. DO NOT EDIT.
package types

type Empty struct {
}

type PageReq struct {
	Page     int `form:"page"`      //页码
	PageSize int `form:"page_size"` //每页条数
}

type Pager struct {
	Page      int `json:"page"`       //页码
	PageSize  int `json:"page_size"`  //每页条数
	PageCount int `json:"page_count"` //总页数
	Total     int `json:"total"`      //总条数
}

type RefreshTokenRsp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
	RefreshAfter int64  `json:"refresh_after"` // 建议客户端刷新token的绝对时间
}

type TestGetReq struct {
	Id int `form:"id"`
}

type TestGetRsp struct {
	Id       int    `json:"id"`
	Phone    string `json:"phone"`
	NickName string `json:"name"`
}

type TestAddReq struct {
	PageReq
	Name string `json:"name"`
	Age  int    `json:"age,optional"`
}

type TestElasticGetReq struct {
	Id string `form:"id"`
}

type TestElasticGetRsp struct {
	Time    string  `json:"time"`
	Trace   string  `json:"trace"`
	Span    string  `json:"span"`
	Route   string  `json:"route"`
	UserId  int     `json:"user_id"`
	Level   string  `json:"level"`
	Type    string  `json:"type"`
	Runtime float64 `json:"runtime"`
	Content string  `json:"content"`
	Caller  string  `json:"caller"`
}

type TestWxWorkAppChatCreateReq struct {
	Name     string   `json:"name"`
	Owner    string   `json:"owner"`
	Userlist []string `json:"userlist"`
	Chatid   string   `json:"chatid,optional"`
}

type TestWxWorkAppChatCreateRsp struct {
	Chatid string `json:"chatid"`
}

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRsp struct {
	AccessToken  string `json:"access_token"`
	AccessExpire int    `json:"access_expire"`
	RefreshAfter int    `json:"refresh_after"` // 建议客户端刷新token的绝对时间
}

type UserInfoReq struct {
	Name string `from:"name"`
}

type UserInfoRsp struct {
	Message string `json:"message"` //返回信息
}

type UserListRsp struct {
	Pager Pager             `json:"pager"`
	List  []UserListRspInfo `json:"list"`
}

type UserListRspInfo struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Remark string `json:"remark"`
}
