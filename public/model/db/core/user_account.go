package core

// UserAccount 用户账号表
type UserAccount struct {
	Id         int    `xorm:"'id' pk autoincr" json:"id"`       // INT(10)
	Phone      string `xorm:"'phone'" json:"phone"`             // default:'';comment:手机号;index;VARCHAR(20)
	UserNumber int    `xorm:"'user_number'" json:"user_number"` // default:0;comment:用户user_number;unique;INT(10)
	Status     int    `xorm:"'status'" json:"status"`           // default:1;comment:状态:1正常,2封禁;index;TINYINT(3)
	Remark     string `xorm:"'remark'" json:"remark"`           // default:'';comment:备注;VARCHAR(255)
	Nickname   string `xorm:"'nickname'" json:"nickname"`       // default:'我是一只喵星人';comment:昵称;index;VARCHAR(32)
	Face       string `xorm:"'face'" json:"face"`               // default:'default.png';comment:用户头像;VARCHAR(255)
	CreateTime int    `xorm:"'create_time'" json:"create_time"` // default:0;comment:创建时间;index;INT(10)
	UpdateTime int    `xorm:"'update_time'" json:"update_time"` // default:0;comment:更新时间;INT(10)
	ChangeTime int    `xorm:"'change_time'" json:"change_time"` // default:0;comment:资料修改时间，;INT(10)
	LastLogin  int    `xorm:"'last_login'" json:"last_login"`   // default:0;comment:最后一次登录时间;INT(10)
	OnlineTime int    `xorm:"'online_time'" json:"online_time"` // default:0;comment:在线时长;INT(10)
	Password   string `xorm:"'password'" json:"password"`       // default:'';comment:密码;VARCHAR(50)

}

// TableName 表名
func (t UserAccount) TableName() string {
	return "user_account"
}
