// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the details of an AppInstanceUserEndpoint . You can update the name and
// AllowMessage values.
func (c *Client) UpdateAppInstanceUserEndpoint(ctx context.Context, params *UpdateAppInstanceUserEndpointInput, optFns ...func(*Options)) (*UpdateAppInstanceUserEndpointOutput, error) {
	if params == nil {
		params = &UpdateAppInstanceUserEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAppInstanceUserEndpoint", params, optFns, c.addOperationUpdateAppInstanceUserEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAppInstanceUserEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAppInstanceUserEndpointInput struct {

	// The ARN of the AppInstanceUser .
	//
	// This member is required.
	AppInstanceUserArn *string

	// The unique identifier of the AppInstanceUserEndpoint .
	//
	// This member is required.
	EndpointId *string

	// Boolean that controls whether the AppInstanceUserEndpoint is opted in to
	// receive messages. ALL indicates the endpoint will receive all messages. NONE
	// indicates the endpoint will receive no messages.
	AllowMessages types.AllowMessages

	// The name of the AppInstanceUserEndpoint .
	Name *string

	noSmithyDocumentSerde
}

type UpdateAppInstanceUserEndpointOutput struct {

	// The ARN of the AppInstanceUser .
	AppInstanceUserArn *string

	// The unique identifier of the AppInstanceUserEndpoint .
	EndpointId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAppInstanceUserEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAppInstanceUserEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAppInstanceUserEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAppInstanceUserEndpoint"); err != nil {
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
	if err = addOpUpdateAppInstanceUserEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAppInstanceUserEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAppInstanceUserEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAppInstanceUserEndpoint",
	}
}
