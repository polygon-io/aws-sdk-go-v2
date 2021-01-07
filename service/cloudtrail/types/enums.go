// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type EventCategory string

// Enum values for EventCategory
const (
	EventCategoryInsight EventCategory = "insight"
)

// Values returns all known values for EventCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EventCategory) Values() []EventCategory {
	return []EventCategory{
		"insight",
	}
}

type InsightType string

// Enum values for InsightType
const (
	InsightTypeApiCallRateInsight InsightType = "ApiCallRateInsight"
)

// Values returns all known values for InsightType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (InsightType) Values() []InsightType {
	return []InsightType{
		"ApiCallRateInsight",
	}
}

type LookupAttributeKey string

// Enum values for LookupAttributeKey
const (
	LookupAttributeKeyEventId      LookupAttributeKey = "EventId"
	LookupAttributeKeyEventName    LookupAttributeKey = "EventName"
	LookupAttributeKeyReadOnly     LookupAttributeKey = "ReadOnly"
	LookupAttributeKeyUsername     LookupAttributeKey = "Username"
	LookupAttributeKeyResourceType LookupAttributeKey = "ResourceType"
	LookupAttributeKeyResourceName LookupAttributeKey = "ResourceName"
	LookupAttributeKeyEventSource  LookupAttributeKey = "EventSource"
	LookupAttributeKeyAccessKeyId  LookupAttributeKey = "AccessKeyId"
)

// Values returns all known values for LookupAttributeKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LookupAttributeKey) Values() []LookupAttributeKey {
	return []LookupAttributeKey{
		"EventId",
		"EventName",
		"ReadOnly",
		"Username",
		"ResourceType",
		"ResourceName",
		"EventSource",
		"AccessKeyId",
	}
}

type ReadWriteType string

// Enum values for ReadWriteType
const (
	ReadWriteTypeReadOnly  ReadWriteType = "ReadOnly"
	ReadWriteTypeWriteOnly ReadWriteType = "WriteOnly"
	ReadWriteTypeAll       ReadWriteType = "All"
)

// Values returns all known values for ReadWriteType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReadWriteType) Values() []ReadWriteType {
	return []ReadWriteType{
		"ReadOnly",
		"WriteOnly",
		"All",
	}
}
