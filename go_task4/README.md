# Go Task4 - 博客系统 API

一个使用 Go 和 Gin 框架构建的博客系统 API，支持用户认证、文章管理和评论功能。

## 功能特性

- ✅ 用户认证（注册、登录）
- 📝 文章管理（CRUD）
- 💬 评论功能
- 🔐 JWT 认证
- 🗄️ 数据库迁移
- 🛡️ 请求验证
- 🔄 错误处理中间件

## 技术栈

- **后端框架**: Gin
- **数据库**: MySQL
- **ORM**: GORM
- **认证**: JWT
- **环境管理**: godotenv

## 环境要求

- Go 1.16+
- MySQL 5.7+
- Git

## 快速开始

### 1. 克隆项目
```bash
git clone <your-repo-url>
cd go_task4
```

### 2. 配置环境变量

复制 `.env.example` 为 `.env` 并修改配置：
```bash
cp .env.example .env
```
编辑 `.env` 文件，设置您的数据库连接信息。

### 3. 安装依赖
```bash
go mod download
```

### 4. 数据库设置
- 创建 MySQL 数据库
- 运行应用，自动执行数据库迁移

### 5. 启动应用
```bash
# 开发模式
go run main.go

# 生产模式
go build -o app
./app
```
服务器将运行在 `http://localhost:8080`

以下是完整的API文档和测试部分内容，涵盖所有9个接口：

## API 文档

### 认证相关接口

#### 1. 用户注册
**URL**: `POST /auth/register`  
**Headers**: `Content-Type: application/json`  
**Body**:
```json
{
    "username": "string (必填，唯一)",
    "email": "string (必填，邮箱格式)",
    "password": "string (必填，最少6位)"
}
```
**成功响应 (201)**:
```json
{
    "success": true,
    "message": "用户注册成功",
    "data": {
        "id": "number",
        "username": "string",
        "email": "string"
    }
}
```
**错误响应**:
- 400: 参数验证失败
- 409: 用户名或邮箱已存在
- 500: 服务器内部错误

#### 2. 用户登录
**URL**: `POST /auth/login`  
**Headers**: `Content-Type: application/json`  
**Body**:
```json
{
    "username": "string (必填)",
    "password": "string (必填)"
}
```
**成功响应 (200)**:
```json
{
    "success": true,
    "message": "登录成功",
    "data": {
        "token": "jwt_token",
        "user": {
            "id": "number",
            "username": "string",
            "email": "string"
        }
    }
}
```
**错误响应**:
- 400: 参数验证失败
- 401: 用户名或密码错误
- 500: 服务器内部错误

### 文章相关接口

#### 3. 获取所有文章
**URL**: `GET /posts`  
**成功响应 (200)**:
```json
{
    "success": true,
    "message": "获取文章列表成功",
    "data": [
        {
            "id": "number",
            "title": "string",
            "content": "string",
            "User": {
                "id": "number",
                "username": "string"
            },
            "Comments": []
        }
    ]
}
```

#### 4. 获取单个文章
**URL**: `GET /posts/:id`  
**成功响应 (200)**:
```json
{
    "success": true,
    "message": "获取文章成功",
    "data": {
        "id": "number",
        "title": "string",
        "content": "string",
        "User": {
            "id": "number",
            "username": "string"
        },
        "Comments": [
            {
                "id": "number",
                "content": "string",
                "User": {
                    "id": "number",
                    "username": "string"
                }
            }
        ]
    }
}
```

#### 5. 创建文章 (需认证)
**URL**: `POST /posts`  
**Headers**: 
```
Authorization: Bearer <token>
Content-Type: application/json
```
**Body**:
```json
{
    "title": "string (必填)",
    "content": "string (必填)"
}
```
**成功响应 (201)**:
```json
{
    "success": true,
    "message": "文章创建成功",
    "data": {
        "id": "number",
        "title": "string",
        "content": "string",
        "UserID": "number"
    }
}
```

#### 6. 更新文章 (需认证)
**URL**: `PUT /posts/:id`  
**Headers**: 
```
Authorization: Bearer <token>
Content-Type: application/json
```
**Body**:
```json
{
    "title": "string (可选)",
    "content": "string (可选)"
}
```
**成功响应 (200)**:
```json
{
    "success": true,
    "message": "文章更新成功",
    "data": {
        "id": "number",
        "title": "string",
        "content": "string"
    }
}
```

#### 7. 删除文章 (需认证)
**URL**: `DELETE /posts/:id`  
**Headers**: `Authorization: Bearer <token>`  
**成功响应 (200)**:
```json
{
    "success": true,
    "message": "文章删除成功",
    "data": null
}
```

### 评论相关接口

#### 8. 获取文章评论
**URL**: `GET /comments/:postId`  
**成功响应 (200)**:
```json
{
    "success": true,
    "message": "获取评论列表成功",
    "data": [
        {
            "id": "number",
            "content": "string",
            "User": {
                "id": "number",
                "username": "string"
            }
        }
    ]
}
```

#### 9. 创建评论 (需认证)
**URL**: `POST /comments/:postId`  
**Headers**: 
```
Authorization: Bearer <token>
Content-Type: application/json
```
**Body**:
```json
{
    "content": "string (必填)"
}
```
**成功响应 (201)**:
```json
{
    "success": true,
    "message": "评论创建成功",
    "data": {
        "id": "number",
        "content": "string",
        "UserID": "number",
        "PostID": "number"
    }
}
```



