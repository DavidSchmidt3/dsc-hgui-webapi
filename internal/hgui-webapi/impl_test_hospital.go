package hgui_webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Copy following section to separate file, uncomment, and implement accordingly
// // Test - Test endpoint
func (this *implTestHospitalAPI) Test(ctx *gin.Context) {
	// return hello world
	ctx.JSON(http.StatusOK, "Hello World")
}
