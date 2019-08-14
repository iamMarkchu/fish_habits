package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["fish_habits/controllers:CategoryController"] = append(beego.GlobalControllerRouter["fish_habits/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Index",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fish_habits/controllers:CategoryController"] = append(beego.GlobalControllerRouter["fish_habits/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Store",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
