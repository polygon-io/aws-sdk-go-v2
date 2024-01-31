// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// CloudWatch logging configuration.
type CloudWatchConfig struct {

	// The log group name.
	//
	// This member is required.
	LogGroupName *string

	// The role ARN.
	//
	// This member is required.
	RoleArn *string

	// S3 configuration for delivering a large amount of data.
	LargeDataDeliveryS3Config *S3Config

	noSmithyDocumentSerde
}

// Summary information for a custom model.
type CustomModelSummary struct {

	// The base model ARN.
	//
	// This member is required.
	BaseModelArn *string

	// The base model name.
	//
	// This member is required.
	BaseModelName *string

	// Creation time of the model.
	//
	// This member is required.
	CreationTime *time.Time

	// The ARN of the custom model.
	//
	// This member is required.
	ModelArn *string

	// The name of the custom model.
	//
	// This member is required.
	ModelName *string

	// Specifies whether to carry out continued pre-training of a model or whether to
	// fine-tune it. For more information, see Custom models (https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html)
	// .
	CustomizationType CustomizationType

	noSmithyDocumentSerde
}

// Information about a foundation model.
type FoundationModelDetails struct {

	// The model ARN.
	//
	// This member is required.
	ModelArn *string

	// The model identifier.
	//
	// This member is required.
	ModelId *string

	// The customization that the model supports.
	CustomizationsSupported []ModelCustomization

	// The inference types that the model supports.
	InferenceTypesSupported []InferenceType

	// The input modalities that the model supports.
	InputModalities []ModelModality

	// Contains details about whether a model version is available or deprecated
	ModelLifecycle *FoundationModelLifecycle

	// The model name.
	ModelName *string

	// The output modalities that the model supports.
	OutputModalities []ModelModality

	// he model's provider name.
	ProviderName *string

	// Indicates whether the model supports streaming.
	ResponseStreamingSupported *bool

	noSmithyDocumentSerde
}

// Details about whether a model version is available or deprecated.
type FoundationModelLifecycle struct {

	// Specifies whether a model version is available ( ACTIVE ) or deprecated ( LEGACY
	// .
	//
	// This member is required.
	Status FoundationModelLifecycleStatus

	noSmithyDocumentSerde
}

// Summary information for a foundation model.
type FoundationModelSummary struct {

	// The ARN of the foundation model.
	//
	// This member is required.
	ModelArn *string

	// The model Id of the foundation model.
	//
	// This member is required.
	ModelId *string

	// Whether the model supports fine-tuning or continual pre-training.
	CustomizationsSupported []ModelCustomization

	// The inference types that the model supports.
	InferenceTypesSupported []InferenceType

	// The input modalities that the model supports.
	InputModalities []ModelModality

	// Contains details about whether a model version is available or deprecated.
	ModelLifecycle *FoundationModelLifecycle

	// The name of the model.
	ModelName *string

	// The output modalities that the model supports.
	OutputModalities []ModelModality

	// The model's provider name.
	ProviderName *string

	// Indicates whether the model supports streaming.
	ResponseStreamingSupported *bool

	noSmithyDocumentSerde
}

// Configuration fields for invokation logging.
type LoggingConfig struct {

	// CloudWatch logging configuration.
	CloudWatchConfig *CloudWatchConfig

	// Set to include embeddings data in the log delivery.
	EmbeddingDataDeliveryEnabled *bool

	// Set to include image data in the log delivery.
	ImageDataDeliveryEnabled *bool

	// S3 configuration for storing log data.
	S3Config *S3Config

	// Set to include text data in the log delivery.
	TextDataDeliveryEnabled *bool

	noSmithyDocumentSerde
}

// Information about one customization job
type ModelCustomizationJobSummary struct {

	// ARN of the base model.
	//
	// This member is required.
	BaseModelArn *string

	// Creation time of the custom model.
	//
	// This member is required.
	CreationTime *time.Time

	// ARN of the customization job.
	//
	// This member is required.
	JobArn *string

	// Name of the customization job.
	//
	// This member is required.
	JobName *string

	// Status of the customization job.
	//
	// This member is required.
	Status ModelCustomizationJobStatus

	// ARN of the custom model.
	CustomModelArn *string

	// Name of the custom model.
	CustomModelName *string

	// Specifies whether to carry out continued pre-training of a model or whether to
	// fine-tune it. For more information, see Custom models (https://docs.aws.amazon.com/bedrock/latest/userguide/custom-models.html)
	// .
	CustomizationType CustomizationType

	// Time that the customization job ended.
	EndTime *time.Time

	// Time that the customization job was last modified.
	LastModifiedTime *time.Time

	noSmithyDocumentSerde
}

// S3 Location of the output data.
type OutputDataConfig struct {

	// The S3 URI where the output data is stored.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// Set of fields associated with a provisioned throughput.
type ProvisionedModelSummary struct {

	// The time that this provisioned throughput was created.
	//
	// This member is required.
	CreationTime *time.Time

	// Desired model ARN.
	//
	// This member is required.
	DesiredModelArn *string

	// Desired model units.
	//
	// This member is required.
	DesiredModelUnits *int32

	// Foundation model ARN.
	//
	// This member is required.
	FoundationModelArn *string

	// The time that this provisioned throughput was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The ARN of the model associated with this provisioned throughput.
	//
	// This member is required.
	ModelArn *string

	// The number of model units allocated.
	//
	// This member is required.
	ModelUnits *int32

	// The ARN of the provisioned throughput.
	//
	// This member is required.
	ProvisionedModelArn *string

	// The name of the provisioned throughput.
	//
	// This member is required.
	ProvisionedModelName *string

	// Status of the provisioned throughput.
	//
	// This member is required.
	Status ProvisionedModelStatus

	// Commitment duration for the provisioned throughput.
	CommitmentDuration CommitmentDuration

	// Commitment expiration time for the provisioned throughput.
	CommitmentExpirationTime *time.Time

	noSmithyDocumentSerde
}

// S3 configuration for storing log data.
type S3Config struct {

	// S3 bucket name.
	//
	// This member is required.
	BucketName *string

	// S3 prefix.
	KeyPrefix *string

	noSmithyDocumentSerde
}

// Definition of the key/value pair for a tag.
type Tag struct {

	// Key for the tag.
	//
	// This member is required.
	Key *string

	// Value for the tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// S3 Location of the training data.
type TrainingDataConfig struct {

	// The S3 URI where the training data is stored.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// Metrics associated with the custom job.
type TrainingMetrics struct {

	// Loss metric associated with the custom job.
	TrainingLoss *float32

	noSmithyDocumentSerde
}

// Array of up to 10 validators.
type ValidationDataConfig struct {

	// Information about the validators.
	//
	// This member is required.
	Validators []Validator

	noSmithyDocumentSerde
}

// Information about a validator.
type Validator struct {

	// The S3 URI where the validation data is stored.
	//
	// This member is required.
	S3Uri *string

	noSmithyDocumentSerde
}

// The metric for the validator.
type ValidatorMetric struct {

	// The validation loss associated with this validator.
	ValidationLoss *float32

	noSmithyDocumentSerde
}

// VPC configuration.
type VpcConfig struct {

	// VPC configuration security group Ids.
	//
	// This member is required.
	SecurityGroupIds []string

	// VPC configuration subnets.
	//
	// This member is required.
	SubnetIds []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
