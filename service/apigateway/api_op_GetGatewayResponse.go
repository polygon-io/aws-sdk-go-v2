// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a GatewayResponse of a specified response type on the given RestApi.
func (c *Client) GetGatewayResponse(ctx context.Context, params *GetGatewayResponseInput, optFns ...func(*Options)) (*GetGatewayResponseOutput, error) {
	if params == nil {
		params = &GetGatewayResponseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGatewayResponse", params, optFns, c.addOperationGetGatewayResponseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGatewayResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets a GatewayResponse of a specified response type on the given RestApi.
type GetGatewayResponseInput struct {

	// The response type of the associated GatewayResponse.
	//
	// This member is required.
	ResponseType types.GatewayResponseType

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	noSmithyDocumentSerde
}

// A gateway response of a given response type and status code, with optional
// response parameters and mapping templates.
type GetGatewayResponseOutput struct {

	// A Boolean flag to indicate whether this GatewayResponse is the default gateway
	// response ( true ) or not ( false ). A default gateway response is one generated
	// by API Gateway without any customization by an API developer.
	DefaultResponse bool

	// Response parameters (paths, query strings and headers) of the GatewayResponse
	// as a string-to-string map of key-value pairs.
	ResponseParameters map[string]string

	// Response templates of the GatewayResponse as a string-to-string map of
	// key-value pairs.
	ResponseTemplates map[string]string

	// The response type of the associated GatewayResponse.
	ResponseType types.GatewayResponseType

	// The HTTP status code for this GatewayResponse.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGatewayResponseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGatewayResponse{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGatewayResponse{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGatewayResponse"); err != nil {
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
	if err = addOpGetGatewayResponseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGatewayResponse(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opGetGatewayResponse(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGatewayResponse",
	}
}
