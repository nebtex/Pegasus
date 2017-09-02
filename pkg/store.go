package pkg

import (
	"github.com/nebtex/hybrids/golang/hybrids"
	consul "github.com/hashicorp/consul/api"
)

type KV interface {
	Put(k []byte, v []byte) (err error)
	Get(k []byte) (v []byte, err error)
}

type Pegasus struct {
	KV KV
}

//at the moment stack only have a resource
func (p *Pegasus) UpsertResource(application []byte, rid []byte, resource hybrids.TableReader) (result hybrids.TableReader, err error) {

	if resource == nil {
		return
	}

	err = p.KV.Put(append(application, rid...), resource.DeepCopy())
	if err != nil {
		return
	}

	//return the new state of the resource
	result = resource
	return

}

//GetResource ...
func (p *Pegasus) GetResource(application []byte, rid []byte) (result hybrids.TableReader, err error) {
	var data []byte
	data, err = p.KV.Get(append(application, rid...))
	if err != nil {
		return
	}

	result = hybrids.TableReaderFromBinary(data, 0)
	return

}

type ConsulKV struct {
	client *consul.Client
	prefix string
}

func (ckv *ConsulKV) Put(k []byte, v []byte) (err error) {
	kvp := &consul.KVPair{}
	kvp.Key = ckv.prefix + "/" + string(k)
	kvp.Value = v
	_, err = ckv.client.KV().Put(kvp, nil)
	return
}

func (ckv *ConsulKV) Get(k []byte) (v []byte, err error) {
	var kvp *consul.KVPair

	kvp, _, err = ckv.client.KV().Get(string(k), nil)
	if err != nil {
		return
	}
	v = kvp.Value
	return
}

func NewConsulKV() (kv *ConsulKV, err error) {
	var client *consul.Client
	kv = &ConsulKV{}
	client, err = consul.NewClient(consul.DefaultConfig())
	if err != nil {
		return
	}
	kv.client = client
	return
}
