package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	Oldpwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"`
	Newpwd string `json:"new_pwd" binding:"required" msg:"请输入新密码"`
}

func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr UpdatePasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	//判断密码是否一致
	if !pwd.CheckPwd(user.Password, cr.Oldpwd) {
		res.FailWithMessage("密码错误", c)
		return
	}
	hashpwd := pwd.HashPwd(cr.Newpwd)
	err = global.DB.Model(&user).Update("password", hashpwd).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("用户密码修改失败", c)
		return
	}
	res.OkWithMessage("密码修改成功", c)
	return
}
