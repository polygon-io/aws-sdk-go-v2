// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Defines an alternate key. This value is optional. A legacy data set might not
// have any alternate key defined but if those alternate keys definitions exist,
// provide them, as some applications will make use of them.
type AlternateKey struct {

	// A strictly positive integer value representing the length of the alternate key.
	//
	// This member is required.
	Length int32

	// A positive integer value representing the offset to mark the start of the
	// alternate key part in the record byte array.
	//
	// This member is required.
	Offset int32

	// Indicates whether the alternate key values are supposed to be unique for the
	// given data set.
	AllowDuplicates bool

	// The name of the alternate key.
	Name *string

	noSmithyDocumentSerde
}

// A subset of the possible application attributes. Used in the application list.
type ApplicationSummary struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationArn *string

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The version of the application.
	//
	// This member is required.
	ApplicationVersion *int32

	// The timestamp when the application was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The type of the target platform for this application.
	//
	// This member is required.
	EngineType EngineType

	// The name of the application.
	//
	// This member is required.
	Name *string

	// The status of the application.
	//
	// This member is required.
	Status ApplicationLifecycle

	// Indicates either an ongoing deployment or if the application has ever deployed
	// successfully.
	DeploymentStatus ApplicationDeploymentLifecycle

	// The description of the application.
	Description *string

	// The unique identifier of the runtime environment that hosts this application.
	EnvironmentId *string

	// The timestamp when you last started the application. Null until the application
	// runs for the first time.
	LastStartTime *time.Time

	// Indicates the status of the latest version of the application.
	VersionStatus ApplicationVersionLifecycle

	noSmithyDocumentSerde
}

// Defines an application version summary.
type ApplicationVersionSummary struct {

	// The application version.
	//
	// This member is required.
	ApplicationVersion *int32

	// The timestamp when the application version was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The status of the application.
	//
	// This member is required.
	Status ApplicationVersionLifecycle

	// The reason for the reported status.
	StatusReason *string

	noSmithyDocumentSerde
}

// Defines the details of a batch job.
//
// The following types satisfy this interface:
//
//	BatchJobDefinitionMemberFileBatchJobDefinition
//	BatchJobDefinitionMemberScriptBatchJobDefinition
type BatchJobDefinition interface {
	isBatchJobDefinition()
}

// Specifies a file containing a batch job definition.
type BatchJobDefinitionMemberFileBatchJobDefinition struct {
	Value FileBatchJobDefinition

	noSmithyDocumentSerde
}

func (*BatchJobDefinitionMemberFileBatchJobDefinition) isBatchJobDefinition() {}

// A script containing a batch job definition.
type BatchJobDefinitionMemberScriptBatchJobDefinition struct {
	Value ScriptBatchJobDefinition

	noSmithyDocumentSerde
}

func (*BatchJobDefinitionMemberScriptBatchJobDefinition) isBatchJobDefinition() {}

// A subset of the possible batch job attributes. Used in the batch job list.
type BatchJobExecutionSummary struct {

	// The unique identifier of the application that hosts this batch job.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier of this execution of the batch job.
	//
	// This member is required.
	ExecutionId *string

	// The timestamp when a particular batch job execution started.
	//
	// This member is required.
	StartTime *time.Time

	// The status of a particular batch job execution.
	//
	// This member is required.
	Status BatchJobExecutionStatus

	// Identifies a specific batch job.
	BatchJobIdentifier BatchJobIdentifier

	// The timestamp when this batch job execution ended.
	EndTime *time.Time

	// The unique identifier of a particular batch job.
	JobId *string

	// The name of a particular batch job.
	JobName *string

	// The type of a particular batch job execution.
	JobType BatchJobType

	//
	ReturnCode *string

	noSmithyDocumentSerde
}

// Identifies a specific batch job.
//
// The following types satisfy this interface:
//
//	BatchJobIdentifierMemberFileBatchJobIdentifier
//	BatchJobIdentifierMemberScriptBatchJobIdentifier
type BatchJobIdentifier interface {
	isBatchJobIdentifier()
}

