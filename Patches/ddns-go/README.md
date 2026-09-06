# ddns-go 未变化日志过滤

当前验证版本：sirpdboy/luci-app-ddns-go 的 ddns-go 6.17.1，源码来自
jeessy2/ddns-go v6.17.1（0fc5cdcfd0b216d80f5c6001274e9b2eaa13a653）。

`WRT-CORE.yml` 依次执行 feeds 更新、`Packages.sh`、`Handles.sh`。
`Handles.sh` 每次把本目录补丁复制到
`wrt/package/luci-app-ddns-go/ddns-go/patches/`，由 OpenWrt 默认的
Build/Prepare 在解包后应用。不会修改下载归档、哈希或上游仓库。
重复执行会覆盖同名补丁，保留软件包的其他补丁。

补丁仅在 `util.Log` 中精确匹配七个原始消息键并提前返回，覆盖 IPv4/IPv6
缓存等待、域名 IP 未变化、因 IP 未变化而跳过请求以及 EdgeOne 源站组未变化。
在翻译前匹配，因此同时适用于中文和英文。没有使用关键词或正则模糊过滤。
其他消息仍调用原来的 `log.Println(LogStr(...))`，包括成功、失败和启动错误；
`LogStr`、DNS 更新逻辑、stdout/stderr 和 procd 配置均保持原样。
这些消息在日志源头被移除，因此其他接收该日志输出的地方也不再显示它们。

软件包或补丁缺失会使 Handles.sh 失败；上游改动导致补丁不兼容时，OpenWrt
的源码准备阶段会失败，需要重新检查补丁，不会无提示地继续生成未过滤固件。
上游新增或改名的消息不会自动被过滤，升级后应复核新增日志键。
若在已有构建目录手动更新补丁，先执行 `make package/ddns-go/clean` 再编译。

## 验证

已验证补丁可套用于 6.17.1，脚本语法、重复安装、软件包目录重建后的安装，
以及缺失软件包时的失败处理。

`Tests/ddns-go/messages_test.go` 收录该版本 84 个日志键，并在中英文下检查：
七类目标消息不输出，其他日志输出与原 `LogStr` 完全一致；未知消息也保留。
可将测试文件临时复制到已应用补丁的 ddns-go 源码 `util/` 下，运行
`go test ./util -run TestUnchangedIPLogFilter -v`，随后移除测试文件。

本地只验证日志模块，没有进行整套 OpenWrt 固件交叉编译或路由器实机验证。
