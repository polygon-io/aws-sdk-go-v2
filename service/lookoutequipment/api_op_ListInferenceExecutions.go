// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all inference executions that have been performed by the specified
// inference scheduler.
func (c *Client) ListInferenceExecutions(ctx context.Context, params *ListInferenceExecutionsInput, optFns ...func(*Options)) (*ListInferenceExecutionsOutput, error) {
	if params == nil {
		params = &ListInferenceExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInferenceExecutions", params, optFns, c.addOperationListInferenceExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInferenceExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInferenceExecutionsInput struct {

	// The name of the inference scheduler for the inference execution listed.
	//
	// This member is required.
	InferenceSchedulerName *string

	// The time reference in the inferenced dataset before which Amazon Lookout for
	// Equipment stopped the inference execution.
	DataEndTimeBefore *time.Time

	// The time reference in the inferenced dataset after which Amazon Lookout for
	// Equipment started the inference execution.
	DataStartTimeAfter *time.Time

	// Specifies the maximum number of inference executions to list.
	MaxResults *int32

	// An opaque pagination token indicating where to continue the listing of
	// inference executions.
	NextToken *string

	// The status of the inference execution.
	Status types.InferenceExecutionStatus

	noSmithyDocumentSerde
}

type ListInferenceExecutionsOutput struct {

	// Provides an array of information about the individual inference executions
	// returned from the ListInferenceExecutions operation, including model used,
	// inference scheduler, data configuration, and so on.
	InferenceExecutionSummaries []types.InferenceExecutionSummary

	// An opaque pagination token indicating where to continue the listing of
	// inference executions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInferenceExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListInferenceExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListInferenceExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListInferenceExecutions"); err != nil {
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
	if err = addOpListInferenceExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInferenceExecutions(options.Region), middleware.Before); err != nil {
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

// ListInferenceExecutionsAPIClient is a client that implements the
// ListInferenceExecutions operation.
type ListInferenceExecutionsAPIClient interface {
	ListInferenceExecutions(context.Context, *ListInferenceExecutionsInput, ...func(*Options)) (*ListInferenceExecutionsOutput, error)
}

var _ ListInferenceExecutionsAPIClient = (*Client)(nil)

// ListInferenceExecutionsPaginatorOptions is the paginator options for
// ListInferenceExecutions
type ListInferenceExecutionsPaginatorOptions struct {
	// Specifies the maximum number of inference executions to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInferenceExecutionsPaginator is a paginator for ListInferenceExecutions
type ListInferenceExecutionsPaginator struct {
	options   ListInferenceExecutionsPaginatorOptions
	client    ListInferenceExecutionsAPIClient
	params    *ListInferenceExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListInferenceExecutionsPaginator returns a new
// ListInferenceExecutionsPaginator
func NewListInferenceExecutionsPaginator(client ListInferenceExecutionsAPIClient, params *ListInferenceExecutionsInput, optFns ...func(*ListInferenceExecutionsPaginatorOptions)) *ListInferenceExecutionsPaginator {
	if params == nil {
		params = &ListInferenceExecutionsInput{}
	}

	options := ListInferenceExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInferenceExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInferenceExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInferenceExecutions page.
func (p *ListInferenceExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInferenceExecutionsOutput, error) {
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

	result, err := p.client.ListInferenceExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInferenceExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListInferenceExecutions",
	}
}
