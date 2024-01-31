// Code generated by smithy-go-codegen DO NOT EDIT.

package verifiedpermissions

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/verifiedpermissions/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Makes a series of decisions about multiple authorization requests for one
// principal or resource. Each request contains the equivalent content of an
// IsAuthorized request: principal, action, resource, and context. Either the
// principal or the resource parameter must be identical across all requests. For
// example, Verified Permissions won't evaluate a pair of requests where bob views
// photo1 and alice views photo2 . Authorization of bob to view photo1 and photo2 ,
// or bob and alice to view photo1 , are valid batches. The request is evaluated
// against all policies in the specified policy store that match the entities that
// you declare. The result of the decisions is a series of Allow or Deny
// responses, along with the IDs of the policies that produced each decision. The
// entities of a BatchIsAuthorized API request can contain up to 100 principals
// and up to 100 resources. The requests of a BatchIsAuthorized API request can
// contain up to 30 requests. The BatchIsAuthorized operation doesn't have its own
// IAM permission. To authorize this operation for Amazon Web Services principals,
// include the permission verifiedpermissions:IsAuthorized in their IAM policies.
func (c *Client) BatchIsAuthorized(ctx context.Context, params *BatchIsAuthorizedInput, optFns ...func(*Options)) (*BatchIsAuthorizedOutput, error) {
	if params == nil {
		params = &BatchIsAuthorizedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchIsAuthorized", params, optFns, c.addOperationBatchIsAuthorizedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchIsAuthorizedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchIsAuthorizedInput struct {

	// Specifies the ID of the policy store. Policies in this policy store will be
	// used to make the authorization decisions for the input.
	//
	// This member is required.
	PolicyStoreId *string

	// An array of up to 30 requests that you want Verified Permissions to evaluate.
	//
	// This member is required.
	Requests []types.BatchIsAuthorizedInputItem

	// Specifies the list of resources and principals and their associated attributes
	// that Verified Permissions can examine when evaluating the policies. You can
	// include only principal and resource entities in this parameter; you can't
	// include actions. You must specify actions in the schema.
	Entities types.EntitiesDefinition

	noSmithyDocumentSerde
}

type BatchIsAuthorizedOutput struct {

	// A series of Allow or Deny decisions for each request, and the policies that
	// produced them.
	//
	// This member is required.
	Results []types.BatchIsAuthorizedOutputItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchIsAuthorizedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpBatchIsAuthorized{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpBatchIsAuthorized{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchIsAuthorized"); err != nil {
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
	if err = addOpBatchIsAuthorizedValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchIsAuthorized(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchIsAuthorized(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchIsAuthorized",
	}
}
