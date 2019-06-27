package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name  string
	Phone string
	Id    int
}

func main() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	Insert(c)
	Find(c)
	// Update(c)
	// First(c)
	//Delete(c)

	//ComplexQuery(c)
}

func Insert(c *mgo.Collection) {
	// 插入
	err := c.Insert(&Person{"Ale", "+55 53 8116 9639", 1},
		&Person{"Cla", "+55 53 8402 8510",2})
	if err != nil {
		log.Fatal(err)
	}

}

func Find(c *mgo.Collection) {
	// 查询
	var results []Person
	// err := c.Find(bson.M{"name":"Ale"}).All(&results)
	// err := c.Find(map[string]interface{}{"name":"Ale"}).All(&results)
	err := c.Find(nil).All(&results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results)
}

func First(c *mgo.Collection) {
	// 查询
	var result Person
	// err := c.Find(bson.M{"name":"Ale"}).One(&result)
	// err := c.Find(map[string]interface{}{"name":"Ale"}).One(&result)
	err := c.Find(nil).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
func Update(c *mgo.Collection) {
	err := c.Update(bson.M{"name": "Ale2"}, &Person{"Ale", "+55 53 8116 9639", 1})
	if err != nil {
		log.Fatal(err)
	}

	Find(c)
}
func Delete(c *mgo.Collection) {
	// 只删第一个
	err := c.Remove(nil)
	if err != nil {
		log.Fatal(err)
	}
	// 删除所有
	info, err := c.RemoveAll(nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(info)
	Find(c)
}

func ComplexQuery(c *mgo.Collection) {
	defer c.RemoveAll(nil)
	Insert(c)
	// in
	// where name in ["Ale", "Cla"]
	var results []Person
	if e := c.Find(bson.M{"name": bson.M{"$in": []string{"Ale", "Cla"}}}).All(&results); e != nil {
		panic(e)
	}
	fmt.Println("in", results)

	// and
	// where name = "Ale" and phone = "+55 53 8116 9639"
	results = nil
	if e := c.Find(bson.M{"name": "Ale", "phone": "+55 53 8116 9639"}).All(&results); e != nil {
		panic(e)
	}
	fmt.Println("and", results)

	// or
	// where name = "Ale" or name = "Cla"
	results = nil
	if e := c.Find(bson.M{"$or": []interface{}{bson.M{"name": "Ale"}, bson.M{"name": "Cla"}}}).Sort("-name").Limit(2).All(&results); e != nil {
		panic(e)
	}
	fmt.Println("or", results)

	// and + or
	// where phone = "+55 53 ..." and (name="Ale" or name= "Cla")
	results = nil
	if e := c.Find(bson.M{"phone": "+55 53 8116 9639", "$or": []interface{}{bson.M{"name": "Ale"}, bson.M{"name": "Cla"}}}).All(&results); e != nil {
		panic(e)
	}
	fmt.Println("and + or", results)

	// !=
	// where name != "Cla"
	//$gt -------- greater than  >
	//
	//$gte --------- gt equal  >=
	//
	//$lt -------- less than  <
	//
	//$lte --------- lt equal  <=
	//
	//$ne ----------- not equal  !=
	//
	//$eq  --------  equal  =
	results = nil

	if e := c.Find(bson.M{"name": bson.M{"$ne": "Cla"}}).All(&results); e != nil {
		panic(e)
	}
	fmt.Println("!=", results)
}
