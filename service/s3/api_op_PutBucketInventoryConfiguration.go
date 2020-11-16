// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This implementation of the PUT operation adds an inventory configuration
// (identified by the inventory ID) to the bucket. You can have up to 1,000
// inventory configurations per bucket. Amazon S3 inventory generates inventories
// of the objects in the bucket on a daily or weekly basis, and the results are
// published to a flat file. The bucket that is inventoried is called the source
// bucket, and the bucket where the inventory flat file is stored is called the
// destination bucket. The destination bucket must be in the same AWS Region as the
// source bucket. When you configure an inventory for a source bucket, you specify
// the destination bucket where you want the inventory to be stored, and whether to
// generate the inventory daily or weekly. You can also configure what object
// metadata to include and whether to inventory all object versions or only current
// versions. For more information, see Amazon S3 Inventory
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) in the
// Amazon Simple Storage Service Developer Guide. You must create a bucket policy
// on the destination bucket to grant permissions to Amazon S3 to write objects to
// the bucket in the defined location. For an example policy, see  Granting
// Permissions for Amazon S3 Inventory and Storage Class Analysis
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html#example-bucket-policies-use-case-9).
// To use this operation, you must have permissions to perform the
// s3:PutInventoryConfiguration action. The bucket owner has this permission by
// default and can grant this permission to others. For more information about
// permissions, see Permissions Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
// Amazon Simple Storage Service Developer Guide. Special Errors
//
// * HTTP 400 Bad
// Request Error
//
// * Code: InvalidArgument
//
// * Cause: Invalid Argument
//
// * HTTP 400
// Bad Request Error
//
// * Code: TooManyConfigurations
//
// * Cause: You are attempting to
// create a new configuration but have already reached the 1,000-configuration
// limit.
//
// * HTTP 403 Forbidden Error
//
// * Code: AccessDenied
//
// * Cause: You are not
// the owner of the specified bucket, or you do not have the
// s3:PutInventoryConfiguration bucket permission to set the configuration on the
// bucket.
//
// Related Resources
//
// * GetBucketInventoryConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html)
//
// *
// DeleteBucketInventoryConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html)
//
// *
// ListBucketInventoryConfigurations
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html)
func (c *Client) PutBucketInventoryConfiguration(ctx context.Context, params *PutBucketInventoryConfigurationInput, optFns ...func(*Options)) (*PutBucketInventoryConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketInventoryConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketInventoryConfiguration", params, optFns, addOperationPutBucketInventoryConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketInventoryConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketInventoryConfigurationInput struct {

	// The name of the bucket where the inventory configuration will be stored.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the inventory configuration.
	//
	// This member is required.
	Id *string

	// Specifies the inventory configuration.
	//
	// This member is required.
	InventoryConfiguration *types.InventoryConfiguration

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type PutBucketInventoryConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketInventoryConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketInventoryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketInventoryConfiguration{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketInventoryConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketInventoryConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketInventoryConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketInventoryConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketInventoryConfiguration",
	}
}

// getPutBucketInventoryConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getPutBucketInventoryConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketInventoryConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketInventoryConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketInventoryConfigurationBucketMember,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		SupportsAccelerate:      true,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
