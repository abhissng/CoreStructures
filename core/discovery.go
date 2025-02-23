package core

import (
	"github.com/abhissng/internal/utils/constant"
	"github.com/abhissng/internal/utils/types"
)

// DiscoveryMessagepayload represents the structure of a discovery message request.
type DiscoveryMessagePayload struct {
	// TODO add additionalInformation Later also add binding to metadata
	// MetaData *structures.MetaData `json:"meta_data" binding:"required"`
	Message  *Message[*Core] `json:"message" binding:"required"`
	MetaData *MetaData       `json:"meta_data"`
}

// NewDiscoveryMessagePayload creates a new DiscoveryMessagePayload
// with the given correlation ID and ("github.com/abhissng/core-structures/core") Core payload
// with action set as execute and status set as pending
func NewDiscoveryMessagePayload(
	correlationId types.CorrelationID,
	core *Core,
) *DiscoveryMessagePayload {
	return &DiscoveryMessagePayload{
		Message:  NewMessage(constant.Execute, constant.Pending, correlationId, core),
		MetaData: NewMetaData(),
	}
}

// AddMetaData adds metadata to the DiscoveryMessage
func (d *DiscoveryMessagePayload) AddMetaData(metaData *MetaData) *DiscoveryMessagePayload {
	d.MetaData = metaData
	return d
}

// AddAction adds action to the DiscoveryMessage
func (d *DiscoveryMessagePayload) AddAction(action types.Action) *DiscoveryMessagePayload {
	d.Message.Action = action
	return d
}

// AddStatus adds status to the DiscoveryMessage
func (d *DiscoveryMessagePayload) AddStatus(status types.Status) *DiscoveryMessagePayload {
	d.Message.Status = status
	return d
}

// // DiscoveryResponse represents the structure of a discovery response.
// type DiscoveryMessageResponse struct {
// 	// TODO add additionalInformation Later also add binding to metadata
// 	// MetaData *structures.MetaData `json:"meta_data" binding:"required"`
// 	Message  *Message[*core.Core] `json:"message"`
// 	MetaData *metadata.MetaData   `json:"meta_data"`
// }

// // NewDiscoveryMessageResponse creates a new DiscoveryMessageResponse
// func NewDiscoveryMessageResponse(
// 	metaData *metadata.MetaData,
// 	action types.Action,
// 	status types.Status,
// 	correlationId types.CorrelationID,
// 	core *core.Core) *DiscoveryMessageResponse {
// 	return &DiscoveryMessageResponse{
// 		Message:  NewMessage(action, status, correlationId, core),
// 		MetaData: metaData.UpdateMetaData(),
// 	}
// }
