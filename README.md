# Web Blog System

一个基于Go语言开发的轻量级Web博客系统，支持文章的创建、编辑、删除和浏览功能。

## 🌟 功能特性

- ✅ 文章管理：创建、编辑、删除、查看文章
- ✅ 分类支持：文章分类管理
- ✅ 响应式设计：适配各种设备屏幕
- ✅ SQLite数据库：轻量级数据存储
- ✅ 模板引擎：使用Go HTML模板渲染页面
- ✅ 环境配置：支持.env文件配置
- ✅ 单元测试：完善的测试用例

## 🛠️ 技术栈

- **后端框架**：Go (Golang)
- **Web框架**：标准库net/http
- **数据库**：SQLite3
- **模板引擎**：Go HTML Templates
- **环境配置**：godotenv

## 📦 安装与运行

### 1. 系统要求

- Go 1.16+ 环境
- SQLite3

### 2. 安装步骤

#### 克隆项目

```bash
git clone https://github.com/wellpain/blog_web_demo.git
cd blog_web_demo
```

#### 安装依赖

```bash
go mod download
```

#### 配置环境变量

创建并编辑 `.env` 文件：

```bash
cp .env.example .env
```

配置内容示例：

```env
# 服务器配置
PORT=8080
HOST=0.0.0.0

# 数据库配置
DATABASE_PATH=blog.db

# 应用配置
APP_NAME=Web Blog
APP_VERSION=1.0.0
```

#### 生成测试数据（可选）

```bash
go run testdata/generate_test_data.go
```

#### 编译并运行

```bash
# 编译
go build -o blog main.go

# 运行
./blog
```

或者直接运行：

```bash
go run main.go
```

### 3. 访问应用

打开浏览器访问：http://localhost:8080

## 📁 项目结构
