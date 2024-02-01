// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists recommendation jobs that satisfy various filters.
func (c *Client) ListInferenceRecommendationsJobs(ctx context.Context, params *ListInferenceRecommendationsJobsInput, optFns ...func(*Options)) (*ListInferenceRecommendationsJobsOutput, error) {
	if params == nil {
		params = &ListInferenceRecommendationsJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInferenceRecommendationsJobs", params, optFns, c.addOperationListInferenceRecommendationsJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInferenceRecommendationsJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInferenceRecommendationsJobsInput struct {

	// A filter that returns only jobs created after the specified time (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only jobs created before the specified time (timestamp).
	CreationTimeBefore *time.Time

	// A filter that returns only jobs that were last modified after the specified
	// time (timestamp).
	LastModifiedTimeAfter *time.Time

	// A filter that returns only jobs that were last modified before the specified
	// time (timestamp).
	LastModifiedTimeBefore *time.Time

	// The maximum number of recommendations to return in the response.
	MaxResults *int32

	// A filter that returns only jobs that were created for this model.
	ModelNameEquals *string

	// A filter that returns only jobs that were created for this versioned model
	// package.
	ModelPackageVersionArnEquals *string

	// A string in the job name. This filter returns only recommendations whose name
	// contains the specified string.
	NameContains *string

	// If the response to a previous ListInferenceRecommendationsJobsRequest request
	// was truncated, the response includes a NextToken . To retrieve the next set of
	// recommendations, use the token in the next request.
	NextToken *string

	// The parameter by which to sort the results.
	SortBy types.ListInferenceRecommendationsJobsSortBy

	// The sort order for the results.
	SortOrder types.SortOrder

	// A filter that retrieves only inference recommendations jobs with a specific
	// status.
	StatusEquals types.RecommendationJobStatus

	noSmithyDocumentSerde
}

type ListInferenceRecommendationsJobsOutput struct {

	// The recommendations created from the Amazon SageMaker Inference Recommender job.
	//
	// This member is required.
	InferenceRecommendationsJobs []types.InferenceRecommendationsJob

	// A token for getting the next set of recommendations, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInferenceRecommendationsJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListInferenceRecommendationsJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListInferenceRecommendationsJobs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInferenceRecommendationsJobs"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInferenceRecommendationsJobs(options.Region), middleware.Before); err != nil {
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

// ListInferenceRecommendationsJobsAPIClient is a client that implements the
// ListInferenceRecommendationsJobs operation.
type ListInferenceRecommendationsJobsAPIClient interface {
	ListInferenceRecommendationsJobs(context.Context, *ListInferenceRecommendationsJobsInput, ...func(*Options)) (*ListInferenceRecommendationsJobsOutput, error)
}

var _ ListInferenceRecommendationsJobsAPIClient = (*Client)(nil)

// ListInferenceRecommendationsJobsPaginatorOptions is the paginator options for
// ListInferenceRecommendationsJobs
type ListInferenceRecommendationsJobsPaginatorOptions struct {
	// The maximum number of recommendations to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInferenceRecommendationsJobsPaginator is a paginator for
// ListInferenceRecommendationsJobs
type ListInferenceRecommendationsJobsPaginator struct {
	options   ListInferenceRecommendationsJobsPaginatorOptions
	client    ListInferenceRecommendationsJobsAPIClient
	params    *ListInferenceRecommendationsJobsInput
	nextToken *string
	firstPage bool
}

// NewListInferenceRecommendationsJobsPaginator returns a new
// ListInferenceRecommendationsJobsPaginator
func NewListInferenceRecommendationsJobsPaginator(client ListInferenceRecommendationsJobsAPIClient, params *ListInferenceRecommendationsJobsInput, optFns ...func(*ListInferenceRecommendationsJobsPaginatorOptions)) *ListInferenceRecommendationsJobsPaginator {
	if params == nil {
		params = &ListInferenceRecommendationsJobsInput{}
	}

	options := ListInferenceRecommendationsJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInferenceRecommendationsJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInferenceRecommendationsJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInferenceRecommendationsJobs page.
func (p *ListInferenceRecommendationsJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInferenceRecommendationsJobsOutput, error) {
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

	result, err := p.client.ListInferenceRecommendationsJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInferenceRecommendationsJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInferenceRecommendationsJobs",
	}
}
