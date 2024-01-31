// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of all mitigation actions that match the specified filter criteria.
// Requires permission to access the ListMitigationActions (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListMitigationActions(ctx context.Context, params *ListMitigationActionsInput, optFns ...func(*Options)) (*ListMitigationActionsOutput, error) {
	if params == nil {
		params = &ListMitigationActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMitigationActions", params, optFns, c.addOperationListMitigationActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMitigationActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMitigationActionsInput struct {

	// Specify a value to limit the result to mitigation actions with a specific
	// action type.
	ActionType types.MitigationActionType

	// The maximum number of results to return at one time. The default is 25.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListMitigationActionsOutput struct {

	// A set of actions that matched the specified filter criteria.
	ActionIdentifiers []types.MitigationActionIdentifier

	// The token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMitigationActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMitigationActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMitigationActions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMitigationActions"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMitigationActions(options.Region), middleware.Before); err != nil {
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

// ListMitigationActionsAPIClient is a client that implements the
// ListMitigationActions operation.
type ListMitigationActionsAPIClient interface {
	ListMitigationActions(context.Context, *ListMitigationActionsInput, ...func(*Options)) (*ListMitigationActionsOutput, error)
}

var _ ListMitigationActionsAPIClient = (*Client)(nil)

// ListMitigationActionsPaginatorOptions is the paginator options for
// ListMitigationActions
type ListMitigationActionsPaginatorOptions struct {
	// The maximum number of results to return at one time. The default is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListMitigationActionsPaginator is a paginator for ListMitigationActions
type ListMitigationActionsPaginator struct {
	options   ListMitigationActionsPaginatorOptions
	client    ListMitigationActionsAPIClient
	params    *ListMitigationActionsInput
	nextToken *string
	firstPage bool
}

// NewListMitigationActionsPaginator returns a new ListMitigationActionsPaginator
func NewListMitigationActionsPaginator(client ListMitigationActionsAPIClient, params *ListMitigationActionsInput, optFns ...func(*ListMitigationActionsPaginatorOptions)) *ListMitigationActionsPaginator {
	if params == nil {
		params = &ListMitigationActionsInput{}
	}

	options := ListMitigationActionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListMitigationActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListMitigationActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListMitigationActions page.
func (p *ListMitigationActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListMitigationActionsOutput, error) {
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

	result, err := p.client.ListMitigationActions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListMitigationActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMitigationActions",
	}
}
