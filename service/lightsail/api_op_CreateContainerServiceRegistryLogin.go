// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a temporary set of log in credentials that you can use to log in to the
// Docker process on your local machine. After you're logged in, you can use the
// native Docker commands to push your local container images to the container
// image registry of your Amazon Lightsail account so that you can use them with
// your Lightsail container service. The log in credentials expire 12 hours after
// they are created, at which point you will need to create a new set of log in
// credentials. You can only push container images to the container service
// registry of your Lightsail account. You cannot pull container images or perform
// any other container image management actions on the container service registry.
// After you push your container images to the container image registry of your
// Lightsail account, use the RegisterContainerImage action to register the pushed
// images to a specific Lightsail container service. This action is not required if
// you install and use the Lightsail Control (lightsailctl) plugin to push
// container images to your Lightsail container service. For more information, see
// Pushing and managing container images on your Amazon Lightsail container
// services (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-pushing-container-images)
// in the Amazon Lightsail Developer Guide.
func (c *Client) CreateContainerServiceRegistryLogin(ctx context.Context, params *CreateContainerServiceRegistryLoginInput, optFns ...func(*Options)) (*CreateContainerServiceRegistryLoginOutput, error) {
	if params == nil {
		params = &CreateContainerServiceRegistryLoginInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContainerServiceRegistryLogin", params, optFns, c.addOperationCreateContainerServiceRegistryLoginMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContainerServiceRegistryLoginOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContainerServiceRegistryLoginInput struct {
	noSmithyDocumentSerde
}

type CreateContainerServiceRegistryLoginOutput struct {

	// An object that describes the log in information for the container service
	// registry of your Lightsail account.
	RegistryLogin *types.ContainerServiceRegistryLogin

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContainerServiceRegistryLoginMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContainerServiceRegistryLogin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContainerServiceRegistryLogin{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateContainerServiceRegistryLogin"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContainerServiceRegistryLogin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContainerServiceRegistryLogin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateContainerServiceRegistryLogin",
	}
}
