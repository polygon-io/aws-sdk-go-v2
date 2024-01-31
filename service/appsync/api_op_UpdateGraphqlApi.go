// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a GraphqlApi object.
func (c *Client) UpdateGraphqlApi(ctx context.Context, params *UpdateGraphqlApiInput, optFns ...func(*Options)) (*UpdateGraphqlApiOutput, error) {
	if params == nil {
		params = &UpdateGraphqlApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGraphqlApi", params, optFns, c.addOperationUpdateGraphqlApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGraphqlApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGraphqlApiInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The new name for the GraphqlApi object.
	//
	// This member is required.
	Name *string

	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []types.AdditionalAuthenticationProvider

	// The new authentication type for the GraphqlApi object.
	AuthenticationType types.AuthenticationType

	// Sets the value of the GraphQL API to enable ( ENABLED ) or disable ( DISABLED )
	// introspection. If no value is provided, the introspection configuration will be
	// set to ENABLED by default. This field will produce an error if the operation
	// attempts to use the introspection feature while this field is disabled. For more
	// information about introspection, see GraphQL introspection (https://graphql.org/learn/introspection/)
	// .
	IntrospectionConfig types.GraphQLApiIntrospectionConfig

	// Configuration for Lambda function authorization.
	LambdaAuthorizerConfig *types.LambdaAuthorizerConfig

	// The Amazon CloudWatch Logs configuration for the GraphqlApi object.
	LogConfig *types.LogConfig

	// The Identity and Access Management service role ARN for a merged API. The
	// AppSync service assumes this role on behalf of the Merged API to validate access
	// to source APIs at runtime and to prompt the AUTO_MERGE to update the merged API
	// endpoint with the source API changes automatically.
	MergedApiExecutionRoleArn *string

	// The OpenID Connect configuration for the GraphqlApi object.
	OpenIDConnectConfig *types.OpenIDConnectConfig

	// The owner contact information for an API resource. This field accepts any
	// string input with a length of 0 - 256 characters.
	OwnerContact *string

	// The maximum depth a query can have in a single request. Depth refers to the
	// amount of nested levels allowed in the body of query. The default value is 0
	// (or unspecified), which indicates there's no depth limit. If you set a limit, it
	// can be between 1 and 75 nested levels. This field will produce a limit error if
	// the operation falls out of bounds. Note that fields can still be set to nullable
	// or non-nullable. If a non-nullable field produces an error, the error will be
	// thrown upwards to the first nullable field available.
	QueryDepthLimit int32

	// The maximum number of resolvers that can be invoked in a single request. The
	// default value is 0 (or unspecified), which will set the limit to 10000 . When
	// specified, the limit value can be between 1 and 10000 . This field will produce
	// a limit error if the operation falls out of bounds.
	ResolverCountLimit int32

	// The new Amazon Cognito user pool configuration for the ~GraphqlApi object.
	UserPoolConfig *types.UserPoolConfig

	// A flag indicating whether to use X-Ray tracing for the GraphqlApi .
	XrayEnabled bool

	noSmithyDocumentSerde
}

type UpdateGraphqlApiOutput struct {

	// The updated GraphqlApi object.
	GraphqlApi *types.GraphqlApi

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGraphqlApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGraphqlApi"); err != nil {
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
	if err = addOpUpdateGraphqlApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGraphqlApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGraphqlApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGraphqlApi",
	}
}
