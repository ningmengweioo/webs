package controller

import (
	"task4/common"
	"task4/config"
	"task4/models"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
		common.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 2. 获取数据库连接
	db := config.GetDB()
	if db == nil {
		common.InternalError(c, "数据库连接失败", nil)
		return
	}

	// 3. 检查用户名或邮箱是否已存在
	var existingUser models.Users
	result := db.Where("name = ? OR email = ?", req.Name, req.Email).First(&existingUser)

	if result.Error == nil {
		// 用户已存在
		if existingUser.Name == req.Name {
			common.Error(c, 400, "用户名已存在")
		} else {
			common.Error(c, 400, "邮箱已被注册")
		}
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		// 数据库查询错误
		common.InternalError(c, "查询用户信息失败", result.Error)
		return
	}

	// 4. 加密密码
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		common.InternalError(c, "密码加密失败", err)
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
		common.InternalError(c, "注册失败", err)
		return
	}

	// 6. 返回成功响应（不包含密码等敏感信息）
	responseUser := gin.H{
		"id":         newUser.ID,
		"username":   newUser.Name,
		"email":      newUser.Email,
		"created_at": newUser.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	common.SuccessWithMsg(c, "注册成功", responseUser)
}

// Login 用户登录接口
func Login(c *gin.Context) {
	// 1. 绑定请求参数
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		common.BadRequest(c, "参数错误: "+err.Error())
		return
	}

	// 2. 获取数据库连接
	db := config.GetDB()
	if db == nil {
		common.InternalError(c, "数据库连接失败", nil)
		return
	}

	// 3. 查询用户信息
	var user models.Users
	result := db.Where("username = ?", req.Username).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			common.Error(c, 400, "用户名或密码错误")
		} else {
			common.InternalError(c, "查询用户信息失败", result.Error)
		}
		return
	}

	// 4. 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		common.Error(c, 400, "用户名或密码错误")
		return
	}

	// 5. 生成JWT令牌（这里需要根据项目的JWT配置来实现）
	// 假设使用了标准的JWT库，以下是示例代码
	// token := generateToken(user.ID, user.Username)
	// 由于没有具体的JWT配置，这里简化返回

	// 6. 返回登录成功响应
	responseData := gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"token": "placeholder_jwt_token", // 实际项目中应替换为真实生成的token
		//expiresAt: time.Now().Add(24 * time.Hour).Unix(), // token过期时间
	}

	common.SuccessWithMsg(c, "登录成功", responseData)
}
