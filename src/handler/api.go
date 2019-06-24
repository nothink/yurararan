// api handler パッケージ
package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mailgun/mailgun-go"

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
	go sendUpdateMail(news)
	return c.JSON(http.StatusAccepted, news)
}

// 更新があった時にメールを投げる
func sendUpdateMail(s []interface{}) {
	mg, err := mailgun.NewMailgunFromEnv()
	if err != nil {
		log.Fatal("NewMailgunFromEnv failed - ", err)
	}

	html := "<body>\n<div>\n"

	for _, key := range s {
		keystr := key.(string)
		html = html + fmt.Sprintf("<a href=\"https://%s\">%s</a><br />\n", keystr, keystr)
		pos := strings.LastIndex(keystr, ".")
		ext := keystr[pos + 1:]
		if ext == "jpg" || ext == "png" || ext == "gif" {
			html = html + fmt.Sprintf("<img src=\"https://%s\" style=\"max-width: 320px;\" /><br />\n", keystr)
		} else if ext == "mp3" || ext == "wav" || ext == "m4a" || ext == "ogg" {
			html = html + fmt.Sprintf("<audio src=\"https://%s\" style=\"max-width: 320px;\" /><br />\n", keystr)
		} else if ext == "mp4" {
			html = html + fmt.Sprintf("<video src=\"https://%s\" style=\"max-width: 320px;\" /><br />\n", keystr)
		}
	}
	html = html + "</div>\n</body>\n"


	msg := mg.NewMessage(
		/* From */ "GRANDPA <grandpa@mail.fukita.org>",
		/* Subject */ "VERENAV updated.",
		/* Body */ "",
		/* To */ "nothink@nothink.jp",
	)
	msg.SetHtml("<html>HTML version of the body</html>")
}
