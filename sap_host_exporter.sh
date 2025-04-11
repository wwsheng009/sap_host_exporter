#!/bin/bash

# 定义版本
VERSION="sap_host_exporter_amd64.tar.gz"

# 检查SAP实例并返回实例映射数组
check_sap_instances() {
    # 提取SAP中的进程信息
    # 创建临时文件存储netstat输出
    local netstat_output=$(netstat -tlnp)

    # 初始化实例映射数组
    declare -g -A instance_map
    declare -g -A instance_number_map

    # 检查32xx端口并提取实例信息
    while read -r line; do
        if [[ $line =~ 320([0-9])[[:space:]].*LISTEN[[:space:]]+([0-9]+)/(.+) ]]; then
            # 确保实例号是两位数，如果是个位数则在前面补0
            local original_instance_number=$(printf "%02d" "${BASH_REMATCH[1]}")
            local process_name=${BASH_REMATCH[3]}
            instance_map[$original_instance_number]=$process_name
            instance_number_map[$original_instance_number]=$original_instance_number
            echo "找到实例: $original_instance_number -> $process_name"
        fi
    done <<< "$netstat_output"

    # 检查5xx13端口
    if ! echo "$netstat_output" | grep -E ":[0-9]+13.*LISTEN.*sapstartsrv" > /dev/null; then
        echo "错误：未检测到sapstartsrv的5xx13端口监听"
        return 1
    fi

    # 检查是否找到任何实例
    if [ ${#instance_map[@]} -eq 0 ]; then
        echo "错误：未找到任何SAP实例监听端口"
        return 1
    fi

    # 将实例映射数组声明为全局变量
    declare -g -A instance_map instance_number_map
    return 0
}

install() {
    # 提示用户输入监控主机地址
    echo "请输入监控主机的IP地址或主机名:"
    read -r HOST
    if [ -z "$HOST" ]; then
        echo "错误：监控主机地址不能为空"
        return 1
    fi

    # 检查服务状态
    # 检查sapstartsrv进程是否在运行
    if ! pgrep -x "sapstartsrv" > /dev/null; then
        echo "错误：未检测到sapstartsrv进程在运行"
        echo "请确保SAP系统正在运行后再安装监控"
        return 1
    fi

    # 检查下载工具是否存在
    if command -v wget &> /dev/null; then
        DOWNLOADER="wget"
    elif command -v curl &> /dev/null; then
        DOWNLOADER="curl"
    else
        echo "错误：系统中既没有安装 wget 也没有安装 curl，请先安装其中一个工具"
        echo "Ubuntu/Debian: sudo apt-get install wget 或 sudo apt-get install curl"
        echo "CentOS/RHEL: sudo yum install wget 或 sudo yum install curl"
        return 1
    fi

    # 创建目录
    sudo mkdir -p /opt/exporter/sap_host_exporter

    # 使用可用的下载工具下载文件
    if [ "$DOWNLOADER" = "wget" ]; then
        wget http://${HOST}/n9e_install_files/${VERSION} -P /opt/exporter/sap_host_exporter
    elif [ "$DOWNLOADER" = "curl" ]; then
        curl -L http://${HOST}/n9e_install_files/${VERSION} -o /opt/exporter/sap_host_exporter/${VERSION}
    fi

    # 解压文件，没有子目录
    sudo tar xvf /opt/exporter/sap_host_exporter/${VERSION} -C /opt/exporter/sap_host_exporter

    # 进入目录
    cd /opt/exporter/sap_host_exporter

    # 检查SAP实例
    check_sap_instances || return 1

    # 复制配置文件并修改
    for instance_number in "${!instance_map[@]}"; do
        instance_id=${instance_map[$instance_number]}
        mapped_number=$(printf "%02d" "${instance_number_map[$instance_number]}")
        echo "实例号: ${instance_number}, 进程名: ${instance_id}"
        # 复制配置文件
        cp default.yaml "${instance_id}.yaml"
        # 修改配置文件
        sed -i "s|sap-control-uds: \"\"|sap-control-uds: \"/tmp/.sapstream5${mapped_number}13\"|" "${instance_id}.yaml"
        # 修改端口配置
        sed -i "s|port: \"9680\"|port: \"97${mapped_number}\"|" "${instance_id}.yaml"
    done

    # 修改服务配置文件
    sed -i 's|default.yaml|%i.yaml|' sap_host_exporter@.service

    # 复制服务文件到系统目录
    sudo cp sap_host_exporter@.service /etc/systemd/system/

    # 重新加载系统服务
    sudo systemctl daemon-reload

    # 启动服务
    for instance_id in "${instance_map[@]}"; do
        sudo systemctl enable sap_host_exporter@${instance_id}
        sudo systemctl start sap_host_exporter@${instance_id}
        echo "已启动服务: sap_host_exporter@${instance_id}"
    done

    return 0
}

# 检查服务状态
status() {
    # 检查SAP实例
    check_sap_instances || return 1

    echo "=== SAP Host Exporter 服务状态 ==="
    found_services=false
    # 打印实例映射信息
    echo "当前实例映射:"
    for instance_number in "${!instance_map[@]}"; do
        echo "实例号: ${instance_number}, 进程名: ${instance_map[$instance_number]}"
    done
    echo "-------------------"
    # 遍历所有实例检查服务状态
    for instance_id in "${instance_map[@]}"; do
        echo "应用实例: ${instance_id}"
        # 构建服务名称
        local service_name="sap_host_exporter@${instance_id}"
        found_services=true
        echo "服务名称: ${service_name}"
        
        # 获取服务状态信息
        local status_output=$(systemctl status "${service_name}" 2>/dev/null)
        local is_active=$(systemctl is-active "${service_name}" 2>/dev/null)
        local is_enabled=$(systemctl is-enabled "${service_name}" 2>/dev/null)
        
        echo "启用状态: ${is_enabled}"
        echo "运行状态: ${is_active}"
        
        # 获取进程信息
        local pid=$(systemctl show -p MainPID "${service_name}" 2>/dev/null | cut -d= -f2)
        if [ "${pid}" != "0" ] && [ -n "${pid}" ]; then
            echo "进程 PID: ${pid}"
        fi
        
        # 显示最近的日志
        echo "最近日志:"
        journalctl -u "${service_name}" -n 3 --no-pager 2>/dev/null || echo "无法获取日志信息"
        echo "-------------------"
    done

    if [ "$found_services" = false ]; then
        echo "未找到任何已安装的SAP Host Exporter服务"
        echo "请先运行安装命令来安装服务"
        return 1
    fi

    return 0
}

# 更新Categraf配置
update_categraf_config() {
    # 检查SAP实例
    check_sap_instances || return 1

    # 尝试定位categraf配置文件
    local categraf_base_dir="/opt/categraf"
    local config_file=""
    
    # 通过服务状态获取配置文件路径
    local service_output=$(systemctl status categraf 2>/dev/null)
    if [ $? -eq 0 ]; then
        # 提取配置文件路径
        local conf_path=$(echo "$service_output" | grep -o '/opt/categraf/categraf-v[0-9.]*-linux-amd64/conf')
        if [ -n "$conf_path" ]; then
            config_file="${conf_path}/input.prometheus/prometheus.toml"
        fi
    fi
    
    # 如果通过服务状态未找到，尝试直接查找文件
    if [ -z "$config_file" ] || [ ! -f "$config_file" ]; then
        # 查找任何版本的配置文件
        config_file=$(find "$categraf_base_dir" -name "prometheus.toml" -path "*/input.prometheus/*" 2>/dev/null | head -n 1)
    fi
    
    if [ -z "$config_file" ] || [ ! -f "$config_file" ]; then
        echo "错误：未找到Categraf配置文件"
        return 1
    fi
    
    echo "找到配置文件：$config_file"
    
    # 检查是否所有URL都已存在
    local all_urls_exist=true
    for instance_number in "${!instance_map[@]}"; do
        mapped_number=$(printf "%02d" "${instance_number_map[$instance_number]}")
        if ! grep -q "http://localhost:97${mapped_number}/metrics" "$config_file"; then
            all_urls_exist=false
            break
        fi
    done
    
    # 如果所有URL都存在，无需更新
    if [ "$all_urls_exist" = true ]; then
        echo "所有实例URL已存在于配置中，无需更新"
        return 0
    fi
    
    # 在文件末尾添加新的配置块
    echo >> "$config_file"
    echo "[[instances]]" >> "$config_file"
    echo "urls = [" >> "$config_file"
    for instance_number in "${!instance_map[@]}"; do
        mapped_number=$(printf "%02d" "${instance_number_map[$instance_number]}")
        echo "    \"http://localhost:97${mapped_number}/metrics\"" >> "$config_file"
    done
    echo "]" >> "$config_file"
    
    # 重启Categraf服务
    sudo systemctl restart categraf
    if [ $? -eq 0 ]; then
        echo "已成功更新Categraf配置并重启服务"
    else
        echo "错误：重启Categraf服务失败"
        return 1
    fi
    
    return 0
}

# 显示主菜单
show_menu() {
    echo "=== SAP Host Exporter 管理菜单 ==="
    echo "1. 检查SAP实例"
    echo "2. 安装SAP Host Exporter"
    echo "3. 检查服务状态"
    echo "4. 更新Categraf配置"
    echo "5. 退出"
    echo "请输入选项 [1-5]: "
}

# 主程序
while true; do
    show_menu
    read -r choice

    case $choice in
        1)
            echo "执行SAP实例检查..."
            check_sap_instances
            echo "按回车键继续..."
            read -r
            ;;
        2)
            echo "开始安装SAP Host Exporter..."
            install
            echo "按回车键继续..."
            read -r
            ;;
        3)
            echo "检查SAP服务状态..."
            status
            echo "按回车键继续..."
            read -r
            ;;
        4)
            echo "更新Categraf配置..."
            update_categraf_config
            echo "按回车键继续..."
            read -r
            ;;
        5)
            echo "退出程序"
            exit 0
            ;;
        *)
            echo "无效的选项，请重新选择"
            sleep 1
            ;;
    esac
    echo
done
