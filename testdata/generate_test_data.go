package main

import (
	"blog/models"
	"blog/utils"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		// Set default environment variables if .env file not found
		os.Setenv("DB_TYPE", "sqlite3")
		os.Setenv("DB_NAME", "blog.db")
	}

	// Initialize database connection
	utils.InitDB()

	// Clear existing data (optional)
	// utils.DB.Delete(&models.Comment{})
	// utils.DB.Delete(&models.Post{})

	// Create multiple test posts
	posts := []models.Post{
		{
			Title:     "Go语言入门指南",
			Content:   "Go语言是一种由Google开发的开源编程语言，以其简洁、高效和并发特性而闻名。本文将介绍Go语言的基本语法和特性。\n\n1. 变量声明\n2. 函数定义\n3. 并发编程\n4. 错误处理\n\nGo语言的设计哲学是'少即是多'，通过减少语言特性来提高代码的可读性和可维护性。",
			Author:    "技术专家",
			CreatedAt: time.Now().Add(-7 * 24 * time.Hour), // 7天前
			UpdatedAt: time.Now().Add(-6 * 24 * time.Hour), // 6天前
		},
		{
			Title:     "Gin框架使用教程",
			Content:   "Gin是Go语言中最流行的Web框架之一，提供了高性能的HTTP路由和中间件支持。本文将详细介绍Gin框架的使用方法。\n\n- 路由定义\n- 中间件使用\n- 模板渲染\n- 表单处理\n- 错误处理\n\nGin框架的优势在于其高性能和简洁的API设计，适合构建各种规模的Web应用。",
			Author:    "Web开发者",
			CreatedAt: time.Now().Add(-3 * 24 * time.Hour), // 3天前
			UpdatedAt: time.Now().Add(-2 * 24 * time.Hour), // 2天前
		},
		{
			Title:     "SQLite数据库简介",
			Content:   "SQLite是一种轻量级的嵌入式关系型数据库，无需单独的服务器进程即可运行。本文将介绍SQLite的基本概念和使用方法。\n\nSQLite的主要特点：\n- 零配置，无需安装和管理\n- 支持标准SQL语法\n- 跨平台兼容\n- 占用资源少\n\nSQLite适合用于移动应用、桌面应用以及小型Web应用的数据存储需求。",
			Author:    "数据库工程师",
			CreatedAt: time.Now().Add(-1 * 24 * time.Hour), // 1天前
			UpdatedAt: time.Now().Add(-1 * 24 * time.Hour), // 1天前
		},
		{
			Title:     "博客应用开发总结",
			Content:   "经过一段时间的开发，我们完成了这个基于Go+Gin+SQLite的博客应用。本文将总结开发过程中的经验和教训。\n\n开发过程中遇到的挑战：\n1. 数据库连接配置\n2. 模板渲染错误处理\n3. 路由设计\n\n最终实现的功能：\n- 文章的增删改查\n- 评论功能\n- 响应式布局\n- 良好的用户体验\n\n通过这个项目，我们学习了如何使用现代Web开发技术构建一个完整的应用。",
			Author:    "项目负责人",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Insert posts into database
	for i, post := range posts {
		err = utils.DB.Create(&post).Error
		if err != nil {
			panic(err)
		}
		println("创建文章:", post.Title)

		// Add comments to some posts
		if i < 2 { // Add comments to first two posts
			comments := []models.Comment{
				{
					PostID:    post.ID,
					Name:      "张三",
					Email:     "zhangsan@example.com",
					Content:   "这篇文章写得很好，学到了很多知识！",
					CreatedAt: time.Now().Add(-6*24*time.Hour + time.Duration(i)*time.Hour),
				},
				{
					PostID:    post.ID,
					Name:      "李四",
					Email:     "lisi@example.com",
					Content:   "感谢分享，期待更多类似的文章！",
					CreatedAt: time.Now().Add(-5*24*time.Hour + time.Duration(i)*time.Hour),
				},
			}

			for _, comment := range comments {
				err = utils.DB.Create(&comment).Error
				if err != nil {
					panic(err)
				}
				println("  - 添加评论:", comment.Name)
			}
		}
	}

	println("\n测试数据生成完成！")
}