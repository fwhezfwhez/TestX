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
}
