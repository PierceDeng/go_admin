package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go_admin/config"
	"go_admin/middleware/cache"
	"go_admin/middleware/common"
	"go_admin/middleware/db"
	"go_admin/middleware/exception"
	"go_admin/model/RespVO/menu"
	"go_admin/model/entity"
	"go_admin/model/reqVO"
	userReqVO "go_admin/model/reqVO/user"
	userRespVO "go_admin/model/respVO/user"
	"go_admin/repository/dept"
	userRepository "go_admin/repository/user"
	"go_admin/service/role"
	"log"

	"github.com/google/uuid"
	"github.com/samber/lo"
)
import resp "go_admin/model/respVO"

type UserService struct {
	*role.RoleService
	*userRepository.SysUserPostRepository
}

func NewUserService() *UserService {
	return &UserService{
		RoleService:           role.NewRoleService(),
		SysUserPostRepository: userRepository.NewSysUserPostRepository(),
	}
}

func (*UserService) Login(vo *reqVO.UserLoginReqVO) *resp.UserLoginRespVO {

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

func (tis *UserService) GetUserInfo(userId uint64) (respVO *resp.UserInfoRespVO) {

	var sysUser entity.SysUser
	result := config.DB.Where("user_id = ?", userId).Take(&sysUser)
	if result != nil && result.Error != nil {
		panic(exception.NewBizException(common.BIZ_ERROR_CODE, "用户不存在"))
	}

	roles := tis.RoleService.GetRolePermission(sysUser)
	perms := tis.RoleService.GetMenuPermission(sysUser)

	return &resp.UserInfoRespVO{
		User:               sysUser,
		Roles:              roles,
		Permissions:        perms,
		IsDefaultModifyPwd: false,
		IsPasswordExpired:  false,
	}

}

func (u *UserService) GetDeptTree(sysDept *entity.SysDept) []*menu.MenuTreeSelect {

	deptList := dept.SelectDeptList(sysDept)
	return u.BuildDeptTree(deptList)
}

func (u *UserService) BuildDeptTree(list []*entity.SysDept) []*menu.MenuTreeSelect {
	trees := u.buildTree(list) // 先构建树
	result := make([]*menu.MenuTreeSelect, 0, len(trees))
	for _, root := range trees {
		result = append(result, u.toTreeSelect(root))
	}
	return result
}

func (*UserService) buildTree(list []*entity.SysDept) []*entity.SysDept {
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

func (u *UserService) toTreeSelect(sysDept *entity.SysDept) *menu.MenuTreeSelect {
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

func (u *UserService) GetUserList(vo *userReqVO.SysUserReqVO) resp.PageResp[entity.SysUser] {

	var r resp.PageResp[entity.SysUser]
	userList, total, _ := userRepository.QueryUserList(*vo)
	r.Rows = userList
	r.Total = total
	r.Code = 200
	return r
}

func (tis *UserService) ChangeUserStatus(vo *userReqVO.ChangeUserStatusReqVo) uint64 {

	config.DB.Model(&entity.SysUser{}).Where("user_id = ?", vo.UserId).UpdateColumn("status", vo.Status)
	return vo.UserId
}

func (tis *UserService) QueryUser(id int) *userRespVO.QueryUserInfoRespVO {

	var sysUser = &entity.SysUser{}
	if id != 0 {
		config.DB.Where("user_id = ?", id).First(&sysUser)
	}

	var userPostList []*entity.SysUserPost
	if id != 0 {
		config.DB.Where("user_id = ?", id).Find(&userPostList)
	}

	var userRoleList []*entity.SysUserRole
	if id != 0 {
		config.DB.Where("user_id = ?", id).Find(&userRoleList)
	}

	postIds := lo.Map(userPostList, func(item *entity.SysUserPost, _ int) int64 {
		return item.PostId
	})

	roleIds := lo.Map(userRoleList, func(item *entity.SysUserRole, _ int) int64 {
		return item.RoleId
	})

	var roles []*entity.SysRole
	config.DB.Where("del_flag = '0'").Find(&roles)

	var posts []*entity.SysPost
	config.DB.Where("status = '0'").Find(&posts)

	var respVO = &userRespVO.QueryUserInfoRespVO{
		SysUser: *sysUser,
		PostIds: postIds,
		RoleIds: roleIds,
		Roles:   roles,
		Posts:   posts,
	}

	return respVO
}

func (tis *UserService) UpdateUser(ctx context.Context, reqVO *userReqVO.UserEditReqVO) uint64 {

	userId := reqVO.UserId

	currentDB := config.DB
	err := db.Transaction(ctx, currentDB, func(txCtx context.Context) error {

		checkBug := &entity.SysUser{}
		fmt.Println(checkBug)
		if delUserErr := tis.RoleService.DelUserRole(txCtx, userId); delUserErr != nil {
			return delUserErr
		}
		if addUserRoleErr := tis.RoleService.AddUserRole(txCtx, userId, reqVO.RoleIds); addUserRoleErr != nil {
			return addUserRoleErr
		}
		if delUserPost := tis.SysUserPostRepository.DelUserPost(txCtx, userId); delUserPost != nil {
			return delUserPost
		}
		if addUserPostErr := tis.SysUserPostRepository.AddUserPost(txCtx, userId, reqVO.PostIds); addUserPostErr != nil {
			return addUserPostErr
		}
		if updateUserErr := db.GetDB(txCtx, currentDB).Where("user_id = ?", userId).Updates(&reqVO.SysUser).Error; updateUserErr != nil {
			return updateUserErr
		}
		return nil
	})
	if err != nil {
		log.Println(err.Error())
		panic(exception.NewBizException(common.BIZ_ERROR_CODE, "用户更新失败"))
	}

	return userId
}
