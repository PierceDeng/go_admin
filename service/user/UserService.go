package user

import (
	"crypto/md5"
	"encoding/hex"
	"go_admin/config"
	"go_admin/middleware/cache"
	"go_admin/middleware/common"
	"go_admin/middleware/exception"
	"go_admin/model/entity"
	req "go_admin/model/reqVO"
	roleSerivce "go_admin/service/role"

	"github.com/google/uuid"
)
import resp "go_admin/model/respVO"

func Login(vo *req.UserLoginReqVO) *resp.UserLoginRespVO {

	var sysUser entity.SysUser
	result := config.DB.Where("user_name = ?", vo.Username).Take(&sysUser)

	if result != nil && result.Error != nil {
		panic(exception.NewBizException(common.BIZ_ERROR_CODE, "用户不存在"))
	}

	sum := md5.Sum([]byte(vo.Password))
	md5strPwd := hex.EncodeToString(sum[:])

	if sysUser.Password != md5strPwd {
		panic(exception.NewBizException(common.BIZ_ERROR_CODE, "密码不正确"))
	}

	var token = uuid.New().String()
	cache.SetSysToken(token, sysUser.UserId)

	return &resp.UserLoginRespVO{
		Token: token,
	}
}

func GetUserInfo(userId uint64) (respVO *resp.UserInfoRespVO) {

	var sysUser entity.SysUser
	result := config.DB.Where("user_id = ?", userId).Take(&sysUser)
	if result != nil && result.Error != nil {
		panic(exception.NewBizException(common.BIZ_ERROR_CODE, "用户不存在"))
	}

	roles := roleSerivce.GetRolePermission(sysUser)
	perms := roleSerivce.GetMenuPermission(sysUser)

	return &resp.UserInfoRespVO{
		User:               sysUser,
		Roles:              roles,
		Permissions:        perms,
		IsDefaultModifyPwd: false,
		IsPasswordExpired:  false,
	}

}
