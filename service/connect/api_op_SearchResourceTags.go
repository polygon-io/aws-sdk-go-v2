// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches tags used in an Amazon Connect instance using optional search criteria.
func (c *Client) SearchResourceTags(ctx context.Context, params *SearchResourceTagsInput, optFns ...func(*Options)) (*SearchResourceTagsOutput, error) {
	if params == nil {
		params = &SearchResourceTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchResourceTags", params, optFns, c.addOperationSearchResourceTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchResourceTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchResourceTagsInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The list of resource types to be used to search tags from. If not provided or
	// if any empty list is provided, this API will search from all supported resource
	// types.
	ResourceTypes []string

	// The search criteria to be used to return tags.
	SearchCriteria *types.ResourceTagsSearchCriteria

	noSmithyDocumentSerde
}

type SearchResourceTagsOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// A list of tags used in the Amazon Connect instance.
	Tags []types.TagSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchResourceTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchResourceTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchResourceTags{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchResourceTags"); err != nil {
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
	if err = addOpSearchResourceTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchResourceTags(options.Region), middleware.Before); err != nil {
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

// SearchResourceTagsAPIClient is a client that implements the SearchResourceTags
// operation.
type SearchResourceTagsAPIClient interface {
	SearchResourceTags(context.Context, *SearchResourceTagsInput, ...func(*Options)) (*SearchResourceTagsOutput, error)
}

var _ SearchResourceTagsAPIClient = (*Client)(nil)

// SearchResourceTagsPaginatorOptions is the paginator options for
// SearchResourceTags
type SearchResourceTagsPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchResourceTagsPaginator is a paginator for SearchResourceTags
type SearchResourceTagsPaginator struct {
	options   SearchResourceTagsPaginatorOptions
	client    SearchResourceTagsAPIClient
	params    *SearchResourceTagsInput
	nextToken *string
	firstPage bool
}

// NewSearchResourceTagsPaginator returns a new SearchResourceTagsPaginator
func NewSearchResourceTagsPaginator(client SearchResourceTagsAPIClient, params *SearchResourceTagsInput, optFns ...func(*SearchResourceTagsPaginatorOptions)) *SearchResourceTagsPaginator {
	if params == nil {
		params = &SearchResourceTagsInput{}
	}

	options := SearchResourceTagsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchResourceTagsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchResourceTagsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchResourceTags page.
func (p *SearchResourceTagsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchResourceTagsOutput, error) {
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

	result, err := p.client.SearchResourceTags(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchResourceTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchResourceTags",
	}
}
