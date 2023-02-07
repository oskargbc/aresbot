package bot_api

import (
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/global"
	"aresbot/internal/aio/models"
	l "aresbot/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Api struct {
}

func NewApi() Api {
	return Api{}
}

func (s *Api) Up() {

	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = l.InfoLogger.Writer()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to AresBot"})
	})

	r.GET("/quicktask", func(c *gin.Context) {
		if constants.WithQuickTask {
			if c.Request.URL.Query().Has("product") && c.Request.URL.Query().Has("store") {
				product := c.Request.URL.Query().Get("product")
				store := c.Request.URL.Query().Get("store")
				size := ""
				mode := ""
				region := ""
				sku := ""
				payment := ""

				if c.Request.URL.Query().Has("size") {
					size = c.Request.URL.Query().Get("size")
				}
				if strings.Contains(product, "otto.de") {
					size = "x"
				}

				if c.Request.URL.Query().Has("mode") {
					mode = c.Request.URL.Query().Get("mode")
				}

				if c.Request.URL.Query().Has("region") {
					region = c.Request.URL.Query().Get("region")
				}

				if c.Request.URL.Query().Has("sku") {
					sku = c.Request.URL.Query().Get("sku")
				}

				if c.Request.URL.Query().Has("payment") {
					payment = c.Request.URL.Query().Get("payment")
				}
				payment = "PAYPAL"

				go createAndRunTask(product, store, size, mode, region, sku, payment)
				c.JSON(http.StatusOK, gin.H{"message": "quicktask created"})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": "missing querys"})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "quicktask server offline"})
		}
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Sorry, Page not found"})
	})

	r.Run("127.0.0.1:11914")
}

func createAndRunTask(product string, store string, size string, mode string, region string, sku string, payment string) {
	task := models.Task{
		ProfileName: global.SettingQuicktaskProfileName,
		ProductURL:  product,
		Size:        size,
		UseProxy:    false,
		Mode:        mode,
		Aco:         false,
		Region:      region,
		Store:       store,
		Keywords:    "",
		Sku:         sku,
		Payment:     payment,
		Profile:     global.SettingQuicktaskProfile,
		Id:          0,
	}

	tasks := []models.Task{}

	for i := 0; i < global.SettingQuicktaskAmount; i++ {
		tasks = append(tasks, task)
	}

	global.ShopRunner.Tasks = tasks

	global.ShopRunner.Run()
}
