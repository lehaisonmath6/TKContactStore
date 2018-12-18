package models

import (
	TType "OpenStars/TrustKeys/ContactStore"
	"OpenStars/TrustKeys/ContactStoreApi/transports"
	"context"
	"errors"
	"fmt"

	"github.com/OpenStars/GoEndpointManager"
	thriftpool "github.com/OpenStars/thriftpool"
)

type TKContactModel struct {
	DataBSHost         string
	DataBSPort         string
	EnableSig          int
	ContactStoreConfig ThriftService
}

func (o *TKContactModel) getDataBSClient() *thriftpool.ThriftSocketClient {
	// mHost, mPort, err := o.ContactStoreConfig.EndpoinMgr.GetEndpoint(o.ContactStoreConfig.ServiceID)
	// fmt.Println("mHost ", mHost, "mPort", mPort, "err", err)
	if true {
		return transports.GetTContactStoreServiceCompactClient(o.DataBSHost, o.DataBSPort)
	}
	return nil
	// return transports.GetTContactStoreServiceCompactClient(mHost, mPort)
}

func (o *TKContactModel) GetContact(pubKey string) (err error, contact TKContact) {
	fmt.Println("Get contact model by Pubkey : ", pubKey)
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
				return nil, contact
			}
			if aRes.ErrorCode == TType.TErrorCode_ENotFound {
				return errors.New("Pubkey not found !"), TKContact{}
			}
			if aRes.ErrorCode == TType.TErrorCode_EUnknown {
				return errors.New("Error Unknown !"), TKContact{}
			}
		}
	}
	return errors.New("Backend error !"), TKContact{}
}

func (o *TKContactModel) AddContact(pubKey string, contact TKContact) error {
	fmt.Println("Put contact model by PubKey : ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		var contactThrift = contact.ToThriftType()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).PutData(context.Background(), TType.TKey(pubKey), &contactThrift)
		if Err == nil {
			switch aRes {
			case TType.TErrorCode_EGood:
				return nil
			case TType.TErrorCode_EUnknown:
				return errors.New("Error Unknown !")
			}
		}
	}
	return errors.New("Backend error !")
}

func (o *TKContactModel) AddItem(pubKey string, item TKContactItem) error {
	fmt.Println("Add person model to contact PubKey : ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		var item = item.ToThriftType()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).AddItem(context.Background(), TType.TKey(pubKey), &item)
		if Err == nil {
			switch aRes {
			case TType.TErrorCode_EGood:
				return nil
			case TType.TErrorCode_EDataExisted:
				return errors.New("Error data existed !")
			case TType.TErrorCode_ENotFound:
				return errors.New("Pubkey not found !")
			case TType.TErrorCode_EUnknown:
				return errors.New("Error Unknown !")
			}
		}
	}
	return errors.New("Backend error !")
}

func (o *TKContactModel) RemoveItem(pubKey string, itemPubKey string) error {
	fmt.Println("Remove person model to contact PubKey : ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).DeleteItem(context.Background(), TType.TKey(pubKey), itemPubKey)
		if Err == nil {
			switch aRes {
			case TType.TErrorCode_ENotFound:
				return errors.New("Pubkey not found !")
			case TType.TErrorCode_EGood:
				return nil
			case TType.TErrorCode_EUnknown:
				return errors.New("Error Unknown !")
			}
		}
	}
	return errors.New("Backend error !")
}

func (o *TKContactModel) EditItem(pubKey string, item TKContactItem) error {
	fmt.Println("Edit person model to contact key ", pubKey)
	aClient := o.getDataBSClient()
	if aClient != nil {
		defer aClient.BackToPool()
		var typeItem = item.ToThriftType()
		aRes, Err := aClient.Client.(*TType.TContactStoreServiceClient).EditItem(context.Background(), TType.TKey(pubKey), &typeItem)
		if Err == nil {
			switch aRes {
			case TType.TErrorCode_EGood:
				return nil
			case TType.TErrorCode_ENotFound:
				return errors.New("Pubkey not found !")
			case TType.TErrorCode_EUnknown:
				return errors.New("Error Unknown !")
			}
		}
	}
	return errors.New("Backend error !")
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

func NewTKContactModel(host, port string, serviceid string, etcdtEndpoint string) TKContactModelIf {

	endpoint := []string{etcdtEndpoint}
	return &TKContactModel{
		DataBSHost: host,
		DataBSPort: port,
		ContactStoreConfig: ThriftService{
			ServiceID:  serviceid,
			Protocol:   "binary",
			EndpoinMgr: GoEndpointManager.NewEtcdEndpointManager(endpoint),
		},
	}
}
