package models

import (
	TType "OpenStars/TrustKeys/ContactStore"
	"fmt"
)

type TKContact struct {
	PubKey      string           `json:"pubKey"`
	ListFriends []*TKContactItem `json:"listFriends"`
}

type TKListPubKey struct {
	ListPubKey []string `json:"listPubKey"`
}

func (this TKContact) ToThriftType() TType.TKContact {
	return TType.TKContact{
		PubKey:     TType.TKey(this.PubKey),
		ListFriend: ToListThriftPointTKContactItem(this.ListFriends),
	}
}

func ToTKContactModelType(contact *TType.TKContact) *TKContact {
	tkContactModel := &TKContact{
		PubKey:      string(contact.PubKey),
		ListFriends: ToListPointTKContactItemModel(contact.ListFriend),
	}
	return tkContactModel
}

func init() {
	fmt.Println("Init model TKContact")
}
