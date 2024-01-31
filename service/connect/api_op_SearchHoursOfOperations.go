// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches the hours of operation in an Amazon Connect instance, with optional
// filtering.
func (c *Client) SearchHoursOfOperations(ctx context.Context, params *SearchHoursOfOperationsInput, optFns ...func(*Options)) (*SearchHoursOfOperationsOutput, error) {
	if params == nil {
		params = &SearchHoursOfOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchHoursOfOperations", params, optFns, c.addOperationSearchHoursOfOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchHoursOfOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchHoursOfOperationsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return hours of operations.
	SearchCriteria *types.HoursOfOperationSearchCriteria

	// Filters to be applied to search results.
	SearchFilter *types.HoursOfOperationSearchFilter

	noSmithyDocumentSerde
}

type SearchHoursOfOperationsOutput struct {

	// The total number of hours of operations which matched your search query.
	ApproximateTotalCount *int64

	// Information about the hours of operations.
	HoursOfOperations []types.HoursOfOperation

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchHoursOfOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchHoursOfOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchHoursOfOperations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchHoursOfOperations"); err != nil {
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
	if err = addOpSearchHoursOfOperationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchHoursOfOperations(options.Region), middleware.Before); err != nil {
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

// SearchHoursOfOperationsAPIClient is a client that implements the
// SearchHoursOfOperations operation.
type SearchHoursOfOperationsAPIClient interface {
	SearchHoursOfOperations(context.Context, *SearchHoursOfOperationsInput, ...func(*Options)) (*SearchHoursOfOperationsOutput, error)
}

var _ SearchHoursOfOperationsAPIClient = (*Client)(nil)

// SearchHoursOfOperationsPaginatorOptions is the paginator options for
// SearchHoursOfOperations
type SearchHoursOfOperationsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchHoursOfOperationsPaginator is a paginator for SearchHoursOfOperations
type SearchHoursOfOperationsPaginator struct {
	options   SearchHoursOfOperationsPaginatorOptions
	client    SearchHoursOfOperationsAPIClient
	params    *SearchHoursOfOperationsInput
	nextToken *string
	firstPage bool
}

// NewSearchHoursOfOperationsPaginator returns a new
// SearchHoursOfOperationsPaginator
func NewSearchHoursOfOperationsPaginator(client SearchHoursOfOperationsAPIClient, params *SearchHoursOfOperationsInput, optFns ...func(*SearchHoursOfOperationsPaginatorOptions)) *SearchHoursOfOperationsPaginator {
	if params == nil {
		params = &SearchHoursOfOperationsInput{}
	}

	options := SearchHoursOfOperationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchHoursOfOperationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchHoursOfOperationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchHoursOfOperations page.
func (p *SearchHoursOfOperationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchHoursOfOperationsOutput, error) {
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

	result, err := p.client.SearchHoursOfOperations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchHoursOfOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchHoursOfOperations",
	}
}
