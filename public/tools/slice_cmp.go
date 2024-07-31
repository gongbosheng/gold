package tools

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	NullOutPut             = errors.New("null out put")
	NoUserID               = errors.New("have no userid")
	NoWorkID               = errors.New("there is no such work order")
	ParameterError         = errors.New("parameter error")
	InsideParameterError   = errors.New("inside parameter error")
	UnSupportHaystackError = errors.New("Type error")
)

// 字符串切片比较
func ArrStrCmp(src []string, dest []string) ([]string, []string) {
	msrc := make(map[string]byte) //按源数组建索引
	mall := make(map[string]byte) //源+目所有元素建索引
	var set []string              //交集
	//1.源数组建立map
	for _, v := range src {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range dest {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) {
			continue
		} else {
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除的，找不到的必定能在目数组中找到，即新加的
	var added, deleted []string
	for v := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}
	return added, deleted
}

// uint切片比较
func ArrUintCmp(src []uint, dest []uint) ([]uint, []uint) {
	msrc := make(map[uint]byte) //按源数组建索引
	mall := make(map[uint]byte) //源+目所有元素建索引
	var set []uint              //交集
	//1.源数组建立map
	for _, v := range src {
		msrc[v] = 0
		mall[v] = 0
	}
	//2.目数组中，存不进去，即重复元素，所有存不进去的集合就是并集
	for _, v := range dest {
		l := len(mall)
		mall[v] = 1
		if l != len(mall) {
			continue
		} else {
			set = append(set, v)
		}
	}
	//3.遍历交集，在并集中找，找到就从并集中删，删完后就是补集（即并-交=所有变化的元素）
	for _, v := range set {
		delete(mall, v)
	}
	//4.此时，mall是补集，所有元素去源中找，找到就是删除
	var added, deleted []uint
	for v := range mall {
		_, exist := msrc[v]
		if exist {
			deleted = append(deleted, v)
		} else {
			added = append(added, v)
		}
	}
	return added, deleted
}

// 将字符串切片转换为uint切片
func SliceToString(src []uint, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(src), " ", delim, -1), "[]")
}

// 将字符串切片转换为uint切片
func StringToSlice(src string, delim string) []uint {
	var dest []uint
	if src == "" {
		return dest
	}
	strs := strings.Split(src, delim)
	for _, v := range strs {
		t, _ := strconv.Atoi(v)
		dest = append(dest, uint(t))
	}
	return dest
}

type SliceInterface interface {
	// Page 数组分页
	Page(page, pageSize, nums int) (sliceStart, sliceEnd int)

	// Remove 数组删除元素
	Remove(haystack interface{}, needle interface{}) (interface{}, error)

	// In 判断元素是否存在
	In(haystack interface{}, needle interface{}) (bool, error)

	// RemoveDuplicateElement 数组去重
	RemoveDuplicateElement(originals interface{}) (interface{}, error)

	// ArrayToString 数组转字符串
	ArrayToString(data []string) (r string)

	// DifferenceSet 比较数组不同：针对[]int	类型
	DifferenceSet(a []int, b []int) []int

	// Intersect 交集：针对[]int  类型
	Intersect(slice1, slice2 []int) []int

	// Intersection 交集：针对[]string 类型
	Intersection(slice1, slice2 []string) []string

	// ArrayInGroupsOf 数组分割
	ArrayInGroupsOf(arr []string, num int) [][]string

	// ConvertToStringSlice 将interface数组转换为字符串数组
	ConvertToStringSlice(interfaces []interface{}) []string
}

type SliceStruct struct {
}

// ArrayToString 数组转string
func (s *SliceStruct) ArrayToString(data []string) (r string) {

	for _, d := range data {
		r += d + ","
	}

	r = r[0 : len(r)-1]
	return r
}

type sliceError struct {
	msg string
}

func (e *sliceError) Error() string {
	return e.msg
}

func Errorf(format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return &sliceError{msg}
}

