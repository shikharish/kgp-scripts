package main

import (
	erp "github.com/metakgp/iitkgp-erp-login-go"

	"github.com/pkg/browser"
)

func main() {
	_, ssoToken := erp.ERPSession()

	browser.OpenURL(erp.HOMEPAGE_URL + "?" + ssoToken)
}
