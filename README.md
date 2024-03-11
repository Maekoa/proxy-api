# proxy-api
只要在代理项目中引入 `github.com/maekoa/proxy-api/xx`，实现 `XXCompatible`
接口，调用 `xx.From(YourProxy).Run()` ，即可获得相应的 REST API，可以搭配
兼容的客户端使用，减少前端造轮子。

## 个人想法
一个代理软件应该实现的：
- 代理
- 出入站集合
- 路由
应该经由 API 交给客户端管理的
- 订阅
- 分流规则
- 连接

因此本项目将处理
- providers 订阅
- 代理组