// Specifies a file associated with a specific batch job.
type BatchJobIdentifierMemberFileBatchJobIdentifier struct {
	Value FileBatchJobIdentifier

	noSmithyDocumentSerde
}

func (*BatchJobIdentifierMemberFileBatchJobIdentifier) isBatchJobIdentifier() {}

// A batch job identifier in which the batch job to run is identified by the script
// name.
type BatchJobIdentifierMemberScriptBatchJobIdentifier struct {
	Value ScriptBatchJobIdentifier

	noSmithyDocumentSerde
}

func (*BatchJobIdentifierMemberScriptBatchJobIdentifier) isBatchJobIdentifier() {}

// Defines a data set.
type DataSet struct {

	// The logical identifier for a specific data set (in mainframe format).
	//
	// This member is required.
	DatasetName *string

	// The type of dataset. The only supported value is VSAM.
	//
	// This member is required.
	DatasetOrg DatasetOrgAttributes

	// The length of a record.
	//
	// This member is required.
	RecordLength *RecordLength

	// The relative location of the data set in the database or file system.
	RelativePath *string

	// The storage type of the data set: database or file system. For Micro Focus,
	// database corresponds to datastore and file system corresponds to EFS/FSX. For
	// Blu Age, there is no support of file system and database corresponds to Blusam.
	StorageType *string

	noSmithyDocumentSerde
}

// Additional details about the data set. Different attributes correspond to
// different data set organizations. The values are populated based on datasetOrg,
// storageType and backend (Blu Age or Micro Focus).
//
// The following types satisfy this interface:
//
//	DatasetDetailOrgAttributesMemberGdg
//	DatasetDetailOrgAttributesMemberVsam
type DatasetDetailOrgAttributes interface {
	isDatasetDetailOrgAttributes()
}

// The generation data group of the data set.
type DatasetDetailOrgAttributesMemberGdg struct {
	Value GdgDetailAttributes

	noSmithyDocumentSerde
}

func (*DatasetDetailOrgAttributesMemberGdg) isDatasetDetailOrgAttributes() {}

// The details of a VSAM data set.
type DatasetDetailOrgAttributesMemberVsam struct {
	Value VsamDetailAttributes

	noSmithyDocumentSerde
}

func (*DatasetDetailOrgAttributesMemberVsam) isDatasetDetailOrgAttributes() {}

// Identifies one or more data sets you want to import with the
// CreateDataSetImportTask operation.
//
// The following types satisfy this interface:
//
//	DataSetImportConfigMemberDataSets
//	DataSetImportConfigMemberS3Location
type DataSetImportConfig interface {
	isDataSetImportConfig()
}

// The data sets.
type DataSetImportConfigMemberDataSets struct {
	Value []DataSetImportItem

	noSmithyDocumentSerde
}

func (*DataSetImportConfigMemberDataSets) isDataSetImportConfig() {}

// The Amazon S3 location of the data sets.
type DataSetImportConfigMemberS3Location struct {
	Value string

	noSmithyDocumentSerde
}

func (*DataSetImportConfigMemberS3Location) isDataSetImportConfig() {}

// Identifies a specific data set to import from an external location.
type DataSetImportItem struct {

	// The data set.
	//
	// This member is required.
	DataSet *DataSet

	// The location of the data set.
	//
	// This member is required.
	ExternalLocation ExternalLocation

	noSmithyDocumentSerde
}

// Represents a summary of data set imports.
type DataSetImportSummary struct {

	// The number of data set imports that have failed.
	//
	// This member is required.
	Failed int32

	// The number of data set imports that are in progress.
	//
	// This member is required.
	InProgress int32

	// The number of data set imports that are pending.
	//
	// This member is required.
	Pending int32

	// The number of data set imports that have succeeded.
	//
	// This member is required.
	Succeeded int32

	// The total number of data set imports.
	//
	// This member is required.
	Total int32

	noSmithyDocumentSerde
}

