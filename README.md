# 联通监控数据抓取

## 来源

[py 源码](back)是妖友的（妖火论坛）。本人转换成 golang 更方便使用

## 入口

联通看家：https://we.wo.cn/web/smart-club-pc-v2/?clientId=1001000001

## 说明

1. 从 [Releases](https://github.com/zgcwkjOpenProject/GO_UnicomMonitor/releases) 下载 **二进制程序** 和 **config.json** 文件
2. 修改配置文件 **config.json**，具体参考 [妖友源码说明](back)。
3. 启动程序，会立刻抓取数据。
4. 视频文件缺少文件头，可以用 ffmpeg 转换或补上缺少文件头。

## 开发

- [x] 摄像头录像
- [x] 支持多个摄像头录像
- [x] 使用 FFmpeg 转码，并能直接播放
- [ ] 设置摄像头录下存储上限
- [ ] 摄像头录音

## 配置说明

```
host -> 监听端口
sleep -> 检查间隔时长（单位：秒）
ffmpeg -> FFmpeg配置
    exec -> FFmpeg路径
    type -> cpu/gpu（可以留空，留空时纯复制）
    gpu -> GPU参数
video -> 摄像头配置
    wsHost -> WebSocket 地址
    paramMsg -> 参数（不要 _paramStr_=）
    name -> 摄像头名称
    size -> 截断文件大小（单位：MB）
    count -> 保留文件数量
```

ps: 没有ffmpeg时，把**整块 Json 删除**即可
