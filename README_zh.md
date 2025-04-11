# SAP 主机导出器

这是一个定制的 Prometheus 导出器，用于监控 SAP 系统（即 SAP NetWeaver 应用程序）。

[![导出器 CI](https://github.com/SUSE/sap_host_exporter/workflows/Exporter%20CI/badge.svg)](https://github.com/SUSE/sap_host_exporter/actions?query=workflow%3A%22Exporter+CI%22)

## 目录

1. [功能特点](#功能特点)
2. [安装](#安装)
3. [使用方法](#使用方法)
   1. [配置](#配置)
   2. [指标](#指标)
   3. [systemd 集成](#systemd-集成)
5. [贡献](#贡献)
   1. [设计](doc/design.md)
   2. [开发](doc/development.md)
6. [许可证](#许可证)

## 功能特点

该导出器是一个无状态的 HTTP 端点。在每个 HTTP 请求中，它通过 SAPControl Web 接口从 SAP 系统拉取运行时数据。

导出的数据包括：

- 启动服务进程
- 排队服务器统计信息
- AS 调度器工作进程队列统计信息

## 安装

本项目可以通过多种方式安装，包括但不限于：

1. [手动克隆和构建](#手动克隆和构建)
2. [Go](#go)
3. [RPM](#rpm)

### 手动克隆和构建

```shell
git clone https://github.com/wwsheng009/sap_host_exporter
cd sap_host_exporter
make build
make install
```

### Go

```shell
go install github.com/wwsheng009/sap_host_exporter@latest
```

## 使用方法

你可以按如下方式运行导出器：

```shell
./sap_host_exporter --sap-control-url $SAP_HOST:$SAP_CONTROL_PORT
```

虽然不是严格要求，但建议在目标 SAP 实例主机上本地运行导出器，并通过 Unix 域套接字连接到 SAPControl Web 服务：

```shell
./sap_host_exporter --sap-control-uds /tmp/.sapstream50013
```

有关 SAPControl 的更多详细信息，请参阅 [SAP 官方文档](https://www.sap.com/documents/2016/09/0a40e60d-8b7c-0010-82c7-eda71af511fa.html) 以正确连接到 SAPControl 服务。

导出器将在默认端口 `9680` 的 `/metrics` 路径下暴露指标。

### 配置

运行时参数可以通过 CLI 标志或配置文件进行配置，这两种方式都是完全可选的。

更多详细信息，请参阅 `sap_host_exporter --help` 的帮助信息。

**注意**：
内置默认值是为最新版本的 SUSE Linux Enterprise 和 openSUSE 量身定制的。

程序将按顺序在当前工作目录、`$HOME/.config`、`/etc` 和 `/usr/etc` 中扫描名为 `sap_host_exporter.(yaml|json|toml)` 的文件。
第一个匹配项具有优先权，而 CLI 标志的优先级高于配置文件。

请参阅 [示例 YAML 配置](doc/sap_host_exporter.yaml) 了解更多详细信息。

### 指标

导出器不会导出任何无法收集的指标，但由于它不关心被监控目标中存在哪些子系统，因此无法收集指标不被视为严重的失败条件。
相反，如果某些收集器无法注册或执行收集周期，日志中将打印出软警告。

有关所有导出指标的详细信息，请参阅 [doc/metrics.md](doc/metrics.md)。

### systemd 集成

RPM 包中提供了一个 [systemd 单元文件](packaging/obs/prometheus-sap_host_exporter.spec)。你可以按常规方式启用和启动它：

```
systemctl --now enable prometheus-sap_host_exporter
```

## 贡献

我们非常欢迎提交拉取请求！

在贡献之前，我们建议查看 [设计文档](doc/design.md) 和 [开发说明](doc/development.md)。

## 许可证

版权所有 2020-2025 SUSE LLC

根据 Apache License 2.0 版本（以下简称"许可证"）获得许可；
除非符合许可证的规定，否则你不得使用此代码仓库。
你可以在以下位置获取许可证副本：

   <https://www.apache.org/licenses/LICENSE-2.0>

除非适用法律要求或书面同意，根据许可证分发的软件是基于"按原样"提供的，
不附带任何明示或暗示的担保或条件。
有关许可证下的特定语言管理权限和限制，请参阅许可证。