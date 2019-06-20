// api handler パッケージ
package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nothink/yurararan/shelf"
)

// ファイルリソースのPOST時に与える構造体
type Resources struct {
	Urls []string `json:"urls"`
}

// resources api の　GET ハンドラ
func GetResources(c echo.Context) error {
	all := shelf.All()

	return c.JSON(http.StatusOK, all)
}

// resources api の POST ハンドラ
func PostResources(c echo.Context) error {
	post := new(Resources)
	if err := c.Bind(post); err != nil {
		return err
	}

	urls := make([]interface{}, 0)
	for _, url := range post.Urls {
		urls = append(urls, url)
	}

	news := shelf.Append(urls)

	if len(news) == 0 {
		return c.JSON(http.StatusNoContent, nil)
	}

	fmt.Println(news)
	return c.JSON(http.StatusAccepted, news)
}
