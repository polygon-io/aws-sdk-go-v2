// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the health checks used when evaluating the health state of the targets
// in the specified target group.
func (c *Client) ModifyTargetGroup(ctx context.Context, params *ModifyTargetGroupInput, optFns ...func(*Options)) (*ModifyTargetGroupOutput, error) {
	if params == nil {
		params = &ModifyTargetGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTargetGroup", params, optFns, c.addOperationModifyTargetGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTargetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTargetGroupInput struct {

	// The Amazon Resource Name (ARN) of the target group.
	//
	// This member is required.
	TargetGroupArn *string

	// Indicates whether health checks are enabled.
	HealthCheckEnabled *bool

	// The approximate amount of time, in seconds, between health checks of an
	// individual target.
	HealthCheckIntervalSeconds *int32

	// [HTTP/HTTPS health checks] The destination for health checks on the targets.
	// [HTTP1 or HTTP2 protocol version] The ping path. The default is /. [GRPC
	// protocol version] The path of a custom health check method with the format
	// /package.service/method. The default is /Amazon Web Services.ALB/healthcheck.
	HealthCheckPath *string

	// The port the load balancer uses when performing health checks on targets.
	HealthCheckPort *string

	// The protocol the load balancer uses when performing health checks on targets.
	// For Application Load Balancers, the default is HTTP. For Network Load Balancers
	// and Gateway Load Balancers, the default is TCP. The TCP protocol is not
	// supported for health checks if the protocol of the target group is HTTP or
	// HTTPS. It is supported for health checks only if the protocol of the target
	// group is TCP, TLS, UDP, or TCP_UDP. The GENEVE, TLS, UDP, and TCP_UDP protocols
	// are not supported for health checks.
	HealthCheckProtocol types.ProtocolEnum

	// [HTTP/HTTPS health checks] The amount of time, in seconds, during which no
	// response means a failed health check.
	HealthCheckTimeoutSeconds *int32

	// The number of consecutive health checks successes required before considering
	// an unhealthy target healthy.
	HealthyThresholdCount *int32

	// [HTTP/HTTPS health checks] The HTTP or gRPC codes to use when checking for a
	// successful response from a target. For target groups with a protocol of TCP,
	// TCP_UDP, UDP or TLS the range is 200-599. For target groups with a protocol of
	// HTTP or HTTPS, the range is 200-499. For target groups with a protocol of
	// GENEVE, the range is 200-399.
	Matcher *types.Matcher

	// The number of consecutive health check failures required before considering the
	// target unhealthy.
	UnhealthyThresholdCount *int32

	noSmithyDocumentSerde
}

type ModifyTargetGroupOutput struct {

	// Information about the modified target group.
	TargetGroups []types.TargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyTargetGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyTargetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyTargetGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyTargetGroup"); err != nil {
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
	if err = addOpModifyTargetGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTargetGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTargetGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyTargetGroup",
	}
}
