// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Merges two specified branches using the three-way merge strategy.
func (c *Client) MergeBranchesByThreeWay(ctx context.Context, params *MergeBranchesByThreeWayInput, optFns ...func(*Options)) (*MergeBranchesByThreeWayOutput, error) {
	if params == nil {
		params = &MergeBranchesByThreeWayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MergeBranchesByThreeWay", params, optFns, c.addOperationMergeBranchesByThreeWayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MergeBranchesByThreeWayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MergeBranchesByThreeWayInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	DestinationCommitSpecifier *string

	// The name of the repository where you want to merge two branches.
	//
	// This member is required.
	RepositoryName *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	SourceCommitSpecifier *string

	// The name of the author who created the commit. This information is used as both
	// the author and committer for the commit.
	AuthorName *string

	// The commit message to include in the commit information for the merge.
	CommitMessage *string

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum

	// If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when
	// resolving conflicts during a merge.
	ConflictResolution *types.ConflictResolution

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum

	// The email address of the person merging the branches. This information is used
	// in the commit information for the merge.
	Email *string

	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If true, a .gitkeep file is created for
	// empty folders. The default is false.
	KeepEmptyFolders bool

	// The branch where the merge is applied.
	TargetBranch *string

	noSmithyDocumentSerde
}

type MergeBranchesByThreeWayOutput struct {

	// The commit ID of the merge in the destination or target branch.
	CommitId *string

	// The tree ID of the merge in the destination or target branch.
	TreeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMergeBranchesByThreeWayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMergeBranchesByThreeWay{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergeBranchesByThreeWay{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "MergeBranchesByThreeWay"); err != nil {
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
	if err = addOpMergeBranchesByThreeWayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMergeBranchesByThreeWay(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMergeBranchesByThreeWay(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "MergeBranchesByThreeWay",
	}
}
