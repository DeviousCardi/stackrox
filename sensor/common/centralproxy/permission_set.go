package centralproxy

import (
	"context"
)

const metricsPermissionSetName = "OpenShift Prometheus Metrics Reader"

type proxyPermissionSetKey struct{}

var proxyPathPermissionSets = map[string]string{
	"/metrics": metricsPermissionSetName,
}

func withMetricsScrape(ctx context.Context) context.Context {
	return context.WithValue(ctx, proxyPermissionSetKey{}, proxyPathPermissionSets["/metrics"])
}

func proxyPermissionSetFromContext(ctx context.Context) string {
	permissionSetName, _ := ctx.Value(proxyPermissionSetKey{}).(string)
	return permissionSetName
}
