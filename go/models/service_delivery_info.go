package models


type ServiceDeliveryInfo struct {
	ResponseTimestamp        time.Time `json:"ResponseTimestamp,omitempty"`
	ProducerRef              string    `json:"ProducerRef,omitempty"`
	RequestMessageIdentifier any       `json:"RequestMessageIdentifier,omitempty"`
	ResponseMessageRef       any       `json:"ResponseMessageRef,omitempty"`
} `json:"ServiceDeliveryInfo,omitempty"`