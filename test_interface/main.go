package main

type User struct{
}
type UserService interface{
	Insert()
	Delete(id int)
	Update(user User)
	QueryAll()
	QueryById(id int)
	QueryByName(name string)
}

type UserServiceImpl struct{
}
func (us UserServiceImpl) Insert(){
}
func (us UserServiceImpl) Delete(id int){
}
func (us UserServiceImpl) Update(user User){
}
func (us UserServiceImpl) QueryAll(){
}
func (us UserServiceImpl) QueryById(id int){
}
func (us UserServiceImpl) QueryByName(name string){
}

var userService UserService
func init(){
	userService = UserServiceImpl{}
}
func main(){
	user :=User{}
	/*
		m := new(sync.Mutex)
		m.Lock()
		userService.Update(user)
		m.Unlock()
	 */
	userService.Update(user)
}
