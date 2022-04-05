package api

import (
	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/models"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
)

// ForkedProvisioningApi always forwards requests to a Grafana backend.
// We do not currently support provisioning of external systems through Grafana's API.
type ForkedProvisioningApi struct {
	svc *ProvisioningSrv
}

// NewForkedProvisioningApi creates a new ForkedProvisioningApi instance.
func NewForkedProvisioningApi(svc *ProvisioningSrv) *ForkedProvisioningApi {
	return &ForkedProvisioningApi{
		svc: svc,
	}
}

func (f *ForkedProvisioningApi) forkRouteGetContactpoints(ctx *models.ReqContext) response.Response {
	return f.svc.RouteGetContactpoints(ctx)
}

func (f *ForkedProvisioningApi) forkRoutePostContactpoints(ctx *models.ReqContext, cp apimodels.Contactpoint) response.Response {
	return f.svc.RoutePostContactpoint(ctx, cp)
}

func (f *ForkedProvisioningApi) forkRoutePutContactpoints(ctx *models.ReqContext, cp apimodels.Contactpoint) response.Response {
	return f.svc.RoutePutContactpoints(ctx, cp)
}

func (f *ForkedProvisioningApi) forkRouteDeleteContactpoints(ctx *models.ReqContext) response.Response {
	return f.svc.RouteDeleteContactpoint(ctx)
}
