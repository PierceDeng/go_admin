package user

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/google/uuid"
	"go_admin/config"
	"go_admin/middleware/cache"
	"go_admin/middleware/exception"
	"go_admin/model/entity"
	req "go_admin/model/reqVO"
)
import resp "go_admin/model/respVO"

func Login(vo req.UserLoginReqVO) resp.UserLoginRespVO {

	var sysUser entity.SysUser
	result := config.DB.Where("user_name = ?", vo.Username).First(&sysUser)

	if result.Error != nil {
		panic(exception.NewBizException(10001, "用户不存在"))
	}

	sum := md5.Sum([]byte(vo.Password))
	md5strPwd := hex.EncodeToString(sum[:])

	if sysUser.Password != md5strPwd {
		panic(exception.NewBizException(10001, "密码不正确"))
	}

	var token = uuid.New().String()
	cache.SetSysToken(token, sysUser.UserId)

	return resp.UserLoginRespVO{
		Token: token,
	}

}

func GetInfo() {

}
