## 编译等命令

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
    go build -ldflags "-s -w" main.go

## 打包静态文件
    安装 go install github.com/go-bindata/go-bindata/go-bindata@latest

    命令 //go:generate go-bindata -fs -nocompress -nomemcopy -o=pkg/assets/assets.go -pkg=assets ./assets

## 部署  // 貌似不能自动重启
    nohup ./main &1>2 & 

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go

    nohup 部署 日志 https://www.cnblogs.com/zhangmingcheng/p/11577967.html

## 开发
    热更新
    go get github.com/pilu/fresh

    在环境变量中配置了go bin的目录正常的话就可以正常执行，进入项目目录执行

        fresh

    可以代替

        go run main.go

## 表字段设计

字段值最好不要是0, ''

    func (r *UserRepository) UpdateById(id uint, t *model.User) (*model.User, error) {
        var ret = new(model.User)
        data := make(map[string]interface{})
        if t.Username != "" {
            data["username"] = t.Username
        }
        if t.Username != "" {
            data["real_name"] = t.RealName
        }
        if t.Username != "" {
            data["password"] = t.Password
        }
        // 在 GORM 中，Updates 方法默认会忽略零值。这是因为在 Go 中，每种类型都有一个零值，例如，int 的零值是 0，string 的零值是空字符串。因此，当你尝试更新一个值为 0 的字段时，GORM 会认为这是一个零值，并选择忽略它。
        // 通过结构体变量更新字段值, gorm库会忽略零值字段。就是字段值等于0, nil, “”, false这些值会被忽略掉，不会更新。如果想更新零值，可以使用map类型替代结构体。
        // 所以这里 Updates 不能使用 t 这个结构体
        data["status"] = t.Status
        // 以下两个字段暂时没有用到
        // data["level"] = t.Level
        // data["parent_id"] = t.ParentID
        _db, err := db.GetMysql()
        if err != nil {
            return nil, err
        }
        err = _db.Model(&ret).Where("id=?", id).Updates(data).Error
        if err == nil {
            ret.ID = id
        }
        return ret, err
    }

## 其他

// google authenticator
https://studygolang.com/articles/25748
https://www.jianshu.com/p/3c4832e3c6d5

// 逆向
https://zhuanlan.zhihu.com/p/26733683


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

## 清理无用包
    go mod tidy

## 升级依赖
    #!/usr/bin/env bash
    set -e
    go env -w GOPROXY=https://goproxy.cn,direct
    go get -u ./...
    go mod tidy
    go test ./...
    go build ./...
    echo "All deps upgraded & verified."