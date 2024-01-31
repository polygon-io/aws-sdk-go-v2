// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the list of open and accepted behavior graph invitations for the
// member account. This operation can only be called by an invited member account.
// Open invitations are invitations that the member account has not responded to.
// The results do not include behavior graphs for which the member account declined
// the invitation. The results also do not include behavior graphs that the member
// account resigned from or was removed from.
func (c *Client) ListInvitations(ctx context.Context, params *ListInvitationsInput, optFns ...func(*Options)) (*ListInvitationsOutput, error) {
	if params == nil {
		params = &ListInvitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInvitations", params, optFns, c.addOperationListInvitationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInvitationsInput struct {

	// The maximum number of behavior graph invitations to return in the response. The
	// total must be less than the overall limit on the number of results to return,
	// which is currently 200.
	MaxResults *int32

	// For requests to retrieve the next page of results, the pagination token that
	// was returned with the previous page of results. The initial request does not
	// include a pagination token.
	NextToken *string

	noSmithyDocumentSerde
}

type ListInvitationsOutput struct {

	// The list of behavior graphs for which the member account has open or accepted
	// invitations.
	Invitations []types.MemberDetail

	// If there are more behavior graphs remaining in the results, then this is the
	// pagination token to use to request the next page of behavior graphs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInvitationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInvitations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInvitations(options.Region), middleware.Before); err != nil {
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

// ListInvitationsAPIClient is a client that implements the ListInvitations
// operation.
type ListInvitationsAPIClient interface {
	ListInvitations(context.Context, *ListInvitationsInput, ...func(*Options)) (*ListInvitationsOutput, error)
}

var _ ListInvitationsAPIClient = (*Client)(nil)

// ListInvitationsPaginatorOptions is the paginator options for ListInvitations
type ListInvitationsPaginatorOptions struct {
	// The maximum number of behavior graph invitations to return in the response. The
	// total must be less than the overall limit on the number of results to return,
	// which is currently 200.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInvitationsPaginator is a paginator for ListInvitations
type ListInvitationsPaginator struct {
	options   ListInvitationsPaginatorOptions
	client    ListInvitationsAPIClient
	params    *ListInvitationsInput
	nextToken *string
	firstPage bool
}

// NewListInvitationsPaginator returns a new ListInvitationsPaginator
func NewListInvitationsPaginator(client ListInvitationsAPIClient, params *ListInvitationsInput, optFns ...func(*ListInvitationsPaginatorOptions)) *ListInvitationsPaginator {
	if params == nil {
		params = &ListInvitationsInput{}
	}

	options := ListInvitationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInvitationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInvitationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInvitations page.
func (p *ListInvitationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInvitationsOutput, error) {
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

	result, err := p.client.ListInvitations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInvitations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInvitations",
	}
}
