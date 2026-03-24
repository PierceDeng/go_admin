package menu

import (
	"go_admin/config"
	"go_admin/middleware/common"
	"go_admin/middleware/exception"
	"go_admin/model/entity"
)

func SelectMenuTreeAll() (m []*entity.SysMenu) {

	result := config.DB.Exec("select distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.`query`, m.route_name, " +
		"m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time " +
		"from sys_menu m where m.menu_type in ('M', 'C') and m.status = 0 " +
		"order by m.parent_id, m.order_num").Find(&m)

	if result.Error != nil {
		panic(exception.NewBizException(common.BIZ_ERROR_CODE, "查询错误"))
	}
	return m

}

func SelectMenuTreeByUserId(userId uint64) (m []*entity.SysMenu) {
	result := config.DB.Exec(""+
		"select distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.`query`"+
		", m.route_name, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon"+
		", m.order_num, m.create_time\n\t\tfrom sys_menu m\n\t\t\t left join sys_role_menu rm on m.menu_id = rm.menu_id\n\t\t\t "+
		"left join sys_user_role ur on rm.role_id = ur.role_id\n\t\t\t left join sys_role ro on ur.role_id = ro.role_id\n\t\t\t "+
		"left join sys_user u on ur.user_id = u.user_id\n\t\twhere u.user_id = #{userId} and m.menu_type in ('M', 'C') and m.status = 0  "+
		"AND ro.status = 0\n\t\torder by m.parent_id, m.order_num"+
		"", userId).Find(&m)

	if result.Error != nil {
		panic(exception.NewBizException(common.BIZ_ERROR_CODE, "查询错误"))
	}
	return m
}

func SelectMenuPermsByRoleId(roleId int64) (m []string) {
	config.DB.Exec("select distinct m.perms\n\t\t"+
		"from sys_menu m\n\t\t\t left join sys_role_menu rm on m.menu_id = rm.menu_id\n\t\t"+
		"where m.status = '0' and rm.role_id = ?", roleId).Find(&m)
	return m
}

func SelectMenuPermsByUserId(id uint64) (m []string) {
	config.DB.Exec("select distinct m.perms\n\t\t"+
		"from sys_menu m\n\t\t\t left join sys_role_menu rm on m.menu_id = rm.menu_id\n\t\t\t"+
		" left join sys_user_role ur on rm.role_id = ur.role_id"+
		"\n\t\t\t left join sys_role r on r.role_id = ur.role_id\n\t\t"+
		"where m.status = '0' and r.status = '0' and ur.user_id = ?", id).Find(&m)
	return m
}
