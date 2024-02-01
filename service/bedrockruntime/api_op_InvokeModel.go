// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Invokes the specified Bedrock model to run inference using the input provided
// in the request body. You use InvokeModel to run inference for text models, image
// models, and embedding models. For more information, see Run inference (https://docs.aws.amazon.com/bedrock/latest/userguide/api-methods-run.html)
// in the Bedrock User Guide. For example requests, see Examples (after the Errors
// section).
func (c *Client) InvokeModel(ctx context.Context, params *InvokeModelInput, optFns ...func(*Options)) (*InvokeModelOutput, error) {
	if params == nil {
		params = &InvokeModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InvokeModel", params, optFns, c.addOperationInvokeModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeModelInput struct {

	// Input data in the format specified in the content-type request header. To see
	// the format and content of this field for different models, refer to Inference
	// parameters (https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html)
	// .
	//
	// This member is required.
	Body []byte

	// Identifier of the model.
	//
	// This member is required.
	ModelId *string

	// The desired MIME type of the inference body in the response. The default value
	// is application/json .
	Accept *string

	// The MIME type of the input data in the request. The default value is
	// application/json .
	ContentType *string

	noSmithyDocumentSerde
}

type InvokeModelOutput struct {

	// Inference response from the model in the format specified in the content-type
	// header field. To see the format and content of this field for different models,
	// refer to Inference parameters (https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html)
	// .
	//
	// This member is required.
	Body []byte

	// The MIME type of the inference result.
	//
	// This member is required.
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInvokeModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvokeModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "InvokeModel"); err != nil {
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
	if err = addOpInvokeModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvokeModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InvokeModel",
	}
}
