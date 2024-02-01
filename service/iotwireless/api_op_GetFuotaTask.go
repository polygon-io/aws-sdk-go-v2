// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a FUOTA task.
func (c *Client) GetFuotaTask(ctx context.Context, params *GetFuotaTaskInput, optFns ...func(*Options)) (*GetFuotaTaskOutput, error) {
	if params == nil {
		params = &GetFuotaTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFuotaTask", params, optFns, c.addOperationGetFuotaTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFuotaTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFuotaTaskInput struct {

	// The ID of a FUOTA task.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetFuotaTaskOutput struct {

	// The arn of a FUOTA task.
	Arn *string

	// Created at timestamp for the resource.
	CreatedAt *time.Time

	// The description of the new resource.
	Description *string

	// The S3 URI points to a firmware update image that is to be used with a FUOTA
	// task.
	FirmwareUpdateImage *string

	// The firmware update role that is to be used with a FUOTA task.
	FirmwareUpdateRole *string

	// The interval for sending fragments in milliseconds, rounded to the nearest
	// second. This interval only determines the timing for when the Cloud sends down
	// the fragments to yor device. There can be a delay for when your device will
	// receive these fragments. This delay depends on the device's class and the
	// communication delay with the cloud.
	FragmentIntervalMS *int32

	// The size of each fragment in bytes. This parameter is supported only for FUOTA
	// tasks with multicast groups.
	FragmentSizeBytes *int32

	// The ID of a FUOTA task.
	Id *string

	// The LoRaWAN information returned from getting a FUOTA task.
	LoRaWAN *types.LoRaWANFuotaTaskGetInfo

	// The name of a FUOTA task.
	Name *string

	// The percentage of the added fragments that are redundant. For example, if the
	// size of the firmware image file is 100 bytes and the fragment size is 10 bytes,
	// with RedundancyPercent set to 50(%), the final number of encoded fragments is
	// (100 / 10) + (100 / 10 * 50%) = 15.
	RedundancyPercent *int32

	// The status of a FUOTA task.
	Status types.FuotaTaskStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFuotaTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFuotaTask"); err != nil {
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
	if err = addOpGetFuotaTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFuotaTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFuotaTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFuotaTask",
	}
}
