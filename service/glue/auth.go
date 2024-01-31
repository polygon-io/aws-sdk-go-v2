// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	smithy "github.com/aws/smithy-go"
	smithyauth "github.com/aws/smithy-go/auth"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func bindAuthParamsRegion(params *AuthResolverParameters, _ interface{}, options Options) {
	params.Region = options.Region
}

type setLegacyContextSigningOptionsMiddleware struct {
}

func (*setLegacyContextSigningOptionsMiddleware) ID() string {
	return "setLegacyContextSigningOptions"
}

func (m *setLegacyContextSigningOptionsMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	rscheme := getResolvedAuthScheme(ctx)
	schemeID := rscheme.Scheme.SchemeID()

	if sn := awsmiddleware.GetSigningName(ctx); sn != "" {
		if schemeID == "aws.auth#sigv4" {
			smithyhttp.SetSigV4SigningName(&rscheme.SignerProperties, sn)
		} else if schemeID == "aws.auth#sigv4a" {
			smithyhttp.SetSigV4ASigningName(&rscheme.SignerProperties, sn)
		}
	}

	if sr := awsmiddleware.GetSigningRegion(ctx); sr != "" {
		if schemeID == "aws.auth#sigv4" {
			smithyhttp.SetSigV4SigningRegion(&rscheme.SignerProperties, sr)
		} else if schemeID == "aws.auth#sigv4a" {
			smithyhttp.SetSigV4ASigningRegions(&rscheme.SignerProperties, []string{sr})
		}
	}

	return next.HandleFinalize(ctx, in)
}

func addSetLegacyContextSigningOptionsMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&setLegacyContextSigningOptionsMiddleware{}, "Signing", middleware.Before)
}

type withAnonymous struct {
	resolver AuthSchemeResolver
}

var _ AuthSchemeResolver = (*withAnonymous)(nil)

func (v *withAnonymous) ResolveAuthSchemes(ctx context.Context, params *AuthResolverParameters) ([]*smithyauth.Option, error) {
	opts, err := v.resolver.ResolveAuthSchemes(ctx, params)
	if err != nil {
		return nil, err
	}

	opts = append(opts, &smithyauth.Option{
		SchemeID: smithyauth.SchemeIDAnonymous,
	})
	return opts, nil
}

func wrapWithAnonymousAuth(options *Options) {
	if _, ok := options.AuthSchemeResolver.(*defaultAuthSchemeResolver); !ok {
		return
	}

	options.AuthSchemeResolver = &withAnonymous{
		resolver: options.AuthSchemeResolver,
	}
}

// AuthResolverParameters contains the set of inputs necessary for auth scheme
// resolution.
type AuthResolverParameters struct {
	// The name of the operation being invoked.
	Operation string

	// The region in which the operation is being invoked.
	Region string
}

func bindAuthResolverParams(operation string, input interface{}, options Options) *AuthResolverParameters {
	params := &AuthResolverParameters{
		Operation: operation,
	}

	bindAuthParamsRegion(params, input, options)

	return params
}

// AuthSchemeResolver returns a set of possible authentication options for an
// operation.
type AuthSchemeResolver interface {
	ResolveAuthSchemes(context.Context, *AuthResolverParameters) ([]*smithyauth.Option, error)
}

type defaultAuthSchemeResolver struct{}

var _ AuthSchemeResolver = (*defaultAuthSchemeResolver)(nil)

func (*defaultAuthSchemeResolver) ResolveAuthSchemes(ctx context.Context, params *AuthResolverParameters) ([]*smithyauth.Option, error) {
	if overrides, ok := operationAuthOptions[params.Operation]; ok {
		return overrides(params), nil
	}
	return serviceAuthOptions(params), nil
}

var operationAuthOptions = map[string]func(*AuthResolverParameters) []*smithyauth.Option{}

func serviceAuthOptions(params *AuthResolverParameters) []*smithyauth.Option {
	return []*smithyauth.Option{
		{
			SchemeID: smithyauth.SchemeIDSigV4,
			SignerProperties: func() smithy.Properties {
				var props smithy.Properties
				smithyhttp.SetSigV4SigningName(&props, "glue")
				smithyhttp.SetSigV4SigningRegion(&props, params.Region)
				return props
			}(),
		},
	}
}

