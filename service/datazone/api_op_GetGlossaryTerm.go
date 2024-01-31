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
	"time"
)

// Gets a business glossary term in Amazon DataZone.
func (c *Client) GetGlossaryTerm(ctx context.Context, params *GetGlossaryTermInput, optFns ...func(*Options)) (*GetGlossaryTermOutput, error) {
	if params == nil {
		params = &GetGlossaryTermInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGlossaryTerm", params, optFns, c.addOperationGetGlossaryTermMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGlossaryTermOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGlossaryTermInput struct {

	// The ID of the Amazon DataZone domain in which this business glossary term
	// exists.
	//
	// This member is required.
	DomainIdentifier *string

	// The ID of the business glossary term.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetGlossaryTermOutput struct {

	// The ID of the Amazon DataZone domain in which this business glossary term
	// exists.
	//
	// This member is required.
	DomainId *string

	// The ID of the business glossary to which this term belongs.
	//
	// This member is required.
	GlossaryId *string

	// The ID of the business glossary term.
	//
	// This member is required.
	Id *string

	// The name of the business glossary term.
	//
	// This member is required.
	Name *string

	// The status of the business glossary term.
	//
	// This member is required.
	Status types.GlossaryTermStatus

	// The timestamp of when the business glossary term was created.
	CreatedAt *time.Time

	// The Amazon DataZone user who created the business glossary.
	CreatedBy *string

	// The long description of the business glossary term.
	LongDescription *string

	// The short decription of the business glossary term.
	ShortDescription *string

	// The relations of the business glossary term.
	TermRelations *types.TermRelations

	// The timestamp of when the business glossary term was updated.
	UpdatedAt *time.Time

	// The Amazon DataZone user who updated the business glossary term.
	UpdatedBy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGlossaryTermMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGlossaryTerm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGlossaryTerm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGlossaryTerm"); err != nil {
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
	if err = addOpGetGlossaryTermValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGlossaryTerm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGlossaryTerm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGlossaryTerm",
	}
}
