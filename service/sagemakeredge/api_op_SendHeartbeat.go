// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakeredge

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakeredge/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use to get the current status of devices registered on SageMaker Edge Manager.
func (c *Client) SendHeartbeat(ctx context.Context, params *SendHeartbeatInput, optFns ...func(*Options)) (*SendHeartbeatOutput, error) {
	if params == nil {
		params = &SendHeartbeatInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendHeartbeat", params, optFns, c.addOperationSendHeartbeatMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendHeartbeatOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendHeartbeatInput struct {

	// Returns the version of the agent.
	//
	// This member is required.
	AgentVersion *string

	// The name of the fleet that the device belongs to.
	//
	// This member is required.
	DeviceFleetName *string

	// The unique name of the device.
	//
	// This member is required.
	DeviceName *string

	// For internal use. Returns a list of SageMaker Edge Manager agent operating
	// metrics.
	AgentMetrics []types.EdgeMetric

	// Returns the result of a deployment on the device.
	DeploymentResult *types.DeploymentResult

	// Returns a list of models deployed on the the device.
	Models []types.Model

	noSmithyDocumentSerde
}

type SendHeartbeatOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendHeartbeatMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendHeartbeat{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendHeartbeat{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendHeartbeat"); err != nil {
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
	if err = addOpSendHeartbeatValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendHeartbeat(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendHeartbeat(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendHeartbeat",
	}
}
