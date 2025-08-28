## 项目简介
本项目提供了一个完整的博客系统后端解决方案，包括：

## 用户注册和登录功能
JWT 认证和授权
文章的 CRUD（创建、读取、更新、删除）操作
评论功能
统一的错误处理
完善的日志记录系统
技术栈
编程语言: Go 1.23.0
Web 框架: Gin 1.10.1
ORM 库: GORM 1.30.1
数据库: MySQL
认证: JWT (JSON Web Token)
容器化: Docker & Docker Compose

## 项目结构
go_task1/task4/
├── common/      # 通用组件（认证、响应处理等）
├── config/      # 配置文件与管理
├── controller/  # 请求处理控制器
├── go.mod       # Go模块依赖声明
├── go.sum       # 依赖版本锁定
├── log/         # 日志文件存储
├── main.go      # 应用入口文件
├── models/      # 数据模型定义
├── router/      # 路由配置
└── service/     # 业务逻辑层


## 环境要求
Go 1.23.0 或更高版本
MySQL 8.0 或更高版本
Docker 和 Docker Compose（推荐部署方式）

## 安装和运行
方式一：本地直接运行
1.克隆项目
    git clone https://github.com/yourusername/blog-backend.git
    cd blog-backend
2.安装依赖
    go mod tidy
    go mod download
3.配置数据库
    创建 MySQL 数据库，并修改 config/config.yaml 文件中的数据库配置：

    mysql:
      host: localhost
      port: 3306
      db_name: blog_db
      user: root
      password: your_password
      charset: utf8mb4

4.运行
    go run main.go
    应用将在 http://localhost:8090 启动，日志将输出到 log/ 目录下。