// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about an Availability Zone.
type AvailabilityZone struct {

	// The name of the Availability Zone.
	Name *string

	noSmithyDocumentSerde
}

// A certificate authority (CA) certificate for an Amazon Web Services account.
type Certificate struct {

	// The Amazon Resource Name (ARN) for the certificate. Example:
	// arn:aws:rds:us-east-1::cert:rds-ca-2019
	CertificateArn *string

	// The unique key that identifies a certificate. Example: rds-ca-2019
	CertificateIdentifier *string

	// The type of the certificate. Example: CA
	CertificateType *string

	// The thumbprint of the certificate.
	Thumbprint *string

	// The starting date-time from which the certificate is valid. Example:
	// 2019-07-31T17:57:09Z
	ValidFrom *time.Time

	// The date-time after which the certificate is no longer valid. Example:
	// 2024-07-31T17:57:09Z
	ValidTill *time.Time

	noSmithyDocumentSerde
}

// Returns the details of the DB instance’s server certificate. For more
// information, see Updating Your Amazon DocumentDB TLS Certificates (https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html)
// and Encrypting Data in Transit (https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html)
// in the Amazon DocumentDB Developer Guide.
type CertificateDetails struct {

	// The CA identifier of the CA certificate used for the DB instance's server
	// certificate.
	CAIdentifier *string

	// The expiration date of the DB instance’s server certificate.
	ValidTill *time.Time

	noSmithyDocumentSerde
}

// The configuration setting for the log types to be enabled for export to Amazon
// CloudWatch Logs for a specific instance or cluster. The EnableLogTypes and
// DisableLogTypes arrays determine which logs are exported (or not exported) to
// CloudWatch Logs. The values within these arrays depend on the engine that is
// being used.
type CloudwatchLogsExportConfiguration struct {

	// The list of log types to disable.
	DisableLogTypes []string

	// The list of log types to enable.
	EnableLogTypes []string

	noSmithyDocumentSerde
}

// Detailed information about a cluster.
type DBCluster struct {

	// Provides a list of the Identity and Access Management (IAM) roles that are
	// associated with the cluster. (IAM) roles that are associated with a cluster
	// grant permission for the cluster to access other Amazon Web Services services on
	// your behalf.
	AssociatedRoles []DBClusterRole

	// Provides the list of Amazon EC2 Availability Zones that instances in the
	// cluster can be created in.
	AvailabilityZones []string

	// Specifies the number of days for which automatic snapshots are retained.
	BackupRetentionPeriod *int32

	// Identifies the clone group to which the DB cluster is associated.
	CloneGroupId *string

	// Specifies the time when the cluster was created, in Universal Coordinated Time
	// (UTC).
	ClusterCreateTime *time.Time

	// The Amazon Resource Name (ARN) for the cluster.
	DBClusterArn *string

	// Contains a user-supplied cluster identifier. This identifier is the unique key
	// that identifies a cluster.
	DBClusterIdentifier *string

	// Provides the list of instances that make up the cluster.
	DBClusterMembers []DBClusterMember

	// Specifies the name of the cluster parameter group for the cluster.
	DBClusterParameterGroup *string

	// Specifies information on the subnet group that is associated with the cluster,
	// including the name, description, and subnets in the subnet group.
	DBSubnetGroup *string

	// The Amazon Web Services Region-unique, immutable identifier for the cluster.
	// This identifier is found in CloudTrail log entries whenever the KMS key for the
	// cluster is accessed.
	DbClusterResourceId *string

	// Specifies whether this cluster can be deleted. If DeletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// DeletionProtection is disabled. DeletionProtection protects clusters from being
	// accidentally deleted.
	DeletionProtection *bool

	// The earliest time to which a database can be restored with point-in-time
	// restore.
	EarliestRestorableTime *time.Time

	// A list of log types that this cluster is configured to export to Amazon
	// CloudWatch Logs.
	EnabledCloudwatchLogsExports []string

	// Specifies the connection endpoint for the primary instance of the cluster.
	Endpoint *string

	// Provides the name of the database engine to be used for this cluster.
	Engine *string

	// Indicates the database engine version.
	EngineVersion *string

	// Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
	HostedZoneId *string

	// If StorageEncrypted is true , the KMS key identifier for the encrypted cluster.
	KmsKeyId *string

	// Specifies the latest time to which a database can be restored with
	// point-in-time restore.
	LatestRestorableTime *time.Time

	// Contains the master user name for the cluster.
	MasterUsername *string

	// Specifies whether the cluster has instances in multiple Availability Zones.
	MultiAZ *bool

	// Specifies the progress of the operation as a percentage.
	PercentProgress *string

	// Specifies the port that the database engine is listening on.
	Port *int32

	// Specifies the daily time range during which automated backups are created if
	// automated backups are enabled, as determined by the BackupRetentionPeriod .
	PreferredBackupWindow *string

	// Specifies the weekly time range during which system maintenance can occur, in
	// Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string

	// Contains one or more identifiers of the secondary clusters that are associated
	// with this cluster.
	ReadReplicaIdentifiers []string

	// The reader endpoint for the cluster. The reader endpoint for a cluster load
	// balances connections across the Amazon DocumentDB replicas that are available in
	// a cluster. As clients request new connections to the reader endpoint, Amazon
	// DocumentDB distributes the connection requests among the Amazon DocumentDB
	// replicas in the cluster. This functionality can help balance your read workload
	// across multiple Amazon DocumentDB replicas in your cluster. If a failover
	// occurs, and the Amazon DocumentDB replica that you are connected to is promoted
	// to be the primary instance, your connection is dropped. To continue sending your
	// read workload to other Amazon DocumentDB replicas in the cluster, you can then
	// reconnect to the reader endpoint.
	ReaderEndpoint *string

	// Contains the identifier of the source cluster if this cluster is a secondary
	// cluster.
	ReplicationSourceIdentifier *string

	// Specifies the current state of this cluster.
	Status *string

	// Specifies whether the cluster is encrypted.
	StorageEncrypted *bool

	// Storage type associated with your cluster Storage type associated with your
	// cluster For information on storage types for Amazon DocumentDB clusters, see
	// Cluster storage configurations in the Amazon DocumentDB Developer Guide. Valid
	// values for storage type - standard | iopt1 Default value is standard
	StorageType *string

	// Provides a list of virtual private cloud (VPC) security groups that the cluster
	// belongs to.
	VpcSecurityGroups []VpcSecurityGroupMembership

	noSmithyDocumentSerde
}

