package simulation

import (
	"os"
)

// log writter
type LogWriter interface {
	AddEntry(OperationEntry)
	PrintLogs()
}

// LogWriter - return a dummy or standard log writer given the testingmode
func NewLogWriter(testingmode bool) LogWriter { _ = "STUB: not implemented"; return *new(LogWriter) }

// log writter
type StandardLogWriter struct {
	OpEntries []OperationEntry `json:"op_entries" yaml:"op_entries"`
}

// add an entry to the log writter
func (lw *StandardLogWriter) AddEntry(opEntry OperationEntry) { _ = "STUB: not implemented"; return }

// PrintLogs - print the logs to a simulation file
func (lw *StandardLogWriter) PrintLogs() { _ = "STUB: not implemented"; return }

func createLogFile() *os.File { _ = "STUB: not implemented"; return nil }

// dummy log writter
type DummyLogWriter struct{}

// do nothing
func (lw *DummyLogWriter) AddEntry(_ OperationEntry) {
	_ = "STUB: not implemented"

	// do nothing
	return
}

func (lw *DummyLogWriter) PrintLogs() { _ = "STUB: not implemented"; return }
