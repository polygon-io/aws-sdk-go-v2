// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of gateway summaries. Use GetGateway to retrieve details of a
// specific gateway. An optional gateway group ARN can be provided to only retrieve
// gateway summaries of gateways that are associated with that gateway group ARN.
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) ListGateways(ctx context.Context, params *ListGatewaysInput, optFns ...func(*Options)) (*ListGatewaysOutput, error) {
	if params == nil {
		params = &ListGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGateways", params, optFns, c.addOperationListGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGatewaysInput struct {

	// The gateway group ARN for which to list gateways.
	GatewayGroupArn *string

	// The maximum number of gateway summaries to return. The default is 50.
	MaxResults *int32

	// The token used to paginate though multiple pages of gateway summaries.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGatewaysOutput struct {

	// The gateways in the list.
	Gateways []types.GatewaySummary

	// The token used to paginate though multiple pages of gateway summaries.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGateways{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListGateways"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGateways(options.Region), middleware.Before); err != nil {
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

// ListGatewaysAPIClient is a client that implements the ListGateways operation.
type ListGatewaysAPIClient interface {
	ListGateways(context.Context, *ListGatewaysInput, ...func(*Options)) (*ListGatewaysOutput, error)
}

var _ ListGatewaysAPIClient = (*Client)(nil)

// ListGatewaysPaginatorOptions is the paginator options for ListGateways
type ListGatewaysPaginatorOptions struct {
	// The maximum number of gateway summaries to return. The default is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGatewaysPaginator is a paginator for ListGateways
type ListGatewaysPaginator struct {
	options   ListGatewaysPaginatorOptions
	client    ListGatewaysAPIClient
	params    *ListGatewaysInput
	nextToken *string
	firstPage bool
}

// NewListGatewaysPaginator returns a new ListGatewaysPaginator
func NewListGatewaysPaginator(client ListGatewaysAPIClient, params *ListGatewaysInput, optFns ...func(*ListGatewaysPaginatorOptions)) *ListGatewaysPaginator {
	if params == nil {
		params = &ListGatewaysInput{}
	}

	options := ListGatewaysPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGatewaysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGatewaysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGateways page.
func (p *ListGatewaysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGatewaysOutput, error) {
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

	result, err := p.client.ListGateways(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListGateways",
	}
}
