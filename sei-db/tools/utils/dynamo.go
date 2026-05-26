package utils

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBClient wraps the DynamoDB service with common operations
type DynamoDBClient struct {
	client *dynamodb.DynamoDB
	table  string
}

// NewDynamoDBClient creates a new DynamoDB client
func NewDynamoDBClient(tableName, awsRegion string) (*DynamoDBClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StateSizeAnalysis represents the analysis data to be stored in DynamoDB
type StateSizeAnalysis struct {
	BlockHeight       int64  `json:"block_height" dynamodbav:"block_height"`
	ModuleName        string `json:"module_name" dynamodbav:"module_name"`
	TotalNumKeys      uint64 `json:"total_num_keys" dynamodbav:"total_num_keys"`
	TotalKeySize      uint64 `json:"total_key_size" dynamodbav:"total_key_size"`
	TotalValueSize    uint64 `json:"total_value_size" dynamodbav:"total_value_size"`
	TotalSize         uint64 `json:"total_size" dynamodbav:"total_size"`
	PrefixBreakdown   string `json:"prefix_breakdown" dynamodbav:"prefix_breakdown"`
	ContractBreakdown string `json:"contract_breakdown" dynamodbav:"contract_breakdown"`
}

// ContractSizeEntry represents individual contract size data
type ContractSizeEntry struct {
	Address   string `json:"address"`
	TotalSize uint64 `json:"total_size"`
	KeyCount  uint64 `json:"key_count"`
}

// PrefixSize is a helper structure kept for completeness (unused here)
type PrefixSize struct {
	KeySize   uint64 `json:"key_size"`
	ValueSize uint64 `json:"value_size"`
	TotalSize uint64 `json:"total_size"`
	KeyCount  uint64 `json:"key_count"`
}

// ExportStateSizeAnalysis exports a single module analysis to DynamoDB
func (d *DynamoDBClient) ExportStateSizeAnalysis(analysis *StateSizeAnalysis) error {
	_ = "STUB: not implemented"
	// Convert to DynamoDB attribute values
	return nil
}

// Write to DynamoDB

// ExportMultipleAnalyses exports multiple analyses sequentially
func (d *DynamoDBClient) ExportMultipleAnalyses(analyses []*StateSizeAnalysis) error {
	_ = "STUB: not implemented"
	return nil
}

// UpdateLatestHeightIfGreater keeps exactly one item in the metadata table using the schema:
// Partition key: keyname (S). Numeric attribute: height (N) stores the latest height.
// Upserts the row keyname = "latest_height" and sets height = :h only if missing or lower.
func (d *DynamoDBClient) UpdateLatestHeightIfGreater(metadataTable string, height int64) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
