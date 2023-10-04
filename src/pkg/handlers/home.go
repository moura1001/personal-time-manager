package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura1001/time-manager/src/pkg/util"
)

func HandleGetHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", util.Map{})
}
