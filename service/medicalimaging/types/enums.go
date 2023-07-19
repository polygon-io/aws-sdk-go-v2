// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DatastoreStatus string

// Enum values for DatastoreStatus
const (
	DatastoreStatusCreating     DatastoreStatus = "CREATING"
	DatastoreStatusCreateFailed DatastoreStatus = "CREATE_FAILED"
	DatastoreStatusActive       DatastoreStatus = "ACTIVE"
	DatastoreStatusDeleting     DatastoreStatus = "DELETING"
	DatastoreStatusDeleted      DatastoreStatus = "DELETED"
)

// Values returns all known values for DatastoreStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatastoreStatus) Values() []DatastoreStatus {
	return []DatastoreStatus{
		"CREATING",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

type ImageSetState string

// Enum values for ImageSetState
const (
	ImageSetStateActive  ImageSetState = "ACTIVE"
	ImageSetStateLocked  ImageSetState = "LOCKED"
	ImageSetStateDeleted ImageSetState = "DELETED"
)

// Values returns all known values for ImageSetState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageSetState) Values() []ImageSetState {
	return []ImageSetState{
		"ACTIVE",
		"LOCKED",
		"DELETED",
	}
}

type ImageSetWorkflowStatus string

// Enum values for ImageSetWorkflowStatus
const (
	ImageSetWorkflowStatusCreated                   ImageSetWorkflowStatus = "CREATED"
	ImageSetWorkflowStatusCopied                    ImageSetWorkflowStatus = "COPIED"
	ImageSetWorkflowStatusCopying                   ImageSetWorkflowStatus = "COPYING"
	ImageSetWorkflowStatusCopyingWithReadOnlyAccess ImageSetWorkflowStatus = "COPYING_WITH_READ_ONLY_ACCESS"
	ImageSetWorkflowStatusCopyFailed                ImageSetWorkflowStatus = "COPY_FAILED"
	ImageSetWorkflowStatusUpdating                  ImageSetWorkflowStatus = "UPDATING"
	ImageSetWorkflowStatusUpdated                   ImageSetWorkflowStatus = "UPDATED"
	ImageSetWorkflowStatusUpdateFailed              ImageSetWorkflowStatus = "UPDATE_FAILED"
	ImageSetWorkflowStatusDeleting                  ImageSetWorkflowStatus = "DELETING"
	ImageSetWorkflowStatusDeleted                   ImageSetWorkflowStatus = "DELETED"
)

// Values returns all known values for ImageSetWorkflowStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ImageSetWorkflowStatus) Values() []ImageSetWorkflowStatus {
	return []ImageSetWorkflowStatus{
		"CREATED",
		"COPIED",
		"COPYING",
		"COPYING_WITH_READ_ONLY_ACCESS",
		"COPY_FAILED",
		"UPDATING",
		"UPDATED",
		"UPDATE_FAILED",
		"DELETING",
		"DELETED",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted  JobStatus = "SUBMITTED"
	JobStatusInProgress JobStatus = "IN_PROGRESS"
	JobStatusCompleted  JobStatus = "COMPLETED"
	JobStatusFailed     JobStatus = "FAILED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
	}
}

type Operator string

// Enum values for Operator
const (
	OperatorEqual   Operator = "EQUAL"
	OperatorBetween Operator = "BETWEEN"
)

// Values returns all known values for Operator. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Operator) Values() []Operator {
	return []Operator{
		"EQUAL",
		"BETWEEN",
	}
}
