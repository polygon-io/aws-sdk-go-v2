// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Web Services organization. The account whose user is calling
// the CreateOrganization operation automatically becomes the management account (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account)
// of the new organization. This operation must be called using credentials from
// the account that is to become the new organization's management account. The
// principal must also have the relevant IAM permissions. By default (or if you set
// the FeatureSet parameter to ALL ), the new organization is created with all
// features enabled and service control policies automatically enabled in the root.
// If you instead choose to create the organization supporting only the
// consolidated billing features by setting the FeatureSet parameter to
// CONSOLIDATED_BILLING , no policy types are enabled by default and you can't use
// organization policies.
func (c *Client) CreateOrganization(ctx context.Context, params *CreateOrganizationInput, optFns ...func(*Options)) (*CreateOrganizationOutput, error) {
	if params == nil {
		params = &CreateOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOrganization", params, optFns, c.addOperationCreateOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOrganizationInput struct {

	// Specifies the feature set supported by the new organization. Each feature set
	// supports different levels of functionality.
	//   - CONSOLIDATED_BILLING : All member accounts have their bills consolidated to
	//   and paid by the management account. For more information, see Consolidated
	//   billing (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-cb-only)
	//   in the Organizations User Guide. The consolidated billing feature subset isn't
	//   available for organizations in the Amazon Web Services GovCloud (US) Region.
	//   - ALL : In addition to all the features supported by the consolidated billing
	//   feature set, the management account can also apply any policy type to any member
	//   account in the organization. For more information, see All features (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#feature-set-all)
	//   in the Organizations User Guide.
	FeatureSet types.OrganizationFeatureSet

	noSmithyDocumentSerde
}

type CreateOrganizationOutput struct {

	// A structure that contains details about the newly created organization.
	Organization *types.Organization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateOrganization"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateOrganization",
	}
}
