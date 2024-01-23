package logkit

const (
	/*
	 * 使用 context 传递
	 */
	TraceName     = "trace"
	SpanName      = "span"
	RequestIdName = "request_id"
	RouteName     = "route"
	UserIdName    = "user_id"

	/*
	 * 使用 WithFields 传递
	 */
	TypeName    = "type"    //日志类型
	RuntimeName = "runtime" //执行时间

	StartTimeName = "start_time" //脚本开始时间

	LogDefault = "default"

	LogHttp     = "http"
	LogRequest  = "request"  //请求
	LogResponse = "response" //返回

	LogRun      = "run"       //脚本
	LogRunStart = "run_start" //脚本运行开始
	LogRunEnd   = "run_end"   //脚本运行开始

	LogXorm          = "xorm"
	LogRedis         = "redis"
	LogMongo         = "mongo"
	LogKafka         = "kafka"
	LogKafkaRead     = "kafka_read"
	LogKafkaWrite    = "kafka_write"
	LogRabbitmq      = "rabbitmq"
	LogRabbitmqRead  = "rabbitmq_read"
	LogRabbitmqWrite = "rabbitmq_write"
	LogES            = "elastic"
)
