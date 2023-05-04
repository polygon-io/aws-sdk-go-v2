// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of your queues in the current region. The response includes a
// maximum of 1,000 results. If you specify a value for the optional
// QueueNamePrefix parameter, only queues with a name that begins with the
// specified value are returned. The listQueues methods supports pagination. Set
// parameter MaxResults in the request to specify the maximum number of results to
// be returned in the response. If you do not set MaxResults , the response
// includes a maximum of 1,000 results. If you set MaxResults and there are
// additional results to display, the response includes a value for NextToken . Use
// NextToken as a parameter in your next request to listQueues to receive the next
// page of results. Cross-account permissions don't apply to this action. For more
// information, see Grant cross-account permissions to a role and a username (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon SQS Developer Guide.
func (c *Client) ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) {
	if params == nil {
		params = &ListQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueues", params, optFns, c.addOperationListQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueuesInput struct {

	// Maximum number of results to include in the response. Value range is 1 to 1000.
	// You must set MaxResults to receive a value for NextToken in the response.
	MaxResults *int32

	// Pagination token to request the next set of results.
	NextToken *string

	// A string to use for filtering the list results. Only those queues whose name
	// begins with the specified string are returned. Queue URLs and names are
	// case-sensitive.
	QueueNamePrefix *string

	noSmithyDocumentSerde
}

// A list of your queues.
type ListQueuesOutput struct {

	// Pagination token to include in the next request. Token value is null if there
	// are no additional results to request, or if you did not set MaxResults in the
	// request.
	NextToken *string

	// A list of queue URLs, up to 1,000 entries, or the value of MaxResults that you
	// sent in the request.
	QueueUrls []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListQueues{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQueues(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListQueuesAPIClient is a client that implements the ListQueues operation.
type ListQueuesAPIClient interface {
	ListQueues(context.Context, *ListQueuesInput, ...func(*Options)) (*ListQueuesOutput, error)
}

var _ ListQueuesAPIClient = (*Client)(nil)

// ListQueuesPaginatorOptions is the paginator options for ListQueues
type ListQueuesPaginatorOptions struct {
	// Maximum number of results to include in the response. Value range is 1 to 1000.
	// You must set MaxResults to receive a value for NextToken in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQueuesPaginator is a paginator for ListQueues
type ListQueuesPaginator struct {
	options   ListQueuesPaginatorOptions
	client    ListQueuesAPIClient
	params    *ListQueuesInput
	nextToken *string
	firstPage bool
}

// NewListQueuesPaginator returns a new ListQueuesPaginator
func NewListQueuesPaginator(client ListQueuesAPIClient, params *ListQueuesInput, optFns ...func(*ListQueuesPaginatorOptions)) *ListQueuesPaginator {
	if params == nil {
		params = &ListQueuesInput{}
	}

	options := ListQueuesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQueuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQueuesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQueues page.
func (p *ListQueuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQueuesOutput, error) {
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

	result, err := p.client.ListQueues(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQueues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "ListQueues",
	}
}
