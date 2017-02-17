package victorops

// Alert the details of an alert
type Alert struct {
	MessageType       string `json:"messageType"`       // The type of alert; INFO, WARNING, ACKNOWLEDGEMENT, CRITICAL, RECOVERY
	EntityID          string `json:"entityId"`          // Identifies the entity (host, service, etc.) this alert was about.
	Timestamp         int64  `json:"timestamp"`         // Timestamp of the alert in seconds since epoch
	StateStartTime    int64  `json:"stateStartTime"`    // The time this entity entered its current state (seconds since epoch)
	StateMessage      string `json:"stateMessage"`      // Any additional status information from the alert item
	MonitoringTool    string `json:"monitoringTool"`    // The name of the monitoring system software (eg. nagios, icinga, sensu, etc.)
	EntityDisplayName string `json:"entityDisplayName"` // Used within VictorOps to display a human-readable name for the entity
	AckMsg            string `json:"ackMsg"`            // A user entered comment for the acknowledgment
	AckAuthor         string `json:"ackAuthor"`         // The user that acknowledged the incident
	Raw               string `json:"raw"`               // The full, raw alert details JSON string (i.e. parse the string into a JSON object)
}

// AlertDetails retrieves the details of an alert that was sent VictorOps by you
func (client *Client) AlertDetails(uuid string) (Alert, error) {
	var alert Alert
	err := client.sendRequest("GET", "/api-public/v1/alerts/"+uuid, nil, &alert)
	return alert, err
}
