package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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

}
