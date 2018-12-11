package main

import (
	appconfig "OpenStars/TrustKeys/ContactStoreApi/appconfig"
	controllers "OpenStars/TrustKeys/ContactStoreApi/controllers"
	models "OpenStars/TrustKeys/ContactStoreApi/models"
	_ "OpenStars/TrustKeys/ContactStoreApi/routers"
	"fmt"
	"os"
	"strconv"

	"github.com/astaxie/beego"
)

func InitWithBSHostPort(bsHost, bsPort string) {
	controllers.SetContactModel(models.NewTKContactModel(bsHost, bsPort))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	appconfig.InitConfig()
	InitWithBSHostPort(appconfig.BIGSETKV_HOST, strconv.Itoa(appconfig.BIGSETKV_PORT))
	fmt.Println("Storage hort : ", appconfig.BIGSETKV_HOST, " port : ", appconfig.BIGSETKV_PORT)
	os.Setenv("HOST", appconfig.RunningHost)
	os.Setenv("PORT", appconfig.ListenPort)
	beego.Run()
}
