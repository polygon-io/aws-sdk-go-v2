// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the resource sets in an account.
func (c *Client) ListResourceSets(ctx context.Context, params *ListResourceSetsInput, optFns ...func(*Options)) (*ListResourceSetsOutput, error) {
	if params == nil {
		params = &ListResourceSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourceSets", params, optFns, c.addOperationListResourceSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourceSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourceSetsInput struct {

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResourceSetsOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// A list of resource sets associated with the account.
	ResourceSets []types.ResourceSetOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResourceSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResourceSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResourceSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResourceSets"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourceSets(options.Region), middleware.Before); err != nil {
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

// ListResourceSetsAPIClient is a client that implements the ListResourceSets
// operation.
type ListResourceSetsAPIClient interface {
	ListResourceSets(context.Context, *ListResourceSetsInput, ...func(*Options)) (*ListResourceSetsOutput, error)
}

var _ ListResourceSetsAPIClient = (*Client)(nil)

// ListResourceSetsPaginatorOptions is the paginator options for ListResourceSets
type ListResourceSetsPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResourceSetsPaginator is a paginator for ListResourceSets
type ListResourceSetsPaginator struct {
	options   ListResourceSetsPaginatorOptions
	client    ListResourceSetsAPIClient
	params    *ListResourceSetsInput
	nextToken *string
	firstPage bool
}

// NewListResourceSetsPaginator returns a new ListResourceSetsPaginator
func NewListResourceSetsPaginator(client ListResourceSetsAPIClient, params *ListResourceSetsInput, optFns ...func(*ListResourceSetsPaginatorOptions)) *ListResourceSetsPaginator {
	if params == nil {
		params = &ListResourceSetsInput{}
	}

	options := ListResourceSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResourceSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResourceSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResourceSets page.
func (p *ListResourceSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResourceSetsOutput, error) {
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

	result, err := p.client.ListResourceSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResourceSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResourceSets",
	}
}
