package models

type TKContactModelIf interface {
	AddContact(pubKey string, contact TKContact) (ok bool)
	GetContact(pubKey string) (ok bool, contact TKContact)
	AddItem(pubKey string, item TKContactItem) (ok bool)
	RemoveItem(pubKey string, itemPubKey string) (ok bool)
	EditItem(pubKey string, person TKContactItem) (ok bool)
	SynContact(pubKey string, listItemPub []string) []TKContactItem
	GetLastConversation(pubKey string, numCon int32) []TKContactItem
}