// Contains information about a data set import task.
type DataSetImportTask struct {

	// The status of the data set import task.
	//
	// This member is required.
	Status DataSetTaskLifecycle

	// A summary of the data set import task.
	//
	// This member is required.
	Summary *DataSetImportSummary

	// The identifier of the data set import task.
	//
	// This member is required.
	TaskId *string

	noSmithyDocumentSerde
}

// Additional details about the data set. Different attributes correspond to
// different data set organizations. The values are populated based on datasetOrg,
// storageType and backend (Blu Age or Micro Focus).
//
// The following types satisfy this interface:
//
//	DatasetOrgAttributesMemberGdg
//	DatasetOrgAttributesMemberVsam
type DatasetOrgAttributes interface {
	isDatasetOrgAttributes()
}

// The generation data group of the data set.
type DatasetOrgAttributesMemberGdg struct {
	Value GdgAttributes

	noSmithyDocumentSerde
}

func (*DatasetOrgAttributesMemberGdg) isDatasetOrgAttributes() {}

// The details of a VSAM data set.
type DatasetOrgAttributesMemberVsam struct {
	Value VsamAttributes

	noSmithyDocumentSerde
}

func (*DatasetOrgAttributesMemberVsam) isDatasetOrgAttributes() {}

// A subset of the possible data set attributes.
type DataSetSummary struct {

	// The name of the data set.
	//
	// This member is required.
	DataSetName *string

	// The timestamp when the data set was created.
	CreationTime *time.Time

	// The type of data set. The only supported value is VSAM.
	DataSetOrg *string

	// The format of the data set.
	Format *string

	// The last time the data set was referenced.
	LastReferencedTime *time.Time

	// The last time the data set was updated.
	LastUpdatedTime *time.Time

	noSmithyDocumentSerde
}

// The application definition for a particular application.
//
// The following types satisfy this interface:
//
//	DefinitionMemberContent
//	DefinitionMemberS3Location
type Definition interface {
	isDefinition()
}

// The content of the application definition. This is a JSON object that contains
// the resource configuration/definitions that identify an application.
type DefinitionMemberContent struct {
	Value string

	noSmithyDocumentSerde
}

func (*DefinitionMemberContent) isDefinition() {}

// The S3 bucket that contains the application definition.
type DefinitionMemberS3Location struct {
	Value string

	noSmithyDocumentSerde
}

func (*DefinitionMemberS3Location) isDefinition() {}

// Contains a summary of a deployed application.
type DeployedVersionSummary struct {

	// The version of the deployed application.
	//
	// This member is required.
	ApplicationVersion *int32

	// The status of the deployment.
	//
	// This member is required.
	Status DeploymentLifecycle

	// The reason for the reported status.
	StatusReason *string

	noSmithyDocumentSerde
}

// A subset of information about a specific deployment.
type DeploymentSummary struct {

	// The unique identifier of the application.
	//
	// This member is required.
	ApplicationId *string

	// The version of the application.
	//
	// This member is required.
	ApplicationVersion *int32

	// The timestamp when the deployment was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The unique identifier of the deployment.
	//
	// This member is required.
	DeploymentId *string

	// The unique identifier of the runtime environment.
	//
	// This member is required.
	EnvironmentId *string

	// The current status of the deployment.
	//
	// This member is required.
	Status DeploymentLifecycle

	// The reason for the reported status.
	StatusReason *string

	noSmithyDocumentSerde
}

// Defines the storage configuration for an Amazon EFS file system.
type EfsStorageConfiguration struct {

	// The file system identifier.
	//
	// This member is required.
	FileSystemId *string

	// The mount point for the file system.
	//
	// This member is required.
	MountPoint *string

	noSmithyDocumentSerde
}

// A subset of information about the engine version for a specific application.
type EngineVersionsSummary struct {

	// The type of target platform for the application.
	//
	// This member is required.
	EngineType *string

	// The version of the engine type used by the application.
	//
	// This member is required.
	EngineVersion *string

	noSmithyDocumentSerde
}

