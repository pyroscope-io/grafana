/*Package api contains base API implementation of unified alerting
 *
 *Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 *
 *Do not manually edit these files, please find ngalert/api/swagger-codegen/ for commands on how to generate them.
 */

package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/api/routing"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/models"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	"github.com/grafana/grafana/pkg/services/ngalert/metrics"
	"github.com/grafana/grafana/pkg/web"
)

type ProvisioningApiForkingService interface {
	RouteDeleteContactpoints(*models.ReqContext) response.Response
	RouteGetContactpoints(*models.ReqContext) response.Response
	RoutePostContactpoints(*models.ReqContext) response.Response
	RoutePutContactpoints(*models.ReqContext) response.Response
}

func (f *ForkedProvisioningApi) RouteDeleteContactpoints(ctx *models.ReqContext) response.Response {
	return f.forkRouteDeleteContactpoints(ctx)
}

func (f *ForkedProvisioningApi) RouteGetContactpoints(ctx *models.ReqContext) response.Response {
	return f.forkRouteGetContactpoints(ctx)
}

func (f *ForkedProvisioningApi) RoutePostContactpoints(ctx *models.ReqContext) response.Response {
	conf := apimodels.Contactpoint{}
	if err := web.Bind(ctx.Req, &conf); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}
	return f.forkRoutePostContactpoints(ctx, conf)
}

func (f *ForkedProvisioningApi) RoutePutContactpoints(ctx *models.ReqContext) response.Response {
	conf := apimodels.Contactpoint{}
	if err := web.Bind(ctx.Req, &conf); err != nil {
		return response.Error(http.StatusBadRequest, "bad request data", err)
	}
	return f.forkRoutePutContactpoints(ctx, conf)
}

func (api *API) RegisterProvisioningApiEndpoints(srv ProvisioningApiForkingService, m *metrics.API) {
	api.RouteRegister.Group("", func(group routing.RouteRegister) {
		group.Delete(
			toMacaronPath("/api/provisioning/contactpoints"),
			api.authorize(http.MethodDelete, "/api/provisioning/contactpoints"),
			metrics.Instrument(
				http.MethodDelete,
				"/api/provisioning/contactpoints",
				srv.RouteDeleteContactpoints,
				m,
			),
		)
		group.Get(
			toMacaronPath("/api/provisioning/contactpoints"),
			api.authorize(http.MethodGet, "/api/provisioning/contactpoints"),
			metrics.Instrument(
				http.MethodGet,
				"/api/provisioning/contactpoints",
				srv.RouteGetContactpoints,
				m,
			),
		)
		group.Post(
			toMacaronPath("/api/provisioning/contactpoints"),
			api.authorize(http.MethodPost, "/api/provisioning/contactpoints"),
			metrics.Instrument(
				http.MethodPost,
				"/api/provisioning/contactpoints",
				srv.RoutePostContactpoints,
				m,
			),
		)
		group.Put(
			toMacaronPath("/api/provisioning/contactpoints"),
			api.authorize(http.MethodPut, "/api/provisioning/contactpoints"),
			metrics.Instrument(
				http.MethodPut,
				"/api/provisioning/contactpoints",
				srv.RoutePutContactpoints,
				m,
			),
		)
	}, middleware.ReqSignedIn)
}
