package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"task4/config"
	"task4/models"

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
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "获取文章列表失败",
			"error":   result.Error.Error(),
		})
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
	fmt.Println(idt)

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

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "获取文章成功",
		"data":    article,
	})

}
