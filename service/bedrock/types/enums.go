// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CommitmentDuration string

// Enum values for CommitmentDuration
const (
	CommitmentDurationOneMonth  CommitmentDuration = "OneMonth"
	CommitmentDurationSixMonths CommitmentDuration = "SixMonths"
)

// Values returns all known values for CommitmentDuration. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CommitmentDuration) Values() []CommitmentDuration {
	return []CommitmentDuration{
		"OneMonth",
		"SixMonths",
	}
}

type CustomizationType string

// Enum values for CustomizationType
const (
	CustomizationTypeFineTuning           CustomizationType = "FINE_TUNING"
	CustomizationTypeContinuedPreTraining CustomizationType = "CONTINUED_PRE_TRAINING"
)

// Values returns all known values for CustomizationType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (CustomizationType) Values() []CustomizationType {
	return []CustomizationType{
		"FINE_TUNING",
		"CONTINUED_PRE_TRAINING",
	}
}

type FineTuningJobStatus string

// Enum values for FineTuningJobStatus
const (
	FineTuningJobStatusInProgress FineTuningJobStatus = "InProgress"
	FineTuningJobStatusCompleted  FineTuningJobStatus = "Completed"
	FineTuningJobStatusFailed     FineTuningJobStatus = "Failed"
	FineTuningJobStatusStopping   FineTuningJobStatus = "Stopping"
	FineTuningJobStatusStopped    FineTuningJobStatus = "Stopped"
)

// Values returns all known values for FineTuningJobStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FineTuningJobStatus) Values() []FineTuningJobStatus {
	return []FineTuningJobStatus{
		"InProgress",
		"Completed",
		"Failed",
		"Stopping",
		"Stopped",
	}
}

type FoundationModelLifecycleStatus string

// Enum values for FoundationModelLifecycleStatus
const (
	FoundationModelLifecycleStatusActive FoundationModelLifecycleStatus = "ACTIVE"
	FoundationModelLifecycleStatusLegacy FoundationModelLifecycleStatus = "LEGACY"
)

// Values returns all known values for FoundationModelLifecycleStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (FoundationModelLifecycleStatus) Values() []FoundationModelLifecycleStatus {
	return []FoundationModelLifecycleStatus{
		"ACTIVE",
		"LEGACY",
	}
}

type InferenceType string

// Enum values for InferenceType
const (
	InferenceTypeOnDemand    InferenceType = "ON_DEMAND"
	InferenceTypeProvisioned InferenceType = "PROVISIONED"
)

// Values returns all known values for InferenceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (InferenceType) Values() []InferenceType {
	return []InferenceType{
		"ON_DEMAND",
		"PROVISIONED",
	}
}

type ModelCustomization string

// Enum values for ModelCustomization
const (
	ModelCustomizationFineTuning           ModelCustomization = "FINE_TUNING"
	ModelCustomizationContinuedPreTraining ModelCustomization = "CONTINUED_PRE_TRAINING"
)

// Values returns all known values for ModelCustomization. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ModelCustomization) Values() []ModelCustomization {
	return []ModelCustomization{
		"FINE_TUNING",
		"CONTINUED_PRE_TRAINING",
	}
}

type ModelCustomizationJobStatus string

// Enum values for ModelCustomizationJobStatus
const (
	ModelCustomizationJobStatusInProgress ModelCustomizationJobStatus = "InProgress"
	ModelCustomizationJobStatusCompleted  ModelCustomizationJobStatus = "Completed"
	ModelCustomizationJobStatusFailed     ModelCustomizationJobStatus = "Failed"
	ModelCustomizationJobStatusStopping   ModelCustomizationJobStatus = "Stopping"
	ModelCustomizationJobStatusStopped    ModelCustomizationJobStatus = "Stopped"
)

// Values returns all known values for ModelCustomizationJobStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ModelCustomizationJobStatus) Values() []ModelCustomizationJobStatus {
	return []ModelCustomizationJobStatus{
		"InProgress",
		"Completed",
		"Failed",
		"Stopping",
		"Stopped",
	}
}

type ModelModality string

// Enum values for ModelModality
const (
	ModelModalityText      ModelModality = "TEXT"
	ModelModalityImage     ModelModality = "IMAGE"
	ModelModalityEmbedding ModelModality = "EMBEDDING"
)

// Values returns all known values for ModelModality. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ModelModality) Values() []ModelModality {
	return []ModelModality{
		"TEXT",
		"IMAGE",
		"EMBEDDING",
	}
}

type ProvisionedModelStatus string

// Enum values for ProvisionedModelStatus
const (
	ProvisionedModelStatusCreating  ProvisionedModelStatus = "Creating"
	ProvisionedModelStatusInService ProvisionedModelStatus = "InService"
	ProvisionedModelStatusUpdating  ProvisionedModelStatus = "Updating"
	ProvisionedModelStatusFailed    ProvisionedModelStatus = "Failed"
)

// Values returns all known values for ProvisionedModelStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ProvisionedModelStatus) Values() []ProvisionedModelStatus {
	return []ProvisionedModelStatus{
		"Creating",
		"InService",
		"Updating",
		"Failed",
	}
}

type SortByProvisionedModels string

// Enum values for SortByProvisionedModels
const (
	SortByProvisionedModelsCreationTime SortByProvisionedModels = "CreationTime"
)

// Values returns all known values for SortByProvisionedModels. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SortByProvisionedModels) Values() []SortByProvisionedModels {
	return []SortByProvisionedModels{
		"CreationTime",
	}
}

type SortJobsBy string

// Enum values for SortJobsBy
const (
	SortJobsByCreationTime SortJobsBy = "CreationTime"
)

// Values returns all known values for SortJobsBy. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortJobsBy) Values() []SortJobsBy {
	return []SortJobsBy{
		"CreationTime",
	}
}

type SortModelsBy string

// Enum values for SortModelsBy
const (
	SortModelsByCreationTime SortModelsBy = "CreationTime"
)

// Values returns all known values for SortModelsBy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SortModelsBy) Values() []SortModelsBy {
	return []SortModelsBy{
		"CreationTime",
	}
}

type SortOrder string

// Enum values for SortOrder
const (
	SortOrderAscending  SortOrder = "Ascending"
	SortOrderDescending SortOrder = "Descending"
)

// Values returns all known values for SortOrder. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SortOrder) Values() []SortOrder {
	return []SortOrder{
		"Ascending",
		"Descending",
	}
}
