#!/bin/bash

PORTS=(3000 8080 1313 1517)


for PORT in "${PORTS[@]}"
do

    PID=$(lsof -t -i TCP:$PORT)
    echo $PORT - $PID 
    if [ -z "$PID" ]; then
        echo "No process found on port $PORT"
    else
        echo "Killing process on port $PORT with PID $PID"
        kill -9 $PID > logs/kill_server_$PORT.log 2>&1 &
    fi
done

echo "All specified servers stopped."
