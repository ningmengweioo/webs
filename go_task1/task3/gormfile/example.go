package gormfile

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

// 题目1：模型定义
type User struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	PostCount int    `gorm:"default:0"` // 文章数量统计字段
	Posts     []Post
}

type Post struct {
	ID            uint      `gorm:"primarykey"`
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text;not null"`
	UserID        *uint     `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL"` // 外键，关联到User
	User          User      // 反向引用
	Comments      []Comment // 一对多关系：一篇文章可以有多个评论
	CommentStatus string    `gorm:"size:20;default:'有评论'"` // 评论状态
}

type Comment struct {
	ID      uint   `gorm:"primarykey"`
	Content string `gorm:"type:text;not null"`
	PostID  *uint  `gorm:"foreignKey:PostID;constraint:OnDelete:SET NULL"` // 外
	//Post    Post   // 反向引用
}

// 题目3：钩子函数
// BeforeCreate Post的创建前钩子函数，自动更新用户的文章数量统计字段
func (p *Post) BeforeCreate(tx *gorm.DB) error {
	// 增加用户的文章数量
	fmt.Println("创建引入钩子")
	if p.UserID != nil {
		tx.Model(&User{ID: *p.UserID}).Update("post_count", gorm.Expr("post_count + ?", 1))
	}
	return nil
}

// AfterDelete Comment的删除后钩子函数，检查文章的评论数量
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	fmt.Println("删除引入钩子")
	var post Post
	tx.First(&post, c.PostID)

	// 查询该文章的评论数量
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", post.ID).Count(&commentCount)

	// 如果评论数量为0，更新文章的评论状态为"无评论"
	if commentCount == 0 {
		tx.Model(&post).Update("comment_status", "无评论")
	}
	return nil
}

// 题目2：关联查询函数
func GetUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Debug().Preload("Comments").Where("user_id = ?", userID).Find(&posts).Error
	return posts, err
}

// 查询评论数量最多的文章信息
func GetPostWithMostComments(db *gorm.DB) (Post, error) {
	var post Post
	var maxCommentsPostID uint
	//var maxCommentsCount int64
	var result struct {
		PostID       uint
		CommentCount int64
	}

	// 获取评论数量最多的文章ID
	err := db.Debug().Table("comments").
		Select("post_id, count(*) as comment_count").
		Group("post_id").
		Order("comment_count desc").
		Limit(1).
		Scan(&result).Error

	// 将结果赋值给目标变量
	maxCommentsPostID = result.PostID
	//maxCommentsCount = result.CommentCount

	if err != nil {
		return post, err
	}

	// 查询该文章的详细信息及其评论
	err = db.Debug().Preload("Comments").Preload("User").First(&post, maxCommentsPostID).Error
	return post, err
}

// 初始化数据库表
func InitTables(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

// Run 运行示例
func Run(db *gorm.DB) {
	InitTables(db)
	//fmt.Println("博客系统表结构创建成功")
	//1.获取用户发表的文章及评论信息
	// posts, err := GetUserPostsWithComments(db, 1)
	// if err != nil {
	// 	panic(err)
	// }
	// // jsonData, err := json.Marshal(posts)
	// jsonData, err := json.MarshalIndent(posts, "", "  ")
	// if err != nil {
	// 	fmt.Printf("JSON编码失败: %v\n", err)
	// 	return
	// }
	// fmt.Println(string(jsonData))

	//2.评论数量最多的文章信息
	info, err := GetPostWithMostComments(db)
	if err != nil {
		panic(err)
	}
	info_json, _ := json.Marshal(info)
	fmt.Println(string(info_json))
	var user User
	db.First(&user)

	//对应的钩子
	//创建文章
	// post := Post{
	// 	Title:   "新文章删除",
	// 	Content: "新内容删除",
	// 	UserID:  &user.ID,
	// }
	// db.Create(&post)
	var post Post
	db.Where("id=?", 5).Model(&Post{}).Scan(&post)

	//创建评论
	comment := Comment{
		Content: "新评论12",
		PostID:  &post.ID,
	}
	db.Create(&comment)
	//删除
	// db.Debug().Delete(&comment, comment.ID)

}
