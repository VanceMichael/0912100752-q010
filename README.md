# 三人篮球赛程服务

服务接收裁判比分事件，计算分组积分与净胜分排名，并通过审计接口追踪封存比赛的更正。Redis 用作短时排名缓存。

## 运行

`docker compose up --build` 启动 API 和 Redis，HTTP 端口为 8084。
