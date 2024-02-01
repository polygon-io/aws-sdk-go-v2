// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action returns the following for a given blockchain network:
//   - Lists all token balances owned by an address (either a contract address or
//     a wallet address).
//   - Lists all token balances for all tokens created by a contract.
//   - Lists all token balances for a given token.
//
// You must always specify the network property of the tokenFilter when using this
// operation.
func (c *Client) ListTokenBalances(ctx context.Context, params *ListTokenBalancesInput, optFns ...func(*Options)) (*ListTokenBalancesOutput, error) {
	if params == nil {
		params = &ListTokenBalancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTokenBalances", params, optFns, c.addOperationListTokenBalancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTokenBalancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTokenBalancesInput struct {

	// The contract address or a token identifier on the blockchain network by which
	// to filter the request. You must specify the contractAddress property of this
	// container when listing tokens minted by a contract. You must always specify the
	// network property of this container when using this operation.
	//
	// This member is required.
	TokenFilter *types.TokenFilter

	// The maximum number of token balances to return.
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The contract or wallet address on the blockchain network by which to filter the
	// request. You must specify the address property of the ownerFilter when listing
	// balances of tokens owned by the address.
	OwnerFilter *types.OwnerFilter

	noSmithyDocumentSerde
}

type ListTokenBalancesOutput struct {

	// An array of TokenBalance objects. Each object contains details about the token
	// balance.
	//
	// This member is required.
	TokenBalances []types.TokenBalance

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTokenBalancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTokenBalances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTokenBalances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListTokenBalances"); err != nil {
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
	if err = addOpListTokenBalancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTokenBalances(options.Region), middleware.Before); err != nil {
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

// ListTokenBalancesAPIClient is a client that implements the ListTokenBalances
// operation.
type ListTokenBalancesAPIClient interface {
	ListTokenBalances(context.Context, *ListTokenBalancesInput, ...func(*Options)) (*ListTokenBalancesOutput, error)
}

var _ ListTokenBalancesAPIClient = (*Client)(nil)

// ListTokenBalancesPaginatorOptions is the paginator options for ListTokenBalances
type ListTokenBalancesPaginatorOptions struct {
	// The maximum number of token balances to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTokenBalancesPaginator is a paginator for ListTokenBalances
type ListTokenBalancesPaginator struct {
	options   ListTokenBalancesPaginatorOptions
	client    ListTokenBalancesAPIClient
	params    *ListTokenBalancesInput
	nextToken *string
	firstPage bool
}

// NewListTokenBalancesPaginator returns a new ListTokenBalancesPaginator
func NewListTokenBalancesPaginator(client ListTokenBalancesAPIClient, params *ListTokenBalancesInput, optFns ...func(*ListTokenBalancesPaginatorOptions)) *ListTokenBalancesPaginator {
	if params == nil {
		params = &ListTokenBalancesInput{}
	}

	options := ListTokenBalancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTokenBalancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTokenBalancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTokenBalances page.
func (p *ListTokenBalancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTokenBalancesOutput, error) {
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

	result, err := p.client.ListTokenBalances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTokenBalances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListTokenBalances",
	}
}
