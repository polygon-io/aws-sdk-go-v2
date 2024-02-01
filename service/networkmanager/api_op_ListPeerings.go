// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the peerings for a core network.
func (c *Client) ListPeerings(ctx context.Context, params *ListPeeringsInput, optFns ...func(*Options)) (*ListPeeringsOutput, error) {
	if params == nil {
		params = &ListPeeringsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPeerings", params, optFns, c.addOperationListPeeringsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPeeringsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPeeringsInput struct {

	// The ID of a core network.
	CoreNetworkId *string

	// Returns a list edge locations for the
	EdgeLocation *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// Returns a list of a peering requests.
	PeeringType types.PeeringType

	// Returns a list of the peering request states.
	State types.PeeringState

	noSmithyDocumentSerde
}

type ListPeeringsOutput struct {

	// The token for the next page of results.
	NextToken *string

	// Lists the transit gateway peerings for the ListPeerings request.
	Peerings []types.Peering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPeeringsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPeerings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPeerings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListPeerings"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPeerings(options.Region), middleware.Before); err != nil {
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

// ListPeeringsAPIClient is a client that implements the ListPeerings operation.
type ListPeeringsAPIClient interface {
	ListPeerings(context.Context, *ListPeeringsInput, ...func(*Options)) (*ListPeeringsOutput, error)
}

var _ ListPeeringsAPIClient = (*Client)(nil)

// ListPeeringsPaginatorOptions is the paginator options for ListPeerings
type ListPeeringsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPeeringsPaginator is a paginator for ListPeerings
type ListPeeringsPaginator struct {
	options   ListPeeringsPaginatorOptions
	client    ListPeeringsAPIClient
	params    *ListPeeringsInput
	nextToken *string
	firstPage bool
}

// NewListPeeringsPaginator returns a new ListPeeringsPaginator
func NewListPeeringsPaginator(client ListPeeringsAPIClient, params *ListPeeringsInput, optFns ...func(*ListPeeringsPaginatorOptions)) *ListPeeringsPaginator {
	if params == nil {
		params = &ListPeeringsInput{}
	}

	options := ListPeeringsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPeeringsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPeeringsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPeerings page.
func (p *ListPeeringsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPeeringsOutput, error) {
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

	result, err := p.client.ListPeerings(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPeerings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListPeerings",
	}
}
