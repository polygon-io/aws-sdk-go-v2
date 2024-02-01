// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Runtime settings for a model that is deployed with an inference component.
func (c *Client) UpdateInferenceComponentRuntimeConfig(ctx context.Context, params *UpdateInferenceComponentRuntimeConfigInput, optFns ...func(*Options)) (*UpdateInferenceComponentRuntimeConfigOutput, error) {
	if params == nil {
		params = &UpdateInferenceComponentRuntimeConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInferenceComponentRuntimeConfig", params, optFns, c.addOperationUpdateInferenceComponentRuntimeConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInferenceComponentRuntimeConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInferenceComponentRuntimeConfigInput struct {

	// Runtime settings for a model that is deployed with an inference component.
	//
	// This member is required.
	DesiredRuntimeConfig *types.InferenceComponentRuntimeConfig

	// The name of the inference component to update.
	//
	// This member is required.
	InferenceComponentName *string

	noSmithyDocumentSerde
}

type UpdateInferenceComponentRuntimeConfigOutput struct {

	// The Amazon Resource Name (ARN) of the inference component.
	//
	// This member is required.
	InferenceComponentArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateInferenceComponentRuntimeConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInferenceComponentRuntimeConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInferenceComponentRuntimeConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateInferenceComponentRuntimeConfig"); err != nil {
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
	if err = addOpUpdateInferenceComponentRuntimeConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInferenceComponentRuntimeConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInferenceComponentRuntimeConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateInferenceComponentRuntimeConfig",
	}
}