// RemoveDuplicateElement 对象去重复
func (s *SliceStruct) RemoveDuplicateElement(originals interface{}) (interface{}, error) {
	temp := map[string]struct{}{}
	switch slice := originals.(type) {
	case []string:
		result := make([]string, 0, len(originals.([]string)))
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := temp[key]; !ok {
				temp[key] = struct{}{}
				result = append(result, item)
			}
		}
		return result, nil
	case []int:
		result := make([]int, 0, len(originals.([]int)))
		for _, item := range slice {
			key := fmt.Sprint(item)
			if _, ok := temp[key]; !ok {
				temp[key] = struct{}{}
				result = append(result, item)
			}
		}
		return result, nil
	default:
		err := Errorf("Unknown type: %T", slice)
		return nil, err
	}
}

// In slice in
func (s *SliceStruct) In(haystack interface{}, needle interface{}) (bool, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				return true, nil
			}
		}
		return false, nil
	}
	return false, UnSupportHaystackError
}

// Remove slice remove
func (s *SliceStruct) Remove(haystack interface{}, needle interface{}) (interface{}, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				switch haystack.(type) {
				case []int:
					haystack = append(haystack.([]int)[:i], haystack.([]int)[i+1:]...)
				case []string:
					haystack = append(haystack.([]string)[:i], haystack.([]string)[i+1:]...)
				default:
					return nil, UnSupportHaystackError
				}
				return haystack, UnSupportHaystackError
			}
		}
	}
	return false, UnSupportHaystackError
}

func (s *SliceStruct) Page(page, pageSize, nums int) (sliceStart, sliceEnd int) {
	if page < 0 {
		page = 1
	}

	if pageSize < 0 {
		pageSize = 20
	}

	if pageSize > nums {
		return 0, nums
	}

	// 总页数
	pageCount := int(math.Ceil(float64(nums) / float64(pageSize)))
	if page > pageCount {
		return 0, 0
	}
	sliceStart = (page - 1) * pageSize
	sliceEnd = sliceStart + pageSize

	if sliceEnd > nums {
		sliceEnd = nums
	}
	return sliceStart, sliceEnd

}

// 求差集
func (s *SliceStruct) DifferenceSet(a []int, b []int) []int {
	var c []int
	temp := map[int]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			c = append(c, val)
		}
	}

	return c
}

// Intersect 求交集
func (s *SliceStruct) Intersect(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func (s *SliceStruct) Intersection(slice1, slice2 []string) []string {
	set := make(map[string]bool)

	// 存储第一个切片中的元素到 map
	for _, value := range slice1 {
		set[value] = true
	}

	// 创建一个切片用于存储交集
	var result []string

	// 遍历第二个切片，如果元素在 map 中存在，则为交集
	for _, value := range slice2 {
		if set[value] {
			result = append(result, value)
		}
	}

	return result
}

// 数组分割
func (s *SliceStruct) ArrayInGroupsOf(arr []string, num int) [][]string {
	max := len(arr)
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]string{arr}
	}
	//获取应该数组分割为多少份
	var quantity int
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]string, 0)
	//声明分割数组的截止下标
	var start, end, i int
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, arr[start:end])
		} else {
			segments = append(segments, arr[start:])
		}
		start = i * num
	}
	return segments
}

func (s *SliceStruct) ConvertToStringSlice(interfaces []interface{}) []string {
	var resultStr []string
	for _, item := range interfaces {
		str, ok := item.(string) // 尝试类型断言
		if ok {
			resultStr = append(resultStr, str)
		} else {
			// 可选: 如果类型断言失败，可以在这里处理错误或者忽略当前项
			fmt.Println("Warning: Conversion to string failed.")
		}
	}
	return resultStr
}

func GenerateRandomString(length int) string {
	// 定义包含数字和字母的字符集
	characters := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机字符串
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = characters[rand.Intn(len(characters))]
	}

	return string(result)
}
