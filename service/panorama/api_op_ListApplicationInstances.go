// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of application instances.
func (c *Client) ListApplicationInstances(ctx context.Context, params *ListApplicationInstancesInput, optFns ...func(*Options)) (*ListApplicationInstancesOutput, error) {
	if params == nil {
		params = &ListApplicationInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationInstances", params, optFns, c.addOperationListApplicationInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationInstancesInput struct {

	// The application instances' device ID.
	DeviceId *string

	// The maximum number of application instances to return in one page of results.
	MaxResults int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	// Only include instances with a specific status.
	StatusFilter types.StatusFilter

	noSmithyDocumentSerde
}

type ListApplicationInstancesOutput struct {

	// A list of application instances.
	ApplicationInstances []types.ApplicationInstance

	// A pagination token that's included if more results are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApplicationInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApplicationInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListApplicationInstances"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationInstances(options.Region), middleware.Before); err != nil {
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

// ListApplicationInstancesAPIClient is a client that implements the
// ListApplicationInstances operation.
type ListApplicationInstancesAPIClient interface {
	ListApplicationInstances(context.Context, *ListApplicationInstancesInput, ...func(*Options)) (*ListApplicationInstancesOutput, error)
}

var _ ListApplicationInstancesAPIClient = (*Client)(nil)

// ListApplicationInstancesPaginatorOptions is the paginator options for
// ListApplicationInstances
type ListApplicationInstancesPaginatorOptions struct {
	// The maximum number of application instances to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationInstancesPaginator is a paginator for ListApplicationInstances
type ListApplicationInstancesPaginator struct {
	options   ListApplicationInstancesPaginatorOptions
	client    ListApplicationInstancesAPIClient
	params    *ListApplicationInstancesInput
	nextToken *string
	firstPage bool
}

// NewListApplicationInstancesPaginator returns a new
// ListApplicationInstancesPaginator
func NewListApplicationInstancesPaginator(client ListApplicationInstancesAPIClient, params *ListApplicationInstancesInput, optFns ...func(*ListApplicationInstancesPaginatorOptions)) *ListApplicationInstancesPaginator {
	if params == nil {
		params = &ListApplicationInstancesInput{}
	}

	options := ListApplicationInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationInstances page.
func (p *ListApplicationInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListApplicationInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListApplicationInstances",
	}
}
