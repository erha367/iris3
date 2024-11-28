#!/bin/bash
# 递归处理文件夹中的图片
function traverse_directory {
    local dir="$1"
    local file
    for file in "$dir"/*; do
        if [[ -f "$file" ]]; then
             iris3 -path="$file"
             echo "$file"
        elif [[ -d "$file" ]]; then
            traverse_directory "$file"
        fi
    done
}

traverse_directory "/Users/yang/Downloads/tmp"
