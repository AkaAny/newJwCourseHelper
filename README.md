# HDU 新正方教务抢课 Plus

## 已有功能

- 自动 cas 登录
- 关键词查课
- 自动抢课
- 同时抢多课
- 支持主修课程、通识选修课程和体育课程

## Feat

- 请求限流
- 多用户支持
- 支持公选课、专业课、体育课


## 教程
### Release
1. 进入[下载页](https://github.com/wujunyi792/newJwCourseHelper/releases)选择对应二进制压缩包下载并解压到文件夹
2. 将主目录的 `config.json.example` 复制一份，重命名为 `config.json`，并修改其中的配置项
3. 配置项修改按照下面配置文件的说明进行修改
4. 下面介绍一下主要的参数(就四个比较重要的)
5. 前两个是user里的staffId和password(分别代表你的**智慧杭电**用户名和智慧杭电密码)
6. 第三个是target，里面可以写入多个课程，每个结构体中有name和type，name是要搜索的课程，这里建议填具体的教学班；type，这个是课程的类型，可以是主修课程、通识选修课、体育，分别对应填入0、1、2
7. 第四个是interval，这个是抢课的间隔时间，单位是秒，建议设置不要低于5秒，因为抢课的时候服务器会有一定的延迟，如果设置的太小，可能会导致抢课失败 
8. 运行可执行程序

### Docker
1. Clone this repo
```shell
git clone https://github.com/wujunyi792/newJwCourseHelper.git
```

将主目录的 `config.json.example` 复制一份，重命名为 `config.json`，并修改其中的配置项

Btw, 使用`HELPER_CONFIG_PATH`环境变量来指定配置文件路径也是被支持的

```shell
docker build -t newjw_course_helper .
docker run -it --volume "$PWD":/usr/src/app newjw_course_helper
```

## 成功实例
![运行截图](./doc/run.jpg)

## 配置文件(以下配置文件仅作为示例)
```
[
  {
    "user": {
      "staffId": "张三", // 智慧杭电用户名 // 通常为学号
      "password": "123456", // 智慧杭电密码
      "auto_auth": true // 启动自动重新登录 (当遇到 Req 302 Session过期的重定向后 再次登录)
    },
    "target": [{
      "name": "(2022-2023-2)-C0392006-01",//课程名称 || 课号 || 教学班
      "type": "" //0主修课程 1通识选修课 2体育
    },{
      "name": "",//课程名称 || 课号 || 教学班
      "type": "" //0主修课程 1通识选修课 2体育
    }], // 搜索关键词，一般使用教学班号，例如“(2022-2023-1)-B05025235-2”
    "errTag": [], // 屏蔽课号
    "rate": 500, // 几ms能发送一次请求
    "ua": "", // 浏览器 user-agent，默认 resty
    "interval": 10 // 执行任务间隔（抢课周期）单位秒
  }
]
```

## TODO
- [x] 体育课
- [ ] 更优雅的错误处理
- [ ] web 控制面板
- [x] 更优雅的启动方式
- [ ] 优化抢课逻辑
- [x] 过期重登