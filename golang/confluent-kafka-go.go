package tools

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// kafak 所有的相关参数，可以按需选取
type KafkaConfig struct {
	// Common configuration
	BootstrapServers                 string `kafka:"bootstrap.servers"`
	ClientID                         string `kafka:"client.id"`
	RackID                           string `kafka:"rack.id"`
	APIVersionRequest                bool   `kafka:"api.version.request"`
	APIVersionFallback               bool   `kafka:"api.version.fallback"`
	APIVersionFallbackMs             int    `kafka:"api.version.fallback.ms"`
	SocketKeepaliveEnable            bool   `kafka:"socket.keepalive.enable"`
	SocketKeepaliveMs                int    `kafka:"socket.keepalive.ms"`
	SocketMaxBytes                   int    `kafka:"socket.max.bytes"`
	SocketNagleEnable                bool   `kafka:"socket.nagle.enable"`
	SocketReceiveBufferBytes         int    `kafka:"socket.receive.buffer.bytes"`
	SocketSendBufferBytes            int    `kafka:"socket.send.buffer.bytes"`
	SocketBlockingMaxMs              int    `kafka:"socket.blocking.max.ms"`
	SocketMaxFails                   int    `kafka:"socket.max.fails"`
	SocketMaxIdleMs                  int    `kafka:"socket.max.idle.ms"`
	EnableSaslOauthBearer            bool   `kafka:"enable.sasl.oauthbearer"`
	EnableSslCertificateVerification bool   `kafka:"enable.ssl.certificate.verification"`

	// Producer configuration
	Acks                             string `kafka:"acks"`
	SecurityProtocol                 string `kafka:"security.protocol"`
	EnableIdempotence                bool   `kafka:"enable.idempotence"`
	MaxInFlightRequestsPerConnection int    `kafka:"max.in.flight.requests.per.connection"`
	Retries                          int    `kafka:"retries"`
	RetryBackoffMs                   int    `kafka:"retry.backoff.ms"`
	QueueBufferingMaxMessages        int    `kafka:"queue.buffering.max.messages"`
	QueueBufferingMaxKBytes          int    `kafka:"queue.buffering.max.kbytes"`
	QueueBufferingMaxMs              int    `kafka:"queue.buffering.max.ms"`
	MessageMaxBytes                  int    `kafka:"message.max.bytes"`
	MessageCopyMaxBytes              int    `kafka:"message.copy.max.bytes"`
	Partitioner                      string `kafka:"partitioner"`
	CompressionCodec                 string `kafka:"compression.codec"`
	CompressionLevel                 int    `kafka:"compression.level"`
	BatchNumMessages                 int    `kafka:"batch.num.messages"`
	BatchSize                        int    `kafka:"batch.size"`
	DeliveryReportEnabled            bool   `kafka:"delivery.report.enabled"`
	DrMsgTimeoutMs                   int    `kafka:"dr.msg.timeout.ms"`
	QueueBufferingMaxMsPerPartition  int    `kafka:"queue.buffering.max.ms.per.partition"`
	PartitionerRandomSeed            int    `kafka:"partitioner.random.seed"`
	LingerMs                         int    `kafka:"linger.ms"`
	DeliveryTimeoutMs                int    `kafka:"delivery.timeout.ms"`
	MaxBlockMs                       int    `kafka:"max.block.ms"`
	MaxQueueBufferingTimeMs          int    `kafka:"max.queue.buffering.time.ms"`
	EnableBackgroundPoll             bool   `kafka:"enable.background.poll"`
	EnableGaplessGuarantee           bool   `kafka:"enable.gapless.guarantee"`

	// Consumer configuration
	EnableAutoCommit                  bool   `kafka:"enable.auto.commit"`
	AutoCommitIntervalMs              int    `kafka:"auto.commit.interval.ms"`
	AutoOffsetReset                   string `kafka:"auto.offset.reset"`
	CheckCrcs                         bool   `kafka:"check.crcs"`
	FetchErrorBackoffMs               int    `kafka:"fetch.error.backoff.ms"`
	FetchMaxBytes                     int    `kafka:"fetch.max.bytes"`
	FetchMaxWaitMs                    int    `kafka:"fetch.max.wait.ms"`
	FetchMinBytes                     int    `kafka:"fetch.min.bytes"`
	FetchWaitMaxMs                    int    `kafka:"fetch.wait.max.ms"`
	GroupID                           string `kafka:"group.id"`
	MaxPartitionFetchBytes            int    `kafka:"max.partition.fetch.bytes"`
	SessionTimeoutMs                  int    `kafka:"session.timeout.ms"`
	HeartbeatIntervalMs               int    `kafka:"heartbeat.interval.ms"`
	MaxPollIntervalMs                 int    `kafka:"max.poll.interval.ms"`
	PartitionAssignmentStrategy       string `kafka:"partition.assignment.strategy"`
	PartitionAssignmentTimeoutMs      int    `kafka:"partition.assignment.timeout.ms"`
	IsolationLevel                    string `kafka:"isolation.level"`
	EnablePartitionEof                bool   `kafka:"enable.partition.eof"`
	CheckBrokerConnected              bool   `kafka:"check.broker.connected"`
	RebalanceTimeoutMs                int    `kafka:"rebalance.timeout.ms"`
	EnableAutoOffsetStore             bool   `kafka:"enable.auto.offset.store"`
	EnableConsumerMetadata            bool   `kafka:"enable.consumer.metadata"`
	ExcludeInternalTopics             bool   `kafka:"exclude.internal.topics"`
	AllowAutoCreateTopics             bool   `kafka:"allow.auto.create.topics"`
	Interceptors                      string `kafka:"interceptors"`
	IsolationLevelReadCommitted       bool   `kafka:"isolation.level.read.committed"`
	IsolationLevelReadUncommitted     bool   `kafka:"isolation.level.read.uncommitted"`
	EnableSaslScramUsername           bool   `kafka:"enable.sasl.scram.username"`
	EnableSaslScramPassword           bool   `kafka:"enable.sasl.scram.password"`
	EnableSaslOAuthBearerToken        bool   `kafka:"enable.sasl.oauthbearer.token"`
	EnableSaslOAuthBearerTokenRefresh bool   `kafka:"enable.sasl.oauthbearer.token.refresh"`
	EnableSslKeyLocation              bool   `kafka:"enable.ssl.key.location"`
	EnableSslCertificateLocation      bool   `kafka:"enable.ssl.certificate.location"`

	// Admin configuration
	AdminRequestTimeoutMs   int `kafka:"admin.request.timeout.ms"`
	AdminOperationTimeoutMs int `kafka:"admin.operation.timeout.ms"`
	AdminRetryBackoffMs     int `kafka:"admin.retry.backoff.ms"`
	AdminMetadataMaxAgeMs   int `kafka:"admin.metadata.max.age.ms"`
}

// 更加tag转换成kafkaConfig的参数map，相对于方便管理
func (config *KafkaConfig) CreateConfigMap() *kafka.ConfigMap {
	configMap := &kafka.ConfigMap{}
	configType := reflect.TypeOf(config)
	configValue := reflect.ValueOf(config)

	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		tag := field.Tag.Get("kafka")
		value := configValue.Field(i).Interface()

		if tag != "" {
			configMap.Set(strings.Split(tag, ",")[0], value)
		}
	}

	return configMap
}

func main() {
	producerConfig := KafkaConfig{
		BootstrapServers:  "127.0.0.1:9092",
		APIVersionRequest: true,
		MessageMaxBytes:   1000000,
		LingerMs:          10,
		Retries:           3,
		RetryBackoffMs:    1000,
		Acks:              "1",
		// 设置其他参数...
	}

	configMap := producerConfig.CreateConfigMap()
	fmt.Println(configMap)
}
