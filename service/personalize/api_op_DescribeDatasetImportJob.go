// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the dataset import job created by CreateDatasetImportJob (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDatasetImportJob.html)
// , including the import job status.
func (c *Client) DescribeDatasetImportJob(ctx context.Context, params *DescribeDatasetImportJobInput, optFns ...func(*Options)) (*DescribeDatasetImportJobOutput, error) {
	if params == nil {
		params = &DescribeDatasetImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDatasetImportJob", params, optFns, c.addOperationDescribeDatasetImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetImportJobInput struct {

	// The Amazon Resource Name (ARN) of the dataset import job to describe.
	//
	// This member is required.
	DatasetImportJobArn *string

	noSmithyDocumentSerde
}

type DescribeDatasetImportJobOutput struct {

	// Information about the dataset import job, including the status. The status is
	// one of the following values:
	//   - CREATE PENDING
	//   - CREATE IN_PROGRESS
	//   - ACTIVE
	//   - CREATE FAILED
	DatasetImportJob *types.DatasetImportJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDatasetImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDatasetImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDatasetImportJob"); err != nil {
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
	if err = addOpDescribeDatasetImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatasetImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDatasetImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDatasetImportJob",
	}
}
