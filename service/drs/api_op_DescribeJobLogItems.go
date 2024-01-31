// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a detailed Job log with pagination.
func (c *Client) DescribeJobLogItems(ctx context.Context, params *DescribeJobLogItemsInput, optFns ...func(*Options)) (*DescribeJobLogItemsOutput, error) {
	if params == nil {
		params = &DescribeJobLogItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeJobLogItems", params, optFns, c.addOperationDescribeJobLogItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeJobLogItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeJobLogItemsInput struct {

	// The ID of the Job for which Job log items will be retrieved.
	//
	// This member is required.
	JobID *string

	// Maximum number of Job log items to retrieve.
	MaxResults *int32

	// The token of the next Job log items to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeJobLogItemsOutput struct {

	// An array of Job log items.
	Items []types.JobLog

	// The token of the next Job log items to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeJobLogItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeJobLogItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeJobLogItems{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeJobLogItems"); err != nil {
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
	if err = addOpDescribeJobLogItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeJobLogItems(options.Region), middleware.Before); err != nil {
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

// DescribeJobLogItemsAPIClient is a client that implements the
// DescribeJobLogItems operation.
type DescribeJobLogItemsAPIClient interface {
	DescribeJobLogItems(context.Context, *DescribeJobLogItemsInput, ...func(*Options)) (*DescribeJobLogItemsOutput, error)
}

var _ DescribeJobLogItemsAPIClient = (*Client)(nil)

// DescribeJobLogItemsPaginatorOptions is the paginator options for
// DescribeJobLogItems
type DescribeJobLogItemsPaginatorOptions struct {
	// Maximum number of Job log items to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeJobLogItemsPaginator is a paginator for DescribeJobLogItems
type DescribeJobLogItemsPaginator struct {
	options   DescribeJobLogItemsPaginatorOptions
	client    DescribeJobLogItemsAPIClient
	params    *DescribeJobLogItemsInput
	nextToken *string
	firstPage bool
}

// NewDescribeJobLogItemsPaginator returns a new DescribeJobLogItemsPaginator
func NewDescribeJobLogItemsPaginator(client DescribeJobLogItemsAPIClient, params *DescribeJobLogItemsInput, optFns ...func(*DescribeJobLogItemsPaginatorOptions)) *DescribeJobLogItemsPaginator {
	if params == nil {
		params = &DescribeJobLogItemsInput{}
	}

	options := DescribeJobLogItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeJobLogItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeJobLogItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeJobLogItems page.
func (p *DescribeJobLogItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeJobLogItemsOutput, error) {
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

	result, err := p.client.DescribeJobLogItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeJobLogItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeJobLogItems",
	}
}
