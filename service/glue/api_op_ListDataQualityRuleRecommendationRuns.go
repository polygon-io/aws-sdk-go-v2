// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the recommendation runs meeting the filter criteria.
func (c *Client) ListDataQualityRuleRecommendationRuns(ctx context.Context, params *ListDataQualityRuleRecommendationRunsInput, optFns ...func(*Options)) (*ListDataQualityRuleRecommendationRunsOutput, error) {
	if params == nil {
		params = &ListDataQualityRuleRecommendationRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataQualityRuleRecommendationRuns", params, optFns, c.addOperationListDataQualityRuleRecommendationRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataQualityRuleRecommendationRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataQualityRuleRecommendationRunsInput struct {

	// The filter criteria.
	Filter *types.DataQualityRuleRecommendationRunFilter

	// The maximum number of results to return.
	MaxResults *int32

	// A paginated token to offset the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDataQualityRuleRecommendationRunsOutput struct {

	// A pagination token, if more results are available.
	NextToken *string

	// A list of DataQualityRuleRecommendationRunDescription objects.
	Runs []types.DataQualityRuleRecommendationRunDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataQualityRuleRecommendationRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDataQualityRuleRecommendationRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDataQualityRuleRecommendationRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDataQualityRuleRecommendationRuns"); err != nil {
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
	if err = addOpListDataQualityRuleRecommendationRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataQualityRuleRecommendationRuns(options.Region), middleware.Before); err != nil {
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

// ListDataQualityRuleRecommendationRunsAPIClient is a client that implements the
// ListDataQualityRuleRecommendationRuns operation.
type ListDataQualityRuleRecommendationRunsAPIClient interface {
	ListDataQualityRuleRecommendationRuns(context.Context, *ListDataQualityRuleRecommendationRunsInput, ...func(*Options)) (*ListDataQualityRuleRecommendationRunsOutput, error)
}

var _ ListDataQualityRuleRecommendationRunsAPIClient = (*Client)(nil)

// ListDataQualityRuleRecommendationRunsPaginatorOptions is the paginator options
// for ListDataQualityRuleRecommendationRuns
type ListDataQualityRuleRecommendationRunsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataQualityRuleRecommendationRunsPaginator is a paginator for
// ListDataQualityRuleRecommendationRuns
type ListDataQualityRuleRecommendationRunsPaginator struct {
	options   ListDataQualityRuleRecommendationRunsPaginatorOptions
	client    ListDataQualityRuleRecommendationRunsAPIClient
	params    *ListDataQualityRuleRecommendationRunsInput
	nextToken *string
	firstPage bool
}

// NewListDataQualityRuleRecommendationRunsPaginator returns a new
// ListDataQualityRuleRecommendationRunsPaginator
func NewListDataQualityRuleRecommendationRunsPaginator(client ListDataQualityRuleRecommendationRunsAPIClient, params *ListDataQualityRuleRecommendationRunsInput, optFns ...func(*ListDataQualityRuleRecommendationRunsPaginatorOptions)) *ListDataQualityRuleRecommendationRunsPaginator {
	if params == nil {
		params = &ListDataQualityRuleRecommendationRunsInput{}
	}

	options := ListDataQualityRuleRecommendationRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataQualityRuleRecommendationRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataQualityRuleRecommendationRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataQualityRuleRecommendationRuns page.
func (p *ListDataQualityRuleRecommendationRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataQualityRuleRecommendationRunsOutput, error) {
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

	result, err := p.client.ListDataQualityRuleRecommendationRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataQualityRuleRecommendationRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDataQualityRuleRecommendationRuns",
	}
}
