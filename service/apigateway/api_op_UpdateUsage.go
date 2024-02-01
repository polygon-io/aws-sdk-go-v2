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

// Grants a temporary extension to the remaining quota of a usage plan associated
// with a specified API key.
func (c *Client) UpdateUsage(ctx context.Context, params *UpdateUsageInput, optFns ...func(*Options)) (*UpdateUsageOutput, error) {
	if params == nil {
		params = &UpdateUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateUsage", params, optFns, c.addOperationUpdateUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The PATCH request to grant a temporary extension to the remaining quota of a
// usage plan associated with a specified API key.
type UpdateUsageInput struct {

	// The identifier of the API key associated with the usage plan in which a
	// temporary extension is granted to the remaining quota.
	//
	// This member is required.
	KeyId *string

	// The Id of the usage plan associated with the usage data.
	//
	// This member is required.
	UsagePlanId *string

	// For more information about supported patch operations, see Patch Operations (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html)
	// .
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// Represents the usage data of a usage plan.
type UpdateUsageOutput struct {

	// The ending date of the usage data.
	EndDate *string

	// The usage data, as daily logs of used and remaining quotas, over the specified
	// time interval indexed over the API keys in a usage plan. For example, {...,
	// "values" : { "{api_key}" : [ [0, 100], [10, 90], [100, 10]]} , where {api_key}
	// stands for an API key value and the daily log entry is of the format [used
	// quota, remaining quota] .
	Items map[string][][]int64

	// The current pagination position in the paged result set.
	Position *string

	// The starting date of the usage data.
	StartDate *string

	// The plan Id associated with this usage data.
	UsagePlanId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUsage{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateUsage"); err != nil {
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
	if err = addOpUpdateUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUsage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateUsage",
	}
}
