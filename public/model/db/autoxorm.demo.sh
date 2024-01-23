#!/bin/bash
# 如果没有安装xorm cmd
# 1.go get xorm.io/xorm
# 2.go get github.com/bailuoyu/xorm-cmd
# windows可以直接复制xorm.exe到go bin目录下
# 将此文件复制为autoxorm.sh到对应数据库目录使用
# sh ./autoxorm.sh {表名}
db_host=""
db_username=""
db_password=""
db_name=""

if [ ! -n "$1" ]; then
  echo "you have not input a comment!"
  exit 0
else
  tableFilterReg=$1
  echo "自动生成数据表结构体..."
fi

datasource="${db_username}:${db_password}@tcp(${db_host})/${db_name}?charset=utf8"
if [ "$tableFilterReg" == "all" ]; then
  echo "要生成全部数据表吗，确认信息开始下一步"
  read answer
  xorm reverse mysql ${datasource} ../xormcmd ./
else
  tableFilterReg="^${tableFilterReg}$"
  echo "表过滤${tableFilterReg}"
  xorm reverse mysql ${datasource} ../xormcmd ./ ${tableFilterReg}
fi
# shellcheck disable=SC2090
