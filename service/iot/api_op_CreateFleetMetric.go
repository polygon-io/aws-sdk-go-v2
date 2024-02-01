// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a fleet metric. Requires permission to access the CreateFleetMetric (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) CreateFleetMetric(ctx context.Context, params *CreateFleetMetricInput, optFns ...func(*Options)) (*CreateFleetMetricOutput, error) {
	if params == nil {
		params = &CreateFleetMetricInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFleetMetric", params, optFns, c.addOperationCreateFleetMetricMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFleetMetricOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFleetMetricInput struct {

	// The field to aggregate.
	//
	// This member is required.
	AggregationField *string

	// The type of the aggregation query.
	//
	// This member is required.
	AggregationType *types.AggregationType

	// The name of the fleet metric to create.
	//
	// This member is required.
	MetricName *string

	// The time in seconds between fleet metric emissions. Range [60(1 min), 86400(1
	// day)] and must be multiple of 60.
	//
	// This member is required.
	Period *int32

	// The search query string.
	//
	// This member is required.
	QueryString *string

	// The fleet metric description.
	Description *string

	// The name of the index to search.
	IndexName *string

	// The query version.
	QueryVersion *string

	// Metadata, which can be used to manage the fleet metric.
	Tags []types.Tag

	// Used to support unit transformation such as milliseconds to seconds. The unit
	// must be supported by CW metric (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html)
	// . Default to null.
	Unit types.FleetMetricUnit

	noSmithyDocumentSerde
}

type CreateFleetMetricOutput struct {

	// The Amazon Resource Name (ARN) of the new fleet metric.
	MetricArn *string

	// The name of the fleet metric to create.
	MetricName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFleetMetricMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFleetMetric{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFleetMetric{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFleetMetric"); err != nil {
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
	if err = addOpCreateFleetMetricValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFleetMetric(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFleetMetric(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFleetMetric",
	}
}
