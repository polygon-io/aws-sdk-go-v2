// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// View a list of environment account connections. For more information, see
// Environment account connections (https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html)
// in the Proton User guide.
func (c *Client) ListEnvironmentAccountConnections(ctx context.Context, params *ListEnvironmentAccountConnectionsInput, optFns ...func(*Options)) (*ListEnvironmentAccountConnectionsOutput, error) {
	if params == nil {
		params = &ListEnvironmentAccountConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentAccountConnections", params, optFns, c.addOperationListEnvironmentAccountConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentAccountConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentAccountConnectionsInput struct {

	// The type of account making the ListEnvironmentAccountConnections request.
	//
	// This member is required.
	RequestedBy types.EnvironmentAccountConnectionRequesterAccountType

	// The environment name that's associated with each listed environment account
	// connection.
	EnvironmentName *string

	// The maximum number of environment account connections to list.
	MaxResults *int32

	// A token that indicates the location of the next environment account connection
	// in the array of environment account connections, after the list of environment
	// account connections that was previously requested.
	NextToken *string

	// The status details for each listed environment account connection.
	Statuses []types.EnvironmentAccountConnectionStatus

	noSmithyDocumentSerde
}

type ListEnvironmentAccountConnectionsOutput struct {

	// An array of environment account connections with details that's returned by
	// Proton.
	//
	// This member is required.
	EnvironmentAccountConnections []types.EnvironmentAccountConnectionSummary

	// A token that indicates the location of the next environment account connection
	// in the array of environment account connections, after the current requested
	// list of environment account connections.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentAccountConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListEnvironmentAccountConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListEnvironmentAccountConnections{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnvironmentAccountConnections"); err != nil {
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
	if err = addOpListEnvironmentAccountConnectionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironmentAccountConnections(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentAccountConnectionsAPIClient is a client that implements the
// ListEnvironmentAccountConnections operation.
type ListEnvironmentAccountConnectionsAPIClient interface {
	ListEnvironmentAccountConnections(context.Context, *ListEnvironmentAccountConnectionsInput, ...func(*Options)) (*ListEnvironmentAccountConnectionsOutput, error)
}

var _ ListEnvironmentAccountConnectionsAPIClient = (*Client)(nil)

// ListEnvironmentAccountConnectionsPaginatorOptions is the paginator options for
// ListEnvironmentAccountConnections
type ListEnvironmentAccountConnectionsPaginatorOptions struct {
	// The maximum number of environment account connections to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentAccountConnectionsPaginator is a paginator for
// ListEnvironmentAccountConnections
type ListEnvironmentAccountConnectionsPaginator struct {
	options   ListEnvironmentAccountConnectionsPaginatorOptions
	client    ListEnvironmentAccountConnectionsAPIClient
	params    *ListEnvironmentAccountConnectionsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentAccountConnectionsPaginator returns a new
// ListEnvironmentAccountConnectionsPaginator
func NewListEnvironmentAccountConnectionsPaginator(client ListEnvironmentAccountConnectionsAPIClient, params *ListEnvironmentAccountConnectionsInput, optFns ...func(*ListEnvironmentAccountConnectionsPaginatorOptions)) *ListEnvironmentAccountConnectionsPaginator {
	if params == nil {
		params = &ListEnvironmentAccountConnectionsInput{}
	}

	options := ListEnvironmentAccountConnectionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentAccountConnectionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentAccountConnectionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironmentAccountConnections page.
func (p *ListEnvironmentAccountConnectionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentAccountConnectionsOutput, error) {
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

	result, err := p.client.ListEnvironmentAccountConnections(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnvironmentAccountConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnvironmentAccountConnections",
	}
}
