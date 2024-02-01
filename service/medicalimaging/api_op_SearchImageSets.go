// Code generated by smithy-go-codegen DO NOT EDIT.

package medicalimaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medicalimaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Search image sets based on defined input attributes. SearchImageSets accepts a
// single search query parameter and returns a paginated response of all image sets
// that have the matching criteria. All range queries must be input as
// (lowerBound, upperBound) . SearchImageSets uses the updatedAt field for sorting
// in decreasing order from latest to oldest.
func (c *Client) SearchImageSets(ctx context.Context, params *SearchImageSetsInput, optFns ...func(*Options)) (*SearchImageSetsOutput, error) {
	if params == nil {
		params = &SearchImageSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchImageSets", params, optFns, c.addOperationSearchImageSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchImageSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchImageSetsInput struct {

	// The identifier of the data store where the image sets reside.
	//
	// This member is required.
	DatastoreId *string

	// The maximum number of results that can be returned in a search.
	MaxResults *int32

	// The token used for pagination of results returned in the response. Use the
	// token returned from the previous request to continue results where the previous
	// request ended.
	NextToken *string

	// The search criteria that filters by applying a maximum of 1 item to
	// SearchByAttribute .
	SearchCriteria *types.SearchCriteria

	noSmithyDocumentSerde
}

type SearchImageSetsOutput struct {

	// The model containing the image set results.
	//
	// This member is required.
	ImageSetsMetadataSummaries []types.ImageSetsMetadataSummary

	// The token for pagination results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchImageSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchImageSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchImageSets{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchImageSets"); err != nil {
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
	if err = addEndpointPrefix_opSearchImageSetsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSearchImageSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchImageSets(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opSearchImageSetsMiddleware struct {
}

func (*endpointPrefix_opSearchImageSetsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opSearchImageSetsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "runtime-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opSearchImageSetsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opSearchImageSetsMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// SearchImageSetsAPIClient is a client that implements the SearchImageSets
// operation.
type SearchImageSetsAPIClient interface {
	SearchImageSets(context.Context, *SearchImageSetsInput, ...func(*Options)) (*SearchImageSetsOutput, error)
}

var _ SearchImageSetsAPIClient = (*Client)(nil)

// SearchImageSetsPaginatorOptions is the paginator options for SearchImageSets
type SearchImageSetsPaginatorOptions struct {
	// The maximum number of results that can be returned in a search.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchImageSetsPaginator is a paginator for SearchImageSets
type SearchImageSetsPaginator struct {
	options   SearchImageSetsPaginatorOptions
	client    SearchImageSetsAPIClient
	params    *SearchImageSetsInput
	nextToken *string
	firstPage bool
}

// NewSearchImageSetsPaginator returns a new SearchImageSetsPaginator
func NewSearchImageSetsPaginator(client SearchImageSetsAPIClient, params *SearchImageSetsInput, optFns ...func(*SearchImageSetsPaginatorOptions)) *SearchImageSetsPaginator {
	if params == nil {
		params = &SearchImageSetsInput{}
	}

	options := SearchImageSetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchImageSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchImageSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchImageSets page.
func (p *SearchImageSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchImageSetsOutput, error) {
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

	result, err := p.client.SearchImageSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchImageSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchImageSets",
	}
}
