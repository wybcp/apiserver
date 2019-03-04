# go api server

基于 gin 框架的 RESTful api 服务器

- viper：配置文件和热更新配置数据
- https 支持
- gorm MySQL 数据库
- 日志记录
- middleware 中间件处理
- 定制错误信息
- jwt token 认证
- api 版本
- Swagger：API 文档构建工具（gin-swagger）

```bash
cp apiserver/conf/config.yaml.example apiserver/conf/config.yaml
```

并配置数据库
