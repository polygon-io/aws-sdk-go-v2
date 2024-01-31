// Code generated by smithy-go-codegen DO NOT EDIT.

package keyspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/keyspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of keyspaces.
func (c *Client) ListKeyspaces(ctx context.Context, params *ListKeyspacesInput, optFns ...func(*Options)) (*ListKeyspacesOutput, error) {
	if params == nil {
		params = &ListKeyspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListKeyspaces", params, optFns, c.addOperationListKeyspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListKeyspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListKeyspacesInput struct {

	// The total number of keyspaces to return in the output. If the total number of
	// keyspaces available is more than the value specified, a NextToken is provided
	// in the output. To resume pagination, provide the NextToken value as an argument
	// of a subsequent API invocation.
	MaxResults *int32

	// The pagination token. To resume pagination, provide the NextToken value as
	// argument of a subsequent API invocation.
	NextToken *string

	noSmithyDocumentSerde
}

type ListKeyspacesOutput struct {

	// A list of keyspaces.
	//
	// This member is required.
	Keyspaces []types.KeyspaceSummary

	// A token to specify where to start paginating. This is the NextToken from a
	// previously truncated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListKeyspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListKeyspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListKeyspaces{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListKeyspaces"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListKeyspaces(options.Region), middleware.Before); err != nil {
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

// ListKeyspacesAPIClient is a client that implements the ListKeyspaces operation.
type ListKeyspacesAPIClient interface {
	ListKeyspaces(context.Context, *ListKeyspacesInput, ...func(*Options)) (*ListKeyspacesOutput, error)
}

var _ ListKeyspacesAPIClient = (*Client)(nil)

// ListKeyspacesPaginatorOptions is the paginator options for ListKeyspaces
type ListKeyspacesPaginatorOptions struct {
	// The total number of keyspaces to return in the output. If the total number of
	// keyspaces available is more than the value specified, a NextToken is provided
	// in the output. To resume pagination, provide the NextToken value as an argument
	// of a subsequent API invocation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListKeyspacesPaginator is a paginator for ListKeyspaces
type ListKeyspacesPaginator struct {
	options   ListKeyspacesPaginatorOptions
	client    ListKeyspacesAPIClient
	params    *ListKeyspacesInput
	nextToken *string
	firstPage bool
}

// NewListKeyspacesPaginator returns a new ListKeyspacesPaginator
func NewListKeyspacesPaginator(client ListKeyspacesAPIClient, params *ListKeyspacesInput, optFns ...func(*ListKeyspacesPaginatorOptions)) *ListKeyspacesPaginator {
	if params == nil {
		params = &ListKeyspacesInput{}
	}

	options := ListKeyspacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListKeyspacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListKeyspacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListKeyspaces page.
func (p *ListKeyspacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListKeyspacesOutput, error) {
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

	result, err := p.client.ListKeyspaces(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListKeyspaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListKeyspaces",
	}
}
