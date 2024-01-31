// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmeetings

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops transcription for the specified meetingId . For more information, refer to
// Using Amazon Chime SDK live transcription  (https://docs.aws.amazon.com/chime-sdk/latest/dg/meeting-transcription.html)
// in the Amazon Chime SDK Developer Guide. By default, Amazon Transcribe may use
// and store audio content processed by the service to develop and improve Amazon
// Web Services AI/ML services as further described in section 50 of the Amazon
// Web Services Service Terms (https://aws.amazon.com/service-terms/) . Using
// Amazon Transcribe may be subject to federal and state laws or regulations
// regarding the recording or interception of electronic communications. It is your
// and your end users’ responsibility to comply with all applicable laws regarding
// the recording, including properly notifying all participants in a recorded
// session or communication that the session or communication is being recorded,
// and obtaining all necessary consents. You can opt out from Amazon Web Services
// using audio content to develop and improve Amazon Web Services AI/ML services by
// configuring an AI services opt out policy using Amazon Web Services
// Organizations.
func (c *Client) StopMeetingTranscription(ctx context.Context, params *StopMeetingTranscriptionInput, optFns ...func(*Options)) (*StopMeetingTranscriptionOutput, error) {
	if params == nil {
		params = &StopMeetingTranscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopMeetingTranscription", params, optFns, c.addOperationStopMeetingTranscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopMeetingTranscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopMeetingTranscriptionInput struct {

	// The unique ID of the meeting for which you stop transcription.
	//
	// This member is required.
	MeetingId *string

	noSmithyDocumentSerde
}

type StopMeetingTranscriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopMeetingTranscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStopMeetingTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStopMeetingTranscription{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StopMeetingTranscription"); err != nil {
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
	if err = addOpStopMeetingTranscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopMeetingTranscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopMeetingTranscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StopMeetingTranscription",
	}
}
