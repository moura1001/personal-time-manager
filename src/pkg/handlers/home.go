package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/moura1001/time-manager/src/pkg/data"
	"github.com/moura1001/time-manager/src/pkg/util"
)

func HandleGetHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", getPageData())
}

func HandleIncreaseTiming(ctx *gin.Context) {
	data.AddTiming(data.TIMING_TYPE_INCREASE)

	ctx.HTML(http.StatusOK, "timing/list", getPageData())
}

func HandleConsumeTiming(ctx *gin.Context) {
	data.AddTiming(data.TIMING_TYPE_CONSUME)

	ctx.HTML(http.StatusOK, "timing/list", getPageData())
}

func HandleClearTimings(ctx *gin.Context) {
	data.ClearTimings()

	ctx.HTML(http.StatusOK, "index", getPageData())
}

func HandleDeleteTiming(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.String(400, "id was not an integer")
	}

	deleted := data.DeleteTiming(id)
	if !deleted {
		ctx.String(404, "id not found")
	}

	ctx.HTML(http.StatusOK, "timing/timer", util.Map{
		"oob": true,
	})
	ctx.HTML(http.StatusOK, "timing/total", util.Map{
		"totalTiming": data.GetFormattedTotalTiming(),
		"oob":         true,
	})
}

func getPageData() util.Map {
	return util.Map{
		"timings":     data.GetTimings(),
		"totalTiming": data.GetFormattedTotalTiming(),
		"oob":         false,
	}
}
