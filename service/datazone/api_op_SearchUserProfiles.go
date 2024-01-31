// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches user profiles in Amazon DataZone.
func (c *Client) SearchUserProfiles(ctx context.Context, params *SearchUserProfilesInput, optFns ...func(*Options)) (*SearchUserProfilesOutput, error) {
	if params == nil {
		params = &SearchUserProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchUserProfiles", params, optFns, c.addOperationSearchUserProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchUserProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchUserProfilesInput struct {

	// The identifier of the Amazon DataZone domain in which you want to search user
	// profiles.
	//
	// This member is required.
	DomainIdentifier *string

	// Specifies the user type for the SearchUserProfiles action.
	//
	// This member is required.
	UserType types.UserSearchType

	// The maximum number of results to return in a single call to SearchUserProfiles .
	// When the number of results to be listed is greater than the value of MaxResults
	// , the response contains a NextToken value that you can use in a subsequent call
	// to SearchUserProfiles to list the next set of results.
	MaxResults *int32

	// When the number of results is greater than the default value for the MaxResults
	// parameter, or if you explicitly specify a value for MaxResults that is less
	// than the number of results, the response includes a pagination token named
	// NextToken . You can specify this NextToken value in a subsequent call to
	// SearchUserProfiles to list the next set of results.
	NextToken *string

	// Specifies the text for which to search.
	SearchText *string

	noSmithyDocumentSerde
}

type SearchUserProfilesOutput struct {

	// The results of the SearchUserProfiles action.
	Items []types.UserProfileSummary

	// When the number of results is greater than the default value for the MaxResults
	// parameter, or if you explicitly specify a value for MaxResults that is less
	// than the number of results, the response includes a pagination token named
	// NextToken . You can specify this NextToken value in a subsequent call to
	// SearchUserProfiles to list the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchUserProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchUserProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchUserProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchUserProfiles"); err != nil {
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
	if err = addOpSearchUserProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchUserProfiles(options.Region), middleware.Before); err != nil {
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

// SearchUserProfilesAPIClient is a client that implements the SearchUserProfiles
// operation.
type SearchUserProfilesAPIClient interface {
	SearchUserProfiles(context.Context, *SearchUserProfilesInput, ...func(*Options)) (*SearchUserProfilesOutput, error)
}

var _ SearchUserProfilesAPIClient = (*Client)(nil)

// SearchUserProfilesPaginatorOptions is the paginator options for
// SearchUserProfiles
type SearchUserProfilesPaginatorOptions struct {
	// The maximum number of results to return in a single call to SearchUserProfiles .
	// When the number of results to be listed is greater than the value of MaxResults
	// , the response contains a NextToken value that you can use in a subsequent call
	// to SearchUserProfiles to list the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchUserProfilesPaginator is a paginator for SearchUserProfiles
type SearchUserProfilesPaginator struct {
	options   SearchUserProfilesPaginatorOptions
	client    SearchUserProfilesAPIClient
	params    *SearchUserProfilesInput
	nextToken *string
	firstPage bool
}

// NewSearchUserProfilesPaginator returns a new SearchUserProfilesPaginator
func NewSearchUserProfilesPaginator(client SearchUserProfilesAPIClient, params *SearchUserProfilesInput, optFns ...func(*SearchUserProfilesPaginatorOptions)) *SearchUserProfilesPaginator {
	if params == nil {
		params = &SearchUserProfilesInput{}
	}

	options := SearchUserProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchUserProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchUserProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchUserProfiles page.
func (p *SearchUserProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchUserProfilesOutput, error) {
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

	result, err := p.client.SearchUserProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchUserProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchUserProfiles",
	}
}
