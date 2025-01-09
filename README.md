# 联通监控数据抓取

## 源码说明

- [py 源码](back)是妖友的（妖火论坛）
- 本源码与[ py 源码](back)有很多地方不同，非纯转换
- **使用本程序时，应当尊重开源协议，保留作者信息**
- 联通看家：https://we.wo.cn/web/smart-club-pc-v2/?clientId=1001000001
- 妖火论坛：https://yaohuo.me

## 使用说明

1. 从 [Releases](https://github.com/zgcwkjOpenProject/GO_UnicomMonitor/releases) 下载 **二进制程序** 和 **config.json** 文件
2. 修改配置文件 **config.json**，具体参考 [妖友源码说明](back)
3. 启动程序

> 配置文件说明

```
host -> 监听端口
sleep -> 检查间隔时长（单位：秒）
ffmpeg -> FFmpeg 配置
    exec -> FFmpeg 路径
    type -> cpu/gpu（留空时，调用 -c copy 命令）
    gpu -> GPU 解码参数
video -> 摄像头配置（数组）
    wsHost -> WebSocket 地址
    paramMsg -> 参数（不要 _paramStr_=）
    name -> 摄像头名称
    size -> 截断文件大小（单位：MB）
    count -> 保留文件数量
```

> ps: 没有 ffmpeg 时，删除整块 Json 即可

## 开发计划

- [x] 摄像头录像
- [ ] 摄像头录音
- [x] 支持多个摄像头录像
- [x] 使用 FFmpeg 转码，并能直接播放
- [ ] 设置摄像头录下存储上限
