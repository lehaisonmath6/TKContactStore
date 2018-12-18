package models

type TKContactModelIf interface {
	AddContact(pubKey string, contact TKContact) (ok error)
	GetContact(pubKey string) (ok error, contact TKContact)
	AddItem(pubKey string, item TKContactItem) (ok error)
	RemoveItem(pubKey string, itemPubKey string) (ok error)
	EditItem(pubKey string, person TKContactItem) (ok error)
	SynContact(pubKey string, listItemPub []string) []TKContactItem
	GetLastConversation(pubKey string, numCon int32) []TKContactItem
}
