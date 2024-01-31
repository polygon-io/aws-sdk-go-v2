// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AWSServiceAccessStatus string

// Enum values for AWSServiceAccessStatus
const (
	AWSServiceAccessStatusEnabled  AWSServiceAccessStatus = "ENABLED"
	AWSServiceAccessStatusDisabled AWSServiceAccessStatus = "DISABLED"
)

// Values returns all known values for AWSServiceAccessStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AWSServiceAccessStatus) Values() []AWSServiceAccessStatus {
	return []AWSServiceAccessStatus{
		"ENABLED",
		"DISABLED",
	}
}

type IndexState string

// Enum values for IndexState
const (
	// Resource Explorer is creating the index.
	IndexStateCreating IndexState = "CREATING"
	// Index is active.
	IndexStateActive IndexState = "ACTIVE"
	// Resource Explorer is deleting the index.
	IndexStateDeleting IndexState = "DELETING"
	// Resource Explorer successfully deleted the index.
	IndexStateDeleted IndexState = "DELETED"
	// Resource Explorer is switching the index type between local and aggregator.
	IndexStateUpdating IndexState = "UPDATING"
)

// Values returns all known values for IndexState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IndexState) Values() []IndexState {
	return []IndexState{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"UPDATING",
	}
}

type IndexType string

// Enum values for IndexType
const (
	// local index
	IndexTypeLocal IndexType = "LOCAL"
	// aggregator index
	IndexTypeAggregator IndexType = "AGGREGATOR"
)

// Values returns all known values for IndexType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IndexType) Values() []IndexType {
	return []IndexType{
		"LOCAL",
		"AGGREGATOR",
	}
}
