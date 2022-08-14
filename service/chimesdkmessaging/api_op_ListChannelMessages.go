// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List all the messages in a channel. Returns a paginated list of ChannelMessages.
// By default, sorted by creation timestamp in descending order. Redacted messages
// appear in the results as empty, since they are only redacted, not deleted.
// Deleted messages do not appear in the results. This action always returns the
// latest version of an edited message. Also, the x-amz-chime-bearer request header
// is mandatory. Use the AppInstanceUserArn of the user that makes the API call as
// the value in the header.
func (c *Client) ListChannelMessages(ctx context.Context, params *ListChannelMessagesInput, optFns ...func(*Options)) (*ListChannelMessagesOutput, error) {
	if params == nil {
		params = &ListChannelMessagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListChannelMessages", params, optFns, c.addOperationListChannelMessagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListChannelMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListChannelMessagesInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The maximum number of messages that you want returned.
	MaxResults *int32

	// The token passed by previous API calls until all requested messages are
	// returned.
	NextToken *string

	// The final or ending time stamp for your requested messages.
	NotAfter *time.Time

	// The initial or starting time stamp for your requested messages.
	NotBefore *time.Time

	// The order in which you want messages sorted. Default is Descending, based on
	// time created.
	SortOrder types.SortOrder

	// The ID of the SubChannel in the request. Only required when listing the messages
	// in a SubChannel that the user belongs to.
	SubChannelId *string

	noSmithyDocumentSerde
}

type ListChannelMessagesOutput struct {

	// The ARN of the channel containing the requested messages.
	ChannelArn *string

	// The information about, and content of, each requested message.
	ChannelMessages []types.ChannelMessageSummary

	// The token passed by previous API calls until all requested messages are
	// returned.
	NextToken *string

	// The ID of the SubChannel in the response.
	SubChannelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListChannelMessagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListChannelMessages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListChannelMessages{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListChannelMessagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListChannelMessages(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListChannelMessagesAPIClient is a client that implements the ListChannelMessages
// operation.
type ListChannelMessagesAPIClient interface {
	ListChannelMessages(context.Context, *ListChannelMessagesInput, ...func(*Options)) (*ListChannelMessagesOutput, error)
}

var _ ListChannelMessagesAPIClient = (*Client)(nil)

// ListChannelMessagesPaginatorOptions is the paginator options for
// ListChannelMessages
type ListChannelMessagesPaginatorOptions struct {
	// The maximum number of messages that you want returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListChannelMessagesPaginator is a paginator for ListChannelMessages
type ListChannelMessagesPaginator struct {
	options   ListChannelMessagesPaginatorOptions
	client    ListChannelMessagesAPIClient
	params    *ListChannelMessagesInput
	nextToken *string
	firstPage bool
}

// NewListChannelMessagesPaginator returns a new ListChannelMessagesPaginator
func NewListChannelMessagesPaginator(client ListChannelMessagesAPIClient, params *ListChannelMessagesInput, optFns ...func(*ListChannelMessagesPaginatorOptions)) *ListChannelMessagesPaginator {
	if params == nil {
		params = &ListChannelMessagesInput{}
	}

	options := ListChannelMessagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListChannelMessagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListChannelMessagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListChannelMessages page.
func (p *ListChannelMessagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListChannelMessagesOutput, error) {
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

	result, err := p.client.ListChannelMessages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListChannelMessages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "ListChannelMessages",
	}
}
