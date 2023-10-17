// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycluster

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycluster/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List routing control names and Amazon Resource Names (ARNs), as well as the
// routing control state for each routing control, along with the control panel
// name and control panel ARN for the routing controls. If you specify a control
// panel ARN, this call lists the routing controls in the control panel. Otherwise,
// it lists all the routing controls in the cluster. A routing control is a simple
// on/off switch in Route 53 ARC that you can use to route traffic to cells. When a
// routing control state is set to ON, traffic flows to a cell. When the state is
// set to OFF, traffic does not flow. Before you can create a routing control, you
// must first create a cluster, and then host the control in a control panel on the
// cluster. For more information, see Create routing control structures (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.create.html)
// in the Amazon Route 53 Application Recovery Controller Developer Guide. You
// access one of the endpoints for the cluster to get or update the routing control
// state to redirect traffic for your application. You must specify Regional
// endpoints when you work with API cluster operations to use this API operation to
// list routing controls in Route 53 ARC. Learn more about working with routing
// controls in the following topics in the Amazon Route 53 Application Recovery
// Controller Developer Guide:
//   - Viewing and updating routing control states (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.update.html)
//   - Working with routing controls in Route 53 ARC (https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.html)
func (c *Client) ListRoutingControls(ctx context.Context, params *ListRoutingControlsInput, optFns ...func(*Options)) (*ListRoutingControlsOutput, error) {
	if params == nil {
		params = &ListRoutingControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoutingControls", params, optFns, c.addOperationListRoutingControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoutingControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoutingControlsInput struct {

	// The Amazon Resource Name (ARN) of the control panel of the routing controls to
	// list.
	ControlPanelArn *string

	// The number of routing controls objects that you want to return with this call.
	// The default value is 500.
	MaxResults *int32

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRoutingControlsOutput struct {

	// The list of routing controls.
	//
	// This member is required.
	RoutingControls []types.RoutingControl

	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoutingControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListRoutingControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListRoutingControls{}, middleware.After)
	if err != nil {
		return err
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = addListRoutingControlsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoutingControls(options.Region), middleware.Before); err != nil {
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
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListRoutingControlsAPIClient is a client that implements the
// ListRoutingControls operation.
type ListRoutingControlsAPIClient interface {
	ListRoutingControls(context.Context, *ListRoutingControlsInput, ...func(*Options)) (*ListRoutingControlsOutput, error)
}

var _ ListRoutingControlsAPIClient = (*Client)(nil)

// ListRoutingControlsPaginatorOptions is the paginator options for
// ListRoutingControls
type ListRoutingControlsPaginatorOptions struct {
	// The number of routing controls objects that you want to return with this call.
	// The default value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoutingControlsPaginator is a paginator for ListRoutingControls
type ListRoutingControlsPaginator struct {
	options   ListRoutingControlsPaginatorOptions
	client    ListRoutingControlsAPIClient
	params    *ListRoutingControlsInput
	nextToken *string
	firstPage bool
}

// NewListRoutingControlsPaginator returns a new ListRoutingControlsPaginator
func NewListRoutingControlsPaginator(client ListRoutingControlsAPIClient, params *ListRoutingControlsInput, optFns ...func(*ListRoutingControlsPaginatorOptions)) *ListRoutingControlsPaginator {
	if params == nil {
		params = &ListRoutingControlsInput{}
	}

	options := ListRoutingControlsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoutingControlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoutingControlsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoutingControls page.
func (p *ListRoutingControlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoutingControlsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListRoutingControls(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRoutingControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-cluster",
		OperationName: "ListRoutingControls",
	}
}

type opListRoutingControlsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opListRoutingControlsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opListRoutingControlsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "route53-recovery-cluster"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "route53-recovery-cluster"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("route53-recovery-cluster")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListRoutingControlsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListRoutingControlsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}
