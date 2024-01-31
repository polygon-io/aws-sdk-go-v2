// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies the health check settings to use when evaluating the health state of
// your EC2 instances. For more information, see Configure Health Checks for Your
// Load Balancer (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-healthchecks.html)
// in the Classic Load Balancers Guide.
func (c *Client) ConfigureHealthCheck(ctx context.Context, params *ConfigureHealthCheckInput, optFns ...func(*Options)) (*ConfigureHealthCheckOutput, error) {
	if params == nil {
		params = &ConfigureHealthCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfigureHealthCheck", params, optFns, c.addOperationConfigureHealthCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfigureHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ConfigureHealthCheck.
type ConfigureHealthCheckInput struct {

	// The configuration information.
	//
	// This member is required.
	HealthCheck *types.HealthCheck

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	noSmithyDocumentSerde
}

// Contains the output of ConfigureHealthCheck.
type ConfigureHealthCheckOutput struct {

	// The updated health check.
	HealthCheck *types.HealthCheck

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationConfigureHealthCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpConfigureHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpConfigureHealthCheck{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ConfigureHealthCheck"); err != nil {
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
	if err = addOpConfigureHealthCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opConfigureHealthCheck(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opConfigureHealthCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ConfigureHealthCheck",
	}
}
