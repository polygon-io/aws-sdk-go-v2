// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a snapshot of a volume. You can take snapshots of your gateway volumes
// on a scheduled or ad hoc basis. This API action enables you to delete a snapshot
// schedule for a volume. For more information, see Backing up your volumes (https://docs.aws.amazon.com/storagegateway/latest/userguide/backing-up-volumes.html)
// . In the DeleteSnapshotSchedule request, you identify the volume by providing
// its Amazon Resource Name (ARN). This operation is only supported for cached
// volume gateway types. To list or delete a snapshot, you must use the Amazon EC2
// API. For more information, go to DescribeSnapshots (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html)
// in the Amazon Elastic Compute Cloud API Reference.
func (c *Client) DeleteSnapshotSchedule(ctx context.Context, params *DeleteSnapshotScheduleInput, optFns ...func(*Options)) (*DeleteSnapshotScheduleOutput, error) {
	if params == nil {
		params = &DeleteSnapshotScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSnapshotSchedule", params, optFns, c.addOperationDeleteSnapshotScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSnapshotScheduleInput struct {

	// The volume which snapshot schedule to delete.
	//
	// This member is required.
	VolumeARN *string

	noSmithyDocumentSerde
}

type DeleteSnapshotScheduleOutput struct {

	// The volume which snapshot schedule was deleted.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSnapshotScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteSnapshotSchedule"); err != nil {
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
	if err = addOpDeleteSnapshotScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSnapshotSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSnapshotSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteSnapshotSchedule",
	}
}
