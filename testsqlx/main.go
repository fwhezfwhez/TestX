package main

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	_ "github.com/lib/pq"
)
type User struct {
	Name string
	ClassId int
	ClassId2 string
}
func main(){
	db, err := sqlx.Open("postgres", "postgres://postgres:123@localhost:5432/test?sslmode=disable")
	SimplePanic(err)
	fmt.Println(db)
	defer db.Close()


	//2.  crud 增删改查
	//2.1 增加，删除，修改
	if true {
		var crud int = 2 //增改删123
		medium := vo.Medium{
			"m1",
			"o1",
			"c1",
			"s1",
			"123",
			"hello",
			"aD3aAZedr7dDlk",
		}
		switch crud {
		case 1:
			//_, err = db.Exec("insert into public.example_user(name,class_id) values($1,$2)", "Ft","5")
			//SimplePanic(err)
			_,err=db.Exec("insert into medium(mediumName,os,category,subcategory,intro,pkgName) values($1,$2,$3,$4,$5,$6)",
				medium.MediumName,medium.Os,medium.Category,medium.SubCategory,medium.Intro,medium.PkgName)
			if err!=nil {
				SimplePanic(err)
			}
		case 2:
			//_, err = db.Exec("update public.example_user set name ='ftx' where name=$1", "ft")
			//SimplePanic(err)
			var a = make(map[string]string)
			a["d"]="c"
			fmt.Println(a["q"])

		case 3:
			_, err = db.Exec("delete from public.example_user where name=$1", "ftx")
			SimplePanic(err)

		case 4:
			type U struct {
				Name string
			}
			var user []U = make([]U,0)
			err = db.Select(&user,"select * from public.example_user where name = 'ft7'")
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println(user)
			//_, err = db.Exec("select * from public.example_user WHERE name='ft7213213123'  limit 9 ")
			//SimplePanic(err)
		}
	}
}
func SimplePanic(err error){
	fmt.Println(err)
}
