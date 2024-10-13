#!/bin/bash


PORTS=(3000 8080 1313 1517)
SERVER_JS_PATH="./server.js"

for PORT in "${PORTS[@]}"
do
    echo "Starting server.js on port $PORT"
    PORT=$PORT nohup node $SERVER_JS_PATH > logs/server_$PORT.log 2>&1 &
done

echo "All servers started."
