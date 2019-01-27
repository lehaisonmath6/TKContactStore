package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"],
        beego.ControllerComments{
            Method: "AddNewContact",
            Router: `/AddNewContact/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"],
        beego.ControllerComments{
            Method: "EditContactItem",
            Router: `/EditContactItem/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"],
        beego.ControllerComments{
            Method: "GetContact",
            Router: `/GetContact/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"],
        beego.ControllerComments{
            Method: "SafeGetContact",
            Router: `/GetSafeContact/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"],
        beego.ControllerComments{
            Method: "PutContactItem",
            Router: `/PutContactItem/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"],
        beego.ControllerComments{
            Method: "RemoveContactItem",
            Router: `/RemoveContactItem/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:TKContactController"],
        beego.ControllerComments{
            Method: "SynContact",
            Router: `/SynContact/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"] = append(beego.GlobalControllerRouter["OpenStars/TrustKeys/ContactStoreApi/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