// Contains a subset of the possible runtime environment attributes. Used in the
// environment list.
type EnvironmentSummary struct {

	// The timestamp when the runtime environment was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The target platform for the runtime environment.
	//
	// This member is required.
	EngineType EngineType

	// The version of the runtime engine.
	//
	// This member is required.
	EngineVersion *string

	// The Amazon Resource Name (ARN) of a particular runtime environment.
	//
	// This member is required.
	EnvironmentArn *string

	// The unique identifier of a particular runtime environment.
	//
	// This member is required.
	EnvironmentId *string

	// The instance type of the runtime environment.
	//
	// This member is required.
	InstanceType *string

	// The name of the runtime environment.
	//
	// This member is required.
	Name *string

	// The status of the runtime environment
	//
	// This member is required.
	Status EnvironmentLifecycle

	noSmithyDocumentSerde
}

// Defines an external storage location.
//
// The following types satisfy this interface:
//
//	ExternalLocationMemberS3Location
type ExternalLocation interface {
	isExternalLocation()
}

// The URI of the Amazon S3 bucket.
type ExternalLocationMemberS3Location struct {
	Value string

	noSmithyDocumentSerde
}

func (*ExternalLocationMemberS3Location) isExternalLocation() {}

// A file containing a batch job definition.
type FileBatchJobDefinition struct {

	// The name of the file containing the batch job definition.
	//
	// This member is required.
	FileName *string

	// The path to the file containing the batch job definition.
	FolderPath *string

	noSmithyDocumentSerde
}

// A batch job identifier in which the batch job to run is identified by the file
// name and the relative path to the file name.
type FileBatchJobIdentifier struct {

	// The file name for the batch job identifier.
	//
	// This member is required.
	FileName *string

	// The relative path to the file name for the batch job identifier.
	FolderPath *string

	noSmithyDocumentSerde
}

// Defines the storage configuration for an Amazon FSx file system.
type FsxStorageConfiguration struct {

	// The file system identifier.
	//
	// This member is required.
	FileSystemId *string

	// The mount point for the file system.
	//
	// This member is required.
	MountPoint *string

	noSmithyDocumentSerde
}

// The required attributes for a generation data group data set. A generation data
// set is one of a collection of successive, historically related, catalogued data
// sets that together are known as a generation data group (GDG). Use this
// structure when you want to import a GDG. For more information on GDG, see
// Generation data sets
// (https://www.ibm.com/docs/en/zos/2.3.0?topic=guide-generation-data-sets).
type GdgAttributes struct {

	// The maximum number of generation data sets, up to 255, in a GDG.
	Limit int32

	// The disposition of the data set in the catalog.
	RollDisposition *string

	noSmithyDocumentSerde
}

// The required attributes for a generation data group data set. A generation data
// set is one of a collection of successive, historically related, catalogued data
// sets that together are known as a generation data group (GDG). Use this
// structure when you want to import a GDG. For more information on GDG, see
// Generation data sets
// (https://www.ibm.com/docs/en/zos/2.3.0?topic=guide-generation-data-sets).
type GdgDetailAttributes struct {

	// The maximum number of generation data sets, up to 255, in a GDG.
	Limit int32

	// The disposition of the data set in the catalog.
	RollDisposition *string

	noSmithyDocumentSerde
}

// Defines the details of a high availability configuration.
type HighAvailabilityConfig struct {

	// The number of instances in a high availability configuration.
	//
	// This member is required.
	DesiredCapacity *int32

	noSmithyDocumentSerde
}

// A subset of the attributes that describe a log group. In CloudWatch a log group
// is a group of log streams that share the same retention, monitoring, and access
// control settings.
type LogGroupSummary struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The type of log.
	//
	// This member is required.
	LogType *string

	noSmithyDocumentSerde
}

