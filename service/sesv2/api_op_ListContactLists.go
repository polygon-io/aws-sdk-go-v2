// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the contact lists available.
func (c *Client) ListContactLists(ctx context.Context, params *ListContactListsInput, optFns ...func(*Options)) (*ListContactListsOutput, error) {
	if params == nil {
		params = &ListContactListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListContactLists", params, optFns, c.addOperationListContactListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListContactListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContactListsInput struct {

	// A string token indicating that there might be additional contact lists
	// available to be listed. Use the token provided in the Response to use in the
	// subsequent call to ListContactLists with the same parameters to retrieve the
	// next page of contact lists.
	NextToken *string

	// Maximum number of contact lists to return at once. Use this parameter to
	// paginate results. If additional contact lists exist beyond the specified limit,
	// the NextToken element is sent in the response. Use the NextToken value in
	// subsequent requests to retrieve additional lists.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListContactListsOutput struct {

	// The available contact lists.
	ContactLists []types.ContactList

	// A string token indicating that there might be additional contact lists
	// available to be listed. Copy this token to a subsequent call to ListContactLists
	// with the same parameters to retrieve the next page of contact lists.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListContactListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListContactLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListContactLists{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListContactLists"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListContactLists(options.Region), middleware.Before); err != nil {
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

// ListContactListsAPIClient is a client that implements the ListContactLists
// operation.
type ListContactListsAPIClient interface {
	ListContactLists(context.Context, *ListContactListsInput, ...func(*Options)) (*ListContactListsOutput, error)
}

var _ ListContactListsAPIClient = (*Client)(nil)

// ListContactListsPaginatorOptions is the paginator options for ListContactLists
type ListContactListsPaginatorOptions struct {
	// Maximum number of contact lists to return at once. Use this parameter to
	// paginate results. If additional contact lists exist beyond the specified limit,
	// the NextToken element is sent in the response. Use the NextToken value in
	// subsequent requests to retrieve additional lists.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListContactListsPaginator is a paginator for ListContactLists
type ListContactListsPaginator struct {
	options   ListContactListsPaginatorOptions
	client    ListContactListsAPIClient
	params    *ListContactListsInput
	nextToken *string
	firstPage bool
}

// NewListContactListsPaginator returns a new ListContactListsPaginator
func NewListContactListsPaginator(client ListContactListsAPIClient, params *ListContactListsInput, optFns ...func(*ListContactListsPaginatorOptions)) *ListContactListsPaginator {
	if params == nil {
		params = &ListContactListsInput{}
	}

	options := ListContactListsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListContactListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListContactListsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListContactLists page.
func (p *ListContactListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListContactListsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListContactLists(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListContactLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListContactLists",
	}
}
