// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActionCategory string

// Enum values for ActionCategory
const (
	ActionCategorySource   ActionCategory = "Source"
	ActionCategoryBuild    ActionCategory = "Build"
	ActionCategoryDeploy   ActionCategory = "Deploy"
	ActionCategoryTest     ActionCategory = "Test"
	ActionCategoryInvoke   ActionCategory = "Invoke"
	ActionCategoryApproval ActionCategory = "Approval"
)

// Values returns all known values for ActionCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionCategory) Values() []ActionCategory {
	return []ActionCategory{
		"Source",
		"Build",
		"Deploy",
		"Test",
		"Invoke",
		"Approval",
	}
}

type ActionConfigurationPropertyType string

// Enum values for ActionConfigurationPropertyType
const (
	ActionConfigurationPropertyTypeString  ActionConfigurationPropertyType = "String"
	ActionConfigurationPropertyTypeNumber  ActionConfigurationPropertyType = "Number"
	ActionConfigurationPropertyTypeBoolean ActionConfigurationPropertyType = "Boolean"
)

// Values returns all known values for ActionConfigurationPropertyType. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ActionConfigurationPropertyType) Values() []ActionConfigurationPropertyType {
	return []ActionConfigurationPropertyType{
		"String",
		"Number",
		"Boolean",
	}
}

type ActionExecutionStatus string

// Enum values for ActionExecutionStatus
const (
	ActionExecutionStatusInProgress ActionExecutionStatus = "InProgress"
	ActionExecutionStatusAbandoned  ActionExecutionStatus = "Abandoned"
	ActionExecutionStatusSucceeded  ActionExecutionStatus = "Succeeded"
	ActionExecutionStatusFailed     ActionExecutionStatus = "Failed"
)

// Values returns all known values for ActionExecutionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActionExecutionStatus) Values() []ActionExecutionStatus {
	return []ActionExecutionStatus{
		"InProgress",
		"Abandoned",
		"Succeeded",
		"Failed",
	}
}

type ActionOwner string

// Enum values for ActionOwner
const (
	ActionOwnerAws        ActionOwner = "AWS"
	ActionOwnerThirdParty ActionOwner = "ThirdParty"
	ActionOwnerCustom     ActionOwner = "Custom"
)

// Values returns all known values for ActionOwner. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ActionOwner) Values() []ActionOwner {
	return []ActionOwner{
		"AWS",
		"ThirdParty",
		"Custom",
	}
}

type ApprovalStatus string

// Enum values for ApprovalStatus
const (
	ApprovalStatusApproved ApprovalStatus = "Approved"
	ApprovalStatusRejected ApprovalStatus = "Rejected"
)

// Values returns all known values for ApprovalStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ApprovalStatus) Values() []ApprovalStatus {
	return []ApprovalStatus{
		"Approved",
		"Rejected",
	}
}

type ArtifactLocationType string

// Enum values for ArtifactLocationType
const (
	ArtifactLocationTypeS3 ArtifactLocationType = "S3"
)

// Values returns all known values for ArtifactLocationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactLocationType) Values() []ArtifactLocationType {
	return []ArtifactLocationType{
		"S3",
	}
}

type ArtifactStoreType string

// Enum values for ArtifactStoreType
const (
	ArtifactStoreTypeS3 ArtifactStoreType = "S3"
)

// Values returns all known values for ArtifactStoreType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ArtifactStoreType) Values() []ArtifactStoreType {
	return []ArtifactStoreType{
		"S3",
	}
}

type BlockerType string

// Enum values for BlockerType
const (
	BlockerTypeSchedule BlockerType = "Schedule"
)

// Values returns all known values for BlockerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (BlockerType) Values() []BlockerType {
	return []BlockerType{
		"Schedule",
	}
}

type EncryptionKeyType string

// Enum values for EncryptionKeyType
const (
	EncryptionKeyTypeKms EncryptionKeyType = "KMS"
)

// Values returns all known values for EncryptionKeyType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionKeyType) Values() []EncryptionKeyType {
	return []EncryptionKeyType{
		"KMS",
	}
}

type ExecutorType string

// Enum values for ExecutorType
const (
	ExecutorTypeJobWorker ExecutorType = "JobWorker"
	ExecutorTypeLambda    ExecutorType = "Lambda"
)

// Values returns all known values for ExecutorType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExecutorType) Values() []ExecutorType {
	return []ExecutorType{
		"JobWorker",
		"Lambda",
	}
}

type FailureType string

// Enum values for FailureType
const (
	FailureTypeJobFailed           FailureType = "JobFailed"
	FailureTypeConfigurationError  FailureType = "ConfigurationError"
	FailureTypePermissionError     FailureType = "PermissionError"
	FailureTypeRevisionOutOfSync   FailureType = "RevisionOutOfSync"
	FailureTypeRevisionUnavailable FailureType = "RevisionUnavailable"
	FailureTypeSystemUnavailable   FailureType = "SystemUnavailable"
)

