// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a model in SageMaker. In the request, you name the model and describe a
// primary container. For the primary container, you specify the Docker image that
// contains inference code, artifacts (from prior training), and a custom
// environment map that the inference code uses when you deploy the model for
// predictions. Use this API to create a model if you want to use SageMaker hosting
// services or run a batch transform job. To host your model, you create an
// endpoint configuration with the CreateEndpointConfig API, and then create an
// endpoint with the CreateEndpoint API. SageMaker then deploys all of the
// containers that you defined for the model in the hosting environment. For an
// example that calls this method when deploying a model to SageMaker hosting
// services, see Create a Model (Amazon Web Services SDK for Python (Boto 3)). (https://docs.aws.amazon.com/sagemaker/latest/dg/realtime-endpoints-deployment.html#realtime-endpoints-deployment-create-model)
// To run a batch transform using your model, you start a job with the
// CreateTransformJob API. SageMaker uses your model and your dataset to get
// inferences which are then saved to a specified S3 location. In the request, you
// also provide an IAM role that SageMaker can assume to access model artifacts and
// docker image for deployment on ML compute hosting instances or for batch
// transform jobs. In addition, you also use the IAM role to manage permissions the
// inference code needs. For example, if the inference code access any other Amazon
// Web Services resources, you grant necessary permissions via this role.
func (c *Client) CreateModel(ctx context.Context, params *CreateModelInput, optFns ...func(*Options)) (*CreateModelOutput, error) {
	if params == nil {
		params = &CreateModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateModel", params, optFns, c.addOperationCreateModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateModelInput struct {

	// The name of the new model.
	//
	// This member is required.
	ModelName *string

	// Specifies the containers in the inference pipeline.
	Containers []types.ContainerDefinition

	// Isolates the model container. No inbound or outbound network calls can be made
	// to or from the model container.
	EnableNetworkIsolation *bool

	// The Amazon Resource Name (ARN) of the IAM role that SageMaker can assume to
	// access model artifacts and docker image for deployment on ML compute instances
	// or for batch transform jobs. Deploying on ML compute instances is part of model
	// hosting. For more information, see SageMaker Roles (https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-roles.html)
	// . To be able to pass this role to SageMaker, the caller of this API must have
	// the iam:PassRole permission.
	ExecutionRoleArn *string

	// Specifies details of how containers in a multi-container endpoint are called.
	InferenceExecutionConfig *types.InferenceExecutionConfig

	// The location of the primary docker image containing inference code, associated
	// artifacts, and custom environment map that the inference code uses when the
	// model is deployed for predictions.
	PrimaryContainer *types.ContainerDefinition

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags []types.Tag

	// A VpcConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_VpcConfig.html)
	// object that specifies the VPC that you want your model to connect to. Control
	// access to and from your model container by configuring the VPC. VpcConfig is
	// used in hosting services and in batch transform. For more information, see
	// Protect Endpoints by Using an Amazon Virtual Private Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/host-vpc.html)
	// and Protect Data in Batch Transform Jobs by Using an Amazon Virtual Private
	// Cloud (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-vpc.html) .
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type CreateModelOutput struct {

	// The ARN of the model created in SageMaker.
	//
	// This member is required.
	ModelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateModel{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateModel"); err != nil {
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
	if err = addOpCreateModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateModel",
	}
}
