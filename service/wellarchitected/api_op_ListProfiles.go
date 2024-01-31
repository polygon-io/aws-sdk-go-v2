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

// List profiles.
func (c *Client) ListProfiles(ctx context.Context, params *ListProfilesInput, optFns ...func(*Options)) (*ListProfilesOutput, error) {
	if params == nil {
		params = &ListProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfiles", params, optFns, c.addOperationListProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProfilesInput struct {

	// The maximum number of results to return for this request.
	MaxResults *int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// An optional string added to the beginning of each profile name returned in the
	// results.
	ProfileNamePrefix *string

	// Profile owner type.
	ProfileOwnerType types.ProfileOwnerType

	noSmithyDocumentSerde
}

type ListProfilesOutput struct {

	// The token to use to retrieve the next set of results.
	NextToken *string

	// Profile summaries.
	ProfileSummaries []types.ProfileSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListProfiles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfiles(options.Region), middleware.Before); err != nil {
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

// ListProfilesAPIClient is a client that implements the ListProfiles operation.
type ListProfilesAPIClient interface {
	ListProfiles(context.Context, *ListProfilesInput, ...func(*Options)) (*ListProfilesOutput, error)
}

var _ ListProfilesAPIClient = (*Client)(nil)

// ListProfilesPaginatorOptions is the paginator options for ListProfiles
type ListProfilesPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProfilesPaginator is a paginator for ListProfiles
type ListProfilesPaginator struct {
	options   ListProfilesPaginatorOptions
	client    ListProfilesAPIClient
	params    *ListProfilesInput
	nextToken *string
	firstPage bool
}

// NewListProfilesPaginator returns a new ListProfilesPaginator
func NewListProfilesPaginator(client ListProfilesAPIClient, params *ListProfilesInput, optFns ...func(*ListProfilesPaginatorOptions)) *ListProfilesPaginator {
	if params == nil {
		params = &ListProfilesInput{}
	}

	options := ListProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProfiles page.
func (p *ListProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProfilesOutput, error) {
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

	result, err := p.client.ListProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListProfiles",
	}
}