// Values returns all known values for FailureType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FailureType) Values() []FailureType {
	return []FailureType{
		"JobFailed",
		"ConfigurationError",
		"PermissionError",
		"RevisionOutOfSync",
		"RevisionUnavailable",
		"SystemUnavailable",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusCreated    JobStatus = "Created"
	JobStatusQueued     JobStatus = "Queued"
	JobStatusDispatched JobStatus = "Dispatched"
	JobStatusInProgress JobStatus = "InProgress"
	JobStatusTimedOut   JobStatus = "TimedOut"
	JobStatusSucceeded  JobStatus = "Succeeded"
	JobStatusFailed     JobStatus = "Failed"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"Created",
		"Queued",
		"Dispatched",
		"InProgress",
		"TimedOut",
		"Succeeded",
		"Failed",
	}
}

type PipelineExecutionStatus string

// Enum values for PipelineExecutionStatus
const (
	PipelineExecutionStatusCancelled  PipelineExecutionStatus = "Cancelled"
	PipelineExecutionStatusInProgress PipelineExecutionStatus = "InProgress"
	PipelineExecutionStatusStopped    PipelineExecutionStatus = "Stopped"
	PipelineExecutionStatusStopping   PipelineExecutionStatus = "Stopping"
	PipelineExecutionStatusSucceeded  PipelineExecutionStatus = "Succeeded"
	PipelineExecutionStatusSuperseded PipelineExecutionStatus = "Superseded"
	PipelineExecutionStatusFailed     PipelineExecutionStatus = "Failed"
)

// Values returns all known values for PipelineExecutionStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PipelineExecutionStatus) Values() []PipelineExecutionStatus {
	return []PipelineExecutionStatus{
		"Cancelled",
		"InProgress",
		"Stopped",
		"Stopping",
		"Succeeded",
		"Superseded",
		"Failed",
	}
}

type StageExecutionStatus string

// Enum values for StageExecutionStatus
const (
	StageExecutionStatusCancelled  StageExecutionStatus = "Cancelled"
	StageExecutionStatusInProgress StageExecutionStatus = "InProgress"
	StageExecutionStatusFailed     StageExecutionStatus = "Failed"
	StageExecutionStatusStopped    StageExecutionStatus = "Stopped"
	StageExecutionStatusStopping   StageExecutionStatus = "Stopping"
	StageExecutionStatusSucceeded  StageExecutionStatus = "Succeeded"
)

// Values returns all known values for StageExecutionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StageExecutionStatus) Values() []StageExecutionStatus {
	return []StageExecutionStatus{
		"Cancelled",
		"InProgress",
		"Failed",
		"Stopped",
		"Stopping",
		"Succeeded",
	}
}

type StageRetryMode string

// Enum values for StageRetryMode
const (
	StageRetryModeFailedActions StageRetryMode = "FAILED_ACTIONS"
	StageRetryModeAllActions    StageRetryMode = "ALL_ACTIONS"
)

// Values returns all known values for StageRetryMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StageRetryMode) Values() []StageRetryMode {
	return []StageRetryMode{
		"FAILED_ACTIONS",
		"ALL_ACTIONS",
	}
}

type StageTransitionType string

// Enum values for StageTransitionType
const (
	StageTransitionTypeInbound  StageTransitionType = "Inbound"
	StageTransitionTypeOutbound StageTransitionType = "Outbound"
)

// Values returns all known values for StageTransitionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (StageTransitionType) Values() []StageTransitionType {
	return []StageTransitionType{
		"Inbound",
		"Outbound",
	}
}

type TriggerType string

// Enum values for TriggerType
const (
	TriggerTypeCreatePipeline         TriggerType = "CreatePipeline"
	TriggerTypeStartPipelineExecution TriggerType = "StartPipelineExecution"
	TriggerTypePollForSourceChanges   TriggerType = "PollForSourceChanges"
	TriggerTypeWebhook                TriggerType = "Webhook"
	TriggerTypeCloudWatchEvent        TriggerType = "CloudWatchEvent"
	TriggerTypePutActionRevision      TriggerType = "PutActionRevision"
)

// Values returns all known values for TriggerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TriggerType) Values() []TriggerType {
	return []TriggerType{
		"CreatePipeline",
		"StartPipelineExecution",
		"PollForSourceChanges",
		"Webhook",
		"CloudWatchEvent",
		"PutActionRevision",
	}
}

type WebhookAuthenticationType string

// Enum values for WebhookAuthenticationType
const (
	WebhookAuthenticationTypeGithubHmac      WebhookAuthenticationType = "GITHUB_HMAC"
	WebhookAuthenticationTypeIp              WebhookAuthenticationType = "IP"
	WebhookAuthenticationTypeUnauthenticated WebhookAuthenticationType = "UNAUTHENTICATED"
)

// Values returns all known values for WebhookAuthenticationType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (WebhookAuthenticationType) Values() []WebhookAuthenticationType {
	return []WebhookAuthenticationType{
		"GITHUB_HMAC",
		"IP",
		"UNAUTHENTICATED",
	}
}
