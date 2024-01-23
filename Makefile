GOBIN 	:= go
PWD 	:= $(shell pwd)
GIT := git

# model相关配置
MODELDIR := ${PWD}/public/model

### help:		查看所有makefile规则
.PHONY: help
help:
	@echo Makefile rules:
	@echo
	@grep -E '^### [-A-Za-z0-9_]+:' Makefile | sed 's/###/   /'

### build:		构建服务(e.g. make build app=greet)
.PHONY: build-rest
build-rest:
	@if [[ "${app}" == "" ]]; then \
		echo "app is empty, you should set!"; \
		exit 1; \
	fi;
	${GOBIN} mod tidy; \
	CGO_ENABLED=0 ${GOBIN} build -o ${PWD}/app/${app}/rest/bin/go-${app}.rest "app/${app}/rest/${app}.go";

.PHONY: build-cmd
build-cmd:
	@if [[ "${app}" == "" ]]; then \
		echo "app is empty, you should set!"; \
		exit 1; \
	fi;
	${GOBIN} mod tidy; \
	CGO_ENABLED=0 ${GOBIN} build -o ${PWD}/app/${app}/cmd/bin/go-${app}.cmd "app/${app}/cmd/${app}.go";

### swag:		构建服务(e.g. make swag app=greet mod=user)
.PHONY: swag
swag:
	@if [[ "${app}" == "" ]]; then \
		echo "app is empty, you should set!"; \
		exit 1; \
	elif [[ "${mod}" == "" ]]; then \
		echo "mod is empty, you should set!"; \
		exit 2; \
	fi;
	@cd ${PWD}/app/${app}/ && goctl api plugin -plugin goctl-swagger="swagger -filename ${mod}.swagger.json" -api ./api/${mod}.api -dir ./swagger

### yapi:		swagger导入Yapi(e.g. make yapi app=greet)
.PHONY: yapi
yapi:
	@if [[ "${app}" == "" ]]; then \
		echo "app is empty, you should set!"; \
		exit 1; \
	elif [[ "${mod}" == "" ]]; then \
		echo "mod is empty, you should set!"; \
		exit 2; \
	fi;
	@cd ${PWD}/app/${app}/swagger;	\
	sh yapi.sh ${mod}

### db-core:		生成model文件(e.g. make db-core table=demo)
.PHONY: db-core
db-core:
	@if [[ "${table}" == "" ]]; then \
		echo "table is empty, you should set!"; \
		exit 1; \
	fi;
	@cd ${MODELDIR}/db/core; \
	sh autoxorm.sh ${table}

### db-admin:		生成model文件(e.g. make db-admin table=demo)
.PHONY: db-admin
db-admin:
	@if [[ "${table}" == "" ]]; then \
		echo "table is empty, you should set!"; \
		exit 1; \
	fi;
	@cd ${MODELDIR}/db/admin; \
	sh autoxorm.sh ${table}

### db-data:		生成model文件(e.g. make db-data table=demo)
.PHONY: db-data
db-data:
	@if [[ "${table}" == "" ]]; then \
		echo "table is empty, you should set!"; \
		exit 1; \
	fi;
	@cd ${MODELDIR}/db/data; \
	sh autoxorm.sh ${table}

### mon-core:		生成MongoDB model文件(e.g. make mon-core type=Demo)
.PHONY: mon-core
mon-core:
	@if [[ "${type}" == "" ]]; then \
		echo "col is empty, you should set!"; \
		exit 1; \
	fi;
	@cd ${MODELDIR}/mongo/core; \
	goctl model mongo -t ${type} -e true

### mon-admin:		生成MongoDB model文件(e.g. make mon-admin type=Demo)
.PHONY: mon-admin
mon-admin:
	@if [[ "${type}" == "" ]]; then \
		echo "col is empty, you should set!"; \
		exit 1; \
	fi;
	@cd ${MODELDIR}/mongo/admin; \
	goctl model mongo -t ${type} -e true