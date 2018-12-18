package controllers

import (
	"OpenStars/TrustKeys/ContactStoreApi/common/util"
	models "OpenStars/TrustKeys/ContactStoreApi/models"
	"fmt"

	"github.com/astaxie/beego"

	"encoding/json"
)

var (
	tkContactModel models.TKContactModelIf
	enableSig      = true
)

type TKContactController struct {
	beego.Controller
}

func SetContactModel(aModel models.TKContactModelIf, enSig int) {
	tkContactModel = aModel
	if enSig > 0 {
		fmt.Println("Enable sign")
		enableSig = true
	} else {
		fmt.Println("Disable sign")
		enableSig = false
	}
}

// @Title Add New Contact
// @Description Add new contact to database
// @Param	pubKey		query	string	true	"Public key of a user"
// @Param	sig		query	string true		"Sign message"
// @Param	mess		query	string	true	"message to sign"
// @Success 200	{string} string
// @Failure 403  Error
// @router /AddNewContact/ [post]
func (o *TKContactController) AddNewContact() {
	fmt.Println("Controller add new contact")
	var (
		pubKey string
		sig    string
		mess   string
	)
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	o.Ctx.Input.Bind(&sig, "sig")
	o.Ctx.Input.Bind(&mess, "mess")

	fmt.Println("pubKey : ", pubKey, "Message  : ", mess, "Sig : ", mess)

	if !util.CheckSignature(pubKey, mess, sig) && enableSig {
		fmt.Println("Error sig message")
		o.Data["json"] = map[string]string{"result": "Error sig"}
	} else if tkContactModel != nil {
		ob := models.TKContact{PubKey: pubKey}
		ok := tkContactModel.AddContact(ob.PubKey, ob)
		if ok == nil {
			o.Data["json"] = map[string]string{"result": "OK"}
		} else {
			o.Data["json"] = map[string]string{"result": ok.Error()}
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}

// @Title Get all contact
// @Description find object by objectid
// @Param	pubKey		query	string	true	"Public key of a user"
// @Success 200	{object} models.TKContact
// @Failure 403 :pubKey not found
// @router /GetContact/ [get]
func (o *TKContactController) GetContact() {

	var pubKey string
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	fmt.Println("Controller get all contact by pubkey ", pubKey)
	if pubKey != "" && tkContactModel != nil {
		ok, ob := tkContactModel.GetContact(pubKey)
		if ok == nil {
			o.Data["json"] = ob
		} else {
			o.Data["json"] = map[string]string{"result": ok.Error()}
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}

// @Title Add Contact Item
// @Description Add contact item to Contact by pubKey
// @Param	pubKey		query	string	true	"Public key of a user"
// @Param	sig		query	string true		"Sign message"
// @Param	body		body	models.TKContactItem true
// @Success 200	{string} string
// @Failure 403  Error
// @router /PutContactItem/ [post]
func (o *TKContactController) PutContactItem() {
	fmt.Println("Controller put contact item ")
	var (
		pubKey string
		sig    string
		mess   string
		ob     models.TKContactItem
	)
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	o.Ctx.Input.Bind(&sig, "sig")
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	if err != nil {
		o.Data["json"] = map[string]string{"result": "Error"}
	} else {
		messByte, _ := json.Marshal(ob)
		fmt.Println(string(messByte))
		fmt.Println(ob)
		mess = string(messByte)
		fmt.Println("pubKey : ", pubKey, "Message  : ", mess, "Sig : ", sig)
		if !util.CheckSignature(pubKey, mess, sig) && enableSig {
			fmt.Println("Error sig message")
			o.Data["json"] = map[string]string{"result": "Error sig"}
		} else {
			if tkContactModel != nil {
				ok := tkContactModel.AddItem(pubKey, ob)
				if ok == nil {
					o.Data["json"] = map[string]string{"result": "OK"}
				} else {
					o.Data["json"] = map[string]string{"result": ok.Error()}
				}
			} else {
				o.Data["json"] = map[string]string{"result": "Error"}
			}
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}

// @Title Remove item in contact
// @Description Remove item in Contact by pubKey
// @Param	pubKey		query	string	true
// @Param	sig		query	string	true
// @Param	pubKeyItem	query	string	true
// @Success 200	{string} string
// @Failure 403 Error
// @router /RemoveContactItem/ [post]
func (o *TKContactController) RemoveContactItem() {
	fmt.Println("Controller remove item contact")
	var (
		pubKey     string
		sig        string
		pubKeyItem string
	)
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	o.Ctx.Input.Bind(&sig, "sig")
	o.Ctx.Input.Bind(&pubKeyItem, "pubKeyItem")

	fmt.Println("pubKey : ", pubKey, "Message  : ", pubKeyItem, "Sig : ", sig)

	if !util.CheckSignature(pubKey, pubKeyItem, sig) && enableSig {
		fmt.Println("Error sig message")
		o.Data["json"] = map[string]string{"result": "Error sig"}
	} else if tkContactModel != nil {
		ok := tkContactModel.RemoveItem(pubKey, pubKeyItem)
		if ok == nil {
			o.Data["json"] = map[string]string{"result": "OK"}
		} else {
			o.Data["json"] = map[string]string{"result": ok.Error()}
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}

// @Title Edit item contact
// @Description Edit person to Contact by pubKey
// @Param	pubKey		query	string	true
// @Param	sig		query	string	true
// @Param	body		body	models.TKContactItem true
// @Success 200	{string} string
// @Failure 403 :pubKey not found
// @router /EditContactItem/ [post]
func (o *TKContactController) EditContactItem() {
	fmt.Println("Controller edit contact item ")
	var (
		pubKey string
		sig    string
		mess   string
		ob     models.TKContactItem
	)
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	o.Ctx.Input.Bind(&sig, "sig")
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	if err != nil {
		o.Data["json"] = map[string]string{"result": "Error"}
	} else {
		messByte, _ := json.Marshal(ob)
		mess = string(messByte)
		fmt.Println("pubKey : ", pubKey, "Message  : ", mess, "Sig : ", mess)
		if !util.CheckSignature(pubKey, mess, sig) && enableSig {
			fmt.Println("Error sig message")
			o.Data["json"] = map[string]string{"result": "Error sig"}
		} else {
			if tkContactModel != nil {
				ok := tkContactModel.EditItem(pubKey, ob)
				if ok == nil {
					o.Data["json"] = map[string]string{"result": "OK"}
				} else {
					o.Data["json"] = map[string]string{"result": ok.Error()}
				}
			} else {
				o.Data["json"] = map[string]string{"result": "Error"}
			}
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}

// @Title Syn contact
// @Description Synchronization contact by pubKey
// @Param	pubKey		query	string	true
// @Param	sig		query	string	true
// @Param	body		body	models.TKListPubKey true
// @Success 200	{object} []models.TKContactItem
// @Failure 403 Not found
// @router /SynContact/ [post]
func (o *TKContactController) SynContact() {
	var (
		pubKey string
		sig    string
		mess   string
		ob     models.TKListPubKey
	)
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	o.Ctx.Input.Bind(&sig, "sig")

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	fmt.Println("sync contact pubkey : ", pubKey, "sig : ", sig, " data ", ob.ListPubKey)
	fmt.Println(ob)

	if err != nil {
		o.Data["json"] = map[string]string{"result": "Error"}
	} else {
		messByte, _ := json.Marshal(ob)
		mess = string(messByte)
		fmt.Println("pubKey : ", pubKey, "Message  : ", mess, "Sig : ", sig)
		if !util.CheckSignature(pubKey, mess, sig) && enableSig {
			fmt.Println("Error sig message")
			o.Data["json"] = map[string]string{"result": "Error sig"}
		} else {
			if tkContactModel != nil {
				listTKitem := tkContactModel.SynContact(pubKey, ob.ListPubKey)
				o.Data["json"] = listTKitem
			} else {
				o.Data["json"] = map[string]string{"result": "Error"}
			}
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}
