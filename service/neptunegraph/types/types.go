// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Details about a graph snapshot.
type GraphSnapshotSummary struct {

	// The ARN of the graph snapshot.
	//
	// This member is required.
	Arn *string

	// The unique identifier of the graph snapshot.
	//
	// This member is required.
	Id *string

	// The snapshot name. For example: my-snapshot-1 . The name must contain from 1 to
	// 63 letters, numbers, or hyphens, and its first character must be a letter. It
	// cannot end with a hyphen or contain two consecutive hyphens.
	//
	// This member is required.
	Name *string

	// The ID of the KMS key used to encrypt and decrypt the snapshot.
	KmsKeyIdentifier *string

	// The time when the snapshot was created.
	SnapshotCreateTime *time.Time

	// The graph identifier for the graph for which a snapshot is to be created.
	SourceGraphId *string

	// The status of the graph snapshot.
	Status SnapshotStatus

	noSmithyDocumentSerde
}

// Summary details about a graph.
type GraphSummary struct {

	// The ARN associated with the graph.
	//
	// This member is required.
	Arn *string

	// The unique identifier of the graph.
	//
	// This member is required.
	Id *string

	// The name of the graph.
	//
	// This member is required.
	Name *string

	// If true , deletion protection is enabled for the graph.
	DeletionProtection *bool

	// The graph endpoint.
	Endpoint *string

	// The ID of the KMS key used to encrypt and decrypt graph data.
	KmsKeyIdentifier *string

	// The number of memory-optimized Neptune Capacity Units (m-NCUs) allocated to the
	// graph.
	ProvisionedMemory *int32

	// If true , the graph has a public endpoint, otherwise not.
	PublicConnectivity *bool

	// The number of replicas for the graph.
	ReplicaCount *int32

	// The status of the graph.
	Status GraphStatus

	noSmithyDocumentSerde
}

// Options for how to perform an import.
//
// The following types satisfy this interface:
//
//	ImportOptionsMemberNeptune
type ImportOptions interface {
	isImportOptions()
}

// Options for importing data from a Neptune database.
type ImportOptionsMemberNeptune struct {
	Value NeptuneImportOptions

	noSmithyDocumentSerde
}

func (*ImportOptionsMemberNeptune) isImportOptions() {}

// Contains details about an import task.
type ImportTaskDetails struct {

	// The number of dictionary entries in the import task.
	//
	// This member is required.
	DictionaryEntryCount *int64

	// The number of errors encountered so far.
	//
	// This member is required.
	ErrorCount *int32

	// The percentage progress so far.
	//
	// This member is required.
	ProgressPercentage *int32

	// Time at which the import task started.
	//
	// This member is required.
	StartTime *time.Time

	// The number of statements in the import task.
	//
	// This member is required.
	StatementCount *int64

	// Status of the import task.
	//
	// This member is required.
	Status *string

	// Seconds elapsed since the import task started.
	//
	// This member is required.
	TimeElapsedSeconds *int64

	// Details about the errors that have been encountered.
	ErrorDetails *string

	noSmithyDocumentSerde
}

// Details about an import task.
type ImportTaskSummary struct {

	// The ARN of the IAM role that will allow access to the data that is to be
	// imported.
	//
	// This member is required.
	RoleArn *string

	// A URL identifying to the location of the data to be imported. This can be an
	// Amazon S3 path, or can point to a Neptune database endpoint or snapshot
	//
	// This member is required.
	Source *string

	// Status of the import task.
	//
	// This member is required.
	Status ImportTaskStatus

	// The unique identifier of the import task.
	//
	// This member is required.
	TaskId *string

	// Specifies the format of S3 data to be imported. Valid values are CSV , which
	// identifies the Gremlin CSV format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-gremlin.html)
	// or OPENCYPHER , which identies the openCypher load format (https://docs.aws.amazon.com/neptune/latest/userguide/bulk-load-tutorial-format-opencypher.html)
	// .
	Format Format

	// The unique identifier of the Neptune Analytics graph.
	GraphId *string

	noSmithyDocumentSerde
}

// Options for how to import Neptune data.
type NeptuneImportOptions struct {

	// The KMS key to use to encrypt data in the S3 bucket where the graph data is
	// exported
	//
	// This member is required.
	S3ExportKmsKeyId *string

	// The path to an S3 bucket from which to import data.
	//
	// This member is required.
	S3ExportPath *string

	// Neptune Analytics supports label-less vertices and no labels are assigned
	// unless one is explicitly provided. Neptune assigns default labels when none is
	// explicitly provided. When importing the data into Neptune Analytics, the default
	// vertex labels can be omitted by setting preserveDefaultVertexLabels to false.
	// Note that if the vertex only has default labels, and has no other properties or
	// edges, then the vertex will effectively not get imported into Neptune Analytics
	// when preserveDefaultVertexLabels is set to false.
	PreserveDefaultVertexLabels *bool

	// Neptune Analytics currently does not support user defined edge ids. The edge
	// ids are not imported by default. They are imported if preserveEdgeIds is set to
	// true, and ids are stored as properties on the relationships with the property
	// name neptuneEdgeId.
	PreserveEdgeIds *bool

	noSmithyDocumentSerde
}

// Details about a private graph endpoint.
type PrivateGraphEndpointSummary struct {

	// The status of the private graph endpoint.
	//
	// This member is required.
	Status PrivateGraphEndpointStatus

	// The subnet IDs associated with the private graph endpoint.
	//
	// This member is required.
	SubnetIds []string

	// The ID of the VPC in which the private graph endpoint is located.
	//
	// This member is required.
	VpcId *string

	// The ID of the VPC endpoint.
	VpcEndpointId *string

	noSmithyDocumentSerde
}

// Specifies the number of dimensions for vector embeddings loaded into the graph.
// Max = 65535
type VectorSearchConfiguration struct {

	// The number of dimensions.
	//
	// This member is required.
	Dimension *int32

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

func (*UnknownUnionMember) isImportOptions() {}
