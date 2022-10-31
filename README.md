# go-wxxcx

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sukaifei/go-wxxcx)
![GitHub last commit](https://img.shields.io/github/last-commit/SuKaifei/go-wxxcx)
![GitHub branch checks state](https://img.shields.io/badge/license-GNU-green)
![GitHub search hit counter](https://img.shields.io/github/search/SuKaiFei/go-wxxcx/go)
![GitHub Repo stars](https://img.shields.io/github/stars/sukaifei/go-wxxcx?style=social)



## 简介

微信小程序后台服务；

## 已上线的小程序


|                     微信小程序码                       |                    赞赏码(觉得不错打个赏😁)                     |
|:-----------------------------------------------:|:-----------------------------------------------------:|
| <img src="./assets/image/mp.jpg" width="200px"> | <img src="./assets/image/zanshang.jpg" width="200px"> |

- 鸡音盒(主线产品，也是总入口)
- 纯真盒
- 夹子音盒
- 时代马戏团
- 表情包庇

## 🪤 编译+部署

### 命令

`make buildlinux scp deploy`

### 平滑部署机制

每次部署新程序时随机启动一个可用的端口，
检测到新端口可用时将`nginx`的代理地址配置更改为新端口，
再把之前启动的进程ID`kill`掉，
这样就能利用`nginx`的反向代理实现的平滑部署。

## 🤝 特别感谢

1. [Kratos](https://github.com/go-kratos/kratos)
1. [gorm](https://github.com/go-gorm/gorm)
1. [nginx](https://github.com/nginx/nginx)
