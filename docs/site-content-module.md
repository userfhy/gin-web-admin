# 内容管理模块说明（官网）

## 1. 目标
- 支撑后台内容运营：分类管理、内容管理、发布下线、列表筛选。
- 支撑前台官网读取：公开分类、公开内容列表、公开内容详情。
- 便于后续扩展：标签、作者、SEO、审核流、多语言等。

## 2. 后端结构（gin-web-admin）

### 2.1 Model
- `app/models/site_category_model.go`
  - 分类实体、分页查询、按 slug 查询、内容数统计。
- `app/models/site_content_model.go`
  - 内容实体（含 `status`、`published_at`）、分页查询、按 slug/id 查询、状态更新。
- `app/models/site_content_category_model.go`
  - 内容-分类关联表（多对多），支持关联替换与按内容查分类。
- `app/models/site_tag_model.go`
  - 标签实体、分页查询、按 slug 查询、内容数统计。
- `app/models/site_content_tag_model.go`
  - 内容-标签关联表（多对多），支持关联替换与按内容查标签。

### 2.2 Service
- `app/service/v1/site_category/site_category_service.go`
  - 分类 CRUD、分类列表（含 `contentCount`）。
- `app/service/v1/site_content/site_content_service.go`
  - 内容 CRUD、内容关联分类/标签、摘要自动生成、快速发布/下线。
- `app/service/v1/site_public/site_public_service.go`
  - 官网公开数据聚合：
  - 公开分类列表
  - 公开内容列表（仅已发布）
  - 公开内容详情（按 id/slug）

### 2.3 Controller
- `app/controllers/v1/site_category/site_category_controller.go`
- `app/controllers/v1/site_content/site_content_controller.go`
- `app/controllers/v1/site_public/site_public_controller.go`
- `app/controllers/v1/site_tag/site_tag_controller.go`

### 2.4 Router
- 后台接口（鉴权）
  - `routers/site_category_router.go`
  - `routers/site_content_router.go`
  - `routers/site_tag_router.go`
- 公开接口（无需登录）
  - `routers/site_public_router.go`
- 主路由注册
  - `routers/router.go`

## 3. 前端结构（web-admin-frontend）

### 3.1 API
- `src/api/siteCategory.ts`
- `src/api/siteContent.ts`

### 3.2 页面
- 内容管理
  - `src/views/site/content/index.vue`
  - `src/views/site/content/form.vue`
  - `src/views/site/content/utils/hook.tsx`
- 类别管理
  - `src/views/site/category/index.vue`
  - `src/views/site/category/form.vue`
  - `src/views/site/category/utils/hook.tsx`

### 3.3 已实现交互
- 内容：
  - 列表筛选（关键词、状态、类别）
  - 新增/编辑（Markdown 正文、关联分类多选）
  - 快速发布/下线（开关）
  - 发布时间展示
  - slug 一键由标题生成
- 类别：
  - 列表筛选、CRUD
  - 每个类别内容数展示

## 4. 接口清单

### 4.1 后台接口（需要 JWT + Casbin）
- 分类
  - `GET /v1/api/site-category`
  - `GET /v1/api/site-category/all`
  - `POST /v1/api/site-category`
  - `PUT /v1/api/site-category/{id}`
  - `DELETE /v1/api/site-category/{id}`
- 内容
  - `GET /v1/api/site-content`
  - `GET /v1/api/site-content/{id}`
  - `POST /v1/api/site-content`
  - `PUT /v1/api/site-content/{id}`
  - `PATCH /v1/api/site-content/{id}/status`
  - `DELETE /v1/api/site-content/{id}`
- 标签
  - `GET /v1/api/site-tag`
  - `GET /v1/api/site-tag/all`
  - `POST /v1/api/site-tag`
  - `PUT /v1/api/site-tag/{id}`
  - `DELETE /v1/api/site-tag/{id}`

### 4.2 公开接口（无需登录）
- `GET /v1/api/site/public/categories`
- `GET /v1/api/site/public/tags`
- `GET /v1/api/site/public/contents`
  - 支持参数：`pageNum`、`pageSize`、`keyword`、`categorySlug`
  - 支持参数：`pageNum`、`pageSize`、`keyword`、`categorySlug`、`tagSlug`
  - `summary` 为截断文本，不返回完整正文
- `GET /v1/api/site/public/contents/{id}`（详情）
- `GET /v1/api/site/public/contents/slug/{slug}`（兼容）

## 5. 数据字段建议

### 5.1 内容（site_content）
- 基础：`title`、`slug`、`summary`、`cover`、`content`
- SEO：`seo_keywords`、`seo_description`
- 发布：`status`（1 发布 / 0 草稿）、`published_at`
- 排序：`sort`

### 5.2 分类（site_category）
- `name`、`slug`、`description`、`status`、`sort`

### 5.3 关联（site_content_category）
- `content_id`、`category_id`（联合主键）

### 5.4 关联（site_content_tag）
- `content_id`、`tag_id`（联合主键）

## 6. 菜单与权限
- 菜单结构：
  - 一级：官网管理（目录）
  - 二级：内容管理（页面）
  - 二级：类别管理（页面）
  - 二级：标签管理（页面）
- 注意：
  - 新增菜单后需写入 `gin_role_menu` 给角色授权。
  - 前端动态路由有本地缓存，菜单变更后建议重新登录或清理 `async-routes`。

## 7. 扩展建议（下一步）
- 标签体系（content_tag + content_tag_rel）
- 作者/来源/发布时间手动控制
- 内容置顶、推荐位、浏览量
- 草稿版本与发布版本分离
- 审核流（编辑 -> 审核 -> 发布）
- SEO URL 规范校验（全局唯一、重定向策略）
- 静态化输出（按 slug 导出 HTML）

## 9. 安全策略（XSS）
- 入库前清洗（后端 service 层统一执行）：
  - 纯文本字段：`title`、`summary`、`cover`、`seoKeywords`、`seoDescription`、`category.name`、`category.description`
  - Markdown 字段：`content`
- 清洗规则：
  - 去除控制字符
  - 去除危险标签（`script/style/iframe/object/embed/link/meta`）
  - 去除危险协议（`javascript:`、`vbscript:`、`data:`）
  - Markdown 内容做 HTML 实体转义后入库（更强防护）
  - 清洗前先执行一次反转义，避免二次编辑导致重复转义（如 `&amp;lt;`）
- 公开列表接口中的 `summary` 会强制截断，避免返回完整正文。
- 建议：前端在渲染 Markdown 为 HTML 时继续启用渲染层白名单策略（双层防护）。

## 8. 联调排查 checklist
- 分类下拉为空：先检查 `GET /v1/api/site-category` 是否有数据、是否有权限。
- 前台列表为空：检查内容 `status` 是否为 1（已发布）。
- 详情 404/不存在：检查 id/slug 是否正确，是否已发布。
- `app/service/v1/site_tag/site_tag_service.go`
  - 标签 CRUD、标签列表（含 `contentCount`）。
