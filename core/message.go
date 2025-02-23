package core

import (
	"time"

	"github.com/abhissng/internal/utils/random"
	"github.com/abhissng/internal/utils/types"
)

// Message represents the structure of a transaction message.
// use core.Core from "github.com/abhissng/core-structures/core" for payload Type
// or you can use your custom Type if needed
type Message[T any] struct {
	CorrelationID  types.CorrelationID `json:"correlation_id"`
	RequestId      types.RequestID     `json:"request_id"`
	Payload        T                   `json:"payload"`
	Status         types.Status        `json:"status"` // "pending", "completed or success", "failed"
	Action         types.Action        `json:"action"` // "execute", "rollback"
	Error          error               `json:"error,omitempty"`
	Timestamp      time.Time           `json:"timestamp"`
	CurrentService string              `json:"current_service,omitempty"`
}

func NewMessage[T any](
	action types.Action,
	status types.Status,
	correlationID types.CorrelationID,
	payload T,
) *Message[T] {
	return &Message[T]{
		CorrelationID: correlationID,
		RequestId:     types.RequestID(random.GenerateUUID()),
		Payload:       payload,
		Status:        status,
		Action:        action,
		Timestamp:     time.Now(),
	}
}
