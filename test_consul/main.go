package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fwhezfwhez/errorx"
	"io/ioutil"
	"net/http"
	"time"
)

var c = http.Client{
	Timeout: 30 * time.Second,
}

func main() {
	// // 可以重复运行
	// RegisterService()

	// 获取所有服务记录
	GetServices()

	// 获取指定id记录
	//GetServiceById("4")
}

type Service struct {
	Name    string            // service name shared in many server instances.
	ID      string            // should be unique, default Name value.
	Tags    []string          // used for filter
	Address string            // service host e.g 127.0.0.1
	Meta    map[string]string //service related map group
	Port    int               // service run on the port
	Kind    string            // The kind of service. Defaults to "" which is a typical Consul service. This value may also be "connect-proxy" for services that are Connect-capable proxies representing another service.
	Check   map[string]interface{}
}

func RegisterService() {
	service := Service{
		Name: "consul_hello_4",
		ID:   "4",
		Tags: []string{
			"hello",
		},
		Address: "127.0.0.1",
		Port:    7612,
		Meta: map[string]string{
			"protocol": "http",
		},
		Check: map[string]interface{}{
			"DeregisterCriticalServiceAfter": "90m",
			"HTTP":                           "http://localhost:7612/ping/",
			"Interval":                       "10s",
		},
	}
	url := "http://localhost:8500/v1/agent/service/register"
	buf, e := json.Marshal(service)
	if e != nil {
		fmt.Println(errorx.Wrap(e).Error())
		return
	}
	req, e := http.NewRequest("PUT", url, bytes.NewReader(buf))
	if e != nil {
		fmt.Println(errorx.Wrap(e).Error())
		return
	}
	resp, e := c.Do(req)
	if e != nil {
		fmt.Println(errorx.Wrap(e).Error())
		return
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		buf, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			fmt.Println(errorx.Wrap(e).Error())
			return
		}
		if buf!=nil && len(buf)!=0 {
			var buff = bytes.NewBuffer(nil)
			e = json.Indent(buff, buf, "  ", "  ")
			if e != nil {
				fmt.Println(errorx.Wrap(e).Error())
				return
			}
			fmt.Println(string(buff.String()))
		}
		return
	}
	fmt.Println("nothing returned")
}

// List all services in consul
func GetServices() {
	url := "http://127.0.0.1:8500/v1/agent/services"
	resp, e := http.Get(url)
	if e != nil {
		panic(e.Error())
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		buf, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e.Error())
		}
		var buff = bytes.NewBuffer(nil)
		json.Indent(buff, buf, "  ", "  ")
		fmt.Println(string(buff.String()))
		return
	}
	fmt.Println("nothing returned")
}

// Get a specific service by service_id
func GetServiceById(serviceId string) {
	url := "http://127.0.0.1:8500/v1/agent/service/" + serviceId
	resp, e := http.Get(url)
	if e != nil {
		panic(e.Error())
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
		buf, e := ioutil.ReadAll(resp.Body)
		if e != nil {
			panic(e.Error())
		}
		if len(buf)!=0 {
			var buff = bytes.NewBuffer(nil)
			json.Indent(buff, buf, "  ", "  ")
			fmt.Println(string(buff.String()))
			return
		}
	}
	fmt.Println("nothing returned")
}


