// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a session for running calculations within a workgroup. The session is
// ready when it reaches an IDLE state.
func (c *Client) StartSession(ctx context.Context, params *StartSessionInput, optFns ...func(*Options)) (*StartSessionOutput, error) {
	if params == nil {
		params = &StartSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSession", params, optFns, c.addOperationStartSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSessionInput struct {

	// Contains engine data processing unit (DPU) configuration settings and parameter
	// mappings.
	//
	// This member is required.
	EngineConfiguration *types.EngineConfiguration

	// The workgroup to which the session belongs.
	//
	// This member is required.
	WorkGroup *string

	// A unique case-sensitive string used to ensure the request to create the session
	// is idempotent (executes only once). If another StartSessionRequest is received,
	// the same response is returned and another session is not created. If a parameter
	// has changed, an error is returned. This token is listed as not required because
	// Amazon Web Services SDKs (for example the Amazon Web Services SDK for Java)
	// auto-generate the token for users. If you are not using the Amazon Web Services
	// SDK or the Amazon Web Services CLI, you must provide this token or the action
	// will fail.
	ClientRequestToken *string

	// The session description.
	Description *string

	// The notebook version. This value is supplied automatically for notebook
	// sessions in the Athena console and is not required for programmatic session
	// access. The only valid notebook version is Athena notebook version 1 . If you
	// specify a value for NotebookVersion , you must also specify a value for
	// NotebookId . See EngineConfiguration$AdditionalConfigs .
	NotebookVersion *string

	// The idle timeout in minutes for the session.
	SessionIdleTimeoutInMinutes *int32

	noSmithyDocumentSerde
}

type StartSessionOutput struct {

	// The session ID.
	SessionId *string

	// The state of the session. A description of each state follows. CREATING - The
	// session is being started, including acquiring resources. CREATED - The session
	// has been started. IDLE - The session is able to accept a calculation. BUSY -
	// The session is processing another task and is unable to accept a calculation.
	// TERMINATING - The session is in the process of shutting down. TERMINATED - The
	// session and its resources are no longer running. DEGRADED - The session has no
	// healthy coordinators. FAILED - Due to a failure, the session and its resources
	// are no longer running.
	State types.SessionState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartSession"); err != nil {
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
	if err = addOpStartSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartSession",
	}
}
