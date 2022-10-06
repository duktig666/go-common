# couples-subtotal

## 目录约束

```
- app 业务
-- module （具体模块业务，例如 admin）
--- apis
--- models
--- router
--- service
- cmd 命令行配置
- common 公共依赖
- config 配置
- docs 文档
- examples 示例
- scripts 脚本文件
- static 静态文件目录
- template go模版文件
- test 测试目录
- pkg 第三方工具包
```

## 运行
### 初始化项目
```shell
./cli-init.sh
```

### 运行
```shell
# cli帮助
./couples-cli help

# 查看版本
./couples-cli version

# 运行http服务
./couples-cli server
./couples-cli server -c config/config-xxx.yaml
```
