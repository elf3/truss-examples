package etcd

import (
	"context"
	"core/lib/help"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/pkg/errors"
)

func NewEtcdRegister(addr []string, opt etcdv3.ClientOptions, serverName, serverAddr string) (*etcdv3.Registrar, error) {
	etcdClient, err := etcdv3.NewClient(context.Background(), addr, opt)
	if err != nil {
		return nil, errors.New("can't connect etcd register srv")
	}
	ip, err := help.GetClientIP()
	if err != nil {
		return nil, errors.New("can't get ip address")
	}
	srvReg := etcdv3.Service{
		// unique key, e.g. "/service/foobar/1.2.3.4:8080"
		Key: fmt.Sprintf("%s/%s", serverName, serverAddr),
		// returned to subscribers, e.g. "http://1.2.3.4:8080"
		Value: ip + serverAddr,
		TTL:   etcdv3.NewTTLOption(3, 10),
	}
	return etcdv3.NewRegistrar(etcdClient, srvReg, log.NewNopLogger()), nil
}
