package state

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type Logs struct {
	Ls []*ethtypes.Log `json:"logs"`
}

func (s *DBImpl) AddLog(l *ethtypes.Log) { _ = "STUB: not implemented"; return }

func (s *DBImpl) GetAllLogs() []*ethtypes.Log { _ = "STUB: not implemented"; return nil }

func (s *DBImpl) GetLogs(common.Hash, uint64, common.Hash) []*ethtypes.Log {
	_ = "STUB: not implemented"
	return nil
}

func (s *DBImpl) Logs() []*ethtypes.Log { _ = "STUB: not implemented"; return nil }
