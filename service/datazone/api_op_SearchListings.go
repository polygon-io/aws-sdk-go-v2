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

// Searches listings in Amazon DataZone.
func (c *Client) SearchListings(ctx context.Context, params *SearchListingsInput, optFns ...func(*Options)) (*SearchListingsOutput, error) {
	if params == nil {
		params = &SearchListingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchListings", params, optFns, c.addOperationSearchListingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchListingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchListingsInput struct {

	// The identifier of the domain in which to search listings.
	//
	// This member is required.
	DomainIdentifier *string

	// Specifies additional attributes for the search.
	AdditionalAttributes []types.SearchOutputAdditionalAttribute

	// Specifies the filters for the search of listings.
	Filters types.FilterClause

	// The maximum number of results to return in a single call to SearchListings .
	// When the number of results to be listed is greater than the value of MaxResults
	// , the response contains a NextToken value that you can use in a subsequent call
	// to SearchListings to list the next set of results.
	MaxResults *int32

	// When the number of results is greater than the default value for the MaxResults
	// parameter, or if you explicitly specify a value for MaxResults that is less
	// than the number of results, the response includes a pagination token named
	// NextToken . You can specify this NextToken value in a subsequent call to
	// SearchListings to list the next set of results.
	NextToken *string

	//
	SearchIn []types.SearchInItem

	// Specifies the text for which to search.
	SearchText *string

	// Specifies the way for sorting the search results.
	Sort *types.SearchSort

	noSmithyDocumentSerde
}

type SearchListingsOutput struct {

	// The results of the SearchListings action.
	Items []types.SearchResultItem

	// When the number of results is greater than the default value for the MaxResults
	// parameter, or if you explicitly specify a value for MaxResults that is less
	// than the number of results, the response includes a pagination token named
	// NextToken . You can specify this NextToken value in a subsequent call to
	// SearchListings to list the next set of results.
	NextToken *string

	// Total number of search results.
	TotalMatchCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchListingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchListings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchListings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchListings"); err != nil {
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
	if err = addOpSearchListingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchListings(options.Region), middleware.Before); err != nil {
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

// SearchListingsAPIClient is a client that implements the SearchListings
// operation.
type SearchListingsAPIClient interface {
	SearchListings(context.Context, *SearchListingsInput, ...func(*Options)) (*SearchListingsOutput, error)
}

var _ SearchListingsAPIClient = (*Client)(nil)

// SearchListingsPaginatorOptions is the paginator options for SearchListings
type SearchListingsPaginatorOptions struct {
	// The maximum number of results to return in a single call to SearchListings .
	// When the number of results to be listed is greater than the value of MaxResults
	// , the response contains a NextToken value that you can use in a subsequent call
	// to SearchListings to list the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchListingsPaginator is a paginator for SearchListings
type SearchListingsPaginator struct {
	options   SearchListingsPaginatorOptions
	client    SearchListingsAPIClient
	params    *SearchListingsInput
	nextToken *string
	firstPage bool
}

// NewSearchListingsPaginator returns a new SearchListingsPaginator
func NewSearchListingsPaginator(client SearchListingsAPIClient, params *SearchListingsInput, optFns ...func(*SearchListingsPaginatorOptions)) *SearchListingsPaginator {
	if params == nil {
		params = &SearchListingsInput{}
	}

	options := SearchListingsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchListingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchListingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchListings page.
func (p *SearchListingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchListingsOutput, error) {
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

	result, err := p.client.SearchListings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchListings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchListings",
	}
}
