package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moura1001/time-manager/src/pkg/data"
	"github.com/moura1001/time-manager/src/pkg/util"
)

func HandleGetHome(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index", util.Map{
		"timings": data.GetTimings(),
	})
}

func HandleIncreaseTiming(ctx *gin.Context) {
	if data.GetTimingType() != "" {
		data.SetLastTimingStop()
	}

	data.AddTiming(data.TIMING_TYPE_INCREASE)

	ctx.HTML(http.StatusOK, "index", util.Map{
		"timings": data.GetTimings(),
	})
}

func HandleConsumeTiming(ctx *gin.Context) {
	if data.GetTimingType() != "" {
		data.SetLastTimingStop()
	}

	data.AddTiming(data.TIMING_TYPE_CONSUME)

	ctx.HTML(http.StatusOK, "index", util.Map{
		"timings": data.GetTimings(),
	})
}

func HandleClearTimings(ctx *gin.Context) {
	data.ClearTimings()

	ctx.HTML(http.StatusOK, "index", util.Map{
		"timings": data.GetTimings(),
	})
}
