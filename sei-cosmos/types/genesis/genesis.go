package genesis

type GenesisImportConfig struct {
	StreamGenesisImport bool
	GenesisStreamFile   string
}

const bufferSize = 100000

func IngestGenesisFileLineByLine(filename string) <-chan string {
	_ = "STUB: not implemented"
	return nil
}
