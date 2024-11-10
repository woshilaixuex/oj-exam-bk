#!/bin/bash
TABLE_NAME=$1
DIR=$2
echo "开始创建库：$TABLE_NAME 的表：$DIR"
goctl model mysql datasource -url="delyr1c:SZLnMWh73LrMGPKz@tcp(185.106.176.190:3306)/csd_website_server" -table="$TABLE_NAME"  -dir="$DIR" -cache=false --style=goZero