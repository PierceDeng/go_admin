package role

import (
	"go_admin/config"
	"go_admin/middleware/common"
	"go_admin/middleware/exception"
	"go_admin/model/entity"
	roleRepository "go_admin/repository/role"
	"go_admin/service/menu"
	"go_admin/utils"
	"strings"

	"gorm.io/gorm"
)

const SUPER_ADMIN = "admin"
const ALL_PERMISSION = "*:*:*"

type RoleService struct {
	*menu.MenuService
}

func NewRoleService() *RoleService {
	return &RoleService{
		MenuService: menu.NewMenuService(),
	}
}

func (r RoleService) GetRolePermission(user entity.SysUser) []string {

	var roles []string
	if user.UserId == 1 {
		roles = append(roles, SUPER_ADMIN)
	} else {
		roles = append(roles, selectRolePermissionByUserId(user.UserId)...)
	}
	roles = utils.UniqueStrings(roles)
	return roles
}

func (r RoleService) GetMenuPermission(user entity.SysUser) []string {

	var permissions []string
	if user.UserId == 1 {
		permissions = append(permissions, ALL_PERMISSION)

	} else {

		var roles []entity.SysRole
		result := config.DB.Where("user_id = ?", user.UserId).Find(&roles)
		if result.Error != nil {
			panic(exception.NewBizException(common.DB_ERROR_CODE, result.Error.Error()))
		}
		if len(roles) > 0 {
			for _, role := range roles {
				if role.Status == "0" && role.IsAdmin() {
					var rolePerms = r.MenuService.SelectMenuPermsByRoleId(role.RoleId)
					role.Permissions = append(role.Permissions, rolePerms...)
					permissions = append(permissions, role.Permissions...)
				}
			}
		} else {
			rolePerms := r.MenuService.SelectMenuPermsByUserId(user.UserId)
			rolePerms = utils.UniqueStrings(rolePerms)
			permissions = append(permissions, rolePerms...)
		}
	}

	return permissions

}

func (r RoleService) DelUserRole(db *gorm.DB, id int) {

	db.Model(&entity.SysUserRole{}).Where("user_id = ?", id).Delete(&entity.SysUserRole{})

}

func selectRolePermissionByUserId(id uint64) []string {

	var roles = roleRepository.SelectRolePermissionByUserId(id)
	var permsSet []string
	for _, item := range roles {
		permsSet = append(permsSet, strings.Split(item.RoleKey, ",")...)
	}
	permsSet = utils.UniqueStrings(permsSet)
	return permsSet
}
