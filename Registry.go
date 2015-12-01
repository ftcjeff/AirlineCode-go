package ftcpops

import (
	"log"
	"time"

	"github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/net/context"
	"github.com/coreos/etcd/client"
)

func RegisterSelf(serviceName, address, port, endpoint string) error {
	myUri := "http://" + address + port + "/" + endpoint

	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)

	key := "/ftcpops/registry/" + serviceName

	log.Printf("Registering %q", serviceName)
	_, err = kapi.Set(context.Background(), key, myUri, nil)

	return err
}

func GetServiceURI(serviceName string) string {
	cfg := client.Config{
		Endpoints:               []string{"http://127.0.0.1:2379"},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)

	key := "/ftcpops/registry/" + serviceName

	resp, err := kapi.Get(context.Background(), key, nil)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Node.Value
}