// Contains information about an instance that is part of a cluster.
type DBClusterMember struct {

	// Specifies the status of the cluster parameter group for this member of the DB
	// cluster.
	DBClusterParameterGroupStatus *string

	// Specifies the instance identifier for this member of the cluster.
	DBInstanceIdentifier *string

	// A value that is true if the cluster member is the primary instance for the
	// cluster and false otherwise.
	IsClusterWriter *bool

	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance.
	PromotionTier *int32

	noSmithyDocumentSerde
}

// Detailed information about a cluster parameter group.
type DBClusterParameterGroup struct {

	// The Amazon Resource Name (ARN) for the cluster parameter group.
	DBClusterParameterGroupArn *string

	// Provides the name of the cluster parameter group.
	DBClusterParameterGroupName *string

	// Provides the name of the parameter group family that this cluster parameter
	// group is compatible with.
	DBParameterGroupFamily *string

	// Provides the customer-specified description for this cluster parameter group.
	Description *string

	noSmithyDocumentSerde
}

// Describes an Identity and Access Management (IAM) role that is associated with
// a cluster.
type DBClusterRole struct {

	// The Amazon Resource Name (ARN) of the IAMrole that is associated with the DB
	// cluster.
	RoleArn *string

	// Describes the state of association between the IAMrole and the cluster. The
	// Status property returns one of the following values:
	//   - ACTIVE - The IAMrole ARN is associated with the cluster and can be used to
	//   access other Amazon Web Services services on your behalf.
	//   - PENDING - The IAMrole ARN is being associated with the cluster.
	//   - INVALID - The IAMrole ARN is associated with the cluster, but the cluster
	//   cannot assume the IAMrole to access other Amazon Web Services services on your
	//   behalf.
	Status *string

	noSmithyDocumentSerde
}

