// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists information about assistants.
func (c *Client) ListAssistants(ctx context.Context, params *ListAssistantsInput, optFns ...func(*Options)) (*ListAssistantsOutput, error) {
	if params == nil {
		params = &ListAssistantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssistants", params, optFns, c.addOperationListAssistantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssistantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssistantsInput struct {

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssistantsOutput struct {

	// Information about the assistants.
	//
	// This member is required.
	AssistantSummaries []types.AssistantSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssistantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAssistants{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAssistants{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAssistants"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssistants(options.Region), middleware.Before); err != nil {
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

// ListAssistantsAPIClient is a client that implements the ListAssistants
// operation.
type ListAssistantsAPIClient interface {
	ListAssistants(context.Context, *ListAssistantsInput, ...func(*Options)) (*ListAssistantsOutput, error)
}

var _ ListAssistantsAPIClient = (*Client)(nil)

// ListAssistantsPaginatorOptions is the paginator options for ListAssistants
type ListAssistantsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssistantsPaginator is a paginator for ListAssistants
type ListAssistantsPaginator struct {
	options   ListAssistantsPaginatorOptions
	client    ListAssistantsAPIClient
	params    *ListAssistantsInput
	nextToken *string
	firstPage bool
}

// NewListAssistantsPaginator returns a new ListAssistantsPaginator
func NewListAssistantsPaginator(client ListAssistantsAPIClient, params *ListAssistantsInput, optFns ...func(*ListAssistantsPaginatorOptions)) *ListAssistantsPaginator {
	if params == nil {
		params = &ListAssistantsInput{}
	}

	options := ListAssistantsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssistantsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssistantsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssistants page.
func (p *ListAssistantsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssistantsOutput, error) {
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

	result, err := p.client.ListAssistants(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssistants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAssistants",
	}
}
