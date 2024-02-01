// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Set the user's multi-factor authentication (MFA) method preference, including
// which MFA factors are activated and if any are preferred. Only one factor can be
// set as preferred. The preferred MFA factor will be used to authenticate a user
// if multiple factors are activated. If multiple options are activated and no
// preference is set, a challenge to choose an MFA option will be returned during
// sign-in. If an MFA type is activated for a user, the user will be prompted for
// MFA during all sign-in attempts unless device tracking is turned on and the
// device has been trusted. If you want MFA to be applied selectively based on the
// assessed risk level of sign-in attempts, deactivate MFA for users and turn on
// Adaptive Authentication for the user pool. Amazon Cognito doesn't evaluate
// Identity and Access Management (IAM) policies in requests for this API
// operation. For this operation, you can't use IAM credentials to authorize
// requests, and you can't grant IAM permissions in policies. For more information
// about authorization models in Amazon Cognito, see Using the Amazon Cognito
// native and OIDC APIs (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
// .
func (c *Client) SetUserMFAPreference(ctx context.Context, params *SetUserMFAPreferenceInput, optFns ...func(*Options)) (*SetUserMFAPreferenceOutput, error) {
	if params == nil {
		params = &SetUserMFAPreferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUserMFAPreference", params, optFns, c.addOperationSetUserMFAPreferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUserMFAPreferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUserMFAPreferenceInput struct {

	// A valid access token that Amazon Cognito issued to the user whose MFA
	// preference you want to set.
	//
	// This member is required.
	AccessToken *string

	// The SMS text message multi-factor authentication (MFA) settings.
	SMSMfaSettings *types.SMSMfaSettingsType

	// The time-based one-time password (TOTP) software token MFA settings.
	SoftwareTokenMfaSettings *types.SoftwareTokenMfaSettingsType

	noSmithyDocumentSerde
}

type SetUserMFAPreferenceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetUserMFAPreferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUserMFAPreference{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUserMFAPreference{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetUserMFAPreference"); err != nil {
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
	if err = addOpSetUserMFAPreferenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetUserMFAPreference(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetUserMFAPreference(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetUserMFAPreference",
	}
}
