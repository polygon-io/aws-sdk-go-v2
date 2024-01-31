// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends an email message. You can use the Amazon Pinpoint Email API to send two
// types of messages:
//   - Simple – A standard email message. When you create this type of message,
//     you specify the sender, the recipient, and the message body, and Amazon Pinpoint
//     assembles the message for you.
//   - Raw – A raw, MIME-formatted email message. When you send this type of
//     email, you have to specify all of the message headers, as well as the message
//     body. You can use this message type to send messages that contain attachments.
//     The message that you specify has to be a valid MIME message.
func (c *Client) SendEmail(ctx context.Context, params *SendEmailInput, optFns ...func(*Options)) (*SendEmailOutput, error) {
	if params == nil {
		params = &SendEmailInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendEmail", params, optFns, c.addOperationSendEmailMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to send an email message.
type SendEmailInput struct {

	// An object that contains the body of the message. You can send either a Simple
	// message or a Raw message.
	//
	// This member is required.
	Content *types.EmailContent

	// An object that contains the recipients of the email message.
	//
	// This member is required.
	Destination *types.Destination

	// The name of the configuration set that you want to use when sending the email.
	ConfigurationSetName *string

	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send using the SendEmail operation. Tags correspond to characteristics of the
	// email that you define, so that you can publish email sending events.
	EmailTags []types.MessageTag

	// The address that Amazon Pinpoint should send bounce and complaint notifications
	// to.
	FeedbackForwardingEmailAddress *string

	// The email address that you want to use as the "From" address for the email. The
	// address that you specify has to be verified.
	FromEmailAddress *string

	// The "Reply-to" email addresses for the message. When the recipient replies to
	// the message, each Reply-to address receives the reply.
	ReplyToAddresses []string

	noSmithyDocumentSerde
}

// A unique message ID that you receive when Amazon Pinpoint accepts an email for
// sending.
type SendEmailOutput struct {

	// A unique identifier for the message that is generated when Amazon Pinpoint
	// accepts the message. It is possible for Amazon Pinpoint to accept a message
	// without sending it. This can happen when the message you're trying to send has
	// an attachment doesn't pass a virus check, or when you send a templated email
	// that contains invalid personalization content, for example.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendEmailMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSendEmail{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSendEmail{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SendEmail"); err != nil {
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
	if err = addOpSendEmailValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendEmail(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendEmail(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SendEmail",
	}
}
