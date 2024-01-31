// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The UpdateNotificationSettings operation creates, updates, disables or
// re-enables notifications for a HIT type. If you call the
// UpdateNotificationSettings operation for a HIT type that already has a
// notification specification, the operation replaces the old specification with a
// new one. You can call the UpdateNotificationSettings operation to enable or
// disable notifications for the HIT type, without having to modify the
// notification specification itself by providing updates to the Active status
// without specifying a new notification specification. To change the Active status
// of a HIT type's notifications, the HIT type must already have a notification
// specification, or one must be provided in the same call to
// UpdateNotificationSettings .
func (c *Client) UpdateNotificationSettings(ctx context.Context, params *UpdateNotificationSettingsInput, optFns ...func(*Options)) (*UpdateNotificationSettingsOutput, error) {
	if params == nil {
		params = &UpdateNotificationSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNotificationSettings", params, optFns, c.addOperationUpdateNotificationSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNotificationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNotificationSettingsInput struct {

	// The ID of the HIT type whose notification specification is being updated.
	//
	// This member is required.
	HITTypeId *string

	// Specifies whether notifications are sent for HITs of this HIT type, according
	// to the notification specification. You must specify either the Notification
	// parameter or the Active parameter for the call to UpdateNotificationSettings to
	// succeed.
	Active *bool

	// The notification specification for the HIT type.
	Notification *types.NotificationSpecification

	noSmithyDocumentSerde
}

type UpdateNotificationSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNotificationSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateNotificationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateNotificationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateNotificationSettings"); err != nil {
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
	if err = addOpUpdateNotificationSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNotificationSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNotificationSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateNotificationSettings",
	}
}
