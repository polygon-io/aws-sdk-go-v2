// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a managed kdb environment for the account.
func (c *Client) CreateKxEnvironment(ctx context.Context, params *CreateKxEnvironmentInput, optFns ...func(*Options)) (*CreateKxEnvironmentOutput, error) {
	if params == nil {
		params = &CreateKxEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateKxEnvironment", params, optFns, c.addOperationCreateKxEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateKxEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateKxEnvironmentInput struct {

	// The KMS key ID to encrypt your data in the FinSpace environment.
	//
	// This member is required.
	KmsKeyId *string

	// The name of the kdb environment that you want to create.
	//
	// This member is required.
	Name *string

	// A token that ensures idempotency. This token expires in 10 minutes.
	ClientToken *string

	// A description for the kdb environment.
	Description *string

	// A list of key-value pairs to label the kdb environment. You can add up to 50
	// tags to your kdb environment.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateKxEnvironmentOutput struct {

	// The timestamp at which the kdb environment was created in FinSpace.
	CreationTimestamp *time.Time

	// A description for the kdb environment.
	Description *string

	// The ARN identifier of the environment.
	EnvironmentArn *string

	// A unique identifier for the kdb environment.
	EnvironmentId *string

	// The KMS key ID to encrypt your data in the FinSpace environment.
	KmsKeyId *string

	// The name of the kdb environment.
	Name *string

	// The status of the kdb environment.
	Status types.EnvironmentStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateKxEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateKxEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateKxEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateKxEnvironment"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addIdempotencyToken_opCreateKxEnvironmentMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateKxEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateKxEnvironment(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateKxEnvironment struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateKxEnvironment) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateKxEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateKxEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateKxEnvironmentInput ")
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
func addIdempotencyToken_opCreateKxEnvironmentMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateKxEnvironment{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateKxEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateKxEnvironment",
	}
}
