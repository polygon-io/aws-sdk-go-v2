// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a service instance. There are a few modes for updating a service
// instance. The deploymentType field defines the mode. You can't update a service
// instance while its deployment status, or the deployment status of a component
// attached to it, is IN_PROGRESS . For more information about components, see
// Proton components (https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html)
// in the Proton User Guide.
func (c *Client) UpdateServiceInstance(ctx context.Context, params *UpdateServiceInstanceInput, optFns ...func(*Options)) (*UpdateServiceInstanceOutput, error) {
	if params == nil {
		params = &UpdateServiceInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateServiceInstance", params, optFns, c.addOperationUpdateServiceInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateServiceInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateServiceInstanceInput struct {

	// The deployment type. It defines the mode for updating a service instance, as
	// follows: NONE In this mode, a deployment doesn't occur. Only the requested
	// metadata parameters are updated. CURRENT_VERSION In this mode, the service
	// instance is deployed and updated with the new spec that you provide. Only
	// requested parameters are updated. Don’t include major or minor version
	// parameters when you use this deployment type. MINOR_VERSION In this mode, the
	// service instance is deployed and updated with the published, recommended
	// (latest) minor version of the current major version in use, by default. You can
	// also specify a different minor version of the current major version in use.
	// MAJOR_VERSION In this mode, the service instance is deployed and updated with
	// the published, recommended (latest) major and minor version of the current
	// template, by default. You can specify a different major version that's higher
	// than the major version in use and a minor version.
	//
	// This member is required.
	DeploymentType types.DeploymentUpdateType

	// The name of the service instance to update.
	//
	// This member is required.
	Name *string

	// The name of the service that the service instance belongs to.
	//
	// This member is required.
	ServiceName *string

	// The client token of the service instance to update.
	ClientToken *string

	// The formatted specification that defines the service instance update.
	//
	// This value conforms to the media type: application/yaml
	Spec *string

	// The major version of the service template to update.
	TemplateMajorVersion *string

	// The minor version of the service template to update.
	TemplateMinorVersion *string

	noSmithyDocumentSerde
}

type UpdateServiceInstanceOutput struct {

	// The service instance summary data that's returned by Proton.
	//
	// This member is required.
	ServiceInstance *types.ServiceInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateServiceInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateServiceInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateServiceInstance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateServiceInstance"); err != nil {
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
	if err = addIdempotencyToken_opUpdateServiceInstanceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateServiceInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateServiceInstance(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateServiceInstance struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateServiceInstance) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateServiceInstance) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateServiceInstanceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateServiceInstanceInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateServiceInstanceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateServiceInstance{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateServiceInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateServiceInstance",
	}
}
