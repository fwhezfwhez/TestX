package main
import(
	"encoding/json"
	"fmt"
)
type User struct{
	Name string
	Age int
}
func main(){
	user :=User{"ft",21}
	bytesUser,_:=json.Marshal(user)
	var user2 = User{}
	json.Unmarshal(bytesUser,&user2)
	fmt.Println(user2)

	type Data struct{
		Jsonrpc string `json:"jsonrpc"`
		Result bool `json:"result"`
		Id int `json:"id"`
	}
	data1:=`{"jsonrpc":"2.0","result":true,"id":1}`
	data2,err := json.Marshal(data1)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(data2)
	dataR:=Data{}
	//dataR2 :=""
	json.Unmarshal(data2,&dataR)
	fmt.Println(string(data2))
}
