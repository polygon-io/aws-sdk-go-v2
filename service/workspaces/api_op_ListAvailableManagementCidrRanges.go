// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of IP address ranges, specified as IPv4 CIDR blocks, that you
// can use for the network management interface when you enable Bring Your Own
// License (BYOL). This operation can be run only by Amazon Web Services accounts
// that are enabled for BYOL. If your account isn't enabled for BYOL, you'll
// receive an AccessDeniedException error. The management network interface is
// connected to a secure Amazon WorkSpaces management network. It is used for
// interactive streaming of the WorkSpace desktop to Amazon WorkSpaces clients, and
// to allow Amazon WorkSpaces to manage the WorkSpace.
func (c *Client) ListAvailableManagementCidrRanges(ctx context.Context, params *ListAvailableManagementCidrRangesInput, optFns ...func(*Options)) (*ListAvailableManagementCidrRangesOutput, error) {
	if params == nil {
		params = &ListAvailableManagementCidrRangesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAvailableManagementCidrRanges", params, optFns, c.addOperationListAvailableManagementCidrRangesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAvailableManagementCidrRangesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableManagementCidrRangesInput struct {

	// The IP address range to search. Specify an IP address range that is compatible
	// with your network and in CIDR notation (that is, specify the range as an IPv4
	// CIDR block).
	//
	// This member is required.
	ManagementCidrRangeConstraint *string

	// The maximum number of items to return.
	MaxResults *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAvailableManagementCidrRangesOutput struct {

	// The list of available IP address ranges, specified as IPv4 CIDR blocks.
	ManagementCidrRanges []string

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAvailableManagementCidrRangesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAvailableManagementCidrRanges{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAvailableManagementCidrRanges{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAvailableManagementCidrRanges"); err != nil {
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
	if err = addOpListAvailableManagementCidrRangesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableManagementCidrRanges(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListAvailableManagementCidrRanges(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAvailableManagementCidrRanges",
	}
}
