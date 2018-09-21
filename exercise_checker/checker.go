package exercise_checker

import (
	"reflect"
	"strings"
	"strconv"
	"fmt"
	"errors"
)
type Checker struct{
}

func (checker *Checker) Validate(obj interface{})(bool,string,error){
	vType := reflect.TypeOf(obj)   //该struct的类型数据集
	vValue :=reflect.ValueOf(obj)  //该字段的值数据集
	var tagTmp string              //迭代中的临时注解变量，避免写进循环多次初始化
	var value interface{}          //迭代中的临时值变量，避免写进循环多次初始化

	var er error
	//整型处理的临时变量
	var intTmp int64			   //临时int变量
	var minInt int64			   //用于存放 2:10中的2
	var maxInt int64				//用于存放2：10中的10
	var arrTmp []string			//如用于把 2:10 切割成 2,10
	var scopeArr []string		//如用于把2:10,15:19 切割成 2:10和15:19

	//浮点处理的临时变量
	var floatTmp float64           //临时float变量
	var minFloat float64		//用于存放 2.1:10.1中的2.1
	var maxFloat float64		//用于存放 2.1:10.1中的10.1

	//字符串处理的临时变量
	var stringTmp string		   //临时string变量
	var regRule string
	var timeFormat string		//有时候传的时间变量不一定是date类型，还需要考虑传来string类型的date数据
	var regexArr []string		//如把

	L:
	for i:=0;i<vValue.NumField();i++{
		tagTmp = vType.Field(i).Tag.Get("validate")
		//缺省以及包含-时，跳过该字段不检查
		if tagTmp == "" || strings.Contains(tagTmp,"-"){
			continue
		}
		value = vValue.Field(i).Interface()
		switch vType.Field(i).Type.String() {
		case "int","int8","int16","int32","int64":
			intTmp = PopInt(value)
			scopeArr = strings.Split(tagTmp,",")
			for i,scope:=range scopeArr{
				if strings.Contains(scope,"["){
					arrTmp = strings.Split(scope[1:len(scopeArr)-1],",")  //将'[1,2,3,4,5]'变成数组[1,2,3,4,5]
					for j,num:=range arrTmp{
						tmp, err := strconv.ParseInt(num, 10, 64)
						if err!=nil{
							return false,fmt.Sprintf("scope[%d] [%s][%d]的值为%s,无法转换为int64",i,scope,j,num),err
						}
						if tmp == intTmp {
							continue L
						}
					}
				}else{
					arrTmp = strings.Split(scope,":")
					if len(arrTmp)==1 {
						return false,fmt.Sprintf("scope[%d]的值应该为number:number number: :number而不是%s,限定取值应该为[number1,number2,...]",i,arrTmp[0]),errors.New("int scope ':' not found")
					}else if len(arrTmp)==2 {
						if arrTmp[0]!=""{
							minInt,er = strconv.ParseInt(arrTmp[0],10,64)
							if er!=nil{
								return false,fmt.Sprintf("scope[%d]的值应该为number:number number: :number而不是%s,%s转换int64失败",i,arrTmp[0],arrTmp[0]),er
							}
							if intTmp<minInt {
								return false,fmt.Sprintf("scope[%d]的最小值为[%d],输入的值为[%d],小于了最小范围",i,minInt,intTmp),nil
							}
						}
						if arrTmp[1]!=""{
							maxInt,er = strconv.ParseInt(arrTmp[1],10,64)
							if er!=nil{
								return false,fmt.Sprintf("scope[%d]的值应该为number:number number: :number而不是%s,%s转换int64失败",i,arrTmp[0],arrTmp[0]),er
							}
							if intTmp >maxInt{
								return false,fmt.Sprintf("scope[%d]的最大值为[%d],输入的值为[%d],大于了最大范围",i,maxInt,intTmp),nil
							}
						}

					}
				}
			}
		case "float64","float32":
			floatTmp = PopFloat(value)
		case "string":
			stringTmp = PopString(value)

		}
	}
	return true,"",nil
}

//确保value是int，int32，int16，int64之一
func PopInt(value interface{})int64{
	var rs int64
	switch v:=value.(type){
	case int:
		rs = int64(v)
	case int8:
		rs = int64(v)
	case int16:
		rs= int64(v)
	case int32:
		rs= int64(v)
	case int64:
		return v
	}
	return rs
}

//确保是float32,float64的输入
func PopFloat(value interface{})float64{
	var rs float64
	switch v:=value.(type){
	case float64:
		return v
	case float32:
		return float64(v)
	}
	return rs
}

//确保是string的输入
func PopString(value interface{})string{
	return value.(string)
}