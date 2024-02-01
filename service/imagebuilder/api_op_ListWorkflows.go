// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists workflow build versions based on filtering parameters.
func (c *Client) ListWorkflows(ctx context.Context, params *ListWorkflowsInput, optFns ...func(*Options)) (*ListWorkflowsOutput, error) {
	if params == nil {
		params = &ListWorkflowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkflows", params, optFns, c.addOperationListWorkflowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkflowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowsInput struct {

	// Specify all or part of the workflow name to streamline results.
	ByName bool

	// Used to streamline search results.
	Filters []types.Filter

	// The maximum items to return in a request.
	MaxResults *int32

	// A token to specify where to start paginating. This is the nextToken from a
	// previously truncated response.
	NextToken *string

	// Used to get a list of workflow build version filtered by the identity of the
	// creator.
	Owner types.Ownership

	noSmithyDocumentSerde
}

type ListWorkflowsOutput struct {

	// The next token used for paginated responses. When this field isn't empty, there
	// are additional elements that the service hasn't included in this request. Use
	// this token with the next request to retrieve additional objects.
	NextToken *string

	// A list of workflow build versions that match the request criteria.
	WorkflowVersionList []types.WorkflowVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkflowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListWorkflows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListWorkflows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListWorkflows"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflows(options.Region), middleware.Before); err != nil {
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

// ListWorkflowsAPIClient is a client that implements the ListWorkflows operation.
type ListWorkflowsAPIClient interface {
	ListWorkflows(context.Context, *ListWorkflowsInput, ...func(*Options)) (*ListWorkflowsOutput, error)
}

var _ ListWorkflowsAPIClient = (*Client)(nil)

// ListWorkflowsPaginatorOptions is the paginator options for ListWorkflows
type ListWorkflowsPaginatorOptions struct {
	// The maximum items to return in a request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkflowsPaginator is a paginator for ListWorkflows
type ListWorkflowsPaginator struct {
	options   ListWorkflowsPaginatorOptions
	client    ListWorkflowsAPIClient
	params    *ListWorkflowsInput
	nextToken *string
	firstPage bool
}

// NewListWorkflowsPaginator returns a new ListWorkflowsPaginator
func NewListWorkflowsPaginator(client ListWorkflowsAPIClient, params *ListWorkflowsInput, optFns ...func(*ListWorkflowsPaginatorOptions)) *ListWorkflowsPaginator {
	if params == nil {
		params = &ListWorkflowsInput{}
	}

	options := ListWorkflowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkflowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkflowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkflows page.
func (p *ListWorkflowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkflowsOutput, error) {
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

	result, err := p.client.ListWorkflows(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkflows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListWorkflows",
	}
}
