// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replaces the current set of policies for the specified load balancer port with
// the specified set of policies. To enable back-end server authentication, use
// SetLoadBalancerPoliciesForBackendServer . For more information about setting
// policies, see Update the SSL Negotiation Configuration (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/ssl-config-update.html)
// , Duration-Based Session Stickiness (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-duration)
// , and Application-Controlled Session Stickiness (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-sticky-sessions.html#enable-sticky-sessions-application)
// in the Classic Load Balancers Guide.
func (c *Client) SetLoadBalancerPoliciesOfListener(ctx context.Context, params *SetLoadBalancerPoliciesOfListenerInput, optFns ...func(*Options)) (*SetLoadBalancerPoliciesOfListenerOutput, error) {
	if params == nil {
		params = &SetLoadBalancerPoliciesOfListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetLoadBalancerPoliciesOfListener", params, optFns, c.addOperationSetLoadBalancerPoliciesOfListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetLoadBalancerPoliciesOfListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetLoadBalancePoliciesOfListener.
type SetLoadBalancerPoliciesOfListenerInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The external port of the load balancer.
	//
	// This member is required.
	LoadBalancerPort int32

	// The names of the policies. This list must include all policies to be enabled.
	// If you omit a policy that is currently enabled, it is disabled. If the list is
	// empty, all current policies are disabled.
	//
	// This member is required.
	PolicyNames []string

	noSmithyDocumentSerde
}

// Contains the output of SetLoadBalancePoliciesOfListener.
type SetLoadBalancerPoliciesOfListenerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetLoadBalancerPoliciesOfListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetLoadBalancerPoliciesOfListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetLoadBalancerPoliciesOfListener{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetLoadBalancerPoliciesOfListener"); err != nil {
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
	if err = addOpSetLoadBalancerPoliciesOfListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetLoadBalancerPoliciesOfListener(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetLoadBalancerPoliciesOfListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetLoadBalancerPoliciesOfListener",
	}
}
