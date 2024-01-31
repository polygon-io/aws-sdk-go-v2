// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about one or more branches in a repository.
func (c *Client) ListBranches(ctx context.Context, params *ListBranchesInput, optFns ...func(*Options)) (*ListBranchesOutput, error) {
	if params == nil {
		params = &ListBranchesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBranches", params, optFns, c.addOperationListBranchesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBranchesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a list branches operation.
type ListBranchesInput struct {

	// The name of the repository that contains the branches.
	//
	// This member is required.
	RepositoryName *string

	// An enumeration token that allows the operation to batch the results.
	NextToken *string

	noSmithyDocumentSerde
}

// Represents the output of a list branches operation.
type ListBranchesOutput struct {

	// The list of branch names.
	Branches []string

	// An enumeration token that returns the batch of the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBranchesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBranches{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBranches{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListBranches"); err != nil {
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
	if err = addOpListBranchesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBranches(options.Region), middleware.Before); err != nil {
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

// ListBranchesAPIClient is a client that implements the ListBranches operation.
type ListBranchesAPIClient interface {
	ListBranches(context.Context, *ListBranchesInput, ...func(*Options)) (*ListBranchesOutput, error)
}

var _ ListBranchesAPIClient = (*Client)(nil)

// ListBranchesPaginatorOptions is the paginator options for ListBranches
type ListBranchesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBranchesPaginator is a paginator for ListBranches
type ListBranchesPaginator struct {
	options   ListBranchesPaginatorOptions
	client    ListBranchesAPIClient
	params    *ListBranchesInput
	nextToken *string
	firstPage bool
}

// NewListBranchesPaginator returns a new ListBranchesPaginator
func NewListBranchesPaginator(client ListBranchesAPIClient, params *ListBranchesInput, optFns ...func(*ListBranchesPaginatorOptions)) *ListBranchesPaginator {
	if params == nil {
		params = &ListBranchesInput{}
	}

	options := ListBranchesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBranchesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBranchesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBranches page.
func (p *ListBranchesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBranchesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListBranches(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBranches(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListBranches",
	}
}
