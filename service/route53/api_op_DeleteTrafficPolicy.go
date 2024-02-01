// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a traffic policy. When you delete a traffic policy, Route 53 sets a
// flag on the policy to indicate that it has been deleted. However, Route 53 never
// fully deletes the traffic policy. Note the following:
//   - Deleted traffic policies aren't listed if you run ListTrafficPolicies (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListTrafficPolicies.html)
//     .
//   - There's no way to get a list of deleted policies.
//   - If you retain the ID of the policy, you can get information about the
//     policy, including the traffic policy document, by running GetTrafficPolicy (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetTrafficPolicy.html)
//     .
func (c *Client) DeleteTrafficPolicy(ctx context.Context, params *DeleteTrafficPolicyInput, optFns ...func(*Options)) (*DeleteTrafficPolicyOutput, error) {
	if params == nil {
		params = &DeleteTrafficPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTrafficPolicy", params, optFns, c.addOperationDeleteTrafficPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTrafficPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a specified traffic policy version.
type DeleteTrafficPolicyInput struct {

	// The ID of the traffic policy that you want to delete.
	//
	// This member is required.
	Id *string

	// The version number of the traffic policy that you want to delete.
	//
	// This member is required.
	Version *int32

	noSmithyDocumentSerde
}

// An empty element.
type DeleteTrafficPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTrafficPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteTrafficPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteTrafficPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteTrafficPolicy"); err != nil {
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
	if err = addOpDeleteTrafficPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTrafficPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTrafficPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteTrafficPolicy",
	}
}
