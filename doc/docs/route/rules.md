---
sidebar_position: 2
---

# 自定义路由 和 DNS

UIF 有强大的路由规则。

## 定义规则

比如说你想要 x.com 网站走 `直连`，到 [路由页面](https://uiforfreedom.github.io/#/route/my)，点击右上角 `添加` 按钮，选择 `域名规则` -> `完全匹配` -> 填入 x.com。

那么此时只要访问 x.com 都会走 `直连`

一些规则的功能：
- `完全匹配`： 一模一样
- `后缀匹配`： 后缀一样
- `关键字匹配`： 存在该关键字
- `正则匹配`： 使用正则

- `数据库`： 已预留常见的域名列表或国家列表，使用`国家ISCODE`代表某个国家。

