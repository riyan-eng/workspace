package infrastructure

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var i18Bundle *i18n.Bundle

func NewLocalizer() {
	i18Bundle = i18n.NewBundle(language.Indonesian)
	i18Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	if _, err := i18Bundle.LoadMessageFile("./locale/en.toml"); err != nil {
		fmt.Printf("localize: %v", err.Error())
		os.Exit(1)
	}
	if _, err := i18Bundle.LoadMessageFile("./locale/id.toml"); err != nil {
		fmt.Printf("localize: %v", err.Error())
		os.Exit(1)
	}
	fmt.Println("localize: load successfully")
}

var i18Localizer *i18n.Localizer

func LocalizerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Accept-Language")
		query := c.Query("lang")
		i18Localizer = i18n.NewLocalizer(i18Bundle, string(header), query)
		c.Next()
	}
}

func Localize(params any) string {
	switch p := params.(type) {
	case string:
		return i18Localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: p,
		})
	default:
		return i18Localizer.MustLocalize(p.(*i18n.LocalizeConfig))
	}
}
