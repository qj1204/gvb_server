# gvb_server 

小新知道个人博客后端项目

在线预览：https://www.xiaoxinqj.fun


## 项目介绍

1. 文章数据直接存储到es数据库，使用redis缓存点赞数据，浏览量，收藏量
2. 支持文章全文搜索，按文章分类、标签搜索，按相关数据进行排序
3. 系统配置项众多，满足定制化需求
4. 日志系统，登录，操作，运行日志，可帮助管理员更好的观察系统运行情况
5. 基于 WebSocket 实现的聊天室，以及私聊功能
6. 采用 JWT 进行鉴权
7. 图片上传支持七牛云/本地

## 技术栈

- Golang
- Gin
- GORM
- MySQL
- Redis
- ElasticSearch
- Nginx: 部署静态资源 + 反向代理
- Docker
- ...

其他:

- QQ登录
- 七牛云对象存储
- ...