package offload

import (
	"context"
	"fmt"
	"runtime"
	"time"

	v4 "github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/segmentio/kafka-go/sasl"
)

const (
	mskIAMVersion      = "2020_10_22"
	mskIAMService      = "kafka-cluster"
	mskIAMAction       = "kafka-cluster:Connect"
	mskIAMVersionKey   = "version"
	mskIAMHostKey      = "host"
	mskIAMUserAgentKey = "user-agent"
	mskIAMActionKey    = "action"
	mskIAMQueryAction  = "Action"
)

var mskIAMUserAgent = fmt.Sprintf("sei-chain/cryptosim/aws_msk_iam/%s", runtime.Version())

type awsMSKIAMMechanism struct {
	signer   *v4.Signer
	region   string
	signTime time.Time
	expiry   time.Duration
}

var _ sasl.Mechanism = (*awsMSKIAMMechanism)(nil)
var _ sasl.StateMachine = (*awsMSKIAMMechanism)(nil)

func newAWSMSKIAMMechanism(cfg KafkaConfig) (sasl.Mechanism, error) {
	_ = "STUB: not implemented"
	return *new(sasl.Mechanism), nil
}

func (m *awsMSKIAMMechanism) Name() string { _ = "STUB: not implemented"; return "" }

func (m *awsMSKIAMMechanism) Start(ctx context.Context) (sasl.StateMachine, []byte, error) {
	_ = "STUB: not implemented"
	return *new(sasl.StateMachine), nil, nil
}

func (m *awsMSKIAMMechanism) Next(context.Context, []byte) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}
