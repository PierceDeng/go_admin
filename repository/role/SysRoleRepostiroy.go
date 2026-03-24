package role

import (
	"go_admin/config"
	"go_admin/model/entity"
)

func SelectRolePermissionByUserId(userId uint64) []*entity.SysRole {

	var roles []*entity.SysRole
	result := config.DB.Exec("select distinct r.role_id, r.role_name, r.role_key, r.role_sort, "+
		"r.data_scope, r.menu_check_strictly, r.dept_check_strictly,"+
		"\n            r.status, r.del_flag, r.create_time, r.remark \n from sys_role r\n\t        "+
		"left join sys_user_role ur on ur.role_id = r.role_id\n\t  left join sys_user u on u.user_id = ur.user_id"+
		"\n\t left join sys_dept d on u.dept_id = d.dept_id WHERE r.del_flag = '0' and ur.user_id = ?", userId).Find(&roles)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	return roles
}
