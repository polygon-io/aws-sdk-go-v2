// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves one or more messages (up to 10), from the specified queue. Using the
// WaitTimeSeconds parameter enables long-poll support. For more information, see
// Amazon SQS Long Polling (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html)
// in the Amazon SQS Developer Guide. Short poll is the default behavior where a
// weighted random set of machines is sampled on a ReceiveMessage call. Thus, only
// the messages on the sampled machines are returned. If the number of messages in
// the queue is small (fewer than 1,000), you most likely get fewer messages than
// you requested per ReceiveMessage call. If the number of messages in the queue
// is extremely small, you might not receive any messages in a particular
// ReceiveMessage response. If this happens, repeat the request. For each message
// returned, the response includes the following:
//   - The message body.
//   - An MD5 digest of the message body. For information about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt)
//     .
//   - The MessageId you received when you sent the message to the queue.
//   - The receipt handle.
//   - The message attributes.
//   - An MD5 digest of the message attributes.
//
// The receipt handle is the identifier you must provide when deleting the
// message. For more information, see Queue and Message Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html)
// in the Amazon SQS Developer Guide. You can provide the VisibilityTimeout
// parameter in your request. The parameter is applied to the messages that Amazon
// SQS returns in the response. If you don't include the parameter, the overall
// visibility timeout for the queue is used for the returned messages. For more
// information, see Visibility Timeout (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
// in the Amazon SQS Developer Guide. A message that isn't deleted or a message
// whose visibility isn't extended before the visibility timeout expires counts as
// a failed receive. Depending on the configuration of the queue, the message might
// be sent to the dead-letter queue. In the future, new attributes might be added.
// If you write code that calls this action, we recommend that you structure your
// code so that it can handle new attributes gracefully.
func (c *Client) ReceiveMessage(ctx context.Context, params *ReceiveMessageInput, optFns ...func(*Options)) (*ReceiveMessageOutput, error) {
	if params == nil {
		params = &ReceiveMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReceiveMessage", params, optFns, c.addOperationReceiveMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReceiveMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReceiveMessageInput struct {

	// The URL of the Amazon SQS queue from which messages are received. Queue URLs
	// and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	// A list of attributes that need to be returned along with each message. These
	// attributes include:
	//   - All – Returns all values.
	//   - ApproximateFirstReceiveTimestamp – Returns the time the message was first
	//   received from the queue ( epoch time (http://en.wikipedia.org/wiki/Unix_time)
	//   in milliseconds).
	//   - ApproximateReceiveCount – Returns the number of times a message has been
	//   received across all queues but not deleted.
	//   - AWSTraceHeader – Returns the X-Ray trace header string.
	//   - SenderId
	//   - For a user, returns the user ID, for example ABCDEFGHI1JKLMNOPQ23R .
	//   - For an IAM role, returns the IAM role ID, for example
	//   ABCDE1F2GH3I4JK5LMNOP:i-a123b456 .
	//   - SentTimestamp – Returns the time the message was sent to the queue ( epoch
	//   time (http://en.wikipedia.org/wiki/Unix_time) in milliseconds).
	//   - SqsManagedSseEnabled – Enables server-side queue encryption using SQS owned
	//   encryption keys. Only one server-side encryption option is supported per queue
	//   (for example, SSE-KMS (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html)
	//   or SSE-SQS (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html)
	//   ).
	//   - MessageDeduplicationId – Returns the value provided by the producer that
	//   calls the SendMessage action.
	//   - MessageGroupId – Returns the value provided by the producer that calls the
	//   SendMessage action. Messages with the same MessageGroupId are returned in
	//   sequence.
	//   - SequenceNumber – Returns the value provided by Amazon SQS.
	AttributeNames []types.QueueAttributeName

	// The maximum number of messages to return. Amazon SQS never returns more
	// messages than this value (however, fewer messages might be returned). Valid
	// values: 1 to 10. Default: 1.
	MaxNumberOfMessages int32

	// The name of the message attribute, where N is the index.
	//   - The name can contain alphanumeric characters and the underscore ( _ ),
	//   hyphen ( - ), and period ( . ).
	//   - The name is case-sensitive and must be unique among all attribute names for
	//   the message.
	//   - The name must not start with AWS-reserved prefixes such as AWS. or Amazon.
	//   (or any casing variants).
	//   - The name must not start or end with a period ( . ), and it should not have
	//   periods in succession ( .. ).
	//   - The name can be up to 256 characters long.
	// When using ReceiveMessage , you can send a list of attribute names to receive,
	// or you can return all of the attributes by specifying All or .* in your
	// request. You can also use all message attributes starting with a prefix, for
	// example bar.* .
	MessageAttributeNames []string

	// This parameter applies only to FIFO (first-in-first-out) queues. The token used
	// for deduplication of ReceiveMessage calls. If a networking issue occurs after a
	// ReceiveMessage action, and instead of a response you receive a generic error, it
	// is possible to retry the same action with an identical ReceiveRequestAttemptId
	// to retrieve the same set of messages, even if their visibility timeout has not
	// yet expired.
	//   - You can use ReceiveRequestAttemptId only for 5 minutes after a
	//   ReceiveMessage action.
	//   - When you set FifoQueue , a caller of the ReceiveMessage action can provide a
	//   ReceiveRequestAttemptId explicitly.
	//   - If a caller of the ReceiveMessage action doesn't provide a
	//   ReceiveRequestAttemptId , Amazon SQS generates a ReceiveRequestAttemptId .
	//   - It is possible to retry the ReceiveMessage action with the same
	//   ReceiveRequestAttemptId if none of the messages have been modified (deleted or
	//   had their visibility changes).
	//   - During a visibility timeout, subsequent calls with the same
	//   ReceiveRequestAttemptId return the same messages and receipt handles. If a
	//   retry occurs within the deduplication interval, it resets the visibility
	//   timeout. For more information, see Visibility Timeout (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	//   in the Amazon SQS Developer Guide. If a caller of the ReceiveMessage action
	//   still processes messages when the visibility timeout expires and messages become
	//   visible, another worker consuming from the same queue can receive the same
	//   messages and therefore process duplicates. Also, if a consumer whose message
	//   processing time is longer than the visibility timeout tries to delete the
	//   processed messages, the action fails with an error. To mitigate this effect,
	//   ensure that your application observes a safe threshold before the visibility
	//   timeout expires and extend the visibility timeout as necessary.
	//   - While messages with a particular MessageGroupId are invisible, no more
	//   messages belonging to the same MessageGroupId are returned until the
	//   visibility timeout expires. You can still receive messages with another
	//   MessageGroupId as long as it is also visible.
	//   - If a caller of ReceiveMessage can't track the ReceiveRequestAttemptId , no
	//   retries work until the original visibility timeout expires. As a result, delays
	//   might occur but the messages in the queue remain in a strict order.
	// The maximum length of ReceiveRequestAttemptId is 128 characters.
	// ReceiveRequestAttemptId can contain alphanumeric characters ( a-z , A-Z , 0-9 )
	// and punctuation ( !"#$%&'()*+,-./:;<=>?@[\]^_`{|}~ ). For best practices of
	// using ReceiveRequestAttemptId , see Using the ReceiveRequestAttemptId Request
	// Parameter (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-receiverequestattemptid-request-parameter.html)
	// in the Amazon SQS Developer Guide.
	ReceiveRequestAttemptId *string

	// The duration (in seconds) that the received messages are hidden from subsequent
	// retrieve requests after being retrieved by a ReceiveMessage request.
	VisibilityTimeout int32

	// The duration (in seconds) for which the call waits for a message to arrive in
	// the queue before returning. If a message is available, the call returns sooner
	// than WaitTimeSeconds . If no messages are available and the wait time expires,
	// the call returns successfully with an empty list of messages. To avoid HTTP
	// errors, ensure that the HTTP response timeout for ReceiveMessage requests is
	// longer than the WaitTimeSeconds parameter. For example, with the Java SDK, you
	// can set HTTP transport settings using the NettyNioAsyncHttpClient (https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/http/nio/netty/NettyNioAsyncHttpClient.html)
	// for asynchronous clients, or the ApacheHttpClient (https://sdk.amazonaws.com/java/api/latest/software/amazon/awssdk/http/apache/ApacheHttpClient.html)
	// for synchronous clients.
	WaitTimeSeconds int32

	noSmithyDocumentSerde
}

// A list of received messages.
type ReceiveMessageOutput struct {

	// A list of messages.
	Messages []types.Message

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReceiveMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpReceiveMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpReceiveMessage{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addValidateReceiveMessageChecksum(stack, options); err != nil {
		return err
	}
	if err = addOpReceiveMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReceiveMessage(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opReceiveMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "ReceiveMessage",
	}
}
