package healthhandler

import (
	"net/http"

	"github.com/MarianoArias/ApiGo/pkg/elastic-search"
	"github.com/MarianoArias/ApiGo/pkg/entity-manager"
	"github.com/gin-gonic/gin"
)

type Response map[string]interface{}

const (
	UP   = "UP"
	DOWN = "DOWN"
)

func HealthHandler(c *gin.Context) {
	generalStatus, elasticsearchStatus, mysqlStatus := UP, UP, UP

	if err1 := elasticsearch.Ping(); err1 != nil {
		generalStatus = DOWN
		elasticsearchStatus = DOWN
	}

	if err2 := entitymanager.Ping(); err2 != nil {
		generalStatus = DOWN
		mysqlStatus = DOWN
	}

	a := Response{
		"status": Response{
			"code": generalStatus,
		},
		"details": Response{
			"elasticsearch": Response{
				"status": Response{
					"code": elasticsearchStatus,
				},
			},
			"mysql": Response{
				"status": Response{
					"code": mysqlStatus,
				},
			},
		},
	}

	c.JSON(http.StatusOK, a)
}
