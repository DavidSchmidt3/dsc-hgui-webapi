/*
 * Hospital guidance API
 *
 * Hospital guidance API
 *
 * API version: 1.0.0
 * Contact: xschmidtd@stuba.sk
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package hgui_webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HospitalGuidanceListAPI interface {

	// internal registration of api routes
	addRoutes(routerGroup *gin.RouterGroup)

	// CreateHospitalGuidance - Create hospital guidance
	CreateHospitalGuidance(ctx *gin.Context)

	// DeleteHospitalGuidance - Delete hospital guidance
	DeleteHospitalGuidance(ctx *gin.Context)

	// GetHospitalGuidance - Get hospital guidance
	GetHospitalGuidance(ctx *gin.Context)

	// GetHospitalGuidanceList - Get hospital guidance list
	GetHospitalGuidanceList(ctx *gin.Context)

	// UpdateHospitalGuidance - Update hospital guidance
	UpdateHospitalGuidance(ctx *gin.Context)
}

// partial implementation of HospitalGuidanceListAPI - all functions must be implemented in add on files
type implHospitalGuidanceListAPI struct {
}

func newHospitalGuidanceListAPI() HospitalGuidanceListAPI {
	return &implHospitalGuidanceListAPI{}
}

func (this *implHospitalGuidanceListAPI) addRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Handle(http.MethodPost, "/guidance-list/:ambulanceId/:guidanceId", this.CreateHospitalGuidance)
	routerGroup.Handle(http.MethodDelete, "/guidance-list/:ambulanceId/:guidanceId", this.DeleteHospitalGuidance)
	routerGroup.Handle(http.MethodGet, "/guidance-list/:ambulanceId/:guidanceId", this.GetHospitalGuidance)
	routerGroup.Handle(http.MethodGet, "/guidance-list/:ambulanceId", this.GetHospitalGuidanceList)
	routerGroup.Handle(http.MethodPut, "/guidance-list/:ambulanceId/:guidanceId", this.UpdateHospitalGuidance)
}

// Copy following section to separate file, uncomment, and implement accordingly
// // CreateHospitalGuidance - Create hospital guidance
// func (this *implHospitalGuidanceListAPI) CreateHospitalGuidance(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // DeleteHospitalGuidance - Delete hospital guidance
// func (this *implHospitalGuidanceListAPI) DeleteHospitalGuidance(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetHospitalGuidance - Get hospital guidance
// func (this *implHospitalGuidanceListAPI) GetHospitalGuidance(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // GetHospitalGuidanceList - Get hospital guidance list
// func (this *implHospitalGuidanceListAPI) GetHospitalGuidanceList(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
// // UpdateHospitalGuidance - Update hospital guidance
// func (this *implHospitalGuidanceListAPI) UpdateHospitalGuidance(ctx *gin.Context) {
//  	ctx.AbortWithStatus(http.StatusNotImplemented)
// }
//