// The information about the maintenance schedule.
type MaintenanceSchedule struct {

	// The time the scheduled maintenance is to end.
	EndTime *time.Time

	// The time the scheduled maintenance is to start.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// The scheduled maintenance for a runtime engine.
type PendingMaintenance struct {

	// The specific runtime engine that the maintenance schedule applies to.
	EngineVersion *string

	// The maintenance schedule for the runtime engine version.
	Schedule *MaintenanceSchedule

	noSmithyDocumentSerde
}

// The primary key for a KSDS data set.
type PrimaryKey struct {

	// A strictly positive integer value representing the length of the primary key.
	//
	// This member is required.
	Length int32

	// A positive integer value representing the offset to mark the start of the
	// primary key in the record byte array.
	//
	// This member is required.
	Offset int32

	// A name for the Primary Key.
	Name *string

	noSmithyDocumentSerde
}

// The length of the records in the data set.
type RecordLength struct {

	// The maximum record length. In case of fixed, both minimum and maximum are the
	// same.
	//
	// This member is required.
	Max int32

	// The minimum record length of a record.
	//
	// This member is required.
	Min int32

	noSmithyDocumentSerde
}

// A batch job definition contained in a script.
type ScriptBatchJobDefinition struct {

	// The name of the script containing the batch job definition.
	//
	// This member is required.
	ScriptName *string

	noSmithyDocumentSerde
}

// A batch job identifier in which the batch job to run is identified by the script
// name.
type ScriptBatchJobIdentifier struct {

	// The name of the script containing the batch job definition.
	//
	// This member is required.
	ScriptName *string

	noSmithyDocumentSerde
}

// Defines the storage configuration for a runtime environment.
//
// The following types satisfy this interface:
//
//	StorageConfigurationMemberEfs
//	StorageConfigurationMemberFsx
type StorageConfiguration interface {
	isStorageConfiguration()
}

// Defines the storage configuration for an Amazon EFS file system.
type StorageConfigurationMemberEfs struct {
	Value EfsStorageConfiguration

	noSmithyDocumentSerde
}

func (*StorageConfigurationMemberEfs) isStorageConfiguration() {}

// Defines the storage configuration for an Amazon FSx file system.
type StorageConfigurationMemberFsx struct {
	Value FsxStorageConfiguration

	noSmithyDocumentSerde
}

func (*StorageConfigurationMemberFsx) isStorageConfiguration() {}

// Contains information about a validation exception field.
type ValidationExceptionField struct {

	// The message of the exception field.
	//
	// This member is required.
	Message *string

	// The name of the exception field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// The attributes of a VSAM type data set.
type VsamAttributes struct {

	// The record format of the data set.
	//
	// This member is required.
	Format *string

	// The alternate key definitions, if any. A legacy dataset might not have any
	// alternate key defined, but if those alternate keys definitions exist, provide
	// them as some applications will make use of them.
	AlternateKeys []AlternateKey

	// Indicates whether indexes for this dataset are stored as compressed values. If
	// you have a large data set (typically > 100 Mb), consider setting this flag to
	// True.
	Compressed bool

	// The character set used by the data set. Can be ASCII, EBCDIC, or unknown.
	Encoding *string

	// The primary key of the data set.
	PrimaryKey *PrimaryKey

	noSmithyDocumentSerde
}

// The attributes of a VSAM type data set.
type VsamDetailAttributes struct {

	// The alternate key definitions, if any. A legacy dataset might not have any
	// alternate key defined, but if those alternate keys definitions exist, provide
	// them as some applications will make use of them.
	AlternateKeys []AlternateKey

	// If set to True, enforces loading the data set into cache before it’s used by the
	// application.
	CacheAtStartup *bool

	// Indicates whether indexes for this dataset are stored as compressed values. If
	// you have a large data set (typically > 100 Mb), consider setting this flag to
	// True.
	Compressed *bool

	// The character set used by the data set. Can be ASCII, EBCDIC, or unknown.
	Encoding *string

	// The primary key of the data set.
	PrimaryKey *PrimaryKey

	// The record format of the data set.
	RecordFormat *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isBatchJobDefinition()         {}
func (*UnknownUnionMember) isBatchJobIdentifier()         {}
func (*UnknownUnionMember) isDatasetDetailOrgAttributes() {}
func (*UnknownUnionMember) isDataSetImportConfig()        {}
func (*UnknownUnionMember) isDatasetOrgAttributes()       {}
func (*UnknownUnionMember) isDefinition()                 {}
func (*UnknownUnionMember) isExternalLocation()           {}
func (*UnknownUnionMember) isStorageConfiguration()       {}
