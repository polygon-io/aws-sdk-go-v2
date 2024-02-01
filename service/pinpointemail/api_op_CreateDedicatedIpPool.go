// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new pool of dedicated IP addresses. A pool can include one or more
// dedicated IP addresses that are associated with your Amazon Pinpoint account.
// You can associate a pool with a configuration set. When you send an email that
// uses that configuration set, Amazon Pinpoint sends it using only the IP
// addresses in the associated pool.
func (c *Client) CreateDedicatedIpPool(ctx context.Context, params *CreateDedicatedIpPoolInput, optFns ...func(*Options)) (*CreateDedicatedIpPoolOutput, error) {
	if params == nil {
		params = &CreateDedicatedIpPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDedicatedIpPool", params, optFns, c.addOperationCreateDedicatedIpPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDedicatedIpPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a new dedicated IP pool.
type CreateDedicatedIpPoolInput struct {

	// The name of the dedicated IP pool.
	//
	// This member is required.
	PoolName *string

	// An object that defines the tags (keys and values) that you want to associate
	// with the pool.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type CreateDedicatedIpPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDedicatedIpPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDedicatedIpPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDedicatedIpPool{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDedicatedIpPool"); err != nil {
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
	if err = addOpCreateDedicatedIpPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDedicatedIpPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDedicatedIpPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDedicatedIpPool",
	}
}
