// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an instance fleet to a running cluster. The instance fleet configuration
// is available only in Amazon EMR releases 4.8.0 and later, excluding 5.0.x.
func (c *Client) AddInstanceFleet(ctx context.Context, params *AddInstanceFleetInput, optFns ...func(*Options)) (*AddInstanceFleetOutput, error) {
	if params == nil {
		params = &AddInstanceFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddInstanceFleet", params, optFns, c.addOperationAddInstanceFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddInstanceFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddInstanceFleetInput struct {

	// The unique identifier of the cluster.
	//
	// This member is required.
	ClusterId *string

	// Specifies the configuration of the instance fleet.
	//
	// This member is required.
	InstanceFleet *types.InstanceFleetConfig

	noSmithyDocumentSerde
}

type AddInstanceFleetOutput struct {

	// The Amazon Resource Name of the cluster.
	ClusterArn *string

	// The unique identifier of the cluster.
	ClusterId *string

	// The unique identifier of the instance fleet.
	InstanceFleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddInstanceFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAddInstanceFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAddInstanceFleet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AddInstanceFleet"); err != nil {
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
	if err = addOpAddInstanceFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddInstanceFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddInstanceFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AddInstanceFleet",
	}
}
