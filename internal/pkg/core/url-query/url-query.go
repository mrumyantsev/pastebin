package urlquery

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrumyantsev/pastebin/internal/pkg/config"
	"github.com/mrumyantsev/pastebin/internal/pkg/core"
)

func Pagination(ctx *gin.Context, cfg *config.Config) core.Pagination {
	qParams := ctx.Request.URL.Query()

	var pg core.Pagination
	var param string

	if param = qParams.Get("page"); param != "" {
		pg.Page, _ = strconv.Atoi(param)
	}

	if param = qParams.Get("limit"); param != "" {
		pg.Limit, _ = strconv.Atoi(param)
	}

	if pg.Limit <= 0 {
		pg.Limit = cfg.ItemsOnPage
	}
	if pg.Limit > cfg.MaxItemsOnPage {
		pg.Limit = cfg.MaxItemsOnPage
	}

	pg.Page = pg.Page*pg.Limit - pg.Limit

	return pg
}
