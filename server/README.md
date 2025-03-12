打包
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app_amd64 main.go

mac m系列
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o app_arm64 main.go

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ go build -tags=musl -o app_linux_amd64 main.go
交叉编译教程
https://www.cnblogs.com/informatics/p/17682616.html

CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=/Users/huangyanglv/Work/WB/x86_64-unknown-linux-gnu/bin/x86_64-unknown-linux-gnu-gcc CXX=/Users/huangyanglv/Work/WB/x86_64-unknown-linux-gnu/bin/x86_64-unknown-linux-gnu-g++ go build -o app_linux_amd64 main.go

[
    {
        "name": "管理层薪资支出",
        "count":2000
    },
    {
        "name": "推荐人奖金支出",
        "count":1000
    },
    {
        "name": "新人奖金支出",
        "count":1000
    },
    {
        "name": "餐补报销支出",
        "count":1000
    },
    {
        "name": "晋升奖金支出",
        "count":1000
    },
    {
        "name": "扶持奖金支出",
        "count":1000
    },
    {
        "name": "USDT充值返点支出",
        "count":1000
    },
    {
        "name": "活动办理福利支出",
        "count":1000
    },
    {
        "name": "实物订购支出",
        "count":1000
    },
    {
        "name": "余创收佣金",
        "count":1000
    }
]