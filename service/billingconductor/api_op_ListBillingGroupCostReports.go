// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A paginated call to retrieve a summary report of actual Amazon Web Services
// charges and the calculated Amazon Web Services charges based on the associated
// pricing plan of a billing group.
func (c *Client) ListBillingGroupCostReports(ctx context.Context, params *ListBillingGroupCostReportsInput, optFns ...func(*Options)) (*ListBillingGroupCostReportsOutput, error) {
	if params == nil {
		params = &ListBillingGroupCostReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBillingGroupCostReports", params, optFns, c.addOperationListBillingGroupCostReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBillingGroupCostReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBillingGroupCostReportsInput struct {

	// The preferred billing period for your report.
	BillingPeriod *string

	// A ListBillingGroupCostReportsFilter to specify billing groups to retrieve
	// reports from.
	Filters *types.ListBillingGroupCostReportsFilter

	// The maximum number of reports to retrieve.
	MaxResults *int32

	// The pagination token that's used on subsequent calls to get reports.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBillingGroupCostReportsOutput struct {

	// A list of BillingGroupCostReportElement retrieved.
	BillingGroupCostReports []types.BillingGroupCostReportElement

	// The pagination token that's used on subsequent calls to get reports.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBillingGroupCostReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBillingGroupCostReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBillingGroupCostReports{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBillingGroupCostReports"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBillingGroupCostReports(options.Region), middleware.Before); err != nil {
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

// ListBillingGroupCostReportsAPIClient is a client that implements the
// ListBillingGroupCostReports operation.
type ListBillingGroupCostReportsAPIClient interface {
	ListBillingGroupCostReports(context.Context, *ListBillingGroupCostReportsInput, ...func(*Options)) (*ListBillingGroupCostReportsOutput, error)
}

var _ ListBillingGroupCostReportsAPIClient = (*Client)(nil)

// ListBillingGroupCostReportsPaginatorOptions is the paginator options for
// ListBillingGroupCostReports
type ListBillingGroupCostReportsPaginatorOptions struct {
	// The maximum number of reports to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBillingGroupCostReportsPaginator is a paginator for
// ListBillingGroupCostReports
type ListBillingGroupCostReportsPaginator struct {
	options   ListBillingGroupCostReportsPaginatorOptions
	client    ListBillingGroupCostReportsAPIClient
	params    *ListBillingGroupCostReportsInput
	nextToken *string
	firstPage bool
}

// NewListBillingGroupCostReportsPaginator returns a new
// ListBillingGroupCostReportsPaginator
func NewListBillingGroupCostReportsPaginator(client ListBillingGroupCostReportsAPIClient, params *ListBillingGroupCostReportsInput, optFns ...func(*ListBillingGroupCostReportsPaginatorOptions)) *ListBillingGroupCostReportsPaginator {
	if params == nil {
		params = &ListBillingGroupCostReportsInput{}
	}

	options := ListBillingGroupCostReportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBillingGroupCostReportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBillingGroupCostReportsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBillingGroupCostReports page.
func (p *ListBillingGroupCostReportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBillingGroupCostReportsOutput, error) {
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

	result, err := p.client.ListBillingGroupCostReports(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBillingGroupCostReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBillingGroupCostReports",
	}
}
