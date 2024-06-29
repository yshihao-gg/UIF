---
sidebar_position: 1
---

# 用 UIF 搭建

请先确保你有属于自己的，并且能访问国外网站的 VPS 服务器。

一般该 VPS 服务器都是 Linux 操作系统

## 1 安装 UIF

### 1.1 安装 `ufw`

防火墙管理工具，方便小白一键操作；如果你知道这种用来干嘛的，也可以自行决定是否安装。

```bash
apt install ufw
```

一些常见命令：

```bash
apt install ufw # 确保已安装 ufw
ufw enable # 确保已启用
ufw allow 9527/tcp # 放行 9413 的 tcp 流量
ufw allow 80/tcp # 放行 80 的 tcp 流量
ufw allow 443/udp # 放行 443 的 udp 流量
```

### 1.2 安装主程序：

支持 Ubuntu 和 Centos 一键安装 （推荐 Ubuntu）：

```bash
bash <(curl https://raw.githubusercontent.com/UIforFreedom/UIF/master/uifd/linux_install.sh)
```

**成功安装后，UIF 所有的东西都会放在目录 `/usr/bin/uif/`，可自行检查是否出错**

- 脚本会自动使用 ufw 开放并占用`80`端口，可以用命令 `netstat -tunlp | grep 80` 检查一下端口情况
- 用 `chmod -R 755 /usr/bin/uif/` 确保有运行权限
- [如何修改 UIF 默认的端口？](../setting)

### 1.3 启动

安装好后，使用下面命令启动 UIF。

```bash
systemctl start uiforfreedom
```

还有一些常见命令：

- `systemctl status uiforfreedom` 查看 UIF 的状态

- `systemctl stop uiforfreedom` 关闭 UIF

![pic alt](../pics/55.gif)

## 2 查看密码

使用命令查看 UIF 自动生成的密码

```bash
cat /usr/bin/uif/uif_key.txt
```

使用上面命令，应该会输出 UIF 的密码；如果显示的是该文件不存在，那么就意味着 UIF 未启动或未正确安装。

## 3 访问 Web

- 如果你已安装了并运行 UIF，那么浏览器打开 [http://127.0.0.1:9527](http://127.0.0.1:9527) 即可看到 UIF 面板。（推荐）

- 你也可以直接访问 [http://ui4freedom.org](http://ui4freedom.org) 也同样可以使用UIF面板。

到 `主页`，把`UIF 接口地址`设置成 `http://服务器的 IP:80`，举个例子就是`http://7.7.7.7:80`（端口地址为上面安装脚本自动占用），

填入上面获得的 `密码`，点击右上角的 `连接后端` 即可。密码会缓存在本机浏览器的 cookie，首次访问才需要填入。

![pic alt](../pics/66.gif)

## 4 设置入站

到 [入站页面](https://uiforfreedom.github.io/#/in/my)，点击右上角 `添加`，选择你喜欢的 `代理协议`，这里以 `hysteria2`，作为演示。

![pic alt](../pics/77.gif)

## 5 复制并导入

到 [设置页面](https://uiforfreedom.github.io/#/settings/uif)，点击 `复制分享链接`。

### 5.1 分享到安卓

到 `Profile` 新建订阅，类型选择 `remote`，粘贴链接即可。

![pic alt](../pics/88.gif)

### 5.2 PC

到 [我的订阅](https://uiforfreedom.github.io/#/out/subscribe)，点击 `添加订阅` 按钮，粘贴链接即可。
