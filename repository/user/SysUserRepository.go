package user

import (
	"go_admin/config"
	"go_admin/model/entity"
	"go_admin/model/reqVO/user"
)

func QueryUserList(query *user.SysUserReqVO) ([]*entity.SysUser, int64, error) {
	var users []*entity.SysUser
	var total int64
	db := config.DB
	// 基础查询（过滤删除标志）
	tx := db.Table("sys_user u").
		Select("u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag, u.login_ip, u.login_date, u.create_by, u.create_time, u.remark, d.dept_name, d.leader").
		Joins("left join sys_dept d on u.dept_id = d.dept_id").
		Where("u.del_flag = ?", "0")

	// 动态添加条件
	if query.UserId != 0 {
		tx = tx.Where("u.user_id = ?", query.UserId)
	}
	if query.Username != "" {
		tx = tx.Where("u.user_name like ?", "%"+query.Username+"%")
	}
	if query.Status != "" {
		tx = tx.Where("u.status = ?", query.Status)
	}
	if query.Phonenumber != "" {
		tx = tx.Where("u.phonenumber like ?", "%"+query.Phonenumber+"%")
	}
	// 日期范围处理（推荐使用 time.Time 类型，可直接比较）
	if query.Param.BeginTime != "" && query.Param.EndTime != "" {
		// 假设日期格式为 "2006-01-02"，将字符串转为 time.Time
		tx = tx.Where("u.create_time BETWEEN ? AND ?", query.Param.BeginTime, query.Param.EndTime)
	} else if query.Param.BeginTime != "" {
		tx = tx.Where("u.create_time >= ?", query.Param.BeginTime)
	} else if query.Param.EndTime != "" {
		tx = tx.Where("u.create_time <= ?", query.Param.EndTime)
	}

	// 部门条件（包含子部门，使用 find_in_set 或递归查询）
	if query.DeptId != 0 && query.DeptId != 100 {
		// 方法1：使用原生 SQL 的 FIND_IN_SET 函数（需要 ancestors 字段格式为 ",1,2,3," 或类似）
		tx = tx.Where("(u.dept_id = ? OR FIND_IN_SET(?, (SELECT GROUP_CONCAT(dept_id) FROM sys_dept WHERE find_in_set(?, ancestors))))",
			query.DeptId, query.DeptId, query.DeptId) // 复杂逻辑，建议封装
	}

	if query.PageSize != 0 {
		// 先统计总数（分页前）
		if err := tx.Count(&total).Error; err != nil {
			return users, 0, err
		}
		tx.Offset((query.PageNum - 1) * query.PageSize).Limit(query.PageSize)
	}

	if err := tx.Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
