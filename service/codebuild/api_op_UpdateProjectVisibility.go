// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the public visibility for a project. The project's build results, logs,
// and artifacts are available to the general public. For more information, see
// Public build projects (https://docs.aws.amazon.com/codebuild/latest/userguide/public-builds.html)
// in the CodeBuild User Guide. The following should be kept in mind when making
// your projects public:
//
//   - All of a project's build results, logs, and artifacts, including builds
//     that were run when the project was private, are available to the general public.
//
//   - All build logs and artifacts are available to the public. Environment
//     variables, source code, and other sensitive information may have been output to
//     the build logs and artifacts. You must be careful about what information is
//     output to the build logs. Some best practice are:
//
//   - Do not store sensitive values in environment variables. We recommend that
//     you use an Amazon EC2 Systems Manager Parameter Store or Secrets Manager to
//     store sensitive values.
//
//   - Follow Best practices for using webhooks (https://docs.aws.amazon.com/codebuild/latest/userguide/webhooks.html#webhook-best-practices)
//     in the CodeBuild User Guide to limit which entities can trigger a build, and do
//     not store the buildspec in the project itself, to ensure that your webhooks are
//     as secure as possible.
//
//   - A malicious user can use public builds to distribute malicious artifacts.
//     We recommend that you review all pull requests to verify that the pull request
//     is a legitimate change. We also recommend that you validate any artifacts with
//     their checksums to make sure that the correct artifacts are being downloaded.
func (c *Client) UpdateProjectVisibility(ctx context.Context, params *UpdateProjectVisibilityInput, optFns ...func(*Options)) (*UpdateProjectVisibilityOutput, error) {
	if params == nil {
		params = &UpdateProjectVisibilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProjectVisibility", params, optFns, c.addOperationUpdateProjectVisibilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProjectVisibilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProjectVisibilityInput struct {

	// The Amazon Resource Name (ARN) of the build project.
	//
	// This member is required.
	ProjectArn *string

	// Specifies the visibility of the project's builds. Possible values are:
	// PUBLIC_READ The project builds are visible to the public. PRIVATE The project
	// builds are not visible to the public.
	//
	// This member is required.
	ProjectVisibility types.ProjectVisibilityType

	// The ARN of the IAM role that enables CodeBuild to access the CloudWatch Logs
	// and Amazon S3 artifacts for the project's builds.
	ResourceAccessRole *string

	noSmithyDocumentSerde
}

type UpdateProjectVisibilityOutput struct {

	// The Amazon Resource Name (ARN) of the build project.
	ProjectArn *string

	// Specifies the visibility of the project's builds. Possible values are:
	// PUBLIC_READ The project builds are visible to the public. PRIVATE The project
	// builds are not visible to the public.
	ProjectVisibility types.ProjectVisibilityType

	// Contains the project identifier used with the public build APIs.
	PublicProjectAlias *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProjectVisibilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProjectVisibility{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProjectVisibility{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateProjectVisibility"); err != nil {
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
	if err = addOpUpdateProjectVisibilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProjectVisibility(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProjectVisibility(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateProjectVisibility",
	}
}
