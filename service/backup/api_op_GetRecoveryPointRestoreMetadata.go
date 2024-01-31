// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a set of metadata key-value pairs that were used to create the backup.
func (c *Client) GetRecoveryPointRestoreMetadata(ctx context.Context, params *GetRecoveryPointRestoreMetadataInput, optFns ...func(*Options)) (*GetRecoveryPointRestoreMetadataOutput, error) {
	if params == nil {
		params = &GetRecoveryPointRestoreMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecoveryPointRestoreMetadata", params, optFns, c.addOperationGetRecoveryPointRestoreMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecoveryPointRestoreMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecoveryPointRestoreMetadataInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Amazon Web Services Region where they are created. They consist of lowercase
	// letters, numbers, and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	//
	// This member is required.
	RecoveryPointArn *string

	// This is the account ID of the specified backup vault.
	BackupVaultAccountId *string

	noSmithyDocumentSerde
}

type GetRecoveryPointRestoreMetadataOutput struct {

	// An ARN that uniquely identifies a backup vault; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault .
	BackupVaultArn *string

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45
	// .
	RecoveryPointArn *string

	// This is the resource type associated with the recovery point.
	ResourceType *string

	// The set of metadata key-value pairs that describe the original configuration of
	// the backed-up resource. These values vary depending on the service that is being
	// restored.
	RestoreMetadata map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecoveryPointRestoreMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecoveryPointRestoreMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecoveryPointRestoreMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetRecoveryPointRestoreMetadata"); err != nil {
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
	if err = addOpGetRecoveryPointRestoreMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecoveryPointRestoreMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRecoveryPointRestoreMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetRecoveryPointRestoreMetadata",
	}
}
