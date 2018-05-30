package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"],
		beego.ControllerComments{
			Method: "MultiGet",
			Router: `/multiget`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"],
		beego.ControllerComments{
			Method: "PutItem",
			Router: `/putitem`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"] = append(beego.GlobalControllerRouter["simplekv/controllers:BigsetKVController"],
		beego.ControllerComments{
			Method: "Remove",
			Router: `/remove`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:ObjectController"] = append(beego.GlobalControllerRouter["simplekv/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:ObjectController"] = append(beego.GlobalControllerRouter["simplekv/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:ObjectController"] = append(beego.GlobalControllerRouter["simplekv/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:ObjectController"] = append(beego.GlobalControllerRouter["simplekv/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:ObjectController"] = append(beego.GlobalControllerRouter["simplekv/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:SimpleKVController"] = append(beego.GlobalControllerRouter["simplekv/controllers:SimpleKVController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:SimpleKVController"] = append(beego.GlobalControllerRouter["simplekv/controllers:SimpleKVController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:UserController"] = append(beego.GlobalControllerRouter["simplekv/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:UserController"] = append(beego.GlobalControllerRouter["simplekv/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:UserController"] = append(beego.GlobalControllerRouter["simplekv/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:UserController"] = append(beego.GlobalControllerRouter["simplekv/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:UserController"] = append(beego.GlobalControllerRouter["simplekv/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:UserController"] = append(beego.GlobalControllerRouter["simplekv/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["simplekv/controllers:UserController"] = append(beego.GlobalControllerRouter["simplekv/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