// Detailed information about a cluster snapshot.
type DBClusterSnapshot struct {

	// Provides the list of Amazon EC2 Availability Zones that instances in the
	// cluster snapshot can be restored in.
	AvailabilityZones []string

	// Specifies the time when the cluster was created, in Universal Coordinated Time
	// (UTC).
	ClusterCreateTime *time.Time

	// Specifies the cluster identifier of the cluster that this cluster snapshot was
	// created from.
	DBClusterIdentifier *string

	// The Amazon Resource Name (ARN) for the cluster snapshot.
	DBClusterSnapshotArn *string

	// Specifies the identifier for the cluster snapshot.
	DBClusterSnapshotIdentifier *string

	// Specifies the name of the database engine.
	Engine *string

	// Provides the version of the database engine for this cluster snapshot.
	EngineVersion *string

	// If StorageEncrypted is true , the KMS key identifier for the encrypted cluster
	// snapshot.
	KmsKeyId *string

	// Provides the master user name for the cluster snapshot.
	MasterUsername *string

	// Specifies the percentage of the estimated data that has been transferred.
	PercentProgress *int32

	// Specifies the port that the cluster was listening on at the time of the
	// snapshot.
	Port *int32

	// Provides the time when the snapshot was taken, in UTC.
	SnapshotCreateTime *time.Time

	// Provides the type of the cluster snapshot.
	SnapshotType *string

	// If the cluster snapshot was copied from a source cluster snapshot, the ARN for
	// the source cluster snapshot; otherwise, a null value.
	SourceDBClusterSnapshotArn *string

	// Specifies the status of this cluster snapshot.
	Status *string

	// Specifies whether the cluster snapshot is encrypted.
	StorageEncrypted *bool

	// Storage type associated with your cluster snapshot For information on storage
	// types for Amazon DocumentDB clusters, see Cluster storage configurations in the
	// Amazon DocumentDB Developer Guide. Valid values for storage type - standard |
	// iopt1 Default value is standard
	StorageType *string

	// Provides the virtual private cloud (VPC) ID that is associated with the cluster
	// snapshot.
	VpcId *string

	noSmithyDocumentSerde
}

// Contains the name and values of a manual cluster snapshot attribute. Manual
// cluster snapshot attributes are used to authorize other Amazon Web Services
// accounts to restore a manual cluster snapshot.
type DBClusterSnapshotAttribute struct {

	// The name of the manual cluster snapshot attribute. The attribute named restore
	// refers to the list of Amazon Web Services accounts that have permission to copy
	// or restore the manual cluster snapshot.
	AttributeName *string

	// The values for the manual cluster snapshot attribute. If the AttributeName
	// field is set to restore , then this element returns a list of IDs of the Amazon
	// Web Services accounts that are authorized to copy or restore the manual cluster
	// snapshot. If a value of all is in the list, then the manual cluster snapshot is
	// public and available for any Amazon Web Services account to copy or restore.
	AttributeValues []string

	noSmithyDocumentSerde
}

// Detailed information about the attributes that are associated with a cluster
// snapshot.
type DBClusterSnapshotAttributesResult struct {

	// The list of attributes and values for the cluster snapshot.
	DBClusterSnapshotAttributes []DBClusterSnapshotAttribute

	// The identifier of the cluster snapshot that the attributes apply to.
	DBClusterSnapshotIdentifier *string

	noSmithyDocumentSerde
}

// Detailed information about an engine version.
type DBEngineVersion struct {

	// The description of the database engine.
	DBEngineDescription *string

	// The description of the database engine version.
	DBEngineVersionDescription *string

	// The name of the parameter group family for the database engine.
	DBParameterGroupFamily *string

	// The name of the database engine.
	Engine *string

	// The version number of the database engine.
	EngineVersion *string

	// The types of logs that the database engine has available for export to Amazon
	// CloudWatch Logs.
	ExportableLogTypes []string

	// A list of the supported CA certificate identifiers. For more information, see
	// Updating Your Amazon DocumentDB TLS Certificates (https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html)
	// and Encrypting Data in Transit (https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html)
	// in the Amazon DocumentDB Developer Guide.
	SupportedCACertificateIdentifiers []string

	// Indicates whether the engine version supports rotating the server certificate
	// without rebooting the DB instance.
	SupportsCertificateRotationWithoutRestart *bool

	// A value that indicates whether the engine version supports exporting the log
	// types specified by ExportableLogTypes to CloudWatch Logs.
	SupportsLogExportsToCloudwatchLogs *bool

	// A list of engine versions that this database engine version can be upgraded to.
	ValidUpgradeTarget []UpgradeTarget

	noSmithyDocumentSerde
}