## Postman 测试指南

### 测试准备
1. 下载并安装 [Postman](https://www.postman.com/downloads/)
2. 启动服务：`go run main.go`
3. 创建新 Collection 命名为 "Blog API"
4. 设置基础 URL 变量：
   - 点击 Collection 的 "Variables" 标签
   - 添加变量 `base_url` 值为 `http://localhost:8080`

### 1. 用户注册
**请求**:
- 方法: POST
- URL: `{{base_url}}/auth/register`
- Headers:
  - `Content-Type: application/json`
- Body (raw/JSON):
```json
{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
}
```

**测试脚本** (Tests 标签):
```javascript
pm.test("Status code is 201", function() {
    pm.response.to.have.status(201);
});
pm.test("Registration successful", function() {
    var jsonData = pm.response.json();
    pm.expect(jsonData.success).to.eql(true);
});
```

### 2. 用户登录
**请求**:
- 方法: POST
- URL: `{{base_url}}/auth/login`
- Headers:
  - `Content-Type: application/json`
- Body:
```json
{
    "username": "testuser",
    "password": "password123"
}
```

**测试脚本**:
```javascript
pm.test("Status code is 200", function() {
    pm.response.to.have.status(200);
});
pm.test("Login successful", function() {
    var jsonData = pm.response.json();
    pm.expect(jsonData.data.token).to.exist;
    // 将 token 保存为环境变量
    pm.environment.set("auth_token", jsonData.data.token);
});
```

### 3. 获取所有文章
**请求**:
- 方法: GET
- URL: `{{base_url}}/posts`

**测试脚本**:
```javascript
pm.test("Status code is 200", function() {
    pm.response.to.have.status(200);
});
```

### 4. 获取单个文章
**请求**:
- 方法: GET
- URL: `{{base_url}}/posts/1` (先确保有ID为1的文章)

**测试脚本**:
```javascript
pm.test("Status code is 200", function() {
    pm.response.to.have.status(200);
});
```

### 5. 创建文章
**请求**:
- 方法: POST
- URL: `{{base_url}}/posts`
- Headers:
  - `Authorization: Bearer {{auth_token}}`
  - `Content-Type: application/json`
- Body:
```json
{
    "title": "Postman测试文章",
    "content": "这是用Postman创建的文章内容"
}
```

**测试脚本**:
```javascript
pm.test("Status code is 201", function() {
    pm.response.to.have.status(201);
});
// 保存文章ID供后续测试使用
pm.test("Save post ID", function() {
    var jsonData = pm.response.json();
    pm.environment.set("post_id", jsonData.data.id);
});
```

### 6. 更新文章
**请求**:
- 方法: PUT
- URL: `{{base_url}}/posts/{{post_id}}`
- Headers:
  - `Authorization: Bearer {{auth_token}}`
  - `Content-Type: application/json`
- Body:
```json
{
    "title": "更新后的标题",
    "content": "更新后的内容"
}
```

**测试脚本**:
```javascript
pm.test("Status code is 200", function() {
    pm.response.to.have.status(200);
});
```

### 7. 删除文章
**请求**:
- 方法: DELETE
- URL: `{{base_url}}/posts/{{post_id}}`
- Headers:
  - `Authorization: Bearer {{auth_token}}`

**测试脚本**:
```javascript
pm.test("Status code is 200", function() {
    pm.response.to.have.status(200);
});
```

### 8. 获取文章评论
**请求**:
- 方法: GET
- URL: `{{base_url}}/comments/{{post_id}}`

**测试脚本**:
```javascript
pm.test("Status code is 200", function() {
    pm.response.to.have.status(200);
});
```

### 9. 创建评论
**请求**:
- 方法: POST
- URL: `{{base_url}}/comments/{{post_id}}`
- Headers:
  - `Authorization: Bearer {{auth_token}}`
  - `Content-Type: application/json`
- Body:
```json
{
    "content": "这是用Postman创建的评论"
}
```

**测试脚本**:
```javascript
pm.test("Status code is 201", function() {
    pm.response.to.have.status(201);
});
// 保存评论ID
pm.test("Save comment ID", function() {
    var jsonData = pm.response.json();
    pm.environment.set("comment_id", jsonData.data.id);
});
```

### 环境变量管理
1. 点击右上角的眼睛图标查看当前环境变量
2. 重要变量：
   - `base_url`: API基础地址
   - `auth_token`: 登录后获得的JWT令牌
   - `post_id`: 创建的文章ID
   - `comment_id`: 创建的评论ID

### 测试流程建议
1. 按顺序运行：注册 → 登录 → 创建文章 → 创建评论 → 其他操作
2. 可以使用Postman的 "Runner" 功能批量执行测试
3. 每次测试前可以点击 "Clear" 清除之前的响应数据

### 导出分享
1. 点击 Collection 的 "..." 选择 "Export" 可导出测试集合
2. 选择 "Collection v2.1" 格式导出为JSON文件
3. 分享给团队成员导入即可使用完整测试环境

### 高级技巧
1. 在Pre-request Script中可以设置动态变量
2. 使用 `pm.collectionVariables` 替代 `pm.environment` 管理集合级变量
3. 添加更多断言验证响应数据结构

## 部署

### 构建应用
```bash
go build -o app
```

### 运行
```bash
# Linux/Mac
./app

# Windows
app.exe
```

