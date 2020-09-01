package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/robherley/lumberjack/es"
	"github.com/robherley/lumberjack/util"

	esapiv6 "github.com/elastic/go-elasticsearch/v6/esapi"
)

// GetProjectLogs gets 100 most recent logs in a project
// @Summary gets most recent logs in a project
// @Description gets most recent logs in a project
// @Accept  json
// @Produce  json
// @Success 200
// @Error 500 {object} util.RequestError
// @Param index path string true "ElasticSearch Index"
// @Param namespace path string true "Project/Namespace"
// @Router /api/v1/logs/{index}/{namespace} [get]
// @Tags logs
func GetProjectLogs(c *gin.Context) {
	es6, err := es.GetClient()
	if err != nil {
		util.Log.Warnln(err.Error())
		util.RespondWithError(c, util.ErrInternalError)
		return
	}

	reqCtx := c.Request.Context()

	// TODO: alternate way of creating body of es request
	// body := `
	// {
	// 	"query": {
	// 		"match": {
	// 			"kubernetes.namespace_name": "rob"
	// 		}
	// 	}
	// }
	// `
	// var bodyMap map[string]interface{}
	// err = json.Unmarshal([]byte(body), &bodyMap)
	// if err != nil {
	// 	util.Log.Warnln(err.Error())
	// 	util.RespondWithError(c, util.ErrInternalError)
	// 	return
	// }

	opts := &es.SearchOpts{
		Params: []func(*esapiv6.SearchRequest){
			es6.Search.WithContext(reqCtx),
			es6.Search.WithIndex(c.Param("index")),
			es6.Search.WithSize(100),
		},
		Body: &map[string]interface{}{
			"sort": map[string]interface{}{"@timestamp": map[string]interface{}{"order": "desc"}},
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"kubernetes.namespace_name": c.Param("namespace"),
				},
			},
		},
	}

	res, err := es.Search(es6, opts)

	if err != nil {
		util.Log.Warnln(err.Error())
		util.RespondWithError(c, util.ErrInternalError)
		return
	}

	c.JSON(200, res)
}
