// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of active App Runner automatic scaling configurations in your
// Amazon Web Services account. You can query the revisions for a specific
// configuration name or the revisions for all active configurations in your
// account. You can optionally query only the latest revision of each requested
// name. To retrieve a full description of a particular configuration revision,
// call and provide one of the ARNs returned by ListAutoScalingConfigurations .
func (c *Client) ListAutoScalingConfigurations(ctx context.Context, params *ListAutoScalingConfigurationsInput, optFns ...func(*Options)) (*ListAutoScalingConfigurationsOutput, error) {
	if params == nil {
		params = &ListAutoScalingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAutoScalingConfigurations", params, optFns, c.addOperationListAutoScalingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAutoScalingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAutoScalingConfigurationsInput struct {

	// The name of the App Runner auto scaling configuration that you want to list. If
	// specified, App Runner lists revisions that share this name. If not specified,
	// App Runner returns revisions of all active configurations.
	AutoScalingConfigurationName *string

	// Set to true to list only the latest revision for each requested configuration
	// name. Set to false to list all revisions for each requested configuration name.
	// Default: true
	LatestOnly bool

	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults , the request
	// retrieves all available results in a single response.
	MaxResults *int32

	// A token from a previous result page. It's used for a paginated request. The
	// request retrieves the next result page. All other parameter values must be
	// identical to the ones that are specified in the initial request. If you don't
	// specify NextToken , the request retrieves the first result page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAutoScalingConfigurationsOutput struct {

	// A list of summary information records for auto scaling configurations. In a
	// paginated request, the request returns up to MaxResults records for each call.
	//
	// This member is required.
	AutoScalingConfigurationSummaryList []types.AutoScalingConfigurationSummary

	// The token that you can pass in a subsequent request to get the next result
	// page. It's returned in a paginated request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAutoScalingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListAutoScalingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListAutoScalingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAutoScalingConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAutoScalingConfigurations(options.Region), middleware.Before); err != nil {
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

// ListAutoScalingConfigurationsAPIClient is a client that implements the
// ListAutoScalingConfigurations operation.
type ListAutoScalingConfigurationsAPIClient interface {
	ListAutoScalingConfigurations(context.Context, *ListAutoScalingConfigurationsInput, ...func(*Options)) (*ListAutoScalingConfigurationsOutput, error)
}

var _ ListAutoScalingConfigurationsAPIClient = (*Client)(nil)

// ListAutoScalingConfigurationsPaginatorOptions is the paginator options for
// ListAutoScalingConfigurations
type ListAutoScalingConfigurationsPaginatorOptions struct {
	// The maximum number of results to include in each response (result page). It's
	// used for a paginated request. If you don't specify MaxResults , the request
	// retrieves all available results in a single response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAutoScalingConfigurationsPaginator is a paginator for
// ListAutoScalingConfigurations
type ListAutoScalingConfigurationsPaginator struct {
	options   ListAutoScalingConfigurationsPaginatorOptions
	client    ListAutoScalingConfigurationsAPIClient
	params    *ListAutoScalingConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListAutoScalingConfigurationsPaginator returns a new
// ListAutoScalingConfigurationsPaginator
func NewListAutoScalingConfigurationsPaginator(client ListAutoScalingConfigurationsAPIClient, params *ListAutoScalingConfigurationsInput, optFns ...func(*ListAutoScalingConfigurationsPaginatorOptions)) *ListAutoScalingConfigurationsPaginator {
	if params == nil {
		params = &ListAutoScalingConfigurationsInput{}
	}

	options := ListAutoScalingConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAutoScalingConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAutoScalingConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAutoScalingConfigurations page.
func (p *ListAutoScalingConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAutoScalingConfigurationsOutput, error) {
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

	result, err := p.client.ListAutoScalingConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAutoScalingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAutoScalingConfigurations",
	}
}
