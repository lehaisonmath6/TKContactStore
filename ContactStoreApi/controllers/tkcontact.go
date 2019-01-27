package controllers

import (
	"OpenStars/TrustKeys/ContactStoreApi/common/util"
	models "OpenStars/TrustKeys/ContactStoreApi/models"
	"fmt"
	"strconv"

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
// @Success 200	{object} models.ErrorResult
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

	var codeRs models.ErrorResult
	if !util.CheckSignature(pubKey, mess, sig) && enableSig {
		fmt.Println("Error sig message")
		codeRs.Result = models.CODE_ErrorSig
		o.Data["json"] = codeRs
	} else if tkContactModel != nil {
		ob := models.TKContact{PubKey: pubKey}
		ok := tkContactModel.AddContact(ob.PubKey, ob)
		if ok == nil {
			codeRs.Result = models.CODE_Ok
			o.Data["json"] = codeRs
		} else {
			codeRs.Result, _ = strconv.Atoi(ok.Error())
			o.Data["json"] = codeRs
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}

// @Title Get all contact
// @Description find object by objectid
// @Param	pubKey		query	string	true	"Public key of a user"
// @Success 200	{object} models.TKContactResult
// @Failure 403 :pubKey not found
// @router /GetContact/ [get]
func (o *TKContactController) GetContact() {

	var pubKey string
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	fmt.Println("Controller get all contact by pubkey ", pubKey)
	if pubKey != "" && tkContactModel != nil {
		ok, ob := tkContactModel.GetContact(pubKey)
		if ok == nil {
			rs := models.TKContactResult{
				ErrorCode: models.CODE_Ok,
				Data:      ob,
			}
			o.Data["json"] = rs
		} else {
			c, _ := strconv.Atoi(ok.Error())
			rs := models.TKContactResult{
				ErrorCode: c,
				Data:      models.TKContact{},
			}
			o.Data["json"] = rs
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
	stringintput := string(o.Ctx.Input.RequestBody)
	fmt.Println(stringintput)
	err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	var codeRs models.ErrorResult
	if err != nil {
		codeRs.Result = models.CODE_UnMarshalFailed
		o.Data["json"] = codeRs
	} else {
		messByte, _ := json.Marshal(ob)
		fmt.Println(messByte)
		mess = string(messByte)
		fmt.Println("pubKey : ", pubKey, "Message  : ", mess, "Sig : ", sig)
		if !util.CheckSignature(pubKey, mess, sig) && enableSig {
			fmt.Println("Error sig message")
			codeRs.Result = models.CODE_ErrorSig
			o.Data["json"] = codeRs
		} else {
			if tkContactModel != nil {
				ok := tkContactModel.AddItem(pubKey, ob)
				if ok == nil {
					codeRs.Result = models.CODE_Ok
					o.Data["json"] = codeRs
				} else {
					codeRs.Result, _ = strconv.Atoi(ok.Error())
					o.Data["json"] = codeRs
				}
			} else {
				codeRs.Result = models.CODE_ServerNull
				o.Data["json"] = codeRs
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
// @Success 200	{object} models.ErrorResult
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

	var codeRs models.ErrorResult
	if !util.CheckSignature(pubKey, pubKeyItem, sig) && enableSig {
		fmt.Println("Error sig message")
		codeRs.Result = models.CODE_ErrorSig
		o.Data["json"] = codeRs
	} else if tkContactModel != nil {
		ok := tkContactModel.RemoveItem(pubKey, pubKeyItem)
		if ok == nil {
			codeRs.Result = models.CODE_Ok
			o.Data["json"] = codeRs
		} else {
			codeRs.Result, _ = strconv.Atoi(ok.Error())
			o.Data["json"] = codeRs
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

	var codeRs models.ErrorResult

	if err != nil {
		codeRs.Result = models.CODE_UnMarshalFailed
		o.Data["json"] = codeRs
	} else {
		messByte, _ := json.Marshal(ob)
		mess = string(messByte)
		fmt.Println("pubKey : ", pubKey, "Message  : ", mess, "Sig : ", mess)
		if !util.CheckSignature(pubKey, mess, sig) && enableSig {
			fmt.Println("Error sig message")
			codeRs.Result = models.CODE_ErrorSig
			o.Data["json"] = codeRs
		} else {
			if tkContactModel != nil {
				ok := tkContactModel.EditItem(pubKey, ob)
				if ok == nil {
					codeRs.Result = models.CODE_Ok
					o.Data["json"] = codeRs
				} else {
					codeRs.Result, _ = strconv.Atoi(ok.Error())
					o.Data["json"] = codeRs
				}
			} else {
				codeRs.Result = models.CODE_ServerNull
				o.Data["json"] = codeRs
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
// @Success 200	{object} models.TKSynContactResult
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
		rs := models.TKSynContactResult{
			ErroCode: models.CODE_UnMarshalFailed,
		}
		o.Data["json"] = rs
	} else {
		messByte, _ := json.Marshal(ob)
		mess = string(messByte)
		fmt.Println("pubKey : ", pubKey, "Message  : ", mess, "Sig : ", sig)
		if !util.CheckSignature(pubKey, mess, sig) && enableSig {
			fmt.Println("Error sig message")
			rs := models.TKSynContactResult{
				ErroCode: models.CODE_ErrorSig,
			}
			o.Data["json"] = rs
		} else {
			if tkContactModel != nil {
				listTKitem := tkContactModel.SynContact(pubKey, ob.ListPubKey)
				rs := models.TKSynContactResult{
					ErroCode: models.CODE_Ok,
					Data:     listTKitem,
				}
				o.Data["json"] = rs
			} else {
				rs := models.TKSynContactResult{
					ErroCode: models.CODE_ServerNull,
				}
				o.Data["json"] = rs
			}
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}

// @Title Get Safe Contact
// @Description get contact by pubKey and check signature
// @Param	pubKey		query	string	true
// @Param	timeStamp		query	string	true
// @Param	sig		query	string true
// @Success 200	{object} models.TKContactResult
// @Failure 403 Not found
// @router /GetSafeContact/ [get]
func (o *TKContactController) SafeGetContact() {
	var pubKey string
	var timeStamp string
	var sig string
	o.Ctx.Input.Bind(&pubKey, "pubKey")
	o.Ctx.Input.Bind(&timeStamp, "timeStamp")
	o.Ctx.Input.Bind(&sig, "sig")

	fmt.Println("Controller safe get all contact by pubkey ", pubKey)
	var rs models.TKContactResult
	if pubKey != "" && tkContactModel != nil && timeStamp != "" && sig != "" {
		if !util.CheckSignature(pubKey, pubKey+timeStamp, sig) && enableSig {
			rs.ErrorCode = models.CODE_ErrorSig
			o.Data["json"] = rs
		} else {
			ok, ob := tkContactModel.GetContact(pubKey)
			if ok == nil {
				rs.Data = ob
				rs.ErrorCode = models.CODE_Ok
			} else {
				rs.ErrorCode, _ = strconv.Atoi(ok.Error())
			}
			o.Data["json"] = rs
		}
	}
	o.ServeJSON()
	fmt.Println("===============END RESPONSE==============")
}
