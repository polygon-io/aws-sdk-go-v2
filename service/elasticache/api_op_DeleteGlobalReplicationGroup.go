// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deleting a Global datastore is a two-step process:
//   - First, you must DisassociateGlobalReplicationGroup to remove the secondary
//     clusters in the Global datastore.
//   - Once the Global datastore contains only the primary cluster, you can use
//     the DeleteGlobalReplicationGroup API to delete the Global datastore while
//     retainining the primary cluster using RetainPrimaryReplicationGroup=true .
//
// Since the Global Datastore has only a primary cluster, you can delete the
// Global Datastore while retaining the primary by setting
// RetainPrimaryReplicationGroup=true . The primary cluster is never deleted when
// deleting a Global Datastore. It can only be deleted when it no longer is
// associated with any Global Datastore. When you receive a successful response
// from this operation, Amazon ElastiCache immediately begins deleting the selected
// resources; you cannot cancel or revert this operation.
func (c *Client) DeleteGlobalReplicationGroup(ctx context.Context, params *DeleteGlobalReplicationGroupInput, optFns ...func(*Options)) (*DeleteGlobalReplicationGroupOutput, error) {
	if params == nil {
		params = &DeleteGlobalReplicationGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGlobalReplicationGroup", params, optFns, c.addOperationDeleteGlobalReplicationGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGlobalReplicationGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGlobalReplicationGroupInput struct {

	// The name of the Global datastore
	//
	// This member is required.
	GlobalReplicationGroupId *string

	// The primary replication group is retained as a standalone replication group.
	//
	// This member is required.
	RetainPrimaryReplicationGroup *bool

	noSmithyDocumentSerde
}

type DeleteGlobalReplicationGroupOutput struct {

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different Amazon region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the
	// secondary cluster.
	//   - The GlobalReplicationGroupIdSuffix represents the name of the Global
	//   datastore, which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *types.GlobalReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteGlobalReplicationGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteGlobalReplicationGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteGlobalReplicationGroup"); err != nil {
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
	if err = addOpDeleteGlobalReplicationGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGlobalReplicationGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGlobalReplicationGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteGlobalReplicationGroup",
	}
}
