// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified tags from the specified resources. When you specify a tag
// key, the action removes both that key and its associated value. The operation
// succeeds even if you attempt to remove tags from a resource that were already
// removed. Note the following:
//   - To remove tags from a resource, you need the necessary permissions for the
//     service that the resource belongs to as well as permissions for removing tags.
//     For more information, see the documentation for the service whose resource you
//     want to untag.
//   - You can only tag resources that are located in the specified Amazon Web
//     Services Region for the calling Amazon Web Services account.
//
// Minimum permissions In addition to the tag:UntagResources permission required
// by this operation, you must also have the remove tags permission defined by the
// service that created the resource. For example, to remove the tags from an
// Amazon EC2 instance using the UntagResources operation, you must have both of
// the following permissions: tag:UntagResource
//
//	ChimeSDKMeetings:DeleteTags
func (c *Client) UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) {
	if params == nil {
		params = &UntagResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagResource", params, optFns, c.addOperationUntagResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagResourceInput struct {

	// The ARN of the resource that you're removing tags from.
	//
	// This member is required.
	ResourceARN *string

	// The tag keys being removed from the resources.
	//
	// This member is required.
	TagKeys []string

	noSmithyDocumentSerde
}

type UntagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUntagResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUntagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUntagResource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UntagResource"); err != nil {
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
	if err = addOpUntagResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUntagResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUntagResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UntagResource",
	}
}
