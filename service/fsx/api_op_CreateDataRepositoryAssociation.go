// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon FSx for Lustre data repository association (DRA). A data
// repository association is a link between a directory on the file system and an
// Amazon S3 bucket or prefix. You can have a maximum of 8 data repository
// associations on a file system. Data repository associations are supported on all
// FSx for Lustre 2.12 and 2.15 file systems, excluding scratch_1 deployment type.
// Each data repository association must have a unique Amazon FSx file system
// directory and a unique S3 bucket or prefix associated with it. You can configure
// a data repository association for automatic import only, for automatic export
// only, or for both. To learn more about linking a data repository to your file
// system, see Linking your file system to an S3 bucket (https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html)
// . CreateDataRepositoryAssociation isn't supported on Amazon File Cache
// resources. To create a DRA on Amazon File Cache, use the CreateFileCache
// operation.
func (c *Client) CreateDataRepositoryAssociation(ctx context.Context, params *CreateDataRepositoryAssociationInput, optFns ...func(*Options)) (*CreateDataRepositoryAssociationOutput, error) {
	if params == nil {
		params = &CreateDataRepositoryAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataRepositoryAssociation", params, optFns, c.addOperationCreateDataRepositoryAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataRepositoryAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataRepositoryAssociationInput struct {

	// The path to the Amazon S3 data repository that will be linked to the file
	// system. The path can be an S3 bucket or prefix in the format
	// s3://myBucket/myPrefix/ . This path specifies where in the S3 data repository
	// files will be imported from or exported to.
	//
	// This member is required.
	DataRepositoryPath *string

	// The globally unique ID of the file system, assigned by Amazon FSx.
	//
	// This member is required.
	FileSystemId *string

	// Set to true to run an import data repository task to import metadata from the
	// data repository to the file system after the data repository association is
	// created. Default is false .
	BatchImportMetaDataOnCreate *bool

	// (Optional) An idempotency token for resource creation, in a string of up to 63
	// ASCII characters. This token is automatically filled on your behalf when you use
	// the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// A path on the file system that points to a high-level directory (such as /ns1/ )
	// or subdirectory (such as /ns1/subdir/ ) that will be mapped 1-1 with
	// DataRepositoryPath . The leading forward slash in the name is required. Two data
	// repository associations cannot have overlapping file system paths. For example,
	// if a data repository is associated with file system path /ns1/ , then you cannot
	// link another data repository with file system path /ns1/ns2 . This path
	// specifies where in your file system files will be exported from or imported to.
	// This file system directory can be linked to only one Amazon S3 bucket, and no
	// other S3 bucket can be linked to the directory. If you specify only a forward
	// slash ( / ) as the file system path, you can link only one data repository to
	// the file system. You can only specify "/" as the file system path for the first
	// data repository associated with a file system.
	FileSystemPath *string

	// For files imported from a data repository, this value determines the stripe
	// count and maximum amount of data per file (in MiB) stored on a single physical
	// disk. The maximum number of disks that a single file can be striped across is
	// limited by the total number of disks that make up the file system. The default
	// chunk size is 1,024 MiB (1 GiB) and can go as high as 512,000 MiB (500 GiB).
	// Amazon S3 objects have a maximum size of 5 TB.
	ImportedFileChunkSize *int32

	// The configuration for an Amazon S3 data repository linked to an Amazon FSx
	// Lustre file system with a data repository association. The configuration defines
	// which file events (new, changed, or deleted files or directories) are
	// automatically imported from the linked data repository to the file system or
	// automatically exported from the file system to the data repository.
	S3 *types.S3DataRepositoryConfiguration

	// A list of Tag values, with a maximum of 50 elements.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateDataRepositoryAssociationOutput struct {

	// The response object returned after the data repository association is created.
	Association *types.DataRepositoryAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataRepositoryAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDataRepositoryAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDataRepositoryAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDataRepositoryAssociation"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateDataRepositoryAssociationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataRepositoryAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataRepositoryAssociation(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpCreateDataRepositoryAssociation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataRepositoryAssociation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataRepositoryAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataRepositoryAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataRepositoryAssociationInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDataRepositoryAssociationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataRepositoryAssociation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataRepositoryAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDataRepositoryAssociation",
	}
}
