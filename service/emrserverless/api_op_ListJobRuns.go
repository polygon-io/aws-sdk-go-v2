// Code generated by smithy-go-codegen DO NOT EDIT.

package emrserverless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emrserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists job runs based on a set of parameters.
func (c *Client) ListJobRuns(ctx context.Context, params *ListJobRunsInput, optFns ...func(*Options)) (*ListJobRunsOutput, error) {
	if params == nil {
		params = &ListJobRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJobRuns", params, optFns, c.addOperationListJobRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJobRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobRunsInput struct {

	// The ID of the application for which to list the job run.
	//
	// This member is required.
	ApplicationId *string

	// The lower bound of the option to filter by creation date and time.
	CreatedAtAfter *time.Time

	// The upper bound of the option to filter by creation date and time.
	CreatedAtBefore *time.Time

	// The maximum number of job runs that can be listed.
	MaxResults *int32

	// The token for the next set of job run results.
	NextToken *string

	// An optional filter for job run states. Note that if this filter contains
	// multiple states, the resulting list will be grouped by the state.
	States []types.JobRunState

	noSmithyDocumentSerde
}

type ListJobRunsOutput struct {

	// The output lists information about the specified job runs.
	//
	// This member is required.
	JobRuns []types.JobRunSummary

	// The output displays the token for the next set of job run results. This is
	// required for pagination and is available as a response of the previous request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListJobRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJobRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobRuns{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListJobRuns"); err != nil {
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
	if err = addOpListJobRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJobRuns(options.Region), middleware.Before); err != nil {
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

// ListJobRunsAPIClient is a client that implements the ListJobRuns operation.
type ListJobRunsAPIClient interface {
	ListJobRuns(context.Context, *ListJobRunsInput, ...func(*Options)) (*ListJobRunsOutput, error)
}

var _ ListJobRunsAPIClient = (*Client)(nil)

// ListJobRunsPaginatorOptions is the paginator options for ListJobRuns
type ListJobRunsPaginatorOptions struct {
	// The maximum number of job runs that can be listed.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJobRunsPaginator is a paginator for ListJobRuns
type ListJobRunsPaginator struct {
	options   ListJobRunsPaginatorOptions
	client    ListJobRunsAPIClient
	params    *ListJobRunsInput
	nextToken *string
	firstPage bool
}

// NewListJobRunsPaginator returns a new ListJobRunsPaginator
func NewListJobRunsPaginator(client ListJobRunsAPIClient, params *ListJobRunsInput, optFns ...func(*ListJobRunsPaginatorOptions)) *ListJobRunsPaginator {
	if params == nil {
		params = &ListJobRunsInput{}
	}

	options := ListJobRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJobRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJobRunsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListJobRuns page.
func (p *ListJobRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJobRunsOutput, error) {
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

	result, err := p.client.ListJobRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJobRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListJobRuns",
	}
}
