# 哔哩哔哩抽奖小程序

一款支持对哔哩哔哩视频或动态评论进行抽奖的小程序，架构上使用 [React](https://reactjs.org) 和 [Flamego](https://flamego.dev) 实现前后端分离。

## 界面预览

<img width="500" src="https://user-images.githubusercontent.com/2946214/221393435-6fe7327e-0b92-4a93-9710-e540f24ef6ce.png">

<img width="500" src="https://user-images.githubusercontent.com/2946214/221393460-beaae548-a846-4365-80a0-e44c787c103e.png">

## 下载使用

1. 前往 [Releases](https://github.com/flamego-examples/bilibili-lottery/releases) 页面下载对应操作系统的二进制
1. 解压并启动程序
1. 浏览器访问 http://localhost:2830

## 本地开发

1. 复制 `.env.example` 文件到 `.env`：
    ```sh
    cd frontend
    cp .env.example .env
    ```
1. 启动前端开发版本：
    ```sh
    pnpm install && pnpm start
    ```
1. 在另一个终端窗口启动后端：
    ```sh
    cd ../backend
    go run ./cmd
    ```
1. 浏览器访问 http://localhost:3000

## 编译部署

1. 构建前端生产版本：
    ```sh
    cd frontend
    pnpm run build
    ```
1. 编译并启动后端：
    ```sh
    cd ../backend
    go build -o bilibili-lottery ./cmd
    FLAMEGO_ENV=production ./bilibili-lottery
    ```
1. 浏览器访问 http://localhost:2830

## 授权许可

本项目采用 MIT 开源授权许可证，完整的授权说明已放置在 [LICENSE](LICENSE) 文件中。
