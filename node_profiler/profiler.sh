#!/bin/bash

if [ $# -lt 1 ]; then 
    echo "Usage: $0 <file_path>"
    exit 1
fi

FILE_PATH=$1

NODE_ENV=production node --prof "$FILE_PATH"

LOG_FILE=$(ls isolate-*.log 2>/dev/null)

if [ -z "$LOG_FILE" ]; then 
    echo "No V8 log file found. Profiling might have failed."
    exit 1
fi

node --prof-process "$LOG_FILE" > processed_profile.txt

if [ -f processed_profile.txt ]; then 
    echo "Processed profile generated: processed_profile.txt"
    rm "$LOG_FILE"
else
    echo "Failed to generate processed profile."
    exit 1
fi
