// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about endpoints for an Amazon Neptune DB cluster. This
// operation can also return information for Amazon RDS clusters and Amazon DocDB
// clusters.
func (c *Client) DescribeDBClusterEndpoints(ctx context.Context, params *DescribeDBClusterEndpointsInput, optFns ...func(*Options)) (*DescribeDBClusterEndpointsOutput, error) {
	if params == nil {
		params = &DescribeDBClusterEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterEndpoints", params, optFns, c.addOperationDescribeDBClusterEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterEndpointsInput struct {

	// The identifier of the endpoint to describe. This parameter is stored as a
	// lowercase string.
	DBClusterEndpointIdentifier *string

	// The DB cluster identifier of the DB cluster associated with the endpoint. This
	// parameter is stored as a lowercase string.
	DBClusterIdentifier *string

	// A set of name-value pairs that define which endpoints to include in the output.
	// The filters are specified as name-value pairs, in the format
	// Name=endpoint_type,Values=endpoint_type1,endpoint_type2,... . Name can be one
	// of: db-cluster-endpoint-type , db-cluster-endpoint-custom-type ,
	// db-cluster-endpoint-id , db-cluster-endpoint-status . Values for the
	// db-cluster-endpoint-type filter can be one or more of: reader , writer , custom
	// . Values for the db-cluster-endpoint-custom-type filter can be one or more of:
	// reader , any . Values for the db-cluster-endpoint-status filter can be one or
	// more of: available , creating , deleting , inactive , modifying .
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBClusterEndpoints
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeDBClusterEndpointsOutput struct {

	// Contains the details of the endpoints associated with the cluster and matching
	// any filter conditions.
	DBClusterEndpoints []types.DBClusterEndpoint

	// An optional pagination token provided by a previous DescribeDBClusterEndpoints
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBClusterEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeDBClusterEndpoints"); err != nil {
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
	if err = addOpDescribeDBClusterEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterEndpoints(options.Region), middleware.Before); err != nil {
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

// DescribeDBClusterEndpointsAPIClient is a client that implements the
// DescribeDBClusterEndpoints operation.
type DescribeDBClusterEndpointsAPIClient interface {
	DescribeDBClusterEndpoints(context.Context, *DescribeDBClusterEndpointsInput, ...func(*Options)) (*DescribeDBClusterEndpointsOutput, error)
}

var _ DescribeDBClusterEndpointsAPIClient = (*Client)(nil)

// DescribeDBClusterEndpointsPaginatorOptions is the paginator options for
// DescribeDBClusterEndpoints
type DescribeDBClusterEndpointsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBClusterEndpointsPaginator is a paginator for
// DescribeDBClusterEndpoints
type DescribeDBClusterEndpointsPaginator struct {
	options   DescribeDBClusterEndpointsPaginatorOptions
	client    DescribeDBClusterEndpointsAPIClient
	params    *DescribeDBClusterEndpointsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBClusterEndpointsPaginator returns a new
// DescribeDBClusterEndpointsPaginator
func NewDescribeDBClusterEndpointsPaginator(client DescribeDBClusterEndpointsAPIClient, params *DescribeDBClusterEndpointsInput, optFns ...func(*DescribeDBClusterEndpointsPaginatorOptions)) *DescribeDBClusterEndpointsPaginator {
	if params == nil {
		params = &DescribeDBClusterEndpointsInput{}
	}

	options := DescribeDBClusterEndpointsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBClusterEndpointsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBClusterEndpointsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBClusterEndpoints page.
func (p *DescribeDBClusterEndpointsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBClusterEndpointsOutput, error) {
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

	result, err := p.client.DescribeDBClusterEndpoints(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDBClusterEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeDBClusterEndpoints",
	}
}
