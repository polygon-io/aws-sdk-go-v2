// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a business glossary term.
func (c *Client) CreateGlossaryTerm(ctx context.Context, params *CreateGlossaryTermInput, optFns ...func(*Options)) (*CreateGlossaryTermOutput, error) {
	if params == nil {
		params = &CreateGlossaryTermInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGlossaryTerm", params, optFns, c.addOperationCreateGlossaryTermMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGlossaryTermOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGlossaryTermInput struct {

	// The ID of the Amazon DataZone domain in which this business glossary term is
	// created.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the business glossary in which this term is created.
	//
	// This member is required.
	GlossaryIdentifier *string

	// The name of this business glossary term.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive identifier that is provided to ensure the idempotency
	// of the request.
	ClientToken *string

	// The long description of this business glossary term.
	LongDescription *string

	// The short description of this business glossary term.
	ShortDescription *string

	// The status of this business glossary term.
	Status types.GlossaryTermStatus

	// The term relations of this business glossary term.
	TermRelations *types.TermRelations

	noSmithyDocumentSerde
}

type CreateGlossaryTermOutput struct {

	// The ID of the Amazon DataZone domain in which this business glossary term is
	// created.
	//
	// This member is required.
	DomainId *string

	// The ID of the business glossary in which this term is created.
	//
	// This member is required.
	GlossaryId *string

	// The ID of this business glossary term.
	//
	// This member is required.
	Id *string

	// The name of this business glossary term.
	//
	// This member is required.
	Name *string

	// The status of this business glossary term.
	//
	// This member is required.
	Status types.GlossaryTermStatus

	// The long description of this business glossary term.
	LongDescription *string

	// The short description of this business glossary term.
	ShortDescription *string

	// The term relations of this business glossary term.
	TermRelations *types.TermRelations

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGlossaryTermMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGlossaryTerm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGlossaryTerm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateGlossaryTerm"); err != nil {
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
	if err = addIdempotencyToken_opCreateGlossaryTermMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateGlossaryTermValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGlossaryTerm(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateGlossaryTerm struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateGlossaryTerm) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateGlossaryTerm) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateGlossaryTermInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateGlossaryTermInput ")
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
func addIdempotencyToken_opCreateGlossaryTermMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateGlossaryTerm{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateGlossaryTerm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateGlossaryTerm",
	}
}
