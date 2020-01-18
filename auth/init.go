package auth

import (
	"github.com/renosyah/AyoLesPortal/util"
)

var cookieConfig *util.CookieConfig

func Init() {
	cookieConfig = util.NewCookieConfig()
}
