package kgp_scripts

import (
	erp_login "github.com/metakgp/iitkgp-erp-login-go"

	"github.com/pkg/browser"
)

func OpenERP() {
	_, ssoToken := erp_login.ERPSession()

	browser.OpenURL(erp_login.HOMEPAGE_URL + "?" + ssoToken)
}
