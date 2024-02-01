// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of specified virtual tapes in the virtual tape shelf
// (VTS). This operation is only supported in the tape gateway type. If a specific
// TapeARN is not specified, Storage Gateway returns a description of all virtual
// tapes found in the VTS associated with your account.
func (c *Client) DescribeTapeArchives(ctx context.Context, params *DescribeTapeArchivesInput, optFns ...func(*Options)) (*DescribeTapeArchivesOutput, error) {
	if params == nil {
		params = &DescribeTapeArchivesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTapeArchives", params, optFns, c.addOperationDescribeTapeArchivesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTapeArchivesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTapeArchivesInput
type DescribeTapeArchivesInput struct {

	// Specifies that the number of virtual tapes described be limited to the
	// specified number.
	Limit *int32

	// An opaque string that indicates the position at which to begin describing
	// virtual tapes.
	Marker *string

	// Specifies one or more unique Amazon Resource Names (ARNs) that represent the
	// virtual tapes you want to describe.
	TapeARNs []string

	noSmithyDocumentSerde
}

// DescribeTapeArchivesOutput
type DescribeTapeArchivesOutput struct {

	// An opaque string that indicates the position at which the virtual tapes that
	// were fetched for description ended. Use this marker in your next request to
	// fetch the next set of virtual tapes in the virtual tape shelf (VTS). If there
	// are no more virtual tapes to describe, this field does not appear in the
	// response.
	Marker *string

	// An array of virtual tape objects in the virtual tape shelf (VTS). The
	// description includes of the Amazon Resource Name (ARN) of the virtual tapes. The
	// information returned includes the Amazon Resource Names (ARNs) of the tapes,
	// size of the tapes, status of the tapes, progress of the description, and tape
	// barcode.
	TapeArchives []types.TapeArchive

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTapeArchivesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTapeArchives{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTapeArchives{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTapeArchives"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTapeArchives(options.Region), middleware.Before); err != nil {
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

// DescribeTapeArchivesAPIClient is a client that implements the
// DescribeTapeArchives operation.
type DescribeTapeArchivesAPIClient interface {
	DescribeTapeArchives(context.Context, *DescribeTapeArchivesInput, ...func(*Options)) (*DescribeTapeArchivesOutput, error)
}

var _ DescribeTapeArchivesAPIClient = (*Client)(nil)

// DescribeTapeArchivesPaginatorOptions is the paginator options for
// DescribeTapeArchives
type DescribeTapeArchivesPaginatorOptions struct {
	// Specifies that the number of virtual tapes described be limited to the
	// specified number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTapeArchivesPaginator is a paginator for DescribeTapeArchives
type DescribeTapeArchivesPaginator struct {
	options   DescribeTapeArchivesPaginatorOptions
	client    DescribeTapeArchivesAPIClient
	params    *DescribeTapeArchivesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTapeArchivesPaginator returns a new DescribeTapeArchivesPaginator
func NewDescribeTapeArchivesPaginator(client DescribeTapeArchivesAPIClient, params *DescribeTapeArchivesInput, optFns ...func(*DescribeTapeArchivesPaginatorOptions)) *DescribeTapeArchivesPaginator {
	if params == nil {
		params = &DescribeTapeArchivesInput{}
	}

	options := DescribeTapeArchivesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTapeArchivesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTapeArchivesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTapeArchives page.
func (p *DescribeTapeArchivesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTapeArchivesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeTapeArchives(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeTapeArchives(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTapeArchives",
	}
}
