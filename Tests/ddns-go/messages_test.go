package util

import (
 "bytes"
 "log"
 "testing"
)

// Captured message keys from ddns-go v6.17.1, including all DNS providers.
func TestUnchangedIPLogFilter(t *testing.T) {
 oldWriter := log.Writer()
 oldFlags := log.Flags()
 defer log.SetOutput(oldWriter)
 defer log.SetFlags(oldFlags)
 defer InitLogLang("en")
 log.SetFlags(0)
 suppressed := map[string]bool{
  "IP %s 没有变化，域名 %s": true,
  "IPv4未改变, 将等待 %d 次后与DNS服务商进行比对": true,
  "IPv6未改变, 将等待 %d 次后与DNS服务商进行比对": true,
  "你的IP %s 没有变化, EdgeOne 源站组 %s": true,
  "你的IP %s 没有变化, 域名 %s": true,
  "你的IPv4未变化, 未触发 %s 请求": true,
  "你的IPv6未变化, 未触发 %s 请求": true,
 }
 keys := []string{
  "%q 帐号密码不正确",
  "%q 登录成功",
  "%q 被禁止从公网访问",
  "%q 配置文件为空, 超过3小时禁止从公网访问",
  "%s 后重试...",
  "Callback的URL不正确",
  "Callback调用失败, 异常信息: %v",
  "Callback调用成功, 域名: %s, IP: %s, 返回数据: %s",
  "IP %s 没有变化，域名 %s",
  "IPv4未改变, 将等待 %d 次后与DNS服务商进行比对",
  "IPv6匹配表达式 %s 不正确! 最小从1开始",
  "IPv6将使用正则表达式 %s 进行匹配",
  "IPv6未改变, 将等待 %d 次后与DNS服务商进行比对",
  "Namecheap 不支持更新 IPv6",
  "Webhook Header不正确: %s",
  "Webhook中的 RequestBody JSON 无效",
  "Webhook调用失败! 异常信息：%s",
  "Webhook调用成功! 返回数据：%s",
  "Webhook配置中的URL不正确",
  "ddns-go 服务卸载失败, 异常信息: %s",
  "ddns-go 服务卸载成功",
  "ddns-go 服务已安装, 无需再次安装",
  "ddns-go 服务未安装, 请先安装服务",
  "dynadot仅支持单域名配置，多个域名请添加更多配置",
  "从网卡中获得IPv4失败! 网卡名: %s",
  "从网卡中获得IPv6失败! 网卡名: %s",
  "从网卡获得IPv4失败",
  "从网卡获得IPv6失败",
  "你的IP %s 没有变化, EdgeOne 源站组 %s",
  "你的IP %s 没有变化, 域名 %s",
  "你的IPv4未变化, 未触发 %s 请求",
  "你的IPv6未变化, 未触发 %s 请求",
  "匹配成功! 匹配到地址: %s",
  "可使用 .\\\\ddns-go.exe -s install 安装服务运行",
  "可使用 sudo ./ddns-go -s install 安装服务运行",
  "启动 ddns-go 服务成功",
  "在DNS服务商中未找到域名: %s",
  "在DNS服务商中未找到根域名: %s",
  "域名 %s 解析未找到，且因添加了参数 %s=%s 导致无法创建。本次更新已被忽略",
  "域名: %s 不正确",
  "域名: %s 解析失败",
  "安装 ddns-go 服务失败, 异常信息: %s",
  "安装 ddns-go 服务成功! 请打开浏览器并进行配置",
  "将不会触发Webhook, 仅在第 3 次失败时触发一次Webhook, 当前失败次数：%d",
  "异常信息: %s",
  "异常信息: %v",
  "数据解析失败, 请刷新页面重试",
  "整理 EdgeOne 源站组记录失败! %s",
  "新增域名解析 %s 失败! 异常信息: %s",
  "新增域名解析 %s 失败! 异常信息: %v",
  "新增域名解析 %s 成功! IP: %s",
  "更新 EdgeOne 源站组 %s 失败! 异常信息: %s",
  "更新 EdgeOne 源站组 %s 成功! IP: %s",
  "更新域名解析 %s 失败! 异常信息: %s",
  "更新域名解析 %s 失败! 异常信息: %v",
  "更新域名解析 %s 成功! IP: %s",
  "未找到第 %d 个IPv6地址! 将使用第一个IPv6地址",
  "未能获取IPv4地址, 将不会更新",
  "未能获取IPv6地址, 将不会更新",
  "本机DNS异常! 将默认使用 %s, 可参考文档通过 -dns 自定义 DNS 服务器",
  "查询 EdgeOne 源站组信息发生异常! %s",
  "查询 EdgeOne 站点信息发生异常! %s",
  "查询域名 %s 信息发生异常! %v",
  "查询域名信息发生异常! %s",
  "查询域名信息发生异常！ %s",
  "没有匹配到任何一个IPv6地址, 将使用第一个地址",
  "用户名 %s 的密码已重置成功! 请重启ddns-go",
  "监听 %s",
  "第 %s 个配置未填写域名",
  "等待网络连接: %s",
  "绑定网卡失败, 将使用默认网卡. 网卡: %s, 网络: %s, 错误: 本地IP无效: %s",
  "绑定网卡失败, 将使用默认网卡. 网卡: %s, 错误: %v",
  "网络已连接",
  "获取%s结果失败! 命令: %s, 标准输出: %q",
  "获取%s结果失败! 未能成功执行命令：%s, 错误：%q, 退出状态码：%s",
  "获取IPv4结果失败! 接口: %s ,返回值: %s",
  "获取IPv6结果失败! 接口: %s ,返回值: %s",
  "设置 SO_BINDTODEVICE 失败, 回退为仅 LocalAddr 绑定. 网卡: %s, 错误: %v",
  "请输入Webhook的URL",
  "通过接口获取IPv4失败! 接口地址: %s",
  "通过接口获取IPv6失败! 接口地址: %s",
  "配置文件 %s 不存在, 可通过-c指定配置文件",
  "配置文件已保存在: %s",
  "重启 ddns-go 服务成功",
  "unknown future message: IP没有变化 but request failed",
 }
 for _, lang := range []string{"en", "zh"} {
  InitLogLang(lang)
  for _, key := range keys {
   var output bytes.Buffer
   log.SetOutput(&output)
   Log(key)
   if suppressed[key] {
    if output.Len() != 0 { t.Errorf("%s: should suppress %q", lang, key) }
   } else if want := LogStr(key) + "\n"; output.String() != want {
    t.Errorf("%s: changed other log %q", lang, key)
   }
  }
 }
}
