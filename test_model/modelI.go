package test_model

import (
	"encoding/json"
	"github.com/fwhezfwhez/errorx"
)
type Block map[string]interface{}
func (b *Block) Update(key string, value interface{})*Block{
	(*b)[key] = value
	return b
}

func (b *Block) Pop(key string)*Block{
	delete(*b, key)
	return b
}

type SerializerI interface{
	Serialize(dest interface{},f func(b Block)(func(b Block)Block,[]string))([]byte,error)
}

type Serializer struct{
}

func (s Serializer) Serialize(dest interface{},f func(b Block)(func(b Block)Block,[]string))([]byte,error){
	//1. 把结构体转换成map
	var m = make(Block,0)
	buf, er :=json.Marshal(dest)
	if er!=nil {
		return nil, errorx.New(er)
	}
	er = json.Unmarshal(buf,&m)
	if er!=nil {
		return nil, errorx.New(er)
	}

	//2. 把map托管给用户自定义操作,收集需要update或pop后的result，以及待pick的
	blockHandler,picked :=f(m)
	//2.1 收集pick
	var filtPicked Block
	if picked ==nil || len(picked) ==0{
		filtPicked = m
	}else{
		filtPicked = make(Block,0)
		for _,v :=range picked{
			vm,ok := m[v]
			if !ok {
				continue
			}
			filtPicked[v] = vm
		}
	}

	//2.2 收集update和pop并json
	buf, er = json.Marshal(blockHandler(filtPicked))
	if er != nil {
		return nil, errorx.Wrap(er)
	}

	return buf, nil
}
