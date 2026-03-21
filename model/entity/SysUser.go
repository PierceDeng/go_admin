package entity

import (
	"time"
)

type SysUser struct {
	UserId        uint64     `gorm:"column:user_id;primaryKey;" json:"userId"`
	DeptId        *int64     `gorm:"column:dept_id;" json:"deptId"`
	Username      string     `gorm:"column:username;not null;" json:"username"`
	Nickname      string     `gorm:"column:nickname;default:'00'" json:"nickname"`
	UserType      string     `gorm:"column:user_type" json:"userType"`
	Email         string     `gorm:"column:email" json:"email"`
	Phonenumber   string     `gorm:"column:phonenumber" json:"phonenumber"`
	Sex           string     `gorm:"column:sex;default:'0'" json:"sex"`           // 用户性别（0男 1女 2未知）
	Avatar        string     `gorm:"column:avatar;default:''" json:"avatar"`      // 头像地址
	Password      string     `gorm:"column:password;default:''" json:"password"`  // 密码
	Status        string     `gorm:"column:status;default:'0'" json:"status"`     // 账号状态（0正常 1停用）
	DelFlag       string     `gorm:"column:del_flag;default:'0'" json:"delFlag"`  // 删除标志（0代表存在 2代表删除）
	LoginIp       string     `gorm:"column:login_ip;default:''" json:"loginIp"`   // 最后登录IP
	LoginDate     *time.Time `gorm:"column:login_date" json:"loginDate"`          // 最后登录时间
	PwdUpdateDate *time.Time `gorm:"column:pwd_update_date" json:"pwdUpdateDate"` // 密码最后更新时间
	CreateBy      string     `gorm:"column:create_by;default:''" json:"createBy"` // 创建者
	CreateTime    *time.Time `gorm:"column:create_time" json:"createTime"`        // 创建时间
	UpdateBy      string     `gorm:"column:update_by;default:''" json:"updateBy"` // 更新者
	UpdateTime    *time.Time `gorm:"column:update_time" json:"updateTime"`        // 更新时间
	Remark        *string    `gorm:"column:remark" json:"remark"`
}

func (SysUser) TableName() string {
	return "sys_user"
}
