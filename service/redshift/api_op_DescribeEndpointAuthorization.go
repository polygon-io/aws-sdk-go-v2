// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes an endpoint authorization.
func (c *Client) DescribeEndpointAuthorization(ctx context.Context, params *DescribeEndpointAuthorizationInput, optFns ...func(*Options)) (*DescribeEndpointAuthorizationOutput, error) {
	if params == nil {
		params = &DescribeEndpointAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEndpointAuthorization", params, optFns, c.addOperationDescribeEndpointAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEndpointAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEndpointAuthorizationInput struct {

	// The Amazon Web Services account ID of either the cluster owner (grantor) or
	// grantee. If Grantee parameter is true, then the Account value is of the grantor.
	Account *string

	// The cluster identifier of the cluster to access.
	ClusterIdentifier *string

	// Indicates whether to check authorization from a grantor or grantee point of
	// view. If true, Amazon Redshift returns endpoint authorizations that you've been
	// granted. If false (default), checks authorization from a grantor point of view.
	Grantee *bool

	// An optional pagination token provided by a previous
	// DescribeEndpointAuthorization request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// the MaxRecords parameter.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a Marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeEndpointAuthorizationOutput struct {

	// The authorizations to an endpoint.
	EndpointAuthorizationList []types.EndpointAuthorization

	// An optional pagination token provided by a previous
	// DescribeEndpointAuthorization request. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// the MaxRecords parameter.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEndpointAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeEndpointAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeEndpointAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeEndpointAuthorization"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEndpointAuthorization(options.Region), middleware.Before); err != nil {
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

// DescribeEndpointAuthorizationAPIClient is a client that implements the
// DescribeEndpointAuthorization operation.
type DescribeEndpointAuthorizationAPIClient interface {
	DescribeEndpointAuthorization(context.Context, *DescribeEndpointAuthorizationInput, ...func(*Options)) (*DescribeEndpointAuthorizationOutput, error)
}

var _ DescribeEndpointAuthorizationAPIClient = (*Client)(nil)

// DescribeEndpointAuthorizationPaginatorOptions is the paginator options for
// DescribeEndpointAuthorization
type DescribeEndpointAuthorizationPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a Marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeEndpointAuthorizationPaginator is a paginator for
// DescribeEndpointAuthorization
type DescribeEndpointAuthorizationPaginator struct {
	options   DescribeEndpointAuthorizationPaginatorOptions
	client    DescribeEndpointAuthorizationAPIClient
	params    *DescribeEndpointAuthorizationInput
	nextToken *string
	firstPage bool
}

// NewDescribeEndpointAuthorizationPaginator returns a new
// DescribeEndpointAuthorizationPaginator
func NewDescribeEndpointAuthorizationPaginator(client DescribeEndpointAuthorizationAPIClient, params *DescribeEndpointAuthorizationInput, optFns ...func(*DescribeEndpointAuthorizationPaginatorOptions)) *DescribeEndpointAuthorizationPaginator {
	if params == nil {
		params = &DescribeEndpointAuthorizationInput{}
	}

	options := DescribeEndpointAuthorizationPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeEndpointAuthorizationPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeEndpointAuthorizationPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeEndpointAuthorization page.
func (p *DescribeEndpointAuthorizationPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeEndpointAuthorizationOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeEndpointAuthorization(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeEndpointAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeEndpointAuthorization",
	}
}
