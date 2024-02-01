// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// To use this API operation, your IAM role must have permissions to perform the
// ListAggregatedUtterances (https://docs.aws.amazon.com/lexv2/latest/APIReference/API_ListAggregatedUtterances.html)
// operation, which provides access to utterance-related analytics. See Viewing
// utterance statistics (https://docs.aws.amazon.com/lexv2/latest/dg/monitoring-utterances.html)
// for the IAM policy to apply to the IAM role. Retrieves summary metrics for the
// utterances in your bot. The following fields are required:
//   - metrics – A list of AnalyticsUtteranceMetric (https://docs.aws.amazon.com/lexv2/latest/APIReference/API_AnalyticsUtteranceMetric.html)
//     objects. In each object, use the name field to specify the metric to
//     calculate, the statistic field to specify whether to calculate the Sum ,
//     Average , or Max number, and the order field to specify whether to sort the
//     results in Ascending or Descending order.
//   - startDateTime and endDateTime – Define a time range for which you want to
//     retrieve results.
//
// Of the optional fields, you can organize the results in the following ways:
//   - Use the filters field to filter the results, the groupBy field to specify
//     categories by which to group the results, and the binBy field to specify time
//     intervals by which to group the results.
//   - Use the maxResults field to limit the number of results to return in a
//     single response and the nextToken field to return the next batch of results if
//     the response does not return the full set of results.
//
// Note that an order field exists in both binBy and metrics . Currently, you can
// specify it in either field, but not in both.
func (c *Client) ListUtteranceMetrics(ctx context.Context, params *ListUtteranceMetricsInput, optFns ...func(*Options)) (*ListUtteranceMetricsOutput, error) {
	if params == nil {
		params = &ListUtteranceMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListUtteranceMetrics", params, optFns, c.addOperationListUtteranceMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListUtteranceMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListUtteranceMetricsInput struct {

	// The identifier for the bot for which you want to retrieve utterance metrics.
	//
	// This member is required.
	BotId *string

	// The date and time that marks the end of the range of time for which you want to
	// see utterance metrics.
	//
	// This member is required.
	EndDateTime *time.Time

	// A list of objects, each of which contains a metric you want to list, the
	// statistic for the metric you want to return, and the method by which to organize
	// the results.
	//
	// This member is required.
	Metrics []types.AnalyticsUtteranceMetric

	// The date and time that marks the beginning of the range of time for which you
	// want to see utterance metrics.
	//
	// This member is required.
	StartDateTime *time.Time

	// A list containing attributes related to the utterance that you want the
	// response to return. The following attributes are possible:
	//   - LastUsedIntent – The last used intent at the time of the utterance.
	Attributes []types.AnalyticsUtteranceAttribute

	// A list of objects, each of which contains specifications for organizing the
	// results by time.
	BinBy []types.AnalyticsBinBySpecification

	// A list of objects, each of which describes a condition by which you want to
	// filter the results.
	Filters []types.AnalyticsUtteranceFilter

	// A list of objects, each of which specifies how to group the results. You can
	// group by the following criteria:
	//   - UtteranceText – The transcription of the utterance.
	//   - UtteranceState – The state of the utterance. The possible states are
	//   detailed in Key definitions (https://docs.aws.amazon.com/analytics-key-definitions-utterances)
	//   in the user guide.
	GroupBy []types.AnalyticsUtteranceGroupBySpecification

	// The maximum number of results to return in each page of results. If there are
	// fewer results than the maximum page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListUtteranceMetrics operation contains more results
	// than specified in the maxResults parameter, a token is returned in the response.
	// Use the returned token in the nextToken parameter of a ListUtteranceMetrics
	// request to return the next page of results. For a complete set of results, call
	// the ListUtteranceMetrics operation until the nextToken returned in the response
	// is null.
	NextToken *string

	noSmithyDocumentSerde
}

type ListUtteranceMetricsOutput struct {

	// The identifier for the bot for which you retrieved utterance metrics.
	BotId *string

	// If the response from the ListUtteranceMetrics operation contains more results
	// than specified in the maxResults parameter, a token is returned in the response.
	// Use the returned token in the nextToken parameter of a ListUtteranceMetrics
	// request to return the next page of results. For a complete set of results, call
	// the ListUtteranceMetrics operation until the nextToken returned in the response
	// is null.
	NextToken *string

	// The results for the utterance metrics.
	Results []types.AnalyticsUtteranceResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListUtteranceMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListUtteranceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListUtteranceMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListUtteranceMetrics"); err != nil {
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
	if err = addOpListUtteranceMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListUtteranceMetrics(options.Region), middleware.Before); err != nil {
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

// ListUtteranceMetricsAPIClient is a client that implements the
// ListUtteranceMetrics operation.
type ListUtteranceMetricsAPIClient interface {
	ListUtteranceMetrics(context.Context, *ListUtteranceMetricsInput, ...func(*Options)) (*ListUtteranceMetricsOutput, error)
}

var _ ListUtteranceMetricsAPIClient = (*Client)(nil)

// ListUtteranceMetricsPaginatorOptions is the paginator options for
// ListUtteranceMetrics
type ListUtteranceMetricsPaginatorOptions struct {
	// The maximum number of results to return in each page of results. If there are
	// fewer results than the maximum page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListUtteranceMetricsPaginator is a paginator for ListUtteranceMetrics
type ListUtteranceMetricsPaginator struct {
	options   ListUtteranceMetricsPaginatorOptions
	client    ListUtteranceMetricsAPIClient
	params    *ListUtteranceMetricsInput
	nextToken *string
	firstPage bool
}

// NewListUtteranceMetricsPaginator returns a new ListUtteranceMetricsPaginator
func NewListUtteranceMetricsPaginator(client ListUtteranceMetricsAPIClient, params *ListUtteranceMetricsInput, optFns ...func(*ListUtteranceMetricsPaginatorOptions)) *ListUtteranceMetricsPaginator {
	if params == nil {
		params = &ListUtteranceMetricsInput{}
	}

	options := ListUtteranceMetricsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListUtteranceMetricsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListUtteranceMetricsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListUtteranceMetrics page.
func (p *ListUtteranceMetricsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListUtteranceMetricsOutput, error) {
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

	result, err := p.client.ListUtteranceMetrics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListUtteranceMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListUtteranceMetrics",
	}
}
