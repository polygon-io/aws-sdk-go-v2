// Code generated by smithy-go-codegen DO NOT EDIT.

package signer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/signer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all available signing profiles in your AWS account. Returns only profiles
// with an ACTIVE status unless the includeCanceled request field is set to true .
// If additional jobs remain to be listed, AWS Signer returns a nextToken value.
// Use this value in subsequent calls to ListSigningJobs to fetch the remaining
// values. You can continue calling ListSigningJobs with your maxResults parameter
// and with new values that Signer returns in the nextToken parameter until all of
// your signing jobs have been returned.
func (c *Client) ListSigningProfiles(ctx context.Context, params *ListSigningProfilesInput, optFns ...func(*Options)) (*ListSigningProfilesOutput, error) {
	if params == nil {
		params = &ListSigningProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSigningProfiles", params, optFns, c.addOperationListSigningProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSigningProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSigningProfilesInput struct {

	// Designates whether to include profiles with the status of CANCELED .
	IncludeCanceled bool

	// The maximum number of profiles to be returned.
	MaxResults *int32

	// Value for specifying the next set of paginated results to return. After you
	// receive a response with truncated results, use this parameter in a subsequent
	// request. Set it to the value of nextToken from the response that you just
	// received.
	NextToken *string

	// Filters results to return only signing jobs initiated for a specified signing
	// platform.
	PlatformId *string

	// Filters results to return only signing jobs with statuses in the specified list.
	Statuses []types.SigningProfileStatus

	noSmithyDocumentSerde
}

type ListSigningProfilesOutput struct {

	// Value for specifying the next set of paginated results to return.
	NextToken *string

	// A list of profiles that are available in the AWS account. This includes
	// profiles with the status of CANCELED if the includeCanceled parameter is set to
	// true .
	Profiles []types.SigningProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSigningProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSigningProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSigningProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSigningProfiles"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSigningProfiles(options.Region), middleware.Before); err != nil {
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

// ListSigningProfilesAPIClient is a client that implements the
// ListSigningProfiles operation.
type ListSigningProfilesAPIClient interface {
	ListSigningProfiles(context.Context, *ListSigningProfilesInput, ...func(*Options)) (*ListSigningProfilesOutput, error)
}

var _ ListSigningProfilesAPIClient = (*Client)(nil)

// ListSigningProfilesPaginatorOptions is the paginator options for
// ListSigningProfiles
type ListSigningProfilesPaginatorOptions struct {
	// The maximum number of profiles to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSigningProfilesPaginator is a paginator for ListSigningProfiles
type ListSigningProfilesPaginator struct {
	options   ListSigningProfilesPaginatorOptions
	client    ListSigningProfilesAPIClient
	params    *ListSigningProfilesInput
	nextToken *string
	firstPage bool
}

// NewListSigningProfilesPaginator returns a new ListSigningProfilesPaginator
func NewListSigningProfilesPaginator(client ListSigningProfilesAPIClient, params *ListSigningProfilesInput, optFns ...func(*ListSigningProfilesPaginatorOptions)) *ListSigningProfilesPaginator {
	if params == nil {
		params = &ListSigningProfilesInput{}
	}

	options := ListSigningProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSigningProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSigningProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSigningProfiles page.
func (p *ListSigningProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSigningProfilesOutput, error) {
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

	result, err := p.client.ListSigningProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSigningProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSigningProfiles",
	}
}
