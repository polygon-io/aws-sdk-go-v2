// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakera2iruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the specified human loop. If the human loop was
// deleted, this operation will return a ResourceNotFoundException error.
func (c *Client) DescribeHumanLoop(ctx context.Context, params *DescribeHumanLoopInput, optFns ...func(*Options)) (*DescribeHumanLoopOutput, error) {
	if params == nil {
		params = &DescribeHumanLoopInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHumanLoop", params, optFns, c.addOperationDescribeHumanLoopMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHumanLoopOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHumanLoopInput struct {

	// The name of the human loop that you want information about.
	//
	// This member is required.
	HumanLoopName *string

	noSmithyDocumentSerde
}

type DescribeHumanLoopOutput struct {

	// The creation time when Amazon Augmented AI created the human loop.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the flow definition.
	//
	// This member is required.
	FlowDefinitionArn *string

	// The Amazon Resource Name (ARN) of the human loop.
	//
	// This member is required.
	HumanLoopArn *string

	// The name of the human loop. The name must be lowercase, unique within the
	// Region in your account, and can have up to 63 characters. Valid characters: a-z,
	// 0-9, and - (hyphen).
	//
	// This member is required.
	HumanLoopName *string

	// The status of the human loop.
	//
	// This member is required.
	HumanLoopStatus types.HumanLoopStatus

	// A failure code that identifies the type of failure. Possible values:
	// ValidationError , Expired , InternalError
	FailureCode *string

	// The reason why a human loop failed. The failure reason is returned when the
	// status of the human loop is Failed .
	FailureReason *string

	// An object that contains information about the output of the human loop.
	HumanLoopOutput *types.HumanLoopOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHumanLoopMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeHumanLoop{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeHumanLoop{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeHumanLoop"); err != nil {
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
	if err = addOpDescribeHumanLoopValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHumanLoop(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHumanLoop(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeHumanLoop",
	}
}
