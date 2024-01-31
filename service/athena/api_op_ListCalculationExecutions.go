// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the calculations that have been submitted to a session in descending
// order. Newer calculations are listed first; older calculations are listed later.
func (c *Client) ListCalculationExecutions(ctx context.Context, params *ListCalculationExecutionsInput, optFns ...func(*Options)) (*ListCalculationExecutionsOutput, error) {
	if params == nil {
		params = &ListCalculationExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCalculationExecutions", params, optFns, c.addOperationListCalculationExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCalculationExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCalculationExecutionsInput struct {

	// The session ID.
	//
	// This member is required.
	SessionId *string

	// The maximum number of calculation executions to return.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// A filter for a specific calculation execution state. A description of each
	// state follows. CREATING - The calculation is in the process of being created.
	// CREATED - The calculation has been created and is ready to run. QUEUED - The
	// calculation has been queued for processing. RUNNING - The calculation is
	// running. CANCELING - A request to cancel the calculation has been received and
	// the system is working to stop it. CANCELED - The calculation is no longer
	// running as the result of a cancel request. COMPLETED - The calculation has
	// completed without error. FAILED - The calculation failed and is no longer
	// running.
	StateFilter types.CalculationExecutionState

	noSmithyDocumentSerde
}

type ListCalculationExecutionsOutput struct {

	// A list of CalculationSummary objects.
	Calculations []types.CalculationSummary

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCalculationExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCalculationExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCalculationExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListCalculationExecutions"); err != nil {
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
	if err = addOpListCalculationExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCalculationExecutions(options.Region), middleware.Before); err != nil {
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

// ListCalculationExecutionsAPIClient is a client that implements the
// ListCalculationExecutions operation.
type ListCalculationExecutionsAPIClient interface {
	ListCalculationExecutions(context.Context, *ListCalculationExecutionsInput, ...func(*Options)) (*ListCalculationExecutionsOutput, error)
}

var _ ListCalculationExecutionsAPIClient = (*Client)(nil)

// ListCalculationExecutionsPaginatorOptions is the paginator options for
// ListCalculationExecutions
type ListCalculationExecutionsPaginatorOptions struct {
	// The maximum number of calculation executions to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCalculationExecutionsPaginator is a paginator for ListCalculationExecutions
type ListCalculationExecutionsPaginator struct {
	options   ListCalculationExecutionsPaginatorOptions
	client    ListCalculationExecutionsAPIClient
	params    *ListCalculationExecutionsInput
	nextToken *string
	firstPage bool
}

// NewListCalculationExecutionsPaginator returns a new
// ListCalculationExecutionsPaginator
func NewListCalculationExecutionsPaginator(client ListCalculationExecutionsAPIClient, params *ListCalculationExecutionsInput, optFns ...func(*ListCalculationExecutionsPaginatorOptions)) *ListCalculationExecutionsPaginator {
	if params == nil {
		params = &ListCalculationExecutionsInput{}
	}

	options := ListCalculationExecutionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCalculationExecutionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCalculationExecutionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCalculationExecutions page.
func (p *ListCalculationExecutionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCalculationExecutionsOutput, error) {
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

	result, err := p.client.ListCalculationExecutions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCalculationExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListCalculationExecutions",
	}
}
