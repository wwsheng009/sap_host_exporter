# Metrics

This document describes the metrics exposed by `sap_host_exporter`.

General notes:
- All the metrics are _namespaced_ with the prefix `sap`, which is followed by a _subsystem_, and both are in turn composed into a _Fully Qualified Name_ (FQN) of each metrics.
- All the metrics and labels _names_ are in snake_case, as conventional with Prometheus. That said, as much as we'll try to keep this consistent throughout the project, the label _values_ may not actually follow this convention, though (e.g. label is a hostname).

### Subsystems

These are the currently implemented subsystems.

1. [SAP Start Service](#sap-start-service)
2. [SAP Enqueue Server](#sap-enqueue-server)

### Appendix

1. [SAP State colors](#sap-state-colors)
2. [Common labels](#common-labels)


## SAP Start Service

The Start Service subsystem collects generic host-related metrics.

1. [`sap_start_service_instances`](#sap_start_service_instances) 
2. [`sap_start_service_processes`](#sap_start_service_processes) 

### `sap_start_service_instances`

The SAP instances in the context of the whole SAP system.

The value of this metric follows the [SAP state colors](#sap-state-colors) convention.   

#### Labels

- `start_priority`: the instance start priority
- `features`: a pipe-separated (`|`) list of features running in the instance  
   e.g. `ABAP|GATEWAY|ICMAN|IGS`  
   
#### Examples

```
# TYPE sap_start_service_instances gauge
sap_start_service_instances{features="MESSAGESERVER|ENQUE",hostname="sapha1as",instance_number="0",start_priority="1"} 2
sap_start_service_instances{features="ENQREP",hostname="sapha1er",instance_number="10",start_priority="0.5"} 2
sap_start_service_instances{features="ABAP|GATEWAY|ICMAN|IGS",hostname="sapha1pas",instance_number="1",start_priority="3"} 2
sap_start_service_instances{features="ABAP|GATEWAY|ICMAN|IGS",hostname="sapha1aas",instance_number="2",start_priority="3"} 2
```

### `sap_start_service_processes`

The processes started by the SAP Start Service.

The value of this metric follows the [SAP state colors](#sap-state-colors) convention.

#### Labels

- `name`: the name of the process.
- `pid`: the PID of the process.
- `status`: a textual status of the process, e.g. `Running`

The total number of lines for this metric will be the cardinality of `pid`.

#### Example

```
# TYPE sap_start_service_processes gauge
sap_start_service_processes{name="enserver",pid="30787",status="Stopping"} 3
sap_start_service_processes{name="msg_server",pid="30786",status="Running"} 2
```


## SAP Enqueue Server

The Enqueue Server (also known as the lock server) is the SAP system component that manages the lock table.

01. [`sap_enqueue_server_arguments_high`](#sap_enqueue_server_arguments_high)
02. [`sap_enqueue_server_arguments_max`](#sap_enqueue_server_arguments_max)
03. [`sap_enqueue_server_arguments_now`](#sap_enqueue_server_arguments_now)
04. [`sap_enqueue_server_arguments_state`](#sap_enqueue_server_arguments_state)
05. [`sap_enqueue_server_backup_requests`](#sap_enqueue_server_backup_requests)
06. [`sap_enqueue_server_cleanup_requests`](#sap_enqueue_server_cleanup_requests)
07. [`sap_enqueue_server_dequeue_all_requests`](#sap_enqueue_server_dequeue_all_requests)
08. [`sap_enqueue_server_dequeue_errors`](#sap_enqueue_server_dequeue_errors)
09. [`sap_enqueue_server_dequeue_requests`](#sap_enqueue_server_dequeue_requests)
10. [`sap_enqueue_server_enqueue_errors`](#sap_enqueue_server_enqueue_errors)
11. [`sap_enqueue_server_enqueue_rejects`](#sap_enqueue_server_enqueue_rejects)
12. [`sap_enqueue_server_enqueue_requests`](#sap_enqueue_server_enqueue_requests)
13. [`sap_enqueue_server_lock_time`](#sap_enqueue_server_lock_time)
14. [`sap_enqueue_server_lock_wait_time`](#sap_enqueue_server_lock_wait_time)
15. [`sap_enqueue_server_locks_high`](#sap_enqueue_server_locks_high)
16. [`sap_enqueue_server_locks_max`](#sap_enqueue_server_locks_max)
17. [`sap_enqueue_server_locks_now`](#sap_enqueue_server_locks_now)
18. [`sap_enqueue_server_locks_state`](#sap_enqueue_server_locks_state)
19. [`sap_enqueue_server_owner_high`](#sap_enqueue_server_owner_high)
20. [`sap_enqueue_server_owner_max`](#sap_enqueue_server_owner_max)
21. [`sap_enqueue_server_owner_now`](#sap_enqueue_server_owner_now)
22. [`sap_enqueue_server_owner_state`](#sap_enqueue_server_owner_state)
23. [`sap_enqueue_server_replication_state`](#sap_enqueue_server_replication_state)
24. [`sap_enqueue_server_reporting_requests`](#sap_enqueue_server_reporting_requests)
25. [`sap_enqueue_server_server_time`](#sap_enqueue_server_server_time)

### `sap_enqueue_server_arguments_high`

Peak number of different lock arguments that have been stored simultaneously in the lock table.

#### Example

```
# TYPE sap_enqueue_server_arguments_high counter
sap_enqueue_server_arguments_high 104
```

### `sap_enqueue_server_arguments_max`

Maximum number of lock arguments that can be stored in the lock table.

#### Example

```
# TYPE sap_enqueue_server_arguments_max counter
sap_enqueue_server_arguments_max 56415
```

### `sap_enqueue_server_arguments_now`

Current number of lock arguments in the lock table.

#### Example

```
# TYPE sap_enqueue_server_arguments_now gauge
sap_enqueue_server_arguments_now 0
```

### `sap_enqueue_server_arguments_state`

General state of lock arguments.

Refer to the [appendix](#sap-state-colors) to know more about the possible values of this metric.

#### Example

```
# TYPE sap_enqueue_server_arguments_state gauge
sap_enqueue_server_arguments_state 2
```

### `sap_enqueue_server_backup_requests`

Number of requests forwarded to the update process.

#### Example

```
# TYPE sap_enqueue_server_backup_requests counter
sap_enqueue_server_backup_requests 0
```

### `sap_enqueue_server_cleanup_requests`

Requests to release of all the locks of an application server.

#### Example

```
# TYPE sap_enqueue_server_cleanup_requests counter
sap_enqueue_server_cleanup_requests 4
```

### `sap_enqueue_server_dequeue_all_requests`

Requests to release of all the locks of an LUW.

#### Example

```
# TYPE sap_enqueue_server_dequeue_all_requests counter
sap_enqueue_server_dequeue_all_requests 150372
```

### `sap_enqueue_server_dequeue_errors`

Lock release errors.

#### Example

```
# TYPE sap_enqueue_server_dequeue_errors counter
sap_enqueue_server_dequeue_errors 0
```

### `sap_enqueue_server_dequeue_requests`

Lock release requests.

#### Example

```
# TYPE sap_enqueue_server_dequeue_requests counter
sap_enqueue_server_dequeue_requests 85213
```

### `sap_enqueue_server_enqueue_errors`

Lock acquisition errors

#### Example

```
# TYPE sap_enqueue_server_enqueue_errors counter
sap_enqueue_server_enqueue_errors 0
```

### `sap_enqueue_server_enqueue_rejects`

Rejected lock requests.

#### Example

```
# TYPE sap_enqueue_server_enqueue_rejects counter
sap_enqueue_server_enqueue_rejects 4
```

### `sap_enqueue_server_enqueue_requests`

Lock acquisition requests.

#### Example

```
# TYPE sap_enqueue_server_enqueue_requests counter
sap_enqueue_server_enqueue_requests 109408
```

### `sap_enqueue_server_lock_time`

Total time spent in lock operations.

#### Example

```
# TYPE sap_enqueue_server_lock_time counter
sap_enqueue_server_lock_time 174.574351
```

### `sap_enqueue_server_lock_wait_time`

Total waiting time of all work processes for accessing lock table.

#### Example

```
# TYPE sap_enqueue_server_lock_wait_time counter
sap_enqueue_server_lock_wait_time 0
```

### `sap_enqueue_server_locks_high`

Peak number of elementary locks that have been stored simultaneously in the lock table.

#### Example

```
# TYPE sap_enqueue_server_locks_high counter
sap_enqueue_server_locks_high 104
```

### `sap_enqueue_server_locks_max`

Maximum number of elementary locks that can be stored in the lock table.

#### Example

```
# TYPE sap_enqueue_server_locks_max gauge
sap_enqueue_server_locks_max 56415
```

### `sap_enqueue_server_locks_now`

Current number of elementary locks in the lock table.

#### Example

```
# TYPE sap_enqueue_server_locks_now gauge
sap_enqueue_server_locks_now 0
```

### `sap_enqueue_server_locks_state`

General state of elementary locks.

Refer to the [appendix](#sap-state-colors) to know more about the possible values of this metric.

#### Example

```
# TYPE sap_enqueue_server_locks_state gauge
sap_enqueue_server_locks_state 2
```

### `sap_enqueue_server_owner_high`

Peak number of lock owners that have been stored simultaneously in the lock table.

#### Example

```
# TYPE sap_enqueue_server_owner_high counter
sap_enqueue_server_owner_high 5
```

### `sap_enqueue_server_owner_max`

Maximum number of lock owner IDs that can be stored in the lock table.

#### Example

```
# TYPE sap_enqueue_server_owner_max gauge
sap_enqueue_server_owner_max 56415
```

### `sap_enqueue_server_owner_now`

Current number of lock owners in the lock table.

#### Example

```
# TYPE sap_enqueue_server_owner_now gauge
sap_enqueue_server_owner_now 0
```

### `sap_enqueue_server_owner_state`

General state of lock owners.

Refer to the [appendix](#sap-state-colors) to know more about the possible values of this metric.

#### Example

```
# TYPE sap_enqueue_server_owner_state gauge
sap_enqueue_server_owner_state 2
```

### `sap_enqueue_server_replication_state`

General state of lock server replication.

Refer to the [appendix](#sap-state-colors) to know more about the possible values of this metric.

#### Example

```
# TYPE sap_enqueue_server_replication_state gauge
sap_enqueue_server_replication_state 2
```

### `sap_enqueue_server_reporting_requests`

Number of reading operations on the lock table.

#### Example

```
# TYPE sap_enqueue_server_reporting_requests counter
sap_enqueue_server_reporting_requests 0
```

### `sap_enqueue_server_server_time`

Total time spent in lock operations by all processes in the enqueue server

#### Example

```
# TYPE sap_enqueue_server_server_time counter
sap_enqueue_server_server_time 0
```


## SAP AS Dispatcher

The Application Server Dispatcher is the component that manages the Work Process queues. We collect a set of queue stats for each type of Work Process queue.

1. [`sap_dispatcher_queue_now`](#sap_dispatcher_queue_now)
2. [`sap_dispatcher_queue_high`](#sap_dispatcher_queue_high)
3. [`sap_dispatcher_queue_max`](#sap_dispatcher_queue_max)
4. [`sap_dispatcher_queue_writes`](#sap_dispatcher_queue_writes)
5. [`sap_dispatcher_queue_reads`](#sap_dispatcher_queue_reads)

### `sap_dispatcher_queue_now`

Work process current queue length

#### Labels

- `type`: the type of the work queue.

#### Example

```
# TYPE sap_dispatcher_queue_now gauge
sap_dispatcher_queue_now{type="ABAP/BTC"} 0
sap_dispatcher_queue_now{type="ABAP/DIA"} 0
sap_dispatcher_queue_now{type="ABAP/ENQ"} 0
sap_dispatcher_queue_now{type="ABAP/NOWP"} 0
sap_dispatcher_queue_now{type="ABAP/SPO"} 0
sap_dispatcher_queue_now{type="ABAP/UP2"} 0
sap_dispatcher_queue_now{type="ABAP/UPD"} 0
sap_dispatcher_queue_now{type="ICM/Intern"} 0
```

### `sap_dispatcher_queue_high` 

Work process highest queue length

#### Labels

- `type`: the type of the work queue.

#### Example

```
# TYPE sap_dispatcher_queue_high counter
sap_dispatcher_queue_high{type="ABAP/BTC"} 2
sap_dispatcher_queue_high{type="ABAP/DIA"} 5
sap_dispatcher_queue_high{type="ABAP/ENQ"} 0
sap_dispatcher_queue_high{type="ABAP/NOWP"} 3
sap_dispatcher_queue_high{type="ABAP/SPO"} 1
sap_dispatcher_queue_high{type="ABAP/UP2"} 1
sap_dispatcher_queue_high{type="ABAP/UPD"} 2
sap_dispatcher_queue_high{type="ICM/Intern"} 1
```

### `sap_dispatcher_queue_max`

Work process maximum queue length

#### Labels

- `type`: the type of the work queue.

#### Example

```
# TYPE sap_dispatcher_queue_max gauge
sap_dispatcher_queue_max{type="ABAP/BTC"} 14000
sap_dispatcher_queue_max{type="ABAP/DIA"} 14000
sap_dispatcher_queue_max{type="ABAP/ENQ"} 14000
sap_dispatcher_queue_max{type="ABAP/NOWP"} 14000
sap_dispatcher_queue_max{type="ABAP/SPO"} 14000
sap_dispatcher_queue_max{type="ABAP/UP2"} 14000
sap_dispatcher_queue_max{type="ABAP/UPD"} 14000
sap_dispatcher_queue_max{type="ICM/Intern"} 6000
```

### `sap_dispatcher_queue_writes`

Work process queue writes

#### Labels

- `type`: the type of the work queue.

#### Example

```
# TYPE sap_dispatcher_queue_writes counter
sap_dispatcher_queue_writes{type="ABAP/BTC"} 11229
sap_dispatcher_queue_writes{type="ABAP/DIA"} 479801
sap_dispatcher_queue_writes{type="ABAP/ENQ"} 0
sap_dispatcher_queue_writes{type="ABAP/NOWP"} 267333
sap_dispatcher_queue_writes{type="ABAP/SPO"} 41171
sap_dispatcher_queue_writes{type="ABAP/UP2"} 3743
sap_dispatcher_queue_writes{type="ABAP/UPD"} 3746
sap_dispatcher_queue_writes{type="ICM/Intern"} 37426
```

### `sap_dispatcher_queue_reads`

Work process queue reads.

#### Labels

- `type`: the type of the work queue.

#### Example

```
# TYPE sap_dispatcher_queue_reads counter
sap_dispatcher_queue_reads{type="ABAP/BTC"} 11229
sap_dispatcher_queue_reads{type="ABAP/DIA"} 479801
sap_dispatcher_queue_reads{type="ABAP/ENQ"} 0
sap_dispatcher_queue_reads{type="ABAP/NOWP"} 267333
sap_dispatcher_queue_reads{type="ABAP/SPO"} 41171
sap_dispatcher_queue_reads{type="ABAP/UP2"} 3743
sap_dispatcher_queue_reads{type="ABAP/UPD"} 3746
sap_dispatcher_queue_reads{type="ICM/Intern"} 37426
```

### `sap_abap_work_processes_btc`

后台作业的工作时间。

#### Labels

- `SID`: 系统 ID
- `instance_name`: 实例名称
- `instance_number`: 实例编号
- `no`: 工作进程编号
- `pid`: 工作进程 PID
- `typ`: 工作进程类型
- `status`: 工作进程状态
- `start`: 工作进程是否已启动
- `action`: 工作进程动作
- `table`: 工作进程表
- `client`: 工作进程客户端
- `program`: 工作进程程序
- `user`: 工作进程用户
- `cpu`: 工作进程 CPU 时间
- `reason`: 工作进程原因
- `sem`: 工作进程信号量
- `err`: 工作进程错误
- `instance_hostname`: 实例主机名

#### Example

```
sap_abap_work_processes_btc{SID="EP1",action="",client="",cpu="1254",err="",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",no="36",pid="6542",program="",reason="",sem="",start="yes",status="Wait",table="",typ="BTC",user=""} 0
sap_abap_work_processes_btc{SID="EP1",action="",client="",cpu="1295",err="",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",no="38",pid="4782",program="",reason="",sem="",start="yes",status="Wait",table="",typ="BTC",user=""} 0
```

### `sap_abap_work_processes_dia`

对话进程工作时间。

#### Labels

#### Example

```
sap_abap_work_processes_dia{SID="EP1",action="",client="",cpu="10175",err="",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",no="13",pid="11825",program="",reason="",sem="",start="yes",status="Wait",table="",typ="DIA",user=""} 0

```

### `sap_abap_work_processes_spo`

打印进程工作时间。

#### Labels

#### Example

```
sap_abap_work_processes_spo{SID="EP1",action="",client="",cpu="139",err="",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",no="46",pid="4790",program="",reason="",sem="",start="yes",status="Wait",table="",typ="SPO",user=""} 0
```

### `sap_abap_work_processes_up2`

后台更新进工作时间。

#### Labels

#### Example

```
sap_abap_work_processes_up2{SID="EP1",action="",client="",cpu="14323",err="",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",no="47",pid="4791",program="",reason="",sem="",start="yes",status="Wait",table="",typ="UP2",user=""} 0
```

### `sap_abap_work_processes_upd`

后台更新进程工作时间。
#### Labels

#### Example

```
sap_abap_work_processes_upd{SID="EP1",action="",client="",cpu="20457",err="",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",no="30",pid="4774",program="",reason="",sem="",start="yes",status="Wait",table="",typ="UPD",user=""} 0
```

### `sap_abap_work_processes_wait_pct`

后台工作进程等待百分比，100% 表示所有工作进程都在等待。

#### Labels
- `typ`: 工作进程类型
- `SID`: 系统 ID
- `instance_name`: 实例名称
- `instance_number`: 实例编号
- `instance_hostname`: 实例主机名


#### Example

```
sap_abap_work_processes_wait_pct{SID="EP1",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",typ="ABAP/BTC"} 100
sap_abap_work_processes_wait_pct{SID="EP1",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",typ="ABAP/DIA"} 100
sap_abap_work_processes_wait_pct{SID="EP1",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",typ="ABAP/SPO"} 100
sap_abap_work_processes_wait_pct{SID="EP1",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",typ="ABAP/UP2"} 100
sap_abap_work_processes_wait_pct{SID="EP1",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",typ="ABAP/UPD"} 100
```

### `sap_abap_work_processes_simple`

后台工作进程信息，剔除了一些无用的信息，值是进程工作时间。

#### Labels

- `typ`: 工作进程类型
- `SID`: 系统 ID
- `instance_name`: 实例名称
- `instance_number`: 实例编号
- `instance_hostname`: 实例主机名
- `no`: 工作进程编号

#### Example

```
sap_abap_work_processes_simple{SID="EP1",instance_hostname="ERPPRD01",instance_name="D00",instance_number="0",no="1",typ="DIA"} 0
```

## Appendix

### SAP State colors

The value of `*_state` metrics is an integer status code that maps to the conventional SAP color-coded names as follows:
- `1`: `GRAY`.
- `2`: `GREEN`.
- `3`: `YELLOW`.
- `4`: `RED`.

### Common labels

The following labels are shared among all the metrics.

- `SID`: the SAP System ID.
- `instance_name`: the SAP instance name.
- `instance_hostname`: the SAP instance virtual hostname. Note, this may differ from the actual hostname of the instance host.
- `instance_number`: the SAP instance number,
