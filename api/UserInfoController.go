package api

import (
	"fmt"
	"go_admin/middleware/common"
	resp "go_admin/model"
	"go_admin/model/entity"
	"go_admin/model/reqVO/user"
	userSerivce "go_admin/service/user"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type UserController struct {
	UserService *userSerivce.UserService
}

var UserControl = NewUserController()

func NewUserController() *UserController {
	return &UserController{
		UserService: userSerivce.NewUserService(),
	}
}

func (u *UserController) GetUserInfo(c *gin.Context) {
	userId, _ := c.Get("userId")
	resp.Ok(c, u.UserService.GetUserInfo(userId.(uint64)))
}

func (u *UserController) GetDeptTree(c *gin.Context) {
	dept, err := common.BindQuery[entity.SysDept](c)
	if err != nil {
		return
	}
	resp.Ok(c, u.UserService.GetDeptTree(dept))
}

func (u *UserController) GetUserList(c *gin.Context) {
	userReqVO, err := common.BindQuery[user.SysUserReqVO](c)
	if err != nil {
		return
	}
	resp.OkWithWrapper(c, u.UserService.GetUserList(userReqVO))
}

func (u *UserController) ChangeUserStatus(c *gin.Context) {
	reqVO, err := common.BindJSON[user.ChangeUserStatusReqVo](c)
	if err != nil {
		return
	}
	resp.Ok(c, u.UserService.ChangeUserStatus(reqVO))
}

func (u *UserController) QueryUser(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Param("userId"))
	resp.Ok(c, u.UserService.QueryUser(userId))
}

func (u *UserController) UpdateUser(c *gin.Context) {

	reqVO, _ := common.BindJSON[user.UserEditReqVO](c)
	resp.Ok(c, u.UserService.UpdateUser(c, reqVO))
}

func (u *UserController) AddUser(c *gin.Context) {
	reqVO, _ := common.BindJSON[user.UserEditReqVO](c)
	resp.Ok(c, u.UserService.AddUser(c, reqVO))
}

func (u *UserController) ImportUserTable(c *gin.Context) {
	// 1. 获取上传文件
	file, _ := c.FormFile("file")

	src, _ := file.Open()
	defer src.Close()

	f, _ := excelize.OpenReader(src)
	defer f.Close()

	sheetName := f.GetSheetName(0)
	rows, _ := f.GetRows(sheetName)

	var users []*entity.SysUser
	for _, row := range rows[1:] {

		name := strings.TrimSpace(row[1])
		phonenumber := strings.TrimSpace(row[2])
		email := strings.TrimSpace(row[3])

		users = append(users, &entity.SysUser{
			Username:    name,
			Phonenumber: phonenumber,
			Email:       email,
		})
	}

	resp.Ok(c, u.UserService.ImportUserTable(c, users))

}

func (u *UserController) QueryUserListForExport(c *gin.Context) {

	userReqVO, _ := common.BindQuery[user.SysUserReqVO](c)
	userList := u.UserService.QueryUserListForExport(userReqVO)

	// 创建 Excel 文件
	f := excelize.NewFile()
	// 设置工作表名称
	sheetName := "用户列表"
	f.SetSheetName("Sheet1", sheetName)

	// 写入表头
	headers := []string{"ID", "姓名", "手机号", "邮箱", "创建时间"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheetName, cell, h)
	}

	// 写入数据
	for idx, user := range userList {
		row := idx + 2 // 数据从第二行开始
		f.SetCellInt(sheetName, fmt.Sprintf("A%d", row), int64(int(user.UserId)))
		f.SetCellStr(sheetName, fmt.Sprintf("B%d", row), user.Username)
		f.SetCellStr(sheetName, fmt.Sprintf("C%d", row), user.Phonenumber)
		f.SetCellStr(sheetName, fmt.Sprintf("D%d", row), user.Email)
		if user.CreateTime != nil {
			f.SetCellStr(sheetName, fmt.Sprintf("E%d", row), user.CreateTime.Format("2006-01-02 15:04:05"))
		}

	}

	// 设置响应头，告诉浏览器下载文件
	fileName := fmt.Sprintf("users_%s.xlsx", time.Now().Format("20060102"))
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")

	// 将 Excel 文件写入响应流
	if err := f.Write(c.Writer); err != nil {
		// 写入失败时，响应头已经设置，无法返回错误JSON，只能记录日志
		fmt.Printf("导出 Excel 失败: %v\n", err)
	}

}

func (u *UserController) DeleteUserBatch(c *gin.Context) {
	idsStr := c.Param("userId") // "1,2,3"
	ids := strings.Split(idsStr, ",")
	var intIds []uint64
	for _, s := range ids {
		id, _ := strconv.ParseUint(s, 10, 64)
		intIds = append(intIds, id)
	}
	resp.Ok(c, u.UserService.DeleteUserBatch(c, intIds))
}

func (u *UserController) ResetUserPwd(c *gin.Context) {
	reqVO, _ := common.BindJSON[user.ResetUserPwdReqVO](c)
	resp.Ok(c, u.UserService.ResetUserPwd(c, reqVO))
}
