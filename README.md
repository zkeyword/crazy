CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

go build -ldflags "-s -w" main.go


// google authenticator
https://studygolang.com/articles/25748
https://www.jianshu.com/p/3c4832e3c6d5

// 逆向
https://zhuanlan.zhihu.com/p/26733683

// 打包静态文件
安装 go install github.com/go-bindata/go-bindata/go-bindata@latest

命令 //go:generate go-bindata -fs -nocompress -nomemcopy -o=pkg/assets/assets.go -pkg=assets ./assets

// 部署  // 貌似不能自动重启
nohup ./main &1>2 & 

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go

nohup 部署 日志 https://www.cnblogs.com/zhangmingcheng/p/11577967.html

// 热更新
go get github.com/pilu/fresh

在环境变量中配置了go bin的目录正常的话就可以正常执行，进入项目目录执行

    fresh

可以代替

    go run main.go

// 
go mod tidy的使用
引用项目需要的依赖增加到go.mod文件。
去掉go.mod文件中项目不需要的依赖。


系统
    用户管理
    角色管理
    菜单管理
    权限分配
    定时任务
    数据字典
    操作日志
    登录日志
    cms管理
    消息管理
        消息模版
        发送短信
        邮件消息
商场管理
    会员管理
    商品类别
    商品管理
    订单管理
    购物车
    banner管理
    收藏列表

手机端
    首页
    分类
    详情
    购物车
    订单