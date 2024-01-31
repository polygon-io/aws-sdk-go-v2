// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListQualificationRequests operation retrieves requests for Qualifications
// of a particular Qualification type. The owner of the Qualification type calls
// this operation to poll for pending requests, and accepts them using the
// AcceptQualification operation.
func (c *Client) ListQualificationRequests(ctx context.Context, params *ListQualificationRequestsInput, optFns ...func(*Options)) (*ListQualificationRequestsOutput, error) {
	if params == nil {
		params = &ListQualificationRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQualificationRequests", params, optFns, c.addOperationListQualificationRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQualificationRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQualificationRequestsInput struct {

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The ID of the QualificationType.
	QualificationTypeId *string

	noSmithyDocumentSerde
}

type ListQualificationRequestsOutput struct {

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of Qualification requests on this page in the filtered results list,
	// equivalent to the number of Qualification requests being returned by this call.
	NumResults *int32

	// The Qualification request. The response includes one QualificationRequest
	// element for each Qualification request returned by the query.
	QualificationRequests []types.QualificationRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListQualificationRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListQualificationRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListQualificationRequests{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListQualificationRequests"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListQualificationRequests(options.Region), middleware.Before); err != nil {
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

// ListQualificationRequestsAPIClient is a client that implements the
// ListQualificationRequests operation.
type ListQualificationRequestsAPIClient interface {
	ListQualificationRequests(context.Context, *ListQualificationRequestsInput, ...func(*Options)) (*ListQualificationRequestsOutput, error)
}

var _ ListQualificationRequestsAPIClient = (*Client)(nil)

// ListQualificationRequestsPaginatorOptions is the paginator options for
// ListQualificationRequests
type ListQualificationRequestsPaginatorOptions struct {
	// The maximum number of results to return in a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListQualificationRequestsPaginator is a paginator for ListQualificationRequests
type ListQualificationRequestsPaginator struct {
	options   ListQualificationRequestsPaginatorOptions
	client    ListQualificationRequestsAPIClient
	params    *ListQualificationRequestsInput
	nextToken *string
	firstPage bool
}

// NewListQualificationRequestsPaginator returns a new
// ListQualificationRequestsPaginator
func NewListQualificationRequestsPaginator(client ListQualificationRequestsAPIClient, params *ListQualificationRequestsInput, optFns ...func(*ListQualificationRequestsPaginatorOptions)) *ListQualificationRequestsPaginator {
	if params == nil {
		params = &ListQualificationRequestsInput{}
	}

	options := ListQualificationRequestsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListQualificationRequestsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListQualificationRequestsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListQualificationRequests page.
func (p *ListQualificationRequestsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListQualificationRequestsOutput, error) {
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

	result, err := p.client.ListQualificationRequests(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListQualificationRequests(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListQualificationRequests",
	}
}
