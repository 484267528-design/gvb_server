package flag

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/pwd"
)

func CreatUser(permissions string) {
	//创建用户的逻辑
	//用户名，昵称，密码，确认密码，邮箱
	var (
		userName string
		nickName string
		passWord string
		rePasswd string
		email    string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入密码：")
	fmt.Scan(&passWord)
	fmt.Printf("请确认密码：")
	fmt.Scan(&rePasswd)
	fmt.Printf("请输入邮箱：")
	fmt.Scan(&email)

	//判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name =?", userName).Error
	if err == nil {
		//存在
		global.Log.Error("用户名已存在，请重新输入")
		return
	}
	//检验俩次密码
	if passWord != rePasswd {
		global.Log.Error("俩次密码不一致，请重新输入")
		return
	}
	//对密码进行hash
	hashPwd := pwd.HashPwd(passWord)

	//头像问题
	avatar := "/uploads/avatar/default.png"
	//入库
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SingEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功！", userName)

}
