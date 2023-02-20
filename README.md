# 哔哩哔哩抽奖小程序

一款支持对哔哩哔哩视频或动态评论进行抽奖的小程序，架构上使用 [React](https://reactjs.org) 和 [Flamego](https://flamego.dev) 实现前后端分离。

## 本地开发

1. 复制 `.env.example` 文件到 `.env`：
    ```sh
    cd frontend
    cp .env.example .env
    ```
1. 启动前端开发版本：
    ```sh
    pnpm start
    ```
1. 在另一个终端窗口启动后端：
    ```sh
    cd ../backend
    go run ./cmd
    ```

## 授权许可

本项目采用 MIT 开源授权许可证，完整的授权说明已放置在 [LICENSE](LICENSE) 文件中。
