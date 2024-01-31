// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The status of the associations for the managed node(s).
func (c *Client) DescribeInstanceAssociationsStatus(ctx context.Context, params *DescribeInstanceAssociationsStatusInput, optFns ...func(*Options)) (*DescribeInstanceAssociationsStatusOutput, error) {
	if params == nil {
		params = &DescribeInstanceAssociationsStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceAssociationsStatus", params, optFns, c.addOperationDescribeInstanceAssociationsStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceAssociationsStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceAssociationsStatusInput struct {

	// The managed node IDs for which you want association status information.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInstanceAssociationsStatusOutput struct {

	// Status information about the association.
	InstanceAssociationStatusInfos []types.InstanceAssociationStatusInfo

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInstanceAssociationsStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstanceAssociationsStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstanceAssociationsStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeInstanceAssociationsStatus"); err != nil {
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
	if err = addOpDescribeInstanceAssociationsStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceAssociationsStatus(options.Region), middleware.Before); err != nil {
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

// DescribeInstanceAssociationsStatusAPIClient is a client that implements the
// DescribeInstanceAssociationsStatus operation.
type DescribeInstanceAssociationsStatusAPIClient interface {
	DescribeInstanceAssociationsStatus(context.Context, *DescribeInstanceAssociationsStatusInput, ...func(*Options)) (*DescribeInstanceAssociationsStatusOutput, error)
}

var _ DescribeInstanceAssociationsStatusAPIClient = (*Client)(nil)

// DescribeInstanceAssociationsStatusPaginatorOptions is the paginator options for
// DescribeInstanceAssociationsStatus
type DescribeInstanceAssociationsStatusPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstanceAssociationsStatusPaginator is a paginator for
// DescribeInstanceAssociationsStatus
type DescribeInstanceAssociationsStatusPaginator struct {
	options   DescribeInstanceAssociationsStatusPaginatorOptions
	client    DescribeInstanceAssociationsStatusAPIClient
	params    *DescribeInstanceAssociationsStatusInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstanceAssociationsStatusPaginator returns a new
// DescribeInstanceAssociationsStatusPaginator
func NewDescribeInstanceAssociationsStatusPaginator(client DescribeInstanceAssociationsStatusAPIClient, params *DescribeInstanceAssociationsStatusInput, optFns ...func(*DescribeInstanceAssociationsStatusPaginatorOptions)) *DescribeInstanceAssociationsStatusPaginator {
	if params == nil {
		params = &DescribeInstanceAssociationsStatusInput{}
	}

	options := DescribeInstanceAssociationsStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstanceAssociationsStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstanceAssociationsStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInstanceAssociationsStatus page.
func (p *DescribeInstanceAssociationsStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstanceAssociationsStatusOutput, error) {
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

	result, err := p.client.DescribeInstanceAssociationsStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInstanceAssociationsStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeInstanceAssociationsStatus",
	}
}
