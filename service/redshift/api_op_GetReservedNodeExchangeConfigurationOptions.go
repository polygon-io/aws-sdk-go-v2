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

// Gets the configuration options for the reserved-node exchange. These options
// include information about the source reserved node and target reserved node
// offering. Details include the node type, the price, the node count, and the
// offering type.
func (c *Client) GetReservedNodeExchangeConfigurationOptions(ctx context.Context, params *GetReservedNodeExchangeConfigurationOptionsInput, optFns ...func(*Options)) (*GetReservedNodeExchangeConfigurationOptionsOutput, error) {
	if params == nil {
		params = &GetReservedNodeExchangeConfigurationOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReservedNodeExchangeConfigurationOptions", params, optFns, c.addOperationGetReservedNodeExchangeConfigurationOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReservedNodeExchangeConfigurationOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReservedNodeExchangeConfigurationOptionsInput struct {

	// The action type of the reserved-node configuration. The action type can be an
	// exchange initiated from either a snapshot or a resize.
	//
	// This member is required.
	ActionType types.ReservedNodeExchangeActionType

	// The identifier for the cluster that is the source for a reserved-node exchange.
	ClusterIdentifier *string

	// An optional pagination token provided by a previous
	// GetReservedNodeExchangeConfigurationOptions request. If this parameter is
	// specified, the response includes only records beyond the marker, up to the value
	// specified by the MaxRecords parameter. You can retrieve the next set of
	// response records by providing the returned marker value in the Marker parameter
	// and retrying the request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a Marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// The identifier for the snapshot that is the source for the reserved-node
	// exchange.
	SnapshotIdentifier *string

	noSmithyDocumentSerde
}

type GetReservedNodeExchangeConfigurationOptionsOutput struct {

	// A pagination token provided by a previous
	// GetReservedNodeExchangeConfigurationOptions request.
	Marker *string

	// the configuration options for the reserved-node exchange. These options include
	// information about the source reserved node and target reserved node. Details
	// include the node type, the price, the node count, and the offering type.
	ReservedNodeConfigurationOptionList []types.ReservedNodeConfigurationOption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReservedNodeExchangeConfigurationOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetReservedNodeExchangeConfigurationOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetReservedNodeExchangeConfigurationOptions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetReservedNodeExchangeConfigurationOptions"); err != nil {
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
	if err = addOpGetReservedNodeExchangeConfigurationOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReservedNodeExchangeConfigurationOptions(options.Region), middleware.Before); err != nil {
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

// GetReservedNodeExchangeConfigurationOptionsAPIClient is a client that
// implements the GetReservedNodeExchangeConfigurationOptions operation.
type GetReservedNodeExchangeConfigurationOptionsAPIClient interface {
	GetReservedNodeExchangeConfigurationOptions(context.Context, *GetReservedNodeExchangeConfigurationOptionsInput, ...func(*Options)) (*GetReservedNodeExchangeConfigurationOptionsOutput, error)
}

var _ GetReservedNodeExchangeConfigurationOptionsAPIClient = (*Client)(nil)

// GetReservedNodeExchangeConfigurationOptionsPaginatorOptions is the paginator
// options for GetReservedNodeExchangeConfigurationOptions
type GetReservedNodeExchangeConfigurationOptionsPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a Marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetReservedNodeExchangeConfigurationOptionsPaginator is a paginator for
// GetReservedNodeExchangeConfigurationOptions
type GetReservedNodeExchangeConfigurationOptionsPaginator struct {
	options   GetReservedNodeExchangeConfigurationOptionsPaginatorOptions
	client    GetReservedNodeExchangeConfigurationOptionsAPIClient
	params    *GetReservedNodeExchangeConfigurationOptionsInput
	nextToken *string
	firstPage bool
}

// NewGetReservedNodeExchangeConfigurationOptionsPaginator returns a new
// GetReservedNodeExchangeConfigurationOptionsPaginator
func NewGetReservedNodeExchangeConfigurationOptionsPaginator(client GetReservedNodeExchangeConfigurationOptionsAPIClient, params *GetReservedNodeExchangeConfigurationOptionsInput, optFns ...func(*GetReservedNodeExchangeConfigurationOptionsPaginatorOptions)) *GetReservedNodeExchangeConfigurationOptionsPaginator {
	if params == nil {
		params = &GetReservedNodeExchangeConfigurationOptionsInput{}
	}

	options := GetReservedNodeExchangeConfigurationOptionsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetReservedNodeExchangeConfigurationOptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetReservedNodeExchangeConfigurationOptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetReservedNodeExchangeConfigurationOptions page.
func (p *GetReservedNodeExchangeConfigurationOptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetReservedNodeExchangeConfigurationOptionsOutput, error) {
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

	result, err := p.client.GetReservedNodeExchangeConfigurationOptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetReservedNodeExchangeConfigurationOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetReservedNodeExchangeConfigurationOptions",
	}
}
