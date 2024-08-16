package nice

type Message struct {
	Eid  string `json:"eid"`
	Kind string `json:"kind"`
	Data any    `json:"data"`
}

type EventMsg struct {
	EventKind  string `json:"event_kind"`
	InputValue string `json:"input_value"`
}
