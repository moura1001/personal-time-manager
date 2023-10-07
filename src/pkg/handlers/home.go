package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura1001/time-manager/src/pkg/data"
	"github.com/moura1001/time-manager/src/pkg/util"
)

func HandleGetHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", getPageData())
}

func HandleIncreaseTiming(ctx *gin.Context) {
	data.AddTiming(data.TIMING_TYPE_INCREASE)

	ctx.HTML(http.StatusOK, "timings", getPageData())
}

func HandleConsumeTiming(ctx *gin.Context) {
	data.AddTiming(data.TIMING_TYPE_CONSUME)

	ctx.HTML(http.StatusOK, "timings", getPageData())
}

func HandleClearTimings(ctx *gin.Context) {
	data.ClearTimings()

	ctx.HTML(http.StatusOK, "index", getPageData())
}

func getPageData() util.Map {
	return util.Map{
		"timings":     data.GetTimings(),
		"totalTiming": data.GetFormattedTotalTiming(),
	}
}
