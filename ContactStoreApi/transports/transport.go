package transports

import (
	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/TrustKeys/ContactStore" //Todo: Fix this
	"fmt"

	"OpenStars/TrustKeys/ContactStore" //Todo: Fix this

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/OpenStars/thriftpool"
)

var (
	mTContactStoreServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (ContactStore.NewTContactStoreServiceClient(c)) }),
		thriftpool.DefaultClose)

	mTContactStoreServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (ContactStore.NewTContactStoreServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift TContactStoreService client ")
}

//GetTContactStoreServiceBinaryClient client by host:port
func GetTContactStoreServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTContactStoreServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTContactStoreServiceCompactClient get compact client by host:port
func GetTContactStoreServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTContactStoreServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
