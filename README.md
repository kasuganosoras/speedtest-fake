# speedtest-fake
装逼用的 Speedtest 脚本，专治网速装逼大佬

- 支持自定义物理位置、IP 地址
- 支持自定义延迟，ISP，地区
- 支持模拟原版测速效果

~~需要 PHP 环境才能运行，没有 php 环境请安装~~

鉴于网友们觉得要装个 PHP 环境太麻烦，我又写了个 Golang 版的

拉取下来后直接 `chmod +x speedtestgo` 即可开始玩耍

## Golang 版使用方法

以下是命令行参数

```
./speedtestgo <地区> <ip> <测速提供商> <测速服务器地区> <距离> <随机ping最小值> <随机ping最大值> <随机上传最小值> <随机上传最大值> <随机下载最小值> <随机下载最大值>
```

示例使用

```
./speedtestgo "China Telecom Guangdong" "183.15.172.10" "W Professional Services Limited" "New Territories" "18.14" 10 100 10000 20000 10000 20000
```

## PHP 版使用方法
拉取脚本

```
cd ~/
git clone https://github.com/kasuganosoras/speedtest-fake
```

修改配置和参数

```
vi speedtest.php
```

尝试运行

```
php speedtest.php
```

修改命令行配置，让效果更逼真

```
vi ~/.bashrc
```

在结尾增加一行，添加一个命令别名

```
alias speedtest='php ~/speedtest-fake/speedtest.php'
```

让配置生效

```
source ~/.bashrc
```

开始装逼测速

```
speedtest
```

## 测速效果图

![img](https://i.natfrp.org/90a86a02e1c7d9e22ec1f886e4f89f66.gif)

### 装逼虽好，不要贪装哦

快用神仙一般的网速去吓死你的小伙伴吧！
