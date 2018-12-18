package models

import (
	TType "OpenStars/TrustKeys/ContactStore"
	"fmt"
)

type TKContactItem struct {
	PubKeyHex    string   `json:"pubKeyHex"`
	DisplayName  string   `json:"displayName"`
	Emails       []string `json:"emails"`
	PhoneNumbers []string `json:"phoneNumbers"`
	Gender       int32    `json:"gender"`

	ProfilePicture string `json:"profilePicture"`
	CoverPicture   string `json:"coverPicture"`
	TimeAdd        string `json:"timeAdd"`
	BirthDay       string `json:"birthDay"`
}

func (this TKContactItem) ToThriftType() TType.TKContactItem {
	typeTKPerson := TType.TKContactItem{
		PubKeyHex:      TType.TKey(this.PubKeyHex),
		DisplayName:    this.DisplayName,
		PhoneNumbers:   this.PhoneNumbers,
		Emails:         this.Emails,
		Gender:         this.Gender,
		ProfilePicture: this.ProfilePicture,
		CoverPicture:   this.CoverPicture,
		TimeAdd:        this.TimeAdd,
		BirthDay:       this.BirthDay,
	}
	return typeTKPerson
}

func (this TKContactItem) ToThriftPointType() *TType.TKContactItem {
	typeTKPerson := &TType.TKContactItem{
		PubKeyHex:      TType.TKey(this.PubKeyHex),
		DisplayName:    this.DisplayName,
		PhoneNumbers:   this.PhoneNumbers,
		Emails:         this.Emails,
		Gender:         this.Gender,
		ProfilePicture: this.ProfilePicture,
		CoverPicture:   this.CoverPicture,
		TimeAdd:        this.TimeAdd,
		BirthDay:       this.BirthDay,
	}
	return typeTKPerson
}

func ToListThriftPointTKContactItem(listPerson []*TKContactItem) []*TType.TKContactItem {
	listThriftTKPerson := make([]*TType.TKContactItem, len(listPerson))
	for i, v := range listPerson {
		listThriftTKPerson[i] = v.ToThriftPointType()
	}
	return listThriftTKPerson
}

func ToModelPointType(this *TType.TKContactItem) *TKContactItem {
	modelTKPerson := &TKContactItem{
		PubKeyHex:      string(this.PubKeyHex),
		DisplayName:    this.DisplayName,
		PhoneNumbers:   this.PhoneNumbers,
		Emails:         this.Emails,
		Gender:         this.Gender,
		ProfilePicture: this.ProfilePicture,
		CoverPicture:   this.CoverPicture,
		TimeAdd:        this.TimeAdd,
		BirthDay:       this.BirthDay,
	}
	return modelTKPerson
}

func ToListPointTKContactItemModel(listPerson []*TType.TKContactItem) []*TKContactItem {
	listPersonModel := make([]*TKContactItem, len(listPerson))
	for i, v := range listPerson {
		listPersonModel[i] = ToModelPointType(v)
	}
	return listPersonModel
}
func init() {
	fmt.Println("Init model TKPerson")
}
