package broker

type Elements struct {
	ClientID  string `json:"client_id"`
	Username  string `json:"username"`
	Topic     string `json:"topic"`
	Payload   string `json:"payload"`
	Timestamp int64  `json:"ts"`
	Size      int32  `json:"size"`
	Action    string `json:"action"`
}

// Bridge is message bridge
type Bridge interface {
	Publish(e *Elements) (bool, error)
}

// Auth is auth interface
type Auth interface {
	CheckConnect(clientID, username, password string) bool
}
