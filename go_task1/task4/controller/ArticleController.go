package controller

import (
	"net/http"
	"strconv"
	"task4/common"
	"task4/config"
	"task4/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {

	var (
		db       = config.GetDB()
		articles []models.Posts
		total    int64
	)

	// 分页参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	offset := (page - 1) * size

	// 查询文章列表和总数

	result := db.Offset(offset).Limit(size).Find(&articles)
	db.Model(&models.Posts{}).Count(&total)

	if result.Error != nil {
		common.UnknownErrorRes(c, "获取文章列表失败", result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取文章列表成功",
		"data":    articles,
		"total":   total,
		"page":    page,
		"size":    size,
	})
}

func GetArticle(c *gin.Context) {
	idt := c.Param("id")
	//若是获取json中的id
	//var id int
	// if err := c.ShouldBindJSON(&id); err == nil {
	// 	fmt.Println(id)
	// }

	var (
		db      = config.GetDB()
		article models.Posts
		//comments []models.Comments
	)
	id, _ := strconv.Atoi(idt)
	if id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
		})
		return
	}
	res := db.Where("id = ?", id).First(&article)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取文章失败",
			"error":   res.Error.Error(),
		})
		return
	}

	common.Success(c, article)

}

// DeleteArticle 删除文章接口（需要认证，只有作者能删除）
func DeleteArticle(c *gin.Context) {
	// 从URL参数中获取文章ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		common.ParamsErrorRes(c, "无效的文章ID")
		return
	}

	// 获取数据库连接
	db := config.GetDB()
	if db == nil {
		common.UnknownErrorRes(c, "数据库连接失败")
		return
	}

	// 获取当前登录用户ID
	userID, exists := c.Get("userID")
	if !exists {
		common.UnknownErrorRes(c, "无法获取用户信息")
		return
	}
	currentUserID := userID.(uint)

	// 查询文章是否存在，并验证是否是作者
	var article models.Posts
	result := db.Where("id = ?", id).First(&article)
	if result.Error != nil {
		common.UnknownErrorRes(c, "文章不存在或已被删除")
		return
	}

	// 检查是否是文章的作者
	if article.UserID != currentUserID {
		common.UnknownErrorRes(c, "没有权限删除此文章")
		return
	}

	// 删除文章
	if err := db.Delete(&article).Error; err != nil {
		common.UnknownErrorRes(c, "删除文章失败", err)
		return
	}

	// 返回成功响应
	common.Success(c, gin.H{"message": "文章删除成功"})
}

// CreateArticle 创建文章接口（需要认证）
func CreateArticle(c *gin.Context) {
	// 获取数据库连接
	db := config.GetDB()
	if db == nil {
		common.UnknownErrorRes(c, "数据库连接失败")
		return
	}

	// 获取当前登录用户ID
	userID, exists := c.Get("userID")
	if !exists {
		common.UnknownErrorRes(c, "无法获取用户信息")
		return
	}
	currentUserID := userID.(uint)

	// 绑定请求参数
	var articleData struct {
		Title   string `json:"title" binding:"required,min=1,max=200"`
		Content string `json:"content" binding:"required,min=1"`
	}

	if err := c.ShouldBindJSON(&articleData); err != nil {
		common.ParamsErrorRes(c, "参数错误: "+err.Error())
		return
	}

	// 创建新文章
	newArticle := models.Posts{
		Title:     articleData.Title,
		Content:   articleData.Content,
		UserID:    currentUserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := db.Create(&newArticle).Error; err != nil {
		common.UnknownErrorRes(c, "创建文章失败", err)
		return
	}

	// 返回成功响应
	common.Success(c, newArticle)
}

// UpdateArticle 更新文章接口（需要认证，只有作者能更新）
func UpdateArticle(c *gin.Context) {
	// 从URL参数中获取文章ID
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		common.ParamsErrorRes(c, "无效的文章ID")
		return
	}

	// 获取数据库连接
	db := config.GetDB()
	if db == nil {
		common.UnknownErrorRes(c, "数据库连接失败")
		return
	}

	// 获取当前登录用户ID
	userID, exists := c.Get("userID")
	if !exists {
		common.UnknownErrorRes(c, "无法获取用户信息")
		return
	}
	currentUserID := userID.(uint)

	// 查询文章是否存在，并验证是否是作者
	var article models.Posts
	result := db.Where("id = ?", id).First(&article)
	if result.Error != nil {
		common.UnknownErrorRes(c, "文章不存在或已被删除")
		return
	}

	// 检查是否是文章的作者
	if article.UserID != currentUserID {
		common.UnknownErrorRes(c, "没有权限更新此文章")
		return
	}

	// 绑定请求体中的更新数据
	var updateData struct {
		Title   string `json:"title" binding:"omitempty,min=1,max=200"`
		Content string `json:"content" binding:"omitempty,min=1"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		common.ParamsErrorRes(c, "参数错误: "+err.Error())
		return
	}

	// 验证是否提供了更新数据
	if updateData.Title == "" && updateData.Content == "" {
		common.ParamsErrorRes(c, "至少需要提供标题或内容进行更新")
		return
	}

	// 更新文章
	updateFields := map[string]interface{}{
		"updated_at": time.Now(),
	}

	if updateData.Title != "" {
		updateFields["title"] = updateData.Title
	}

	if updateData.Content != "" {
		updateFields["content"] = updateData.Content
	}

	if err := db.Model(&article).Updates(updateFields).Error; err != nil {
		common.UnknownErrorRes(c, "更新文章失败", err)
		return
	}

	// 查询更新后的文章信息
	updatedArticle := models.Posts{}
	db.Where("id = ?", id).First(&updatedArticle)

	// 返回成功响应
	common.Success(c, updatedArticle)
}
