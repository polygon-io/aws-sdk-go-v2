// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a topic refresh schedule.
func (c *Client) DeleteTopicRefreshSchedule(ctx context.Context, params *DeleteTopicRefreshScheduleInput, optFns ...func(*Options)) (*DeleteTopicRefreshScheduleOutput, error) {
	if params == nil {
		params = &DeleteTopicRefreshScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTopicRefreshSchedule", params, optFns, c.addOperationDeleteTopicRefreshScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTopicRefreshScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTopicRefreshScheduleInput struct {

	// The Amazon Web Services account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the dataset.
	//
	// This member is required.
	DatasetId *string

	// The ID of the topic that you want to modify. This ID is unique per Amazon Web
	// Services Region for each Amazon Web Services account.
	//
	// This member is required.
	TopicId *string

	noSmithyDocumentSerde
}

type DeleteTopicRefreshScheduleOutput struct {

	// The Amazon Resource Name (ARN) of the dataset.
	DatasetArn *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The Amazon Resource Name (ARN) of the topic.
	TopicArn *string

	// The ID of the topic that you want to modify. This ID is unique per Amazon Web
	// Services Region for each Amazon Web Services account.
	TopicId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTopicRefreshScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteTopicRefreshSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteTopicRefreshSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteTopicRefreshSchedule"); err != nil {
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
	if err = addOpDeleteTopicRefreshScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTopicRefreshSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTopicRefreshSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteTopicRefreshSchedule",
	}
}
