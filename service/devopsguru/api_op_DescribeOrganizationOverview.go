// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns an overview of your organization's history based on the specified time
// range. The overview includes the total reactive and proactive insights.
func (c *Client) DescribeOrganizationOverview(ctx context.Context, params *DescribeOrganizationOverviewInput, optFns ...func(*Options)) (*DescribeOrganizationOverviewOutput, error) {
	if params == nil {
		params = &DescribeOrganizationOverviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationOverview", params, optFns, c.addOperationDescribeOrganizationOverviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationOverviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationOverviewInput struct {

	// The start of the time range passed in. The start time granularity is at the day
	// level. The floor of the start time is used. Returned information occurred after
	// this day.
	//
	// This member is required.
	FromTime *time.Time

	// The ID of the Amazon Web Services account.
	AccountIds []string

	// The ID of the organizational unit.
	OrganizationalUnitIds []string

	// The end of the time range passed in. The start time granularity is at the day
	// level. The floor of the start time is used. Returned information occurred before
	// this day. If this is not specified, then the current day is used.
	ToTime *time.Time

	noSmithyDocumentSerde
}

type DescribeOrganizationOverviewOutput struct {

	// An integer that specifies the number of open proactive insights in your Amazon
	// Web Services account.
	//
	// This member is required.
	ProactiveInsights int32

	// An integer that specifies the number of open reactive insights in your Amazon
	// Web Services account.
	//
	// This member is required.
	ReactiveInsights int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationOverviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOrganizationOverview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOrganizationOverview{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeOrganizationOverview"); err != nil {
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
	if err = addOpDescribeOrganizationOverviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationOverview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOrganizationOverview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeOrganizationOverview",
	}
}
