// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// If you provide a value for PerformedBy.UserArn you must also have
// connect:DescribeUser (https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeUser.html)
// permission on the User ARN resource that you provide Creates a case in the
// specified Cases domain. Case system and custom fields are taken as an array
// id/value pairs with a declared data types. The following fields are required
// when creating a case:
//   - customer_id - You must provide the full customer profile ARN in this format:
//     arn:aws:profile:your_AWS_Region:your_AWS_account
//     ID:domains/your_profiles_domain_name/profiles/profile_ID
//   - title
func (c *Client) CreateCase(ctx context.Context, params *CreateCaseInput, optFns ...func(*Options)) (*CreateCaseOutput, error) {
	if params == nil {
		params = &CreateCaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCase", params, optFns, c.addOperationCreateCaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCaseInput struct {

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// An array of objects with field ID (matching ListFields/DescribeField) and value
	// union data.
	//
	// This member is required.
	Fields []types.FieldValue

	// A unique identifier of a template.
	//
	// This member is required.
	TemplateId *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If not provided, the Amazon Web Services SDK populates this
	// field. For more information about idempotency, see Making retries safe with
	// idempotent APIs (https://aws.amazon.com/builders-library/making-retries-safe-with-idempotent-APIs/)
	// .
	ClientToken *string

	// Represents the identity of the person who performed the action.
	PerformedBy types.UserUnion

	noSmithyDocumentSerde
}

type CreateCaseOutput struct {

	// The Amazon Resource Name (ARN) of the case.
	//
	// This member is required.
	CaseArn *string

	// A unique identifier of the case.
	//
	// This member is required.
	CaseId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCase"); err != nil {
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
	if err = addIdempotencyToken_opCreateCaseMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateCaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCase(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateCase struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateCase) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateCase) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateCaseInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateCaseInput ")
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
func addIdempotencyToken_opCreateCaseMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateCase{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateCase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCase",
	}
}
