# Share-Text

**Share-Text** 是一款基于 Web 的共享剪切板内容服务，旨在为相同公网 IP 或局域网内的设备提供文本内容共享功能。支持加密共享，阅后即焚，且所有共享内容将在 10 分钟后自动删除。该项目使用 Go 开发，数据存储使用
SQLite 数据库。

## 功能特点

- **公网和局域网共享**：支持相同公网 IP 和相同局域网 IP 的设备共享剪切板内容。
- **加密共享**：支持加密文本内容，内容通过加密后不受 IP 限制，且需要密码才能访问。
- **阅后即焚**：加密内容在成功读取后即被销毁，确保隐私。
- **自动过期**：所有共享的文本内容将在 10 分钟后自动删除，增强安全性。

## 安装与运行

- 源码：

    ```bash
    git clone https://github.com/systemmin/share-text.git
    cd share-text
    go run main.go
    ```

- Window：

  [releases](https://github.com/systemmin/share-text/releases) 页面下载，双击 exe 文件或命令行启动。

    ```bash
    ./share-text.exe
    ```

- Linux：

  [releases](https://github.com/systemmin/share-text/releases) 页面下载，需要添加可执行权限。

   ```bash
   chmoc +x share-text
   ./share-text
   # 后台运行
   nohup ./share-text &
   ```

**备注**：其他平台自行编译

## 命令

```bash
share-text -h
Usage of share-text:
  -p string
        指定端口 0-65535 (default "9999")
```

## 示例

[演示地址](https://dtking.cn/share/)


![share-text](https://raw.githubusercontent.com/systemmin/share-text/refs/heads/master/images/share-text.png)