// Detailed information about an instance.
type DBInstance struct {

	// Does not apply. This parameter does not apply to Amazon DocumentDB. Amazon
	// DocumentDB does not perform minor version upgrades regardless of the value set.
	AutoMinorVersionUpgrade *bool

	// Specifies the name of the Availability Zone that the instance is located in.
	AvailabilityZone *string

	// Specifies the number of days for which automatic snapshots are retained.
	BackupRetentionPeriod *int32

	// The identifier of the CA certificate for this DB instance.
	CACertificateIdentifier *string

	// The details of the DB instance's server certificate.
	CertificateDetails *CertificateDetails

	// A value that indicates whether to copy tags from the DB instance to snapshots
	// of the DB instance. By default, tags are not copied.
	CopyTagsToSnapshot *bool

	// Contains the name of the cluster that the instance is a member of if the
	// instance is a member of a cluster.
	DBClusterIdentifier *string

	// The Amazon Resource Name (ARN) for the instance.
	DBInstanceArn *string

	// Contains the name of the compute and memory capacity class of the instance.
	DBInstanceClass *string

	// Contains a user-provided database identifier. This identifier is the unique key
	// that identifies an instance.
	DBInstanceIdentifier *string

	// Specifies the current state of this database.
	DBInstanceStatus *string

	// Specifies information on the subnet group that is associated with the instance,
	// including the name, description, and subnets in the subnet group.
	DBSubnetGroup *DBSubnetGroup

	// The Amazon Web Services Region-unique, immutable identifier for the instance.
	// This identifier is found in CloudTrail log entries whenever the KMS key for the
	// instance is accessed.
	DbiResourceId *string

	// A list of log types that this instance is configured to export to CloudWatch
	// Logs.
	EnabledCloudwatchLogsExports []string

	// Specifies the connection endpoint.
	Endpoint *Endpoint

	// Provides the name of the database engine to be used for this instance.
	Engine *string

	// Indicates the database engine version.
	EngineVersion *string

	// Provides the date and time that the instance was created.
	InstanceCreateTime *time.Time

	// If StorageEncrypted is true , the KMS key identifier for the encrypted instance.
	KmsKeyId *string

	// Specifies the latest time to which a database can be restored with
	// point-in-time restore.
	LatestRestorableTime *time.Time

	// Specifies that changes to the instance are pending. This element is included
	// only when changes are pending. Specific changes are identified by subelements.
	PendingModifiedValues *PendingModifiedValues

	// Set to true if Amazon RDS Performance Insights is enabled for the DB instance,
	// and otherwise false .
	PerformanceInsightsEnabled *bool

	// The KMS key identifier for encryption of Performance Insights data. The KMS key
	// ID is the Amazon Resource Name (ARN), KMS key identifier, or the KMS key alias
	// for the KMS encryption key.
	PerformanceInsightsKMSKeyId *string

	// Specifies the daily time range during which automated backups are created if
	// automated backups are enabled, as determined by the BackupRetentionPeriod .
	PreferredBackupWindow *string

	// Specifies the weekly time range during which system maintenance can occur, in
	// Universal Coordinated Time (UTC).
	PreferredMaintenanceWindow *string

	// A value that specifies the order in which an Amazon DocumentDB replica is
	// promoted to the primary instance after a failure of the existing primary
	// instance.
	PromotionTier *int32

	// Not supported. Amazon DocumentDB does not currently support public endpoints.
	// The value of PubliclyAccessible is always false .
	PubliclyAccessible *bool

	// The status of a read replica. If the instance is not a read replica, this is
	// blank.
	StatusInfos []DBInstanceStatusInfo

	// Specifies whether or not the instance is encrypted.
	StorageEncrypted *bool

	// Provides a list of VPC security group elements that the instance belongs to.
	VpcSecurityGroups []VpcSecurityGroupMembership

	noSmithyDocumentSerde
}

