## 一、要求概述
使用 Go 语言结合 Gin 框架和 GORM 库开发一个个人博客系统的后端，实现博客文章的基本管理功能，包括文章的创建、读取、更新和删除（CRUD）操作，同时支持用户认证和简单的评论功能。

## 二、项目初始化
1. 1.
   创建一个新的 Go 项目，使用 go mod init 初始化项目依赖管理
2. 2.
   安装必要的库，如 Gin 框架、GORM 以及数据库驱动（如 MySQL 或 SQLite）
## 三、数据库设计与模型定义
### 3.1 数据库表结构设计
至少包含以下几个表：

- users 表 ：存储用户信息，包括 id、username、password、email 等字段
- posts 表 ：存储博客文章信息，包括 id、title、content、user_id（关联 users 表的 id）、created_at、updated_at 等字段
- comments 表 ：存储文章评论信息，包括 id、content、user_id（关联 users 表的 id）、post_id（关联 posts 表的 id）、created_at 等字段
### 3.2 模型定义
使用 GORM 定义对应的 Go 模型结构体

## 四、用户认证与授权
1. 1.
   实现用户注册和登录功能
   - 用户注册时需要对密码进行加密存储
   - 登录时验证用户输入的用户名和密码
2. 2.
   使用 JWT（JSON Web Token）实现用户认证和授权
   - 用户登录成功后返回一个 JWT
   - 后续的需要认证的接口需要验证该 JWT 的有效性
## 五、文章管理功能
1. 1.
   实现文章的创建功能
   - 只有已认证的用户才能创建文章
   - 创建文章时需要提供文章的标题和内容
2. 2.
   实现文章的读取功能
   - 支持获取所有文章列表
   - 支持获取单个文章的详细信息
3. 3.
   实现文章的更新功能
   - 只有文章的作者才能更新自己的文章
4. 4.
   实现文章的删除功能
   - 只有文章的作者才能删除自己的文章
## 六、评论功能
1. 1.
   实现评论的创建功能
   - 已认证的用户可以对文章发表评论
2. 2.
   实现评论的读取功能
   - 支持获取某篇文章的所有评论列表
## 七、错误处理与日志记录
1. 1.
   对可能出现的错误进行统一处理，如：
   - 数据库连接错误
   - 用户认证失败
   - 文章或评论不存在等
   - 返回合适的 HTTP 状态码和错误信息
2. 2.
   使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护

## 测试数据
-- 添加测试文章，关联已创建的用户
INSERT INTO posts (title, content, user_id, created_at, updated_at) VALUES
('Go语言入门指南', '这是一篇关于Go语言基础知识的入门文章，包括变量、数据类型、控制流等内容。Go语言是由Google开发的开源编程语言，它具有静态类型、并发特性和垃圾回收机制。', 1, NOW(), NOW()),
('使用Gin框架构建Web应用', '本文介绍如何使用Gin框架快速构建高性能的Web应用。Gin是一个用Go语言编写的Web框架，它具有轻量级、高性能的特点，适合构建各种Web服务。', 1, NOW(), NOW()),
('GORM数据库操作详解', 'GORM是Go语言中最流行的ORM库之一，本文详细介绍了如何使用GORM进行数据库的增删改查操作，以及如何处理关联关系等高级特性。', 2, NOW(), NOW()),
('Go并发编程实践', 'Go语言的一大特色是其内置的并发支持，本文通过实例讲解了goroutine和channel的使用方法，以及如何避免常见的并发问题。', 3, NOW(), NOW()),
('Go项目结构最佳实践', '一个良好的项目结构对于代码的可维护性至关重要。本文分享了Go项目的最佳目录结构设计，以及如何组织代码以提高可读性和可扩展性。', 2, NOW(), NOW());

-- 添加测试用户
INSERT INTO users (username, password, email, created_at, updated_at) VALUES
('admin', '$2a$10$mN9uSfK1P6j7Jd7V8g9hKOpQwErTyUiPoAsDfGhJkLzXcVbNm', 'admin@example.com', NOW(), NOW()),
('user1', '$2a$10$aBcDeFgHiJkLmNoPqRsTuVwXyZ1234567890', 'user1@example.com', NOW(), NOW()),
('user2', '$2a$10$zYxWvUtSrQpOnMlKjIhGfEdCbA987654321', 'user2@example.com', NOW(), NOW());