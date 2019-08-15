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

    beego.GlobalControllerRouter["fish_habits/controllers:CategoryController"] = append(beego.GlobalControllerRouter["fish_habits/controllers:CategoryController"],
        beego.ControllerComments{
            Method: "Remove",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fish_habits/controllers:HabitController"] = append(beego.GlobalControllerRouter["fish_habits/controllers:HabitController"],
        beego.ControllerComments{
            Method: "Store",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fish_habits/controllers:HabitController"] = append(beego.GlobalControllerRouter["fish_habits/controllers:HabitController"],
        beego.ControllerComments{
            Method: "StoreUserHabit",
            Router: `/:id/user/:uid`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fish_habits/controllers:HabitController"] = append(beego.GlobalControllerRouter["fish_habits/controllers:HabitController"],
        beego.ControllerComments{
            Method: "RemoveUserHabit",
            Router: `/:id/user/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["fish_habits/controllers:HabitController"] = append(beego.GlobalControllerRouter["fish_habits/controllers:HabitController"],
        beego.ControllerComments{
            Method: "Sign",
            Router: `/:id/user/:uid/sign`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
