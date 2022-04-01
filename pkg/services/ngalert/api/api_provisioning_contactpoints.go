package api

import (
	"net/http"

	"github.com/grafana/grafana/pkg/api/response"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/models"
	apimodels "github.com/grafana/grafana/pkg/services/ngalert/api/tooling/definitions"
	domain "github.com/grafana/grafana/pkg/services/ngalert/models"
	"github.com/grafana/grafana/pkg/web"
)

type ProvisioningSrv struct {
	log                 log.Logger
	contactpointService ContactpointService
}

type ContactpointService interface {
	GetContactPoints(orgID int64) ([]domain.EmbeddedContactPoint, error)
	CreateContactPoint(orgID int64, contactPoint domain.EmbeddedContactPoint) (domain.EmbeddedContactPoint, error)
	UpdateContactPoint(orgID int64, contactPoint domain.EmbeddedContactPoint) error
	DeleteContactPoint(orgID int64, uid string) error
}

func (srv *ProvisioningSrv) RouteGetContactpoints(c *models.ReqContext) response.Response {
	cps, err := srv.contactpointService.GetContactPoints(c.OrgId)
	if err != nil {
		return ErrResp(http.StatusInternalServerError, err, "")
	}
	return response.JSON(http.StatusOK, cps)
}

func (srv *ProvisioningSrv) RoutePostContactpoint(c *models.ReqContext, cp apimodels.Contactpoint) response.Response {
	contactPoint, err := srv.contactpointService.CreateContactPoint(c.OrgId, domain.EmbeddedContactPoint(cp))
	if err != nil {
		return ErrResp(http.StatusInternalServerError, err, "")
	}
	return response.JSON(http.StatusOK, contactPoint)
}

func (srv *ProvisioningSrv) RoutePutContactpoints(c *models.ReqContext, cp apimodels.Contactpoint) response.Response {
	err := srv.contactpointService.UpdateContactPoint(c.OrgId, domain.EmbeddedContactPoint(cp))
	if err != nil {
		return ErrResp(http.StatusInternalServerError, err, "")
	}
	return response.JSON(http.StatusOK, "")
}

func (srv *ProvisioningSrv) RouteDeleteContactpoint(c *models.ReqContext) response.Response {
	cpID := web.Params(c.Req)[":ID"]
	err := srv.contactpointService.DeleteContactPoint(c.OrgId, cpID)
	if err != nil {
		return ErrResp(http.StatusInternalServerError, err, "")
	}
	return response.JSON(http.StatusOK, "")
}
