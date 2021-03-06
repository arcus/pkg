package http

import (
	"net/http"
	"strings"

	"github.com/arcus/pkg/service"
)

const (
	AuthorizationHeader = "authorization"

	ContextPrincipalIdHeader    = "x-arcus-principal-id"
	ContextPrincipalRoleHeader  = "x-arcus-principal-role"
	ContextPrincipalTypeHeader  = "x-arcus-principal-type"
	ContextPrincipalNameHeader  = "x-arcus-principal-name"
	ContextPrincipalEmailHeader = "x-arcus-principal-email"

	ContextAccessTokenHeader = "x-arcus-access-token"
)

func parseToken(s string) string {
	idx := strings.IndexByte(s, ' ')
	if idx < 0 {
		return ""
	}
	if strings.ToLower(s[:idx]) != "bearer" {
		return ""
	}
	return s[idx+1:]
}

func ContextFromRequest(req *http.Request) *service.Context {
	// If this was already set (in upstream middleware) then just use as is.
	ctx, ok := req.Context().(*service.Context)
	if ok {
		return ctx
	}

	ctx = service.WithContext(req.Context())

	ctx.PrincipalId = req.Header.Get(ContextPrincipalIdHeader)
	ctx.PrincipalRole = req.Header.Get(ContextPrincipalRoleHeader)
	ctx.PrincipalType = req.Header.Get(ContextPrincipalTypeHeader)
	ctx.PrincipalEmail = req.Header.Get(ContextPrincipalEmailHeader)
	ctx.PrincipalName = req.Header.Get(ContextPrincipalNameHeader)

	ctx.AccessToken = req.Header.Get(ContextAccessTokenHeader)
	// Fallback to authorization header if not set.
	if ctx.AccessToken == "" {
		ctx.AccessToken = parseToken(req.Header.Get(AuthorizationHeader))
	}

	return ctx
}

func ContextToRequest(ctx *service.Context, req *http.Request) {
	req.Header.Set(ContextPrincipalIdHeader, ctx.PrincipalId)
	req.Header.Set(ContextPrincipalRoleHeader, ctx.PrincipalRole)
	req.Header.Set(ContextPrincipalTypeHeader, ctx.PrincipalType)
	req.Header.Set(ContextPrincipalEmailHeader, ctx.PrincipalEmail)
	req.Header.Set(ContextPrincipalNameHeader, ctx.PrincipalName)
	req.Header.Set(ContextAccessTokenHeader, ctx.AccessToken)
	req.Header.Set(AuthorizationHeader, "Bearer "+ctx.AccessToken)
}
