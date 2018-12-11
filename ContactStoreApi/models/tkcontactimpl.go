package models

import (
	TType "OpenStars/TrustKeys/ContactStore"
	"OpenStars/TrustKeys/ContactStoreApi/transports"
	"context"
	"fmt"

	thriftpool "github.com/OpenStars/thriftpool"
)

type TKContactModel struct {
	DataBSHost string
	DataBSPort string
}

func (o *TKContactModel) getDataBSClient() *thriftpool.ThriftSocketClient {
	return transports.GetTContactStoreServiceCompactClient(o.DataBSHost, o.DataBSPort)
}

func (o *TKContactModel) GetContact(pubKey string) (ok bool, contact TKContact) {
	fmt.Println("Get contact model by pubkey ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		var contact = TKContact{
			PubKey: pubKey,
		}
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).GetData(context.Background(), TType.TKey(pubKey))
		if aRes != nil && Err == nil {
			if aRes.ErrorCode == TType.TErrorCode_EGood {
				contact.ListFriends = ToListPointTKContactItemModel(aRes.Data)
				return true, contact
			}
		}
	}
	return false, TKContact{}
}

func (o *TKContactModel) AddContact(pubKey string, contact TKContact) (ok bool) {
	fmt.Println("Put contact model by pubKey ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		var contactThrift = contact.ToThriftType()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).PutData(context.Background(), TType.TKey(pubKey), &contactThrift)
		if Err == nil && aRes == TType.TErrorCode_EGood {
			return true
		}
	}
	return false
}

func (o *TKContactModel) AddItem(pubKey string, item TKContactItem) (ok bool) {
	fmt.Println("Add person model to contact key ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		var item = item.ToThriftType()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).AddItem(context.Background(), TType.TKey(pubKey), &item)
		if Err == nil && aRes == TType.TErrorCode_EGood {
			return true
		}
	}
	return false
}

func (o *TKContactModel) RemoveItem(pubKey string, itemPubKey string) (ok bool) {
	fmt.Println("Remove person model to contact key ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).DeleteItem(context.Background(), TType.TKey(pubKey), itemPubKey)
		if Err == nil && aRes == TType.TErrorCode_EGood {
			return true
		}
	}
	return false
}

func (o *TKContactModel) EditItem(pubKey string, item TKContactItem) (ok bool) {
	fmt.Println("Edit person model to contact key ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		var typeItem = item.ToThriftType()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).EditItem(context.Background(), TType.TKey(pubKey), &typeItem)
		if Err == nil && aRes == TType.TErrorCode_EGood {
			return true
		}
	}
	return false
}

func (o *TKContactModel) GetLastConversation(pubKey string, numCon int32) []TKContactItem {
	fmt.Println("Get last convertsation  model to contact key ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).GetLastConversation(context.Background(), TType.TKey(pubKey), numCon)
		if Err == nil && aRes != nil {

			var listPointPerson = ToListPointTKContactItemModel(aRes.Data)
			var listPerson = make([]TKContactItem, len(listPointPerson))
			for i, v := range listPointPerson {
				listPerson[i] = *v
			}
			return listPerson
		}
	}
	return nil
}

func (o *TKContactModel) SynContact(pubKey string, listItemPub []string) []TKContactItem {
	fmt.Println("SynContact  model to contact key ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).SynContact(context.Background(), TType.TKey(pubKey), listItemPub)
		if Err == nil && aRes != nil {
			var listPointPerson = ToListPointTKContactItemModel(aRes.Data)
			var listPerson = make([]TKContactItem, len(listPointPerson))
			for i, v := range listPointPerson {
				listPerson[i] = *v
			}
			return listPerson
		}
	}
	return nil
}

func NewTKContactModel(host, port string) TKContactModelIf {
	return &TKContactModel{
		DataBSHost: host,
		DataBSPort: port,
	}
}