// Provides a list of status information for an instance.
type DBInstanceStatusInfo struct {

	// Details of the error if there is an error for the instance. If the instance is
	// not in an error state, this value is blank.
	Message *string

	// A Boolean value that is true if the instance is operating normally, or false if
	// the instance is in an error state.
	Normal *bool

	// Status of the instance. For a StatusType of read replica, the values can be
	// replicating , error, stopped , or terminated .
	Status *string

	// This value is currently " read replication ."
	StatusType *string

	noSmithyDocumentSerde
}

// Detailed information about a subnet group.
type DBSubnetGroup struct {

	// The Amazon Resource Name (ARN) for the DB subnet group.
	DBSubnetGroupArn *string

	// Provides the description of the subnet group.
	DBSubnetGroupDescription *string

	// The name of the subnet group.
	DBSubnetGroupName *string

	// Provides the status of the subnet group.
	SubnetGroupStatus *string

	// Detailed information about one or more subnets within a subnet group.
	Subnets []Subnet

	// Provides the virtual private cloud (VPC) ID of the subnet group.
	VpcId *string

	noSmithyDocumentSerde
}

// Network information for accessing a cluster or instance. Client programs must
// specify a valid endpoint to access these Amazon DocumentDB resources.
type Endpoint struct {

	// Specifies the DNS address of the instance.
	Address *string

	// Specifies the ID that Amazon Route 53 assigns when you create a hosted zone.
	HostedZoneId *string

	// Specifies the port that the database engine is listening on.
	Port *int32

	noSmithyDocumentSerde
}

// Contains the result of a successful invocation of the
// DescribeEngineDefaultClusterParameters operation.
type EngineDefaults struct {

	// The name of the cluster parameter group family to return the engine parameter
	// information for.
	DBParameterGroupFamily *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The parameters of a particular cluster parameter group family.
	Parameters []Parameter

	noSmithyDocumentSerde
}

// Detailed information about an event.
type Event struct {

	// Specifies the date and time of the event.
	Date *time.Time

	// Specifies the category for the event.
	EventCategories []string

	// Provides the text of this event.
	Message *string

	// The Amazon Resource Name (ARN) for the event.
	SourceArn *string

	// Provides the identifier for the source of the event.
	SourceIdentifier *string

	// Specifies the source type for this event.
	SourceType SourceType

	noSmithyDocumentSerde
}

// An event source type, accompanied by one or more event category names.
type EventCategoriesMap struct {

	// The event categories for the specified source type.
	EventCategories []string

	// The source type that the returned categories belong to.
	SourceType *string

	noSmithyDocumentSerde
}

// Detailed information about an event to which you have subscribed.
type EventSubscription struct {

	// The Amazon DocumentDB event notification subscription ID.
	CustSubscriptionId *string

	// The Amazon Web Services customer account that is associated with the Amazon
	// DocumentDB event notification subscription.
	CustomerAwsId *string

	// A Boolean value indicating whether the subscription is enabled. A value of true
	// indicates that the subscription is enabled.
	Enabled *bool

	// A list of event categories for the Amazon DocumentDB event notification
	// subscription.
	EventCategoriesList []string

	// The Amazon Resource Name (ARN) for the event subscription.
	EventSubscriptionArn *string

	// The topic ARN of the Amazon DocumentDB event notification subscription.
	SnsTopicArn *string

	// A list of source IDs for the Amazon DocumentDB event notification subscription.
	SourceIdsList []string

	// The source type for the Amazon DocumentDB event notification subscription.
	SourceType *string

	// The status of the Amazon DocumentDB event notification subscription.
	// Constraints: Can be one of the following: creating , modifying , deleting ,
	// active , no-permission , topic-not-exist The no-permission status indicates
	// that Amazon DocumentDB no longer has permission to post to the SNS topic. The
	// topic-not-exist status indicates that the topic was deleted after the
	// subscription was created.
	Status *string

	// The time at which the Amazon DocumentDB event notification subscription was
	// created.
	SubscriptionCreationTime *string

	noSmithyDocumentSerde
}

// A named set of filter values, used to return a more specific list of results.
// You can use a filter to match a set of resources by specific criteria, such as
// IDs. Wildcards are not supported in filters.
type Filter struct {

	// The name of the filter. Filter names are case sensitive.
	//
	// This member is required.
	Name *string

	// One or more filter values. Filter values are case sensitive.
	//
	// This member is required.
	Values []string

	noSmithyDocumentSerde
}

