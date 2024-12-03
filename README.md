qBittorrent Manager 是一个用于管理和监控 qBittorrent 种子的 Go 程序，用于登录 qBittorrent Web UI，获取种子信息，并根据上传率（Ratio）自动删除不符合条件的种子。

# 功能

主要功能：
- 登录到 qBittorrent Web UI
- 获取当前所有种子的状态
- 根据设定的上传率增长条件自动删除种子
- 本地保存种子记录，以便后续比对上传率
- 支持保护期设置，防止在保护期内删除种子

# 用法

运行程序时，可以使用以下命令行参数：

```bash
./qBittorrent-manager -url <qb_url> -username <username> -password <password> -recordFile <record_file> -ratioIncrease <ratio> -protectionPeriod <days> -try <true|false>
```


- `-url`：qBittorrent Web UI 的 URL，默认值为 `http://localhost:8080`
- `-username`：qBittorrent Web UI 的用户名，默认值为 `admin`
- `-password`：qBittorrent Web UI 的密码，默认值为 `adminadmin`
- `-recordFile`：用于保存种子记录的文件名，默认值为 `torrent-records.json`
- `-ratioIncrease`：每次检查时，上传率的最小增长值，默认值为 `0.5`
- `-protectionPeriod`：种子的保护期（以天为单位），默认值为 `7`，在此期间内即使上传率不佳也不会被删除
- `-try`：如果设置为 `true`，则仅显示将要删除的种子，但不会真的执行删除操作，方便测试

运行例子：

```bash
./qBittorrent-manager -url http://localhost:8080 -username admin -password adminadmin -recordFile torrent-records.json -ratioIncrease 0.5 -protectionPeriod 7 -try false
```

# 二次开发
确保开发环境中已安装 Go 语言，可以按照 [Go 官方文档](https://golang.org/doc/install) 中的说明进行安装

将此项目克隆到本地：

```bash
git clone <repository-url>
cd <repository-directory>
```

使用 vscode 进行开发，则 `.vscode/lanunch.json` 如下：

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Torrent Manager",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/torrent-manager/main.go",
            "console": "internalConsole",
            "args": [
                "--url=http://192.168.1.1:8080",
                "--username=admin",
                "--password=admin",
                "--try"
            ]
        }
    ]
}
```

修改代码后，使用以下命令编译程序：
   
```bash
go build -o qBittorrent-manager
```

# 其他

- 请确保有正确的 qBittorrent Web UI 登录凭据
- 本程序将会读取和写入 `torrent-records.json` 文件，请确保有适当的文件权限
- 在删除种子之前，可以使用 `-try` 参数来预览将要删除的种子

# 贡献

欢迎任何人对本项目提出改进建议或贡献代码！请提交 Pull Request 或在问题追踪器中报告问题

# 许可证

本项目采用 [MIT 许可证](LICENSE)
