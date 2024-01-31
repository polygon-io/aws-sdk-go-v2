// Code generated by smithy-go-codegen DO NOT EDIT.

package emrcontainers

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emrcontainers/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Generate a session token to connect to a managed endpoint.
func (c *Client) GetManagedEndpointSessionCredentials(ctx context.Context, params *GetManagedEndpointSessionCredentialsInput, optFns ...func(*Options)) (*GetManagedEndpointSessionCredentialsOutput, error) {
	if params == nil {
		params = &GetManagedEndpointSessionCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetManagedEndpointSessionCredentials", params, optFns, c.addOperationGetManagedEndpointSessionCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetManagedEndpointSessionCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetManagedEndpointSessionCredentialsInput struct {

	// Type of the token requested. Currently supported and default value of this
	// field is “TOKEN.”
	//
	// This member is required.
	CredentialType *string

	// The ARN of the managed endpoint for which the request is submitted.
	//
	// This member is required.
	EndpointIdentifier *string

	// The IAM Execution Role ARN that will be used by the job run.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The ARN of the Virtual Cluster which the Managed Endpoint belongs to.
	//
	// This member is required.
	VirtualClusterIdentifier *string

	// The client idempotency token of the job run request.
	ClientToken *string

	// Duration in seconds for which the session token is valid. The default duration
	// is 15 minutes and the maximum is 12 hours.
	DurationInSeconds *int32

	// String identifier used to separate sections of the execution logs uploaded to
	// S3.
	LogContext *string

	noSmithyDocumentSerde
}

type GetManagedEndpointSessionCredentialsOutput struct {

	// The structure containing the session credentials.
	Credentials types.Credentials

	// The date and time when the session token will expire.
	ExpiresAt *time.Time

	// The identifier of the session token returned.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetManagedEndpointSessionCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetManagedEndpointSessionCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetManagedEndpointSessionCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetManagedEndpointSessionCredentials"); err != nil {
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
	if err = addIdempotencyToken_opGetManagedEndpointSessionCredentialsMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetManagedEndpointSessionCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetManagedEndpointSessionCredentials(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpGetManagedEndpointSessionCredentials struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpGetManagedEndpointSessionCredentials) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpGetManagedEndpointSessionCredentials) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*GetManagedEndpointSessionCredentialsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *GetManagedEndpointSessionCredentialsInput ")
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
func addIdempotencyToken_opGetManagedEndpointSessionCredentialsMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpGetManagedEndpointSessionCredentials{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opGetManagedEndpointSessionCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetManagedEndpointSessionCredentials",
	}
}