// A data type representing an Amazon DocumentDB global cluster.
type GlobalCluster struct {

	// The default database name within the new global cluster.
	DatabaseName *string

	// The deletion protection setting for the new global cluster.
	DeletionProtection *bool

	// The Amazon DocumentDB database engine used by the global cluster.
	Engine *string

	// Indicates the database engine version.
	EngineVersion *string

	// The Amazon Resource Name (ARN) for the global cluster.
	GlobalClusterArn *string

	// Contains a user-supplied global cluster identifier. This identifier is the
	// unique key that identifies a global cluster.
	GlobalClusterIdentifier *string

	// The list of cluster IDs for secondary clusters within the global cluster.
	// Currently limited to one item.
	GlobalClusterMembers []GlobalClusterMember

	// The Amazon Web Services Region-unique, immutable identifier for the global
	// database cluster. This identifier is found in CloudTrail log entries whenever
	// the KMS customer master key (CMK) for the cluster is accessed.
	GlobalClusterResourceId *string

	// Specifies the current state of this global cluster.
	Status *string

	// The storage encryption setting for the global cluster.
	StorageEncrypted *bool

	noSmithyDocumentSerde
}

// A data structure with information about any primary and secondary clusters
// associated with an Amazon DocumentDB global clusters.
type GlobalClusterMember struct {

	// The Amazon Resource Name (ARN) for each Amazon DocumentDB cluster.
	DBClusterArn *string

	// Specifies whether the Amazon DocumentDB cluster is the primary cluster (that
	// is, has read-write capability) for the Amazon DocumentDB global cluster with
	// which it is associated.
	IsWriter *bool

	// The Amazon Resource Name (ARN) for each read-only secondary cluster associated
	// with the Aurora global cluster.
	Readers []string

	noSmithyDocumentSerde
}

// The options that are available for an instance.
type OrderableDBInstanceOption struct {

	// A list of Availability Zones for an instance.
	AvailabilityZones []AvailabilityZone

	// The instance class for an instance.
	DBInstanceClass *string

	// The engine type of an instance.
	Engine *string

	// The engine version of an instance.
	EngineVersion *string

	// The license model for an instance.
	LicenseModel *string

	// The storage type to associate with the DB cluster
	StorageType *string

	// Indicates whether an instance is in a virtual private cloud (VPC).
	Vpc *bool

	noSmithyDocumentSerde
}

// Detailed information about an individual parameter.
type Parameter struct {

	// Specifies the valid range of values for the parameter.
	AllowedValues *string

	// Indicates when to apply parameter updates.
	ApplyMethod ApplyMethod

	// Specifies the engine-specific parameters type.
	ApplyType *string

	// Specifies the valid data type for the parameter.
	DataType *string

	// Provides a description of the parameter.
	Description *string

	// Indicates whether ( true ) or not ( false ) the parameter can be modified. Some
	// parameters have security or operational implications that prevent them from
	// being changed.
	IsModifiable *bool

	// The earliest engine version to which the parameter can apply.
	MinimumEngineVersion *string

	// Specifies the name of the parameter.
	ParameterName *string

	// Specifies the value of the parameter.
	ParameterValue *string

	// Indicates the source of the parameter value.
	Source *string

	noSmithyDocumentSerde
}

// A list of the log types whose configuration is still pending. These log types
// are in the process of being activated or deactivated.
type PendingCloudwatchLogsExports struct {

	// Log types that are in the process of being enabled. After they are enabled,
	// these log types are exported to Amazon CloudWatch Logs.
	LogTypesToDisable []string

	// Log types that are in the process of being deactivated. After they are
	// deactivated, these log types aren't exported to CloudWatch Logs.
	LogTypesToEnable []string

	noSmithyDocumentSerde
}

// Provides information about a pending maintenance action for a resource.
type PendingMaintenanceAction struct {

	// The type of pending maintenance action that is available for the resource.
	Action *string

	// The date of the maintenance window when the action is applied. The maintenance
	// action is applied to the resource during its first maintenance window after this
	// date. If this date is specified, any next-maintenance opt-in requests are
	// ignored.
	AutoAppliedAfterDate *time.Time

	// The effective date when the pending maintenance action is applied to the
	// resource.
	CurrentApplyDate *time.Time

	// A description providing more detail about the maintenance action.
	Description *string

	// The date when the maintenance action is automatically applied. The maintenance
	// action is applied to the resource on this date regardless of the maintenance
	// window for the resource. If this date is specified, any immediate opt-in
	// requests are ignored.
	ForcedApplyDate *time.Time

	// Indicates the type of opt-in request that has been received for the resource.
	OptInStatus *string

	noSmithyDocumentSerde
}

