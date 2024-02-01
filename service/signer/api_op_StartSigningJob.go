// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initiates a signing job to be performed on the code provided. Signing jobs are
// viewable by the ListSigningJobs operation for two years after they are
// performed. Note the following requirements:
//   - You must create an Amazon S3 source bucket. For more information, see
//     Creating a Bucket (http://docs.aws.amazon.com/AmazonS3/latest/gsg/CreatingABucket.html)
//     in the Amazon S3 Getting Started Guide.
//   - Your S3 source bucket must be version enabled.
//   - You must create an S3 destination bucket. AWS Signer uses your S3
//     destination bucket to write your signed code.
//   - You specify the name of the source and destination buckets when calling the
//     StartSigningJob operation.
//   - You must also specify a request token that identifies your request to
//     Signer.
//
// You can call the DescribeSigningJob and the ListSigningJobs actions after you
// call StartSigningJob . For a Java example that shows how to use this action, see
// StartSigningJob (https://docs.aws.amazon.com/signer/latest/developerguide/api-startsigningjob.html)
// .
func (c *Client) StartSigningJob(ctx context.Context, params *StartSigningJobInput, optFns ...func(*Options)) (*StartSigningJobOutput, error) {
	if params == nil {
		params = &StartSigningJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSigningJob", params, optFns, c.addOperationStartSigningJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSigningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSigningJobInput struct {

	// String that identifies the signing request. All calls after the first that use
	// this token return the same response as the first call.
	//
	// This member is required.
	ClientRequestToken *string

	// The S3 bucket in which to save your signed object. The destination contains the
	// name of your bucket and an optional prefix.
	//
	// This member is required.
	Destination *types.Destination

	// The name of the signing profile.
	//
	// This member is required.
	ProfileName *string

	// The S3 bucket that contains the object to sign or a BLOB that contains your raw
	// code.
	//
	// This member is required.
	Source *types.Source

	// The AWS account ID of the signing profile owner.
	ProfileOwner *string

	noSmithyDocumentSerde
}

type StartSigningJobOutput struct {

	// The ID of your signing job.
	JobId *string

	// The AWS account ID of the signing job owner.
	JobOwner *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSigningJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSigningJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSigningJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSigningJob"); err != nil {
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
	if err = addIdempotencyToken_opStartSigningJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSigningJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSigningJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartSigningJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSigningJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSigningJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSigningJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSigningJobInput ")
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
func addIdempotencyToken_opStartSigningJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSigningJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSigningJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSigningJob",
	}
}
