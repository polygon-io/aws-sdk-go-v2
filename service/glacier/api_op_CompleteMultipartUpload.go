// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You call this operation to inform Amazon S3 Glacier (Glacier) that all the
// archive parts have been uploaded and that Glacier can now assemble the archive
// from the uploaded parts. After assembling and saving the archive to the vault,
// Glacier returns the URI path of the newly created archive resource. Using the
// URI path, you can then access the archive. After you upload an archive, you
// should save the archive ID returned to retrieve the archive at a later point.
// You can also get the vault inventory to obtain a list of archive IDs in a vault.
// For more information, see InitiateJob . In the request, you must include the
// computed SHA256 tree hash of the entire archive you have uploaded. For
// information about computing a SHA256 tree hash, see Computing Checksums (https://docs.aws.amazon.com/amazonglacier/latest/dev/checksum-calculations.html)
// . On the server side, Glacier also constructs the SHA256 tree hash of the
// assembled archive. If the values match, Glacier saves the archive to the vault;
// otherwise, it returns an error, and the operation fails. The ListParts
// operation returns a list of parts uploaded for a specific multipart upload. It
// includes checksum information for each uploaded part that can be used to debug a
// bad checksum issue. Additionally, Glacier also checks for any missing content
// ranges when assembling the archive, if missing content ranges are found, Glacier
// returns an error and the operation fails. Complete Multipart Upload is an
// idempotent operation. After your first successful complete multipart upload, if
// you call the operation again within a short period, the operation will succeed
// and return the same archive ID. This is useful in the event you experience a
// network issue that causes an aborted connection or receive a 500 server error,
// in which case you can repeat your Complete Multipart Upload request and get the
// same archive ID without creating duplicate archives. Note, however, that after
// the multipart upload completes, you cannot call the List Parts operation and the
// multipart upload will not appear in List Multipart Uploads response, even if
// idempotent complete is possible. An AWS account has full permission to perform
// all operations (actions). However, AWS Identity and Access Management (IAM)
// users don't have any permissions by default. You must grant them explicit
// permission to perform specific actions. For more information, see Access
// Control Using AWS Identity and Access Management (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html)
// . For conceptual information and underlying REST API, see Uploading Large
// Archives in Parts (Multipart Upload) (https://docs.aws.amazon.com/amazonglacier/latest/dev/uploading-archive-mpu.html)
// and Complete Multipart Upload (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-multipart-complete-upload.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) CompleteMultipartUpload(ctx context.Context, params *CompleteMultipartUploadInput, optFns ...func(*Options)) (*CompleteMultipartUploadOutput, error) {
	if params == nil {
		params = &CompleteMultipartUploadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CompleteMultipartUpload", params, optFns, c.addOperationCompleteMultipartUploadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CompleteMultipartUploadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to complete a multipart upload operation. This informs Amazon
// Glacier that all the archive parts have been uploaded and Amazon S3 Glacier
// (Glacier) can now assemble the archive from the uploaded parts. After assembling
// and saving the archive to the vault, Glacier returns the URI path of the newly
// created archive resource.
type CompleteMultipartUploadInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single ' - ' (hyphen),
	// in which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The upload ID of the multipart upload.
	//
	// This member is required.
	UploadId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string

	// The total size, in bytes, of the entire archive. This value should be the sum
	// of all the sizes of the individual parts that you uploaded.
	ArchiveSize *string

	// The SHA256 tree hash of the entire archive. It is the tree hash of SHA256 tree
	// hash of the individual parts. If the value you specify in the request does not
	// match the SHA256 tree hash of the final assembled archive as computed by Amazon
	// S3 Glacier (Glacier), Glacier returns an error and the request fails.
	Checksum *string

	noSmithyDocumentSerde
}

// Contains the Amazon S3 Glacier response to your request. For information about
// the underlying REST API, see Upload Archive (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-post.html)
// . For conceptual information, see Working with Archives in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-archives.html)
// .
type CompleteMultipartUploadOutput struct {

	// The ID of the archive. This value is also included as part of the location.
	ArchiveId *string

	// The checksum of the archive computed by Amazon S3 Glacier.
	Checksum *string

	// The relative URI path of the newly added archive resource.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCompleteMultipartUploadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCompleteMultipartUpload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCompleteMultipartUpload{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CompleteMultipartUpload"); err != nil {
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
	if err = addOpCompleteMultipartUploadValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCompleteMultipartUpload(options.Region), middleware.Before); err != nil {
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
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
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

func newServiceMetadataMiddleware_opCompleteMultipartUpload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CompleteMultipartUpload",
	}
}
