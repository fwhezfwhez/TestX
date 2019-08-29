package test_practice

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/*

从10000个数中，获取最小的10个:

- 生成1个包含了10000个数据的数组
- 将大数组切割成多个，长度只为100的小数组
- 对各个小数组，获取前十个最小值
- 所有最小值合并，并再取最小值十个。
 */

// 生成包含了M个随机数据的单个大数组
func GenMNumberArr(M int, seed time.Time) []int {
	var rs = make([]int, 0, M)
	rand.Seed(seed.UnixNano())
	for i := 0; i < M; i++ {
		rs = append(rs, rand.Intn(1029381))
	}
	return rs
}

// 生成了长度20的随机数组
func TestGenMNumberArr(t *testing.T) {
	rs := GenMNumberArr(20, time.Now())
	fmt.Println(len(rs))
	fmt.Println(rs)
}

// 从数组中，选出最小的K个数据,并升序排列成数组
func GetMinK(k int, src []int) []int {
	if k > len(src) {
		panic(fmt.Errorf("min k must smaller than src length, but got k '%d', src.len '%d'", k, len(src)))
	}
	var tmp int
	for i := 0; i < k; i++ {
		for j := i + 1; j < len(src); j++ {
			if src[i] > src[j] {
				tmp = src[i]
				src[i] = src[j]
				src[j] = tmp
			}
		}
	}
	return src[:k]
}

// 从数组中，选取了最小的4个
func TestGetMinK(t *testing.T) {
	rs := GetMinK(4, []int{1, 3, 4, 5, 6, 12, 3, 2, 15, 199})
	fmt.Println(rs)
}

// 将大数组，切割成多个小数组，每个小数组容量为vol
func SplitArr(arr []int, vol int) [][]int {
	var rs = make([][]int, 0, 100)
	var tmp []int
L:
	for i, _ := range arr {
		if (i+1)%vol == 0 {
			tmp = arr[0 : i+1]
			rs = append(rs, tmp)
			arr = arr[i+1:]
			goto L
		} else {
			if i == len(arr)-1 {
				rs = append(rs, arr)
			}
			continue
		}
	}
	return rs
}

func TestSplitArr(t *testing.T) {
	rs := SplitArr([]int{1, 24, 1, 2, 3, 4, 11, 12, 19, 11, 17, 19, 22, 23}, 5)
	fmt.Println(rs)
}

func TestGetMin10NumberFrom10000(t *testing.T) {
	arr10000 := GenMNumberArr(10000, time.Now())

	subArrs := SplitArr(arr10000, 100)

	minArr := make([]int, 0, 100)
	for i, _ := range subArrs {
		minArr = append(minArr, GetMinK(10, subArrs[i])...)
	}

	rs :=GetMinK(10, minArr)
	fmt.Println(rs)
}
