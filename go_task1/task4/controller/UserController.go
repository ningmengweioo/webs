package controller

import (
	"task4/common"
	"task4/config"
	"task4/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register 用户注册接口
func Register(c *gin.Context) {
	// 1. 绑定请求参数
	var req struct {
		Name     string `json:"name" binding:"required,min=3,max=20"`
		Password string `json:"password" binding:"required,min=6"`
		Email    string `json:"email" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		common.ParamsErrorRes(c, "参数错误: "+err.Error())
		return
	}

	// 2. 获取数据库连接
	db := config.GetDB()
	if db == nil {
		common.UnknownErrorRes(c, "数据库连接失败", nil)
		return
	}
	// 3. 检查用户名或邮箱是否已存在
	var count int64
	err := db.Table("users").Where("name = ? OR email = ?", req.Name, req.Email).Count(&count).Error

	if err != nil {
		common.UnknownErrorRes(c, "检查用户名或邮箱是否已存在失败", err)
		return
	}
	if count > 0 {
		common.UnknownErrorRes(c, "检查用户名或邮箱已存在")
	}

	// 4. 加密密码
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		common.UnknownErrorRes(c, "密码加密失败", err)
		return
	}

	// 5. 创建新用户
	newUser := models.Users{
		Name:      req.Name,
		Password:  string(passwordHash),
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&newUser).Error; err != nil {
		common.UnknownErrorRes(c, "注册失败", err)
		return
	}

	// 6. 返回成功响应（不包含密码等敏感信息）
	responseUser := gin.H{
		"id":         newUser.ID,
		"name":       newUser.Name,
		"email":      newUser.Email,
		"created_at": newUser.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	common.Success(c, responseUser)
}

// Login 用户登录接口
func Login(c *gin.Context) {

	// 1. 绑定请求参数
	var req struct {
		Name     string `json:"name" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		common.ParamsErrorRes(c, "参数错误 ", err.Error())
		return
	}

	// 2. 获取数据库连接
	db := config.GetDB()
	if db == nil {
		common.UnknownErrorRes(c, "数据库连接失败", nil)
		return
	}

	// 3. 查询用户信息
	var user models.Users
	error := db.Where("name = ?", req.Name).First(&user).Error

	if error != nil {
		common.UnknownErrorRes(c, "查询用户信息失败", error)
		return
	}

	// 4. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		common.UnknownErrorRes(c, "用户名或密码错误", error)
		return
	}

	// 5. 生成JWT令牌（这里需要根据项目的JWT配置来实现）
	token, err := common.GenerateToken(user.ID, user.Name)
	if err != nil {
		common.UnknownErrorRes(c, "生成token失败", err)
		return
	}

	// 6. 返回登录成功响应
	res := make(map[string]interface{})
	res["id"] = user.ID
	res["name"] = user.Name
	res["email"] = user.Email
	res["token"] = token
	res["expiresAt"] = time.Now().Add(24 * time.Hour).Unix()

	common.Success(c, res)
}
