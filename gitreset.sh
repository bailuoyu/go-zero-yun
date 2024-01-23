#!/bin/bash
branch=$(git rev-parse --abbrev-ref HEAD)
echo "Now Origin Branch: ${branch}"
git reset --hard origin/"${branch}"
echo "确认分支信息开始下一步"
# shellcheck disable=SC2162
read answer
git pull
