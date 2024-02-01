// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or updates an authentication method for an application.
func (c *Client) PutApplicationAuthenticationMethod(ctx context.Context, params *PutApplicationAuthenticationMethodInput, optFns ...func(*Options)) (*PutApplicationAuthenticationMethodOutput, error) {
	if params == nil {
		params = &PutApplicationAuthenticationMethodInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutApplicationAuthenticationMethod", params, optFns, c.addOperationPutApplicationAuthenticationMethodMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutApplicationAuthenticationMethodOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutApplicationAuthenticationMethodInput struct {

	// Specifies the ARN of the application with the authentication method to add or
	// update.
	//
	// This member is required.
	ApplicationArn *string

	// Specifies a structure that describes the authentication method to add or
	// update. The structure type you provide is determined by the
	// AuthenticationMethodType parameter.
	//
	// This member is required.
	AuthenticationMethod types.AuthenticationMethod

	// Specifies the type of the authentication method that you want to add or update.
	//
	// This member is required.
	AuthenticationMethodType types.AuthenticationMethodType

	noSmithyDocumentSerde
}

type PutApplicationAuthenticationMethodOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutApplicationAuthenticationMethodMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutApplicationAuthenticationMethod{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutApplicationAuthenticationMethod{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutApplicationAuthenticationMethod"); err != nil {
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
	if err = addOpPutApplicationAuthenticationMethodValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutApplicationAuthenticationMethod(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutApplicationAuthenticationMethod(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutApplicationAuthenticationMethod",
	}
}
