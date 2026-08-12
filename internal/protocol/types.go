package protocol

type Beacon struct {
	ID       string `json:"id"`
	Hostname string `json:"hostname"`
}

type Command struct {
	Cmd string `json:"cmd"`
}

type Result struct {
	AgentID string `json:"agent_id"`
	Output  string `json:"output"`
	Error   string `json:"error,omitempty"`
}
