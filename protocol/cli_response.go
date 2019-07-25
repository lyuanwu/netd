package protocol

// CliResponse ...
type CliResponse struct {
	Retcode int
	Message string
	CmdsStd map[string]string
}
