// Code generated by smithy-go-codegen DO NOT EDIT.

package trustedadvisor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/trustedadvisor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List Resources of a Recommendation within an Organization. This API only
// supports prioritized recommendations.
func (c *Client) ListOrganizationRecommendationResources(ctx context.Context, params *ListOrganizationRecommendationResourcesInput, optFns ...func(*Options)) (*ListOrganizationRecommendationResourcesOutput, error) {
	if params == nil {
		params = &ListOrganizationRecommendationResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationRecommendationResources", params, optFns, c.addOperationListOrganizationRecommendationResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationRecommendationResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationRecommendationResourcesInput struct {

	// The AWS Organization organization's Recommendation identifier
	//
	// This member is required.
	OrganizationRecommendationIdentifier *string

	// An account affected by this organization recommendation
	AffectedAccountId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The AWS Region code of the resource
	RegionCode *string

	// The status of the resource
	Status types.ResourceStatus

	noSmithyDocumentSerde
}

type ListOrganizationRecommendationResourcesOutput struct {

	// A list of Recommendation Resources
	//
	// This member is required.
	OrganizationRecommendationResourceSummaries []types.OrganizationRecommendationResourceSummary

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrganizationRecommendationResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrganizationRecommendationResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrganizationRecommendationResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListOrganizationRecommendationResources"); err != nil {
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
	if err = addOpListOrganizationRecommendationResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationRecommendationResources(options.Region), middleware.Before); err != nil {
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

// ListOrganizationRecommendationResourcesAPIClient is a client that implements
// the ListOrganizationRecommendationResources operation.
type ListOrganizationRecommendationResourcesAPIClient interface {
	ListOrganizationRecommendationResources(context.Context, *ListOrganizationRecommendationResourcesInput, ...func(*Options)) (*ListOrganizationRecommendationResourcesOutput, error)
}

var _ ListOrganizationRecommendationResourcesAPIClient = (*Client)(nil)

// ListOrganizationRecommendationResourcesPaginatorOptions is the paginator
// options for ListOrganizationRecommendationResources
type ListOrganizationRecommendationResourcesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrganizationRecommendationResourcesPaginator is a paginator for
// ListOrganizationRecommendationResources
type ListOrganizationRecommendationResourcesPaginator struct {
	options   ListOrganizationRecommendationResourcesPaginatorOptions
	client    ListOrganizationRecommendationResourcesAPIClient
	params    *ListOrganizationRecommendationResourcesInput
	nextToken *string
	firstPage bool
}

// NewListOrganizationRecommendationResourcesPaginator returns a new
// ListOrganizationRecommendationResourcesPaginator
func NewListOrganizationRecommendationResourcesPaginator(client ListOrganizationRecommendationResourcesAPIClient, params *ListOrganizationRecommendationResourcesInput, optFns ...func(*ListOrganizationRecommendationResourcesPaginatorOptions)) *ListOrganizationRecommendationResourcesPaginator {
	if params == nil {
		params = &ListOrganizationRecommendationResourcesInput{}
	}

	options := ListOrganizationRecommendationResourcesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrganizationRecommendationResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrganizationRecommendationResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrganizationRecommendationResources page.
func (p *ListOrganizationRecommendationResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrganizationRecommendationResourcesOutput, error) {
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

	result, err := p.client.ListOrganizationRecommendationResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOrganizationRecommendationResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListOrganizationRecommendationResources",
	}
}