type resolveAuthSchemeMiddleware struct {
	operation string
	options   Options
}

func (*resolveAuthSchemeMiddleware) ID() string {
	return "ResolveAuthScheme"
}

func (m *resolveAuthSchemeMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	params := bindAuthResolverParams(m.operation, getOperationInput(ctx), m.options)
	options, err := m.options.AuthSchemeResolver.ResolveAuthSchemes(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("resolve auth scheme: %w", err)
	}

	scheme, ok := m.selectScheme(options)
	if !ok {
		return out, metadata, fmt.Errorf("could not select an auth scheme")
	}

	ctx = setResolvedAuthScheme(ctx, scheme)
	return next.HandleFinalize(ctx, in)
}

func (m *resolveAuthSchemeMiddleware) selectScheme(options []*smithyauth.Option) (*resolvedAuthScheme, bool) {
	for _, option := range options {
		if option.SchemeID == smithyauth.SchemeIDAnonymous {
			return newResolvedAuthScheme(smithyhttp.NewAnonymousScheme(), option), true
		}

		for _, scheme := range m.options.AuthSchemes {
			if scheme.SchemeID() != option.SchemeID {
				continue
			}

			if scheme.IdentityResolver(m.options) != nil {
				return newResolvedAuthScheme(scheme, option), true
			}
		}
	}

	return nil, false
}

type resolvedAuthSchemeKey struct{}

type resolvedAuthScheme struct {
	Scheme             smithyhttp.AuthScheme
	IdentityProperties smithy.Properties
	SignerProperties   smithy.Properties
}

func newResolvedAuthScheme(scheme smithyhttp.AuthScheme, option *smithyauth.Option) *resolvedAuthScheme {
	return &resolvedAuthScheme{
		Scheme:             scheme,
		IdentityProperties: option.IdentityProperties,
		SignerProperties:   option.SignerProperties,
	}
}

func setResolvedAuthScheme(ctx context.Context, scheme *resolvedAuthScheme) context.Context {
	return middleware.WithStackValue(ctx, resolvedAuthSchemeKey{}, scheme)
}

func getResolvedAuthScheme(ctx context.Context) *resolvedAuthScheme {
	v, _ := middleware.GetStackValue(ctx, resolvedAuthSchemeKey{}).(*resolvedAuthScheme)
	return v
}

type getIdentityMiddleware struct {
	options Options
}

func (*getIdentityMiddleware) ID() string {
	return "GetIdentity"
}

func (m *getIdentityMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	rscheme := getResolvedAuthScheme(ctx)
	if rscheme == nil {
		return out, metadata, fmt.Errorf("no resolved auth scheme")
	}

	resolver := rscheme.Scheme.IdentityResolver(m.options)
	if resolver == nil {
		return out, metadata, fmt.Errorf("no identity resolver")
	}

	identity, err := resolver.GetIdentity(ctx, rscheme.IdentityProperties)
	if err != nil {
		return out, metadata, fmt.Errorf("get identity: %w", err)
	}

	ctx = setIdentity(ctx, identity)
	return next.HandleFinalize(ctx, in)
}

type identityKey struct{}

func setIdentity(ctx context.Context, identity smithyauth.Identity) context.Context {
	return middleware.WithStackValue(ctx, identityKey{}, identity)
}

func getIdentity(ctx context.Context) smithyauth.Identity {
	v, _ := middleware.GetStackValue(ctx, identityKey{}).(smithyauth.Identity)
	return v
}

type signRequestMiddleware struct {
}

func (*signRequestMiddleware) ID() string {
	return "Signing"
}

func (m *signRequestMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", in.Request)
	}

	rscheme := getResolvedAuthScheme(ctx)
	if rscheme == nil {
		return out, metadata, fmt.Errorf("no resolved auth scheme")
	}

	identity := getIdentity(ctx)
	if identity == nil {
		return out, metadata, fmt.Errorf("no identity")
	}

	signer := rscheme.Scheme.Signer()
	if signer == nil {
		return out, metadata, fmt.Errorf("no signer")
	}

	if err := signer.SignRequest(ctx, req, identity, rscheme.SignerProperties); err != nil {
		return out, metadata, fmt.Errorf("sign request: %w", err)
	}

	return next.HandleFinalize(ctx, in)
}
