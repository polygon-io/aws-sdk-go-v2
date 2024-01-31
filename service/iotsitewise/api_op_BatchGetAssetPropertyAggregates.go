// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets aggregated values (for example, average, minimum, and maximum) for one or
// more asset properties. For more information, see Querying aggregates (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/query-industrial-data.html#aggregates)
// in the IoT SiteWise User Guide.
func (c *Client) BatchGetAssetPropertyAggregates(ctx context.Context, params *BatchGetAssetPropertyAggregatesInput, optFns ...func(*Options)) (*BatchGetAssetPropertyAggregatesOutput, error) {
	if params == nil {
		params = &BatchGetAssetPropertyAggregatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetAssetPropertyAggregates", params, optFns, c.addOperationBatchGetAssetPropertyAggregatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetAssetPropertyAggregatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetAssetPropertyAggregatesInput struct {

	// The list of asset property aggregate entries for the batch get request. You can
	// specify up to 16 entries per request.
	//
	// This member is required.
	Entries []types.BatchGetAssetPropertyAggregatesEntry

	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 1 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 4000.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type BatchGetAssetPropertyAggregatesOutput struct {

	// A list of the errors (if any) associated with the batch request. Each error
	// entry contains the entryId of the entry that failed.
	//
	// This member is required.
	ErrorEntries []types.BatchGetAssetPropertyAggregatesErrorEntry

	// A list of entries that were not processed by this batch request. because these
	// entries had been completely processed by previous paginated requests. Each
	// skipped entry contains the entryId of the entry that skipped.
	//
	// This member is required.
	SkippedEntries []types.BatchGetAssetPropertyAggregatesSkippedEntry

	// A list of entries that were processed successfully by this batch request. Each
	// success entry contains the entryId of the entry that succeeded and the latest
	// query result.
	//
	// This member is required.
	SuccessEntries []types.BatchGetAssetPropertyAggregatesSuccessEntry

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetAssetPropertyAggregatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetAssetPropertyAggregates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetAssetPropertyAggregates{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetAssetPropertyAggregates"); err != nil {
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
	if err = addEndpointPrefix_opBatchGetAssetPropertyAggregatesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchGetAssetPropertyAggregatesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetAssetPropertyAggregates(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchGetAssetPropertyAggregatesMiddleware struct {
}

func (*endpointPrefix_opBatchGetAssetPropertyAggregatesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchGetAssetPropertyAggregatesMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opBatchGetAssetPropertyAggregatesMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opBatchGetAssetPropertyAggregatesMiddleware{}, "ResolveEndpointV2", middleware.After)
}

// BatchGetAssetPropertyAggregatesAPIClient is a client that implements the
// BatchGetAssetPropertyAggregates operation.
type BatchGetAssetPropertyAggregatesAPIClient interface {
	BatchGetAssetPropertyAggregates(context.Context, *BatchGetAssetPropertyAggregatesInput, ...func(*Options)) (*BatchGetAssetPropertyAggregatesOutput, error)
}

var _ BatchGetAssetPropertyAggregatesAPIClient = (*Client)(nil)

// BatchGetAssetPropertyAggregatesPaginatorOptions is the paginator options for
// BatchGetAssetPropertyAggregates
type BatchGetAssetPropertyAggregatesPaginatorOptions struct {
	// The maximum number of results to return for each paginated request. A result
	// set is returned in the two cases, whichever occurs first.
	//   - The size of the result set is equal to 1 MB.
	//   - The number of data points in the result set is equal to the value of
	//   maxResults . The maximum value of maxResults is 4000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// BatchGetAssetPropertyAggregatesPaginator is a paginator for
// BatchGetAssetPropertyAggregates
type BatchGetAssetPropertyAggregatesPaginator struct {
	options   BatchGetAssetPropertyAggregatesPaginatorOptions
	client    BatchGetAssetPropertyAggregatesAPIClient
	params    *BatchGetAssetPropertyAggregatesInput
	nextToken *string
	firstPage bool
}

// NewBatchGetAssetPropertyAggregatesPaginator returns a new
// BatchGetAssetPropertyAggregatesPaginator
func NewBatchGetAssetPropertyAggregatesPaginator(client BatchGetAssetPropertyAggregatesAPIClient, params *BatchGetAssetPropertyAggregatesInput, optFns ...func(*BatchGetAssetPropertyAggregatesPaginatorOptions)) *BatchGetAssetPropertyAggregatesPaginator {
	if params == nil {
		params = &BatchGetAssetPropertyAggregatesInput{}
	}

	options := BatchGetAssetPropertyAggregatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &BatchGetAssetPropertyAggregatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *BatchGetAssetPropertyAggregatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next BatchGetAssetPropertyAggregates page.
func (p *BatchGetAssetPropertyAggregatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*BatchGetAssetPropertyAggregatesOutput, error) {
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

	result, err := p.client.BatchGetAssetPropertyAggregates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opBatchGetAssetPropertyAggregates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetAssetPropertyAggregates",
	}
}
