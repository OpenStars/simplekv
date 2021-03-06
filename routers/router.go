// @APIVersion 1.0.0
// @Title Simple KV API
// @Description beego has a very cool tools to autogenerate documents for your API and SimpleKV used it
// @Contact thanhnt@123xe.vn
// @TermsOfServiceUrl https://123xe.vn/cms/dieukhoan
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"simplekv/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/object",
		// 	beego.NSInclude(
		// 		&controllers.ObjectController{},
		// 	),
		// ),
		// beego.NSNamespace("/user",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
		beego.NSNamespace("/simplekv",
			beego.NSInclude(
				&controllers.SimpleKVController{},
			),
		),
		beego.NSNamespace("/bigsetkv",
			beego.NSInclude(
				&controllers.BigsetKVController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
