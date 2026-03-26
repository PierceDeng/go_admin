package user

import (
	"crypto/md5"
	"encoding/hex"
	"go_admin/config"
	"go_admin/middleware/cache"
	"go_admin/middleware/common"
	"go_admin/middleware/exception"
	"go_admin/model/RespVO/menu"
	"go_admin/model/entity"
	req "go_admin/model/reqVO"
	"go_admin/model/reqVO/user"
	"go_admin/repository/dept"
	userRepository "go_admin/repository/user"
	roleSerivce "go_admin/service/role"

	"github.com/google/uuid"
)
import resp "go_admin/model/respVO"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (UserService) Login(vo *req.UserLoginReqVO) *resp.UserLoginRespVO {

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

func (UserService) GetUserInfo(userId uint64) (respVO *resp.UserInfoRespVO) {

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

func (u UserService) GetDeptTree(sysDept *entity.SysDept) []*menu.MenuTreeSelect {

	deptList := dept.SelectDeptList(sysDept)
	return u.BuildDeptTree(deptList)
}

func (u UserService) BuildDeptTree(list []*entity.SysDept) []*menu.MenuTreeSelect {
	trees := u.buildTree(list) // 先构建树
	result := make([]*menu.MenuTreeSelect, 0, len(trees))
	for _, root := range trees {
		result = append(result, u.toTreeSelect(root))
	}
	return result
}

func (UserService) buildTree(list []*entity.SysDept) []*entity.SysDept {
	// 1. 建立 id -> node 映射，并初始化 Children 切片（防止 nil）
	nodeMap := make(map[int64]*entity.SysDept)
	for _, d := range list {
		if d == nil {
			continue
		}
		nodeMap[d.DeptId] = d
		d.Children = []*entity.SysDept{} // 清空原有子节点
	}

	var roots []*entity.SysDept
	// 2. 建立父子关系
	for _, m := range list {
		if m == nil {
			continue
		}
		// 根节点条件：ParentId == 0（可根据实际业务调整）
		if m.ParentId == 0 {
			roots = append(roots, m)
		} else {
			if parent, ok := nodeMap[m.ParentId]; ok {
				parent.Children = append(parent.Children, m)
			} else {
				// 父节点不存在，也作为根节点（兜底）
				roots = append(roots, m)
			}
		}
	}
	return roots
}

func (u UserService) toTreeSelect(sysDept *entity.SysDept) *menu.MenuTreeSelect {
	if sysDept == nil {
		return nil
	}
	ts := &menu.MenuTreeSelect{
		ID:       sysDept.DeptId,
		Label:    sysDept.DeptName,
		Children: []*menu.MenuTreeSelect{},
	}
	for _, child := range sysDept.Children {
		ts.Children = append(ts.Children, u.toTreeSelect(child))
	}
	return ts
}

func (u UserService) GetUserList(vo *user.SysUserReqVO) resp.PageResp[entity.SysUser] {

	var r resp.PageResp[entity.SysUser]
	userList, total, _ := userRepository.QueryUserList(*vo)
	r.Rows = userList
	r.Total = total
	r.Code = 200
	return r
}