// One or more modified settings for an instance. These modified settings have
// been requested, but haven't been applied yet.
type PendingModifiedValues struct {

	// Contains the new AllocatedStorage size for then instance that will be applied
	// or is currently being applied.
	AllocatedStorage *int32

	// Specifies the pending number of days for which automated backups are retained.
	BackupRetentionPeriod *int32

	// Specifies the identifier of the certificate authority (CA) certificate for the
	// DB instance.
	CACertificateIdentifier *string

	// Contains the new DBInstanceClass for the instance that will be applied or is
	// currently being applied.
	DBInstanceClass *string

	// Contains the new DBInstanceIdentifier for the instance that will be applied or
	// is currently being applied.
	DBInstanceIdentifier *string

	// The new subnet group for the instance.
	DBSubnetGroupName *string

	// Indicates the database engine version.
	EngineVersion *string

	// Specifies the new Provisioned IOPS value for the instance that will be applied
	// or is currently being applied.
	Iops *int32

	// The license model for the instance. Valid values: license-included ,
	// bring-your-own-license , general-public-license
	LicenseModel *string

	// Contains the pending or currently in-progress change of the master credentials
	// for the instance.
	MasterUserPassword *string

	// Indicates that the Single-AZ instance is to change to a Multi-AZ deployment.
	MultiAZ *bool

	// A list of the log types whose configuration is still pending. These log types
	// are in the process of being activated or deactivated.
	PendingCloudwatchLogsExports *PendingCloudwatchLogsExports

	// Specifies the pending port for the instance.
	Port *int32

	// Specifies the storage type to be associated with the instance.
	StorageType *string

	noSmithyDocumentSerde
}

// Represents the output of ApplyPendingMaintenanceAction .
type ResourcePendingMaintenanceActions struct {

	// A list that provides details about the pending maintenance actions for the
	// resource.
	PendingMaintenanceActionDetails []PendingMaintenanceAction

	// The Amazon Resource Name (ARN) of the resource that has pending maintenance
	// actions.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

// Detailed information about a subnet.
type Subnet struct {

	// Specifies the Availability Zone for the subnet.
	SubnetAvailabilityZone *AvailabilityZone

	// Specifies the identifier of the subnet.
	SubnetIdentifier *string

	// Specifies the status of the subnet.
	SubnetStatus *string

	noSmithyDocumentSerde
}

// Metadata assigned to an Amazon DocumentDB resource consisting of a key-value
// pair.
type Tag struct {

	// The required name of the tag. The string value can be from 1 to 128 Unicode
	// characters in length and can't be prefixed with " aws: " or " rds: ". The string
	// can contain only the set of Unicode letters, digits, white space, '_', '.', '/',
	// '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Key *string

	// The optional value of the tag. The string value can be from 1 to 256 Unicode
	// characters in length and can't be prefixed with " aws: " or " rds: ". The string
	// can contain only the set of Unicode letters, digits, white space, '_', '.', '/',
	// '=', '+', '-' (Java regex: "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-]*)$").
	Value *string

	noSmithyDocumentSerde
}

// The version of the database engine that an instance can be upgraded to.
type UpgradeTarget struct {

	// A value that indicates whether the target version is applied to any source DB
	// instances that have AutoMinorVersionUpgrade set to true .
	AutoUpgrade *bool

	// The version of the database engine that an instance can be upgraded to.
	Description *string

	// The name of the upgrade target database engine.
	Engine *string

	// The version number of the upgrade target database engine.
	EngineVersion *string

	// A value that indicates whether a database engine is upgraded to a major version.
	IsMajorVersionUpgrade *bool

	noSmithyDocumentSerde
}

// Used as a response element for queries on virtual private cloud (VPC) security
// group membership.
type VpcSecurityGroupMembership struct {

	// The status of the VPC security group.
	Status *string

	// The name of the VPC security group.
	VpcSecurityGroupId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
