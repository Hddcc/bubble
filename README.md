# GoList - 简单的待办事项应用

一个简洁的待办事项（To-Do List）Web 应用，后端使用 Go 语言。

## 功能特性

- **创建**：添加新的待办事项。
- **查看**：浏览所有待办事项。
- **更新**：将事项标记为已完成。
- **删除**：从列表中移除事项。

## 技术栈

- **后端**：Go
  - **Web 框架**：Gin
  - **ORM**：GORM
- **数据库**：MySQL

## 快速开始

### 环境依赖

- Go (建议 1.18 或更高版本)
- 一个正在运行的 MySQL 服务

### 安装与运行

1.  **克隆仓库：**
    ```bash
    git clone https://github.com/Hddcc/GoList.git
    cd GoList
    ```

2.  **配置数据库：**
    打开 `dao/mysql.go` 文件，将其中的 DSN 字符串修改为您的 MySQL 连接信息。
    ```go
    dsn := "user:password@(host:port)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
    ```
    您还需要手动创建 DSN 中指定的数据库。

3.  **安装依赖：**
    此命令将下载所有必需的 Go 模块。
    ```bash
    go mod tidy
    ```

4.  **运行程序：**
    ```bash
    go run main.go
    ```
    程序默认将在 `http://127.0.0.1:8080` 启动。

## 项目结构

```
.
├── controller/   # 处理请求逻辑
├── dao/          # 数据访问对象 (数据库连接)
├── main.go       # 程序入口文件
├── models/       # 数据结构与数据库操作
├── routers/      # API 路由配置
├── static/       # 编译后的前端静态资源 (JS, CSS)
└── templates/    # 主页 index.html 模板
```
