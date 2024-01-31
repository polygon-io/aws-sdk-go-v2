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

// Lists Amazon DataZone subscription requests.
func (c *Client) ListSubscriptionRequests(ctx context.Context, params *ListSubscriptionRequestsInput, optFns ...func(*Options)) (*ListSubscriptionRequestsOutput, error) {
	if params == nil {
		params = &ListSubscriptionRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubscriptionRequests", params, optFns, c.addOperationListSubscriptionRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubscriptionRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubscriptionRequestsInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the subscription request approver's project.
	ApproverProjectId *string

	// The maximum number of subscription requests to return in a single call to
	// ListSubscriptionRequests . When the number of subscription requests to be listed
	// is greater than the value of MaxResults , the response contains a NextToken
	// value that you can use in a subsequent call to ListSubscriptionRequests to list
	// the next set of subscription requests.
	MaxResults *int32

	// When the number of subscription requests is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of subscription requests, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListSubscriptionRequests to list the next set of
	// subscription requests.
	NextToken *string

	// The identifier of the project for the subscription requests.
	OwningProjectId *string

	// Specifies the way to sort the results of this action.
	SortBy types.SortKey

	// Specifies the sort order for the results of this action.
	SortOrder types.SortOrder

	// Specifies the status of the subscription requests.
	Status types.SubscriptionRequestStatus

	// The identifier of the subscribed listing.
	SubscribedListingId *string

	noSmithyDocumentSerde
}

type ListSubscriptionRequestsOutput struct {

	// The results of the ListSubscriptionRequests action.
	//
	// This member is required.
	Items []types.SubscriptionRequestSummary

	// When the number of subscription requests is greater than the default value for
	// the MaxResults parameter, or if you explicitly specify a value for MaxResults
	// that is less than the number of subscription requests, the response includes a
	// pagination token named NextToken . You can specify this NextToken value in a
	// subsequent call to ListSubscriptionRequests to list the next set of
	// subscription requests.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubscriptionRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSubscriptionRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSubscriptionRequests{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSubscriptionRequests"); err != nil {
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
	if err = addOpListSubscriptionRequestsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptionRequests(options.Region), middleware.Before); err != nil {
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

// ListSubscriptionRequestsAPIClient is a client that implements the
// ListSubscriptionRequests operation.
type ListSubscriptionRequestsAPIClient interface {
	ListSubscriptionRequests(context.Context, *ListSubscriptionRequestsInput, ...func(*Options)) (*ListSubscriptionRequestsOutput, error)
}

var _ ListSubscriptionRequestsAPIClient = (*Client)(nil)

// ListSubscriptionRequestsPaginatorOptions is the paginator options for
// ListSubscriptionRequests
type ListSubscriptionRequestsPaginatorOptions struct {
	// The maximum number of subscription requests to return in a single call to
	// ListSubscriptionRequests . When the number of subscription requests to be listed
	// is greater than the value of MaxResults , the response contains a NextToken
	// value that you can use in a subsequent call to ListSubscriptionRequests to list
	// the next set of subscription requests.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubscriptionRequestsPaginator is a paginator for ListSubscriptionRequests
type ListSubscriptionRequestsPaginator struct {
	options   ListSubscriptionRequestsPaginatorOptions
	client    ListSubscriptionRequestsAPIClient
	params    *ListSubscriptionRequestsInput
	nextToken *string
	firstPage bool
}

// NewListSubscriptionRequestsPaginator returns a new
// ListSubscriptionRequestsPaginator
func NewListSubscriptionRequestsPaginator(client ListSubscriptionRequestsAPIClient, params *ListSubscriptionRequestsInput, optFns ...func(*ListSubscriptionRequestsPaginatorOptions)) *ListSubscriptionRequestsPaginator {
	if params == nil {
		params = &ListSubscriptionRequestsInput{}
	}

	options := ListSubscriptionRequestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubscriptionRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubscriptionRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubscriptionRequests page.
func (p *ListSubscriptionRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubscriptionRequestsOutput, error) {
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

	result, err := p.client.ListSubscriptionRequests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSubscriptionRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSubscriptionRequests",
	}
}
