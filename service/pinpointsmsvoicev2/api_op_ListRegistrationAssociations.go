// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retreive all of the origination identies that are associated with a
// registration.
func (c *Client) ListRegistrationAssociations(ctx context.Context, params *ListRegistrationAssociationsInput, optFns ...func(*Options)) (*ListRegistrationAssociationsOutput, error) {
	if params == nil {
		params = &ListRegistrationAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRegistrationAssociations", params, optFns, c.addOperationListRegistrationAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRegistrationAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRegistrationAssociationsInput struct {

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// An array of RegistrationAssociationFilter to apply to the results that are
	// returned.
	Filters []types.RegistrationAssociationFilter

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRegistrationAssociationsOutput struct {

	// The Amazon Resource Name (ARN) for the registration.
	//
	// This member is required.
	RegistrationArn *string

	// An array of RegistrationAssociationMetadata objects.
	//
	// This member is required.
	RegistrationAssociations []types.RegistrationAssociationMetadata

	// The unique identifier for the registration.
	//
	// This member is required.
	RegistrationId *string

	// The type of registration form. The list of RegistrationTypes can be found using
	// the DescribeRegistrationTypeDefinitions action.
	//
	// This member is required.
	RegistrationType *string

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRegistrationAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRegistrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRegistrationAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRegistrationAssociations"); err != nil {
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
	if err = addOpListRegistrationAssociationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRegistrationAssociations(options.Region), middleware.Before); err != nil {
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

// ListRegistrationAssociationsAPIClient is a client that implements the
// ListRegistrationAssociations operation.
type ListRegistrationAssociationsAPIClient interface {
	ListRegistrationAssociations(context.Context, *ListRegistrationAssociationsInput, ...func(*Options)) (*ListRegistrationAssociationsOutput, error)
}

var _ ListRegistrationAssociationsAPIClient = (*Client)(nil)

// ListRegistrationAssociationsPaginatorOptions is the paginator options for
// ListRegistrationAssociations
type ListRegistrationAssociationsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRegistrationAssociationsPaginator is a paginator for
// ListRegistrationAssociations
type ListRegistrationAssociationsPaginator struct {
	options   ListRegistrationAssociationsPaginatorOptions
	client    ListRegistrationAssociationsAPIClient
	params    *ListRegistrationAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListRegistrationAssociationsPaginator returns a new
// ListRegistrationAssociationsPaginator
func NewListRegistrationAssociationsPaginator(client ListRegistrationAssociationsAPIClient, params *ListRegistrationAssociationsInput, optFns ...func(*ListRegistrationAssociationsPaginatorOptions)) *ListRegistrationAssociationsPaginator {
	if params == nil {
		params = &ListRegistrationAssociationsInput{}
	}

	options := ListRegistrationAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRegistrationAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRegistrationAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRegistrationAssociations page.
func (p *ListRegistrationAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRegistrationAssociationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRegistrationAssociations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRegistrationAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRegistrationAssociations",
	}
}
