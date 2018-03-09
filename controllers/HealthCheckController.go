package controllers

import (
	"net/http"

	"github.com/trio-purnomo/go-rest-starter/helpers"
)

// InitHealthCheckController is
func InitHealthCheckController() *HealthCheckController {
	healthCheckController := new(HealthCheckController)
	return healthCheckController
}

// HealthCheckController is
type HealthCheckController struct{}

// HealthCheck is
func (h *HealthCheckController) HealthCheck(res http.ResponseWriter, req *http.Request) {
	helpers.Response(res, http.StatusOK, "ok", "test")
	return
}
