// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List profile shares.
func (c *Client) ListProfileShares(ctx context.Context, params *ListProfileSharesInput, optFns ...func(*Options)) (*ListProfileSharesOutput, error) {
	if params == nil {
		params = &ListProfileSharesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfileShares", params, optFns, c.addOperationListProfileSharesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfileSharesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProfileSharesInput struct {

	// The profile ARN.
	//
	// This member is required.
	ProfileArn *string

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// The Amazon Web Services account ID, organization ID, or organizational unit
	// (OU) ID with which the profile is shared.
	SharedWithPrefix *string

	// The status of the share request.
	Status types.ShareStatus

	noSmithyDocumentSerde
}

type ListProfileSharesOutput struct {

	// The token to use to retrieve the next set of results.
	NextToken *string

	// Profile share summaries.
	ProfileShareSummaries []types.ProfileShareSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfileSharesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfileShares{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfileShares{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProfileShares"); err != nil {
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
	if err = addOpListProfileSharesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfileShares(options.Region), middleware.Before); err != nil {
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

// ListProfileSharesAPIClient is a client that implements the ListProfileShares
// operation.
type ListProfileSharesAPIClient interface {
	ListProfileShares(context.Context, *ListProfileSharesInput, ...func(*Options)) (*ListProfileSharesOutput, error)
}

var _ ListProfileSharesAPIClient = (*Client)(nil)

// ListProfileSharesPaginatorOptions is the paginator options for ListProfileShares
type ListProfileSharesPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProfileSharesPaginator is a paginator for ListProfileShares
type ListProfileSharesPaginator struct {
	options   ListProfileSharesPaginatorOptions
	client    ListProfileSharesAPIClient
	params    *ListProfileSharesInput
	nextToken *string
	firstPage bool
}

// NewListProfileSharesPaginator returns a new ListProfileSharesPaginator
func NewListProfileSharesPaginator(client ListProfileSharesAPIClient, params *ListProfileSharesInput, optFns ...func(*ListProfileSharesPaginatorOptions)) *ListProfileSharesPaginator {
	if params == nil {
		params = &ListProfileSharesInput{}
	}

	options := ListProfileSharesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProfileSharesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProfileSharesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProfileShares page.
func (p *ListProfileSharesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProfileSharesOutput, error) {
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

	result, err := p.client.ListProfileShares(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProfileShares(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProfileShares",
	}
}
