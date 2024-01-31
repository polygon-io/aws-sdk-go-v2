// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new custom action that can be used in all pipelines associated with
// the Amazon Web Services account. Only used for custom actions.
func (c *Client) CreateCustomActionType(ctx context.Context, params *CreateCustomActionTypeInput, optFns ...func(*Options)) (*CreateCustomActionTypeOutput, error) {
	if params == nil {
		params = &CreateCustomActionTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCustomActionType", params, optFns, c.addOperationCreateCustomActionTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCustomActionTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a CreateCustomActionType operation.
type CreateCustomActionTypeInput struct {

	// The category of the custom action, such as a build action or a test action.
	//
	// This member is required.
	Category types.ActionCategory

	// The details of the input artifact for the action, such as its commit ID.
	//
	// This member is required.
	InputArtifactDetails *types.ArtifactDetails

	// The details of the output artifact of the action, such as its commit ID.
	//
	// This member is required.
	OutputArtifactDetails *types.ArtifactDetails

	// The provider of the service used in the custom action, such as CodeDeploy.
	//
	// This member is required.
	Provider *string

	// The version identifier of the custom action.
	//
	// This member is required.
	Version *string

	// The configuration properties for the custom action. You can refer to a name in
	// the configuration properties of the custom action within the URL templates by
	// following the format of {Config:name}, as long as the configuration property is
	// both required and not secret. For more information, see Create a Custom Action
	// for a Pipeline (https://docs.aws.amazon.com/codepipeline/latest/userguide/how-to-create-custom-action.html)
	// .
	ConfigurationProperties []types.ActionConfigurationProperty

	// URLs that provide users information about this custom action.
	Settings *types.ActionTypeSettings

	// The tags for the custom action.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Represents the output of a CreateCustomActionType operation.
type CreateCustomActionTypeOutput struct {

	// Returns information about the details of an action type.
	//
	// This member is required.
	ActionType *types.ActionType

	// Specifies the tags applied to the custom action.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCustomActionTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCustomActionType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCustomActionType{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateCustomActionType"); err != nil {
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
	if err = addOpCreateCustomActionTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCustomActionType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCustomActionType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateCustomActionType",
	}
}
