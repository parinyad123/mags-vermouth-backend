package routes

import (
	"vermouth-backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", apigo)
	router.GET("/data", controllers.GET_tmsystem)

	// dashboard chart
	router.GET("/datadynamic", controllers.GET_dynamic_tmsystem)
	router.GET("/anochart/:satname/:tmname/:freq", controllers.GET_tmgraphTH1)
	// router.GET("/csvdownload/:ana_tb/:ano_tb/:s_utc/:e_utc", controllers.POST_csvdownload)
	router.POST("/csvdownload", controllers.POST_csvdownload)

	// dashboard summary
	router.GET("/anomalyweekly", controllers.GET_anomalyweekly)
	// router.GET("/dailyfilter", controllers.GET_dailyfilter) // not use
	router.GET("/reportdailyfilter", controllers.GET_reportdailyfilter)
	router.GET("/reportalldaily_proviousmonth", controllers.GET_reportalldaily_proviousmonth)
	router.POST("/postreportdaily", controllers.POST_reportdaily)

	// dashboard chart new
	router.GET("/THEOSchartfilter", controllers.GET_THEOSchartfilters)
	router.GET("/THEOS_chartanomaly/:satname/:tmname/:freq", controllers.GET_THEOS_chartanomaly)
	// router.POST("THEOSpostDownloadcsv/",controllers.POST_downloadCSV)
	// router.GET("/THEOSgetDownloadcsv/:satname/:tmname/:freq/:analysis_table/:anomaly_table/:start_utc/:end_utc",controllers.GET_THEOSDownloadcsv)
	router.GET("/GET_THEOSDownload_staticscsv/:satname/:tmname/:freq/:analysis_table/:anomaly_table/:start_utc/:end_utc",controllers.GET_THEOSDownload_staticscsv)


	router.NoRoute(NotFound)
}

func apigo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"massege": "API success",
	})
	return
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  400,
		"massege": "Route Not Found",
	})
	return
}
