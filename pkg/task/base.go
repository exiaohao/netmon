package task

type AgentTask struct {
	taskType	int
	taskParam	interface{}
}

type TaskList struct {
	agentID	string
	Tasks	[]AgentTask
}
