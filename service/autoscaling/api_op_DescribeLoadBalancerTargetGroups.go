// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API operation is superseded by DescribeTrafficSources , which can describe
// multiple traffic sources types. We recommend using DetachTrafficSources to
// simplify how you manage traffic sources. However, we continue to support
// DescribeLoadBalancerTargetGroups . You can use both the original
// DescribeLoadBalancerTargetGroups API operation and DescribeTrafficSources on
// the same Auto Scaling group. Gets information about the Elastic Load Balancing
// target groups for the specified Auto Scaling group. To determine the attachment
// status of the target group, use the State element in the response. When you
// attach a target group to an Auto Scaling group, the initial State value is
// Adding . The state transitions to Added after all Auto Scaling instances are
// registered with the target group. If Elastic Load Balancing health checks are
// enabled for the Auto Scaling group, the state transitions to InService after at
// least one Auto Scaling instance passes the health check. When the target group
// is in the InService state, Amazon EC2 Auto Scaling can terminate and replace
// any instances that are reported as unhealthy. If no registered instances pass
// the health checks, the target group doesn't enter the InService state. Target
// groups also have an InService state if you attach them in the
// CreateAutoScalingGroup API call. If your target group state is InService , but
// it is not working properly, check the scaling activities by calling
// DescribeScalingActivities and take any corrective actions necessary. For help
// with failed health checks, see Troubleshooting Amazon EC2 Auto Scaling: Health
// checks (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Use
// Elastic Load Balancing to distribute traffic across the instances in your Auto
// Scaling group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
// in the Amazon EC2 Auto Scaling User Guide. You can use this operation to
// describe target groups that were attached by using
// AttachLoadBalancerTargetGroups , but not for target groups that were attached by
// using AttachTrafficSources .
func (c *Client) DescribeLoadBalancerTargetGroups(ctx context.Context, params *DescribeLoadBalancerTargetGroupsInput, optFns ...func(*Options)) (*DescribeLoadBalancerTargetGroupsOutput, error) {
	if params == nil {
		params = &DescribeLoadBalancerTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBalancerTargetGroups", params, optFns, c.addOperationDescribeLoadBalancerTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoadBalancerTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoadBalancerTargetGroupsInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100 .
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLoadBalancerTargetGroupsOutput struct {

	// Information about the target groups.
	LoadBalancerTargetGroups []types.LoadBalancerTargetGroupState

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoadBalancerTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancerTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancerTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLoadBalancerTargetGroups"); err != nil {
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
	if err = addOpDescribeLoadBalancerTargetGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancerTargetGroups(options.Region), middleware.Before); err != nil {
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

// DescribeLoadBalancerTargetGroupsAPIClient is a client that implements the
// DescribeLoadBalancerTargetGroups operation.
type DescribeLoadBalancerTargetGroupsAPIClient interface {
	DescribeLoadBalancerTargetGroups(context.Context, *DescribeLoadBalancerTargetGroupsInput, ...func(*Options)) (*DescribeLoadBalancerTargetGroupsOutput, error)
}

var _ DescribeLoadBalancerTargetGroupsAPIClient = (*Client)(nil)

// DescribeLoadBalancerTargetGroupsPaginatorOptions is the paginator options for
// DescribeLoadBalancerTargetGroups
type DescribeLoadBalancerTargetGroupsPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLoadBalancerTargetGroupsPaginator is a paginator for
// DescribeLoadBalancerTargetGroups
type DescribeLoadBalancerTargetGroupsPaginator struct {
	options   DescribeLoadBalancerTargetGroupsPaginatorOptions
	client    DescribeLoadBalancerTargetGroupsAPIClient
	params    *DescribeLoadBalancerTargetGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLoadBalancerTargetGroupsPaginator returns a new
// DescribeLoadBalancerTargetGroupsPaginator
func NewDescribeLoadBalancerTargetGroupsPaginator(client DescribeLoadBalancerTargetGroupsAPIClient, params *DescribeLoadBalancerTargetGroupsInput, optFns ...func(*DescribeLoadBalancerTargetGroupsPaginatorOptions)) *DescribeLoadBalancerTargetGroupsPaginator {
	if params == nil {
		params = &DescribeLoadBalancerTargetGroupsInput{}
	}

	options := DescribeLoadBalancerTargetGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLoadBalancerTargetGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLoadBalancerTargetGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLoadBalancerTargetGroups page.
func (p *DescribeLoadBalancerTargetGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLoadBalancerTargetGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeLoadBalancerTargetGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeLoadBalancerTargetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLoadBalancerTargetGroups",
	}
}
