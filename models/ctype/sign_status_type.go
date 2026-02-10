package ctype

import "encoding/json"

type SignStatus int

const (
	SingQQ    SignStatus = 1 //qq登录
	SingGite  SignStatus = 2 //Github
	SingEmail SignStatus = 3 //邮箱登录

)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}
func (s SignStatus) String() string {
	var str string
	switch s {
	case SingQQ:
		str = "QQ登录"
	case SingGite:
		str = "Github"
	case SingEmail:
		str = "邮箱登录"
	default:
		str = "其他"
	}
	return str
}
