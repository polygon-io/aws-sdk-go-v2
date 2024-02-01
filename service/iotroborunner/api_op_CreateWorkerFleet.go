// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants permission to create a worker fleet
func (c *Client) CreateWorkerFleet(ctx context.Context, params *CreateWorkerFleetInput, optFns ...func(*Options)) (*CreateWorkerFleetOutput, error) {
	if params == nil {
		params = &CreateWorkerFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkerFleet", params, optFns, c.addOperationCreateWorkerFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkerFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkerFleetInput struct {

	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string

	// Site ARN.
	//
	// This member is required.
	Site *string

	// JSON blob containing additional fixed properties regarding the worker fleet
	AdditionalFixedProperties *string

	// Token used for detecting replayed requests. Replayed requests will not be
	// performed multiple times.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateWorkerFleetOutput struct {

	// Full ARN of the worker fleet.
	//
	// This member is required.
	Arn *string

	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Filters access by the worker fleet's identifier
	//
	// This member is required.
	Id *string

	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkerFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkerFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkerFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateWorkerFleet"); err != nil {
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
	if err = addIdempotencyToken_opCreateWorkerFleetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorkerFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkerFleet(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorkerFleet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorkerFleet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorkerFleet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWorkerFleetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorkerFleetInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateWorkerFleetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorkerFleet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorkerFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateWorkerFleet",
	}
}
