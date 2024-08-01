#!/bin/sh

CMD="./$2 -conf $3"

# 打印传递的参数以进行调试
echo "Script parameters:"
echo "CMD: $CMD"
echo "Parameter 1 (command): $1"
echo "Parameter 2 (process name): $2"
echo "Parameter 3 (config file): $3"

start(){
    local process_name="$1"
    local config_file="$2"
    local cmd="./$process_name -conf $config_file"

    echo "starting $cmd..."
    ps -ef | grep conf | grep "$process_name" | grep "$config_file" | grep -v grep | grep -v /bin/bash | grep -v $$
    num=$(ps -ef | grep conf | grep "$process_name" | grep "$config_file" | grep -v grep | grep -v /bin/bash | grep -v $$ | wc -l)
    if [ $num -eq 0 ]
    then
        nohup $cmd > "$process_name.log" 2>&1 &
        if [ $? -ne 0 ]
        then
            echo "start failed, please check the log!"
            exit $?
        else
            echo "start success"
        fi
    else
        echo "$cmd is already running"
    fi
}

stop(){
    local process_name="$1"
    local config_file="$2"
    local cmd="./$process_name -conf $config_file"

    echo "stopping $cmd..."
    echo "Matching process with: '$process_name' and config: '$config_file'"
    echo "Executing command: ps -ef | grep \"$process_name\" | grep \"$config_file\" | grep conf | grep -v grep | grep -v $$ | awk '{print \$2}'"
    
    # 打印每个步骤的输出
    ps -ef | grep "$process_name" | grep "$config_file" | grep conf | grep -v grep | grep -v $$ | awk '{print $2}'
    PIDS=$(ps -ef | grep "$process_name" | grep "$config_file" | grep conf | grep -v grep | grep -v $$ | awk '{print $2}')
    echo "Query result: $PIDS"
    
    if [ -z "$PIDS" ]; then
        echo "stop failed, maybe $cmd isn't running"
        exit 1
    else
        # 遍历每个PID并终止
        for PID in $PIDS; do
            echo "Killing process with PID: $PID"
            kill -9 "$PID"
            if [ $? -eq 0 ]; then
                echo "Stopped process with PID $PID"
            else
                echo "Failed to stop process with PID $PID"
            fi
        done
        echo "stop success"
    fi
}

restart(){
    stop "$1" "$2"
    start "$1" "$2"
}

status(){
    local process_name="$1"
    local config_file="$2"
    local cmd="./$process_name -conf $config_file"
    ps -ef | grep "$process_name" | grep "$config_file" | grep conf | grep -v grep | grep -v $$
    num=$(ps -ef | grep "$process_name" | grep "$config_file" | grep conf | grep -v grep | grep -v $$ | wc -l)
    if [ $num -eq 0 ]
    then
        echo "$cmd isn't running"
    else
        echo "$cmd is running"
    fi
}

case $1 in
    start) start "$2" "$3" ;;
    stop) stop "$2" "$3" ;;
    restart) restart "$2" "$3" ;;
    status) status "$2" "$3" ;;
    *) echo "Usage: $0 {start|stop|restart|status} <process_name> <config_file>" ;;
esac

exit 0

