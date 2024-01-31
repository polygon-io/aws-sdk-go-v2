// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get position configuration for a given resource. This action is no longer
// supported. Calls to retrieve the position configuration should use the
// GetResourcePosition (https://docs.aws.amazon.com/iot-wireless/2020-11-22/apireference/API_GetResourcePosition.html)
// API operation instead.
//
// Deprecated: This operation is no longer supported.
func (c *Client) GetPositionConfiguration(ctx context.Context, params *GetPositionConfigurationInput, optFns ...func(*Options)) (*GetPositionConfigurationOutput, error) {
	if params == nil {
		params = &GetPositionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPositionConfiguration", params, optFns, c.addOperationGetPositionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPositionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPositionConfigurationInput struct {

	// Resource identifier used in a position configuration.
	//
	// This member is required.
	ResourceIdentifier *string

	// Resource type of the resource for which position configuration is retrieved.
	//
	// This member is required.
	ResourceType types.PositionResourceType

	noSmithyDocumentSerde
}

type GetPositionConfigurationOutput struct {

	// The position data destination that describes the AWS IoT rule that processes
	// the device's position data for use by AWS IoT Core for LoRaWAN.
	Destination *string

	// The wrapper for the solver configuration details object.
	Solvers *types.PositionSolverDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPositionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPositionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPositionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPositionConfiguration"); err != nil {
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
	if err = addOpGetPositionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPositionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPositionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPositionConfiguration",
	}
}
