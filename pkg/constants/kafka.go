package constants

const (
	KafkaReadMinBytes      = 512 * B
	KafkaReadMaxBytes      = 1 * MB
	KafkaRetries           = 3
	DefaultReaderGroupID   = "r"
	DefaultTimeRetainHours = 6 // 6小时

	DefaultConsumerChanCap         = 20
	DefaultKafkaProductorSyncWrite = false

	DefaultKafkaNumPartitions     = -1
	DefaultKafkaReplicationFactor = -1
)

// CartService
const (
	KafkaInteractLikeTopic    = "interact_like"    // Kafka的话题
	KafkaInteractCommentTopic = "interact_comment" // Kafka的话题
	KafkaInteractPartitionNum = 10                 // Kafka的分区数
	KafkaInteractConsumerNum  = 10                 // Kafka的并发消费者数
	KafkaInteractLikeGroupId  = "interacts"        // Kafka的订阅组id
)
