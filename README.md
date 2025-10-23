# Travel Planner - 智能旅行规划系统

## 项目概述

Travel Planner 是一个基于 AI 的智能旅行规划系统，集成了多模态 AI 能力、地图服务和交通信息查询功能。系统采用微服务架构，包含前端界面和后端 API 服务，为用户提供个性化的旅行路线规划和景点推荐服务。

## 项目结构

```
travel-planner/
├── amap/                    # 高德地图前端应用 (Vue 3 + Vite)
│   ├── src/
│   │   ├── components/
│   │   │   └── MapContainer.vue    # 地图容器组件
│   │   ├── App.vue                 # 主应用组件
│   │   └── main.js                 # 入口文件
│   ├── package.json
│   └── vite.config.js
├── frontend/                 # 主前端应用 (Vue 3 + Vite)
│   ├── src/
│   │   ├── components/
│   │   │   └── HelloWorld.vue      # 示例组件
│   │   ├── App.vue                 # 主应用组件
│   │   └── main.js                 # 入口文件
│   └── package.json
├── go-agent/                # Go 语言后端服务
│   ├── biz/                 # 业务逻辑层
│   │   ├── agent/           # AI 智能体模块
│   │   │   ├── agent.go     # 智能体初始化
│   │   │   └── runner.go    # 运行器管理
│   │   ├── config/          # 配置管理
│   │   │   ├── mcp.go       # MCP 工具配置
│   │   │   ├── model.go     # AI 模型配置
│   │   │   ├── database.go  # 数据库配置
│   │   │   └── redis.go     # Redis 配置
│   │   ├── handler/         # HTTP 处理器
│   │   │   ├── basic.go     # 通用处理器
│   │   │   └── chat.go      # 聊天处理器
│   │   ├── router/          # 路由定义
│   │   │   ├── chat.go      # 聊天路由
│   │   │   ├── account.go   # 用户路由
│   │   │   └── register.go  # 路由注册
│   │   ├── service/         # 业务服务层
│   │   │   └── chat.go      # 聊天服务
│   │   └── param/           # 参数定义
│   │       └── chat.go      # 聊天参数
│   ├── test/                # 测试代码
│   │   ├── chatmodelagent.go
│   │   └── streamagent.go
│   ├── main.go              # 应用入口
│   ├── go.mod               # Go 模块依赖
│   └── build.sh             # 构建脚本
├── README.md                # 项目说明文档
└── result.md                # 示例结果文档
```

## 核心功能

### 1. 智能旅行规划
- **路线规划**: 支持驾车、高铁等多种出行方式规划
- **多方案对比**: 提供耗时、费用、舒适度等多维度对比
- **实时交通信息**: 集成高德地图和12306 API

### 2. 景点推荐系统
- **个性化推荐**: 基于用户偏好推荐景点
- **周边搜索**: 支持基于位置的周边景点发现
- **详细信息**: 提供景点介绍、门票信息等

### 3. AI 对话助手
- **多轮对话**: 支持上下文相关的智能对话
- **流式响应**: 实时返回AI生成内容
- **多模态支持**: 支持文本和图片输入

## 技术栈

### 后端技术
- **框架**: CloudWeGo Hertz (高性能 HTTP 框架)
- **AI 引擎**: CloudWeGo Eino (AI 智能体框架)
- **数据库**: MySQL + GORM (ORM 框架)
- **缓存**: Redis
- **向量数据库**: Milvus (用于语义搜索)
- **AI 模型**: 
  - 通义千问 (Ark)
  - OpenAI 兼容模型
  - 视觉模型支持

### 前端技术
- **框架**: Vue 3 + Vite
- **地图**: 高德地图 JS API
- **UI 组件**: 原生 Vue 组件

### 外部服务集成
- **地图服务**: 高德地图 API
- **交通信息**: 12306 火车票查询
- **搜索服务**: Jina AI 搜索
- **网络搜索**: Bing 搜索

## 快速开始

### 环境要求
- Go 1.23+
- Node.js 16+
- MySQL 8.0+
- Redis 6.0+

### 后端启动

1. 配置环境变量
```bash
cd go-agent
cp .env.example .env
# 编辑 .env 文件配置数据库和 API 密钥
```

2. 安装依赖
```bash
go mod tidy
```

3. 构建并运行
```bash
./build.sh
./hertz_service
```

### 前端启动

1. 安装依赖
```bash
cd frontend
npm install
```

2. 开发模式运行
```bash
npm run dev
```

3. 地图应用运行
```bash
cd amap
npm install
npm run dev
```

## API 接口

### 聊天接口
- `POST /chat/user/stream` - 流式聊天接口
- `POST /chat/user/conversation` - 创建对话
- `GET /chat/user/conversations` - 获取对话列表

### 文件上传
- `POST /chat/user/files` - 文件上传接口

### 用户认证
- `POST /auth/public/login` - 用户登录
- `POST /auth/public/register` - 用户注册

## 配置说明

### AI 模型配置
系统支持多种 AI 模型，通过配置文件管理：
- 通义千问模型 (Ark)
- OpenAI 兼容模型
- 视觉模型
- 嵌入模型

### MCP 工具配置
系统通过 MCP (Model Context Protocol) 集成外部工具：
- 高德地图工具
- 12306 火车票工具
- 百度地图工具
- Jina AI 搜索工具
- Bing 搜索工具

## 项目特色

1. **智能体架构**: 基于 CloudWeGo Eino 框架构建的智能体系统
2. **多工具集成**: 通过 MCP 协议灵活集成外部服务
3. **流式响应**: 支持实时流式 AI 响应
4. **多模态支持**: 支持文本和图片输入处理
5. **模块化设计**: 清晰的业务分层和模块划分

## 开发指南

### 添加新的 AI 工具
1. 在 `biz/config/mcp.go` 中添加新的 MCP 服务器配置
2. 在相应的智能体中集成新工具
3. 更新系统提示词以支持新功能

### 扩展业务功能
1. 在 `biz/service/` 中添加新的服务类
2. 在 `biz/handler/` 中添加对应的处理器
3. 在 `biz/router/` 中注册新的路由

## 许可证

本项目采用 MIT 许可证。

## 贡献指南

欢迎提交 Issue 和 Pull Request 来改进项目。

## 联系方式

如有问题或建议，请通过以下方式联系：
- 项目 Issues: [GitHub Issues]
- 邮箱: [项目维护者邮箱]

---

*最后更新: 2025年*