// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotroborunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants permission to update a destination
func (c *Client) UpdateDestination(ctx context.Context, params *UpdateDestinationInput, optFns ...func(*Options)) (*UpdateDestinationOutput, error) {
	if params == nil {
		params = &UpdateDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDestination", params, optFns, c.addOperationUpdateDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDestinationInput struct {

	// Destination ARN.
	//
	// This member is required.
	Id *string

	// JSON document containing additional fixed properties regarding the destination
	AdditionalFixedProperties *string

	// Human friendly name of the resource.
	Name *string

	// State of the destination.
	State types.DestinationState

	noSmithyDocumentSerde
}

type UpdateDestinationOutput struct {

	// Destination ARN.
	//
	// This member is required.
	Arn *string

	// Filters access by the destination's identifier
	//
	// This member is required.
	Id *string

	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string

	// State of the destination.
	//
	// This member is required.
	State types.DestinationState

	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// JSON document containing additional fixed properties regarding the destination
	AdditionalFixedProperties *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDestination{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDestination"); err != nil {
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
	if err = addOpUpdateDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDestination",
	}
}
