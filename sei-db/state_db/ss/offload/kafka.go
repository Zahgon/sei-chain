package offload

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"github.com/segmentio/kafka-go/sasl"

	dbproto "github.com/sei-protocol/sei-chain/sei-db/proto"
)

const kafkaOptionNone = "none"

type KafkaConfig struct {
	Brokers       []string
	Topic         string
	ClientID      string
	Region        string
	Async         bool
	RequiredAcks  string
	Compression   string
	BatchSize     int
	BatchTimeout  time.Duration
	BatchBytes    int
	TLSEnabled    bool
	SASLMechanism string
}

func (c *KafkaConfig) ApplyDefaults() { _ = "STUB: not implemented"; return }

func (c *KafkaConfig) Validate() error { _ = "STUB: not implemented"; return nil }

type kafkaStream struct {
	writer  *kafka.Writer
	durable bool
}

var _ Stream = (*kafkaStream)(nil)

func NewKafkaStream(cfg KafkaConfig) (Stream, error) {
	_ = "STUB: not implemented"
	return *new(Stream), nil
}

func (k *kafkaStream) Publish(ctx context.Context, entry *dbproto.ChangelogEntry) (Ack, error) {
	_ = "STUB: not implemented"
	return *new(Ack), nil
}

func (k *kafkaStream) Close() error { _ = "STUB: not implemented"; return nil }

func kafkaRequiredAcks(requiredAcks string) kafka.RequiredAcks {
	_ = "STUB: not implemented"
	return *new(kafka.RequiredAcks)
}

func kafkaCompression(name string) compress.Compression {
	_ = "STUB: not implemented"
	return *new(compress.Compression)
}

func kafkaSASLMechanism(cfg KafkaConfig) (sasl.Mechanism, error) {
	_ = "STUB: not implemented"
	return *new(sasl.Mechanism), nil
}
