# speedtest-fake
装逼用的 Speedtest 脚本，专治网速装逼大佬

- 支持自定义物理位置、IP 地址
- 支持自定义延迟，ISP，地区
- 支持模拟原版测速效果

需要 PHP 环境才能运行，没有 php 环境请安装

## 使用方法
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
alias speedtest='php speedtest.php'
```

让配置生效

```
source ~/.bashrc
```

开始装逼测速

```
speedtest
```

### 装逼虽好，不要贪装哦

快用神仙一般的网速去吓死你的小伙伴吧！
