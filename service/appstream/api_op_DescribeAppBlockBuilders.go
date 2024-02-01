// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list that describes one or more app block builders.
func (c *Client) DescribeAppBlockBuilders(ctx context.Context, params *DescribeAppBlockBuildersInput, optFns ...func(*Options)) (*DescribeAppBlockBuildersOutput, error) {
	if params == nil {
		params = &DescribeAppBlockBuildersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppBlockBuilders", params, optFns, c.addOperationDescribeAppBlockBuildersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppBlockBuildersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppBlockBuildersInput struct {

	// The maximum size of each page of results. The maximum value is 25.
	MaxResults *int32

	// The names of the app block builders.
	Names []string

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeAppBlockBuildersOutput struct {

	// The list that describes one or more app block builders.
	AppBlockBuilders []types.AppBlockBuilder

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppBlockBuildersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAppBlockBuilders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAppBlockBuilders{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeAppBlockBuilders"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppBlockBuilders(options.Region), middleware.Before); err != nil {
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

// DescribeAppBlockBuildersAPIClient is a client that implements the
// DescribeAppBlockBuilders operation.
type DescribeAppBlockBuildersAPIClient interface {
	DescribeAppBlockBuilders(context.Context, *DescribeAppBlockBuildersInput, ...func(*Options)) (*DescribeAppBlockBuildersOutput, error)
}

var _ DescribeAppBlockBuildersAPIClient = (*Client)(nil)

// DescribeAppBlockBuildersPaginatorOptions is the paginator options for
// DescribeAppBlockBuilders
type DescribeAppBlockBuildersPaginatorOptions struct {
	// The maximum size of each page of results. The maximum value is 25.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeAppBlockBuildersPaginator is a paginator for DescribeAppBlockBuilders
type DescribeAppBlockBuildersPaginator struct {
	options   DescribeAppBlockBuildersPaginatorOptions
	client    DescribeAppBlockBuildersAPIClient
	params    *DescribeAppBlockBuildersInput
	nextToken *string
	firstPage bool
}

// NewDescribeAppBlockBuildersPaginator returns a new
// DescribeAppBlockBuildersPaginator
func NewDescribeAppBlockBuildersPaginator(client DescribeAppBlockBuildersAPIClient, params *DescribeAppBlockBuildersInput, optFns ...func(*DescribeAppBlockBuildersPaginatorOptions)) *DescribeAppBlockBuildersPaginator {
	if params == nil {
		params = &DescribeAppBlockBuildersInput{}
	}

	options := DescribeAppBlockBuildersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeAppBlockBuildersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeAppBlockBuildersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeAppBlockBuilders page.
func (p *DescribeAppBlockBuildersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeAppBlockBuildersOutput, error) {
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

	result, err := p.client.DescribeAppBlockBuilders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeAppBlockBuilders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeAppBlockBuilders",
	}
}
