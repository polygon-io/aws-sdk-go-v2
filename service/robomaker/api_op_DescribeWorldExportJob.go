// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a world export job.
func (c *Client) DescribeWorldExportJob(ctx context.Context, params *DescribeWorldExportJobInput, optFns ...func(*Options)) (*DescribeWorldExportJobOutput, error) {
	if params == nil {
		params = &DescribeWorldExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorldExportJob", params, optFns, c.addOperationDescribeWorldExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorldExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorldExportJobInput struct {

	// The Amazon Resource Name (arn) of the world export job to describe.
	//
	// This member is required.
	Job *string

	noSmithyDocumentSerde
}

type DescribeWorldExportJobOutput struct {

	// The Amazon Resource Name (ARN) of the world export job.
	Arn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The time, in milliseconds since the epoch, when the world export job was
	// created.
	CreatedAt *time.Time

	// The failure code of the world export job if it failed: InternalServiceError
	// Internal service error. LimitExceeded The requested resource exceeds the maximum
	// number allowed, or the number of concurrent stream requests exceeds the maximum
	// number allowed. ResourceNotFound The specified resource could not be found.
	// RequestThrottled The request was throttled. InvalidInput An input parameter in
	// the request is not valid.
	FailureCode types.WorldExportJobErrorCode

	// The reason why the world export job failed.
	FailureReason *string

	// The IAM role that the world export process uses to access the Amazon S3 bucket
	// and put the export.
	IamRole *string

	// The output location.
	OutputLocation *types.OutputLocation

	// The status of the world export job. Pending The world export job request is
	// pending. Running The world export job is running. Completed The world export job
	// completed. Failed The world export job failed. See failureCode and failureReason
	// for more information. Canceled The world export job was cancelled. Canceling The
	// world export job is being cancelled.
	Status types.WorldExportJobStatus

	// A map that contains tag keys and tag values that are attached to the world
	// export job.
	Tags map[string]string

	// A list of Amazon Resource Names (arns) that correspond to worlds to be exported.
	Worlds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorldExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeWorldExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeWorldExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeWorldExportJob"); err != nil {
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
	if err = addOpDescribeWorldExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorldExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeWorldExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeWorldExportJob",
	}
}
