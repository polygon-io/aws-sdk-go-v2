// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This operation lists details about a partner event source that is shared with
// your account.
func (c *Client) DescribeEventSource(ctx context.Context, params *DescribeEventSourceInput, optFns ...func(*Options)) (*DescribeEventSourceOutput, error) {
	if params == nil {
		params = &DescribeEventSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventSource", params, optFns, c.addOperationDescribeEventSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventSourceInput struct {

	// The name of the partner event source to display the details of.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeEventSourceOutput struct {

	// The ARN of the partner event source.
	Arn *string

	// The name of the SaaS partner that created the event source.
	CreatedBy *string

	// The date and time that the event source was created.
	CreationTime *time.Time

	// The date and time that the event source will expire if you do not create a
	// matching event bus.
	ExpirationTime *time.Time

	// The name of the partner event source.
	Name *string

	// The state of the event source. If it is ACTIVE, you have already created a
	// matching event bus for this event source, and that event bus is active. If it is
	// PENDING, either you haven't yet created a matching event bus, or that event bus
	// is deactivated. If it is DELETED, you have created a matching event bus, but the
	// event source has since been deleted.
	State types.EventSourceState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEventSource"); err != nil {
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
	if err = addOpDescribeEventSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEventSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEventSource",
	}
}
