// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Writes the configuration for the specified environment blueprint in Amazon
// DataZone.
func (c *Client) PutEnvironmentBlueprintConfiguration(ctx context.Context, params *PutEnvironmentBlueprintConfigurationInput, optFns ...func(*Options)) (*PutEnvironmentBlueprintConfigurationOutput, error) {
	if params == nil {
		params = &PutEnvironmentBlueprintConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEnvironmentBlueprintConfiguration", params, optFns, c.addOperationPutEnvironmentBlueprintConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEnvironmentBlueprintConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEnvironmentBlueprintConfigurationInput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainIdentifier *string

	// Specifies the enabled Amazon Web Services Regions.
	//
	// This member is required.
	EnabledRegions []string

	// The identifier of the environment blueprint.
	//
	// This member is required.
	EnvironmentBlueprintIdentifier *string

	// The ARN of the manage access role.
	ManageAccessRoleArn *string

	// The ARN of the provisioning role.
	ProvisioningRoleArn *string

	// The regional parameters in the environment blueprint.
	RegionalParameters map[string]map[string]string

	noSmithyDocumentSerde
}

type PutEnvironmentBlueprintConfigurationOutput struct {

	// The identifier of the Amazon DataZone domain.
	//
	// This member is required.
	DomainId *string

	// The identifier of the environment blueprint.
	//
	// This member is required.
	EnvironmentBlueprintId *string

	// The timestamp of when the environment blueprint was created.
	CreatedAt *time.Time

	// Specifies the enabled Amazon Web Services Regions.
	EnabledRegions []string

	// The ARN of the manage access role.
	ManageAccessRoleArn *string

	// The ARN of the provisioning role.
	ProvisioningRoleArn *string

	// The regional parameters in the environment blueprint.
	RegionalParameters map[string]map[string]string

	// The timestamp of when the environment blueprint was updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEnvironmentBlueprintConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEnvironmentBlueprintConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEnvironmentBlueprintConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutEnvironmentBlueprintConfiguration"); err != nil {
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
	if err = addOpPutEnvironmentBlueprintConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEnvironmentBlueprintConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEnvironmentBlueprintConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutEnvironmentBlueprintConfiguration",
	}
}
