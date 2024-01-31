// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of deployment jobs for a fleet. You can optionally provide
// filters to retrieve specific deployment jobs. This API will no longer be
// supported as of May 2, 2022. Use it to remove resources that were created for
// Deployment Service.
//
// Deprecated: Support for the AWS RoboMaker application deployment feature has
// ended. For additional information, see
// https://docs.aws.amazon.com/robomaker/latest/dg/fleets.html.
func (c *Client) ListDeploymentJobs(ctx context.Context, params *ListDeploymentJobsInput, optFns ...func(*Options)) (*ListDeploymentJobsOutput, error) {
	if params == nil {
		params = &ListDeploymentJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeploymentJobs", params, optFns, c.addOperationListDeploymentJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentJobsInput struct {

	// Optional filters to limit results. The filter names status and fleetName are
	// supported. When filtering, you must use the complete value of the filtered item.
	// You can use up to three filters, but they must be for the same named item. For
	// example, if you are looking for items with the status InProgress or the status
	// Pending .
	Filters []types.Filter

	// When this parameter is used, ListDeploymentJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListDeploymentJobs
	// request with the returned nextToken value. This value can be between 1 and 200.
	// If this parameter is not used, then ListDeploymentJobs returns up to 200
	// results and a nextToken value if applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListDeploymentJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDeploymentJobsOutput struct {

	// A list of deployment jobs that meet the criteria of the request.
	DeploymentJobs []types.DeploymentJob

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListDeploymentJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeploymentJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeploymentJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeploymentJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDeploymentJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentJobs(options.Region), middleware.Before); err != nil {
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

// ListDeploymentJobsAPIClient is a client that implements the ListDeploymentJobs
// operation.
type ListDeploymentJobsAPIClient interface {
	ListDeploymentJobs(context.Context, *ListDeploymentJobsInput, ...func(*Options)) (*ListDeploymentJobsOutput, error)
}

var _ ListDeploymentJobsAPIClient = (*Client)(nil)

// ListDeploymentJobsPaginatorOptions is the paginator options for
// ListDeploymentJobs
type ListDeploymentJobsPaginatorOptions struct {
	// When this parameter is used, ListDeploymentJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListDeploymentJobs
	// request with the returned nextToken value. This value can be between 1 and 200.
	// If this parameter is not used, then ListDeploymentJobs returns up to 200
	// results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeploymentJobsPaginator is a paginator for ListDeploymentJobs
type ListDeploymentJobsPaginator struct {
	options   ListDeploymentJobsPaginatorOptions
	client    ListDeploymentJobsAPIClient
	params    *ListDeploymentJobsInput
	nextToken *string
	firstPage bool
}

// NewListDeploymentJobsPaginator returns a new ListDeploymentJobsPaginator
func NewListDeploymentJobsPaginator(client ListDeploymentJobsAPIClient, params *ListDeploymentJobsInput, optFns ...func(*ListDeploymentJobsPaginatorOptions)) *ListDeploymentJobsPaginator {
	if params == nil {
		params = &ListDeploymentJobsInput{}
	}

	options := ListDeploymentJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeploymentJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeploymentJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeploymentJobs page.
func (p *ListDeploymentJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeploymentJobsOutput, error) {
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

	result, err := p.client.ListDeploymentJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeploymentJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDeploymentJobs",
	}
}
