package eventparser

import (
	"encoding/json"
	"fmt"
	"time"
)

// EventParser adds 
type EventParser interface {
	GetEventRequestPayload(string) (EventRequestPayload, error)
	GetEventRequestAsJSON(EventRequestPayload) []byte
}

// EventRequestPayload represents a POST request's body which is sent to Event-Service
type EventRequestPayload struct {
	EventType        string `json:"event-type"`
	EventTypeVersion string `json:"event-type-version"`
	EventID          string `json:"event-id"` //uuid should be generated automatically if send empty
	EventTime        string `json:"event-time"`
	Data             string `json:"data"` //github webhook json payload
}

// GetEventRequestPayload generates structure which 
func GetEventRequestPayload(eventType, eventTypeVersion, eventID, data string) (EventRequestPayload, error) {

	res := EventRequestPayload{
		eventType,
		eventTypeVersion,
		eventID,
		time.Now().Format(time.RFC3339),
		data}

	return res, nil

}


// GetEventRequestAsJSON here description
func GetEventRequestAsJSON(eventRequestPayload EventRequestPayload) ([]byte, error) {
	r, err := json.MarshalIndent(eventRequestPayload, "", "  ")
	if err != nil {
		fmt.Println("error: ")
	}
	return r, nil
}
