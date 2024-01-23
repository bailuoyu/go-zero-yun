#!/bin/bash
# 需要安装nodejs
# 然后安装yapi插件 npm install -g yapi-cli
# 将此文件复制为yapi.sh使用
# token替换为对应的token
# merge参数说明见：
# https://hellosean1025.github.io/yapi/documents/data.html#%e9%80%9a%e8%bf%87%e5%91%bd%e4%bb%a4%e8%a1%8c%e5%af%bc%e5%85%a5%e6%8e%a5%e5%8f%a3%e6%95%b0%e6%8d%ae
token=yourtoekn
merge=mergin

echo "自动导入yapi文档..."

echo "{\"type\":\"swagger\",\"token\":\"${token}\",\"file\":\"./${1}.swagger.json\",\"merge\":\"${merge}\",\"server\":\"http://yapi.aiyi.live/\"}" > "yapi-import.json"
yapi import
