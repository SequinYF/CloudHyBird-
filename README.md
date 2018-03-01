# CloudHyBird-
基于go的云同步终端应用

## 技术

基于 Go 语言编写，利用 fsnotify 进行系统监控，设置 group 监控列表，当有事件发生时交由回调函数进行采集处理后发送至服务端。数据交换格式为 JSON，网络连接采用 Go Socket ，通过同步并发实现服务端文件的更新。

## 功能
实时监控目录下文件的变化，并同步到服务端。

#### 服务端：
1. 更新上传文件
2. 下载文件
3. 列出文件列表

#### 客户端
1. 设置监控目录
2. 文件监控
3. 将更新文件上传

## 结构

#### 传输端进程
  传输进程用于服务器与客户端数据的传输
#### 守护进程
  后台监控文件，触发事件后交由传输进程进行文件传输
#### 控制进程
  控制进程用于与用户进行交互，包括监控目录的设置，信息展示的窗口。
