package set

import (
	"bytes"
	"fmt"
)

type HashSet struct {
	m map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e) //第一个参数为目标字典类型，第二个参数为要删除的那个键值对的键
}

func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

func (set *HashSet) Len() int {
	return len(set.m)
}

func (set *HashSet) Same(other *HashSet) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m) //获取HashSet中字段m的长度，即m中包含元素的数量
	//初始化一个[]interface{}类型的变量snapshot来存储m的值中的元素值
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	//按照既定顺序将迭代值设置到快照值(变量snapshot的值)的指定元素位置上,这一过程并不会创建任何新值。
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else { //m的值中的元素数量有所增加，使得实际迭代的次数大于先前初始化的快照值的长度
			snapshot = append(snapshot, key) //使用append函数向快照值追加元素值。
		}
		actualLen++ //实际迭代的次数
	}
	//对于已被初始化的[]interface{}类型的切片值来说，未被显示初始化的元素位置上的值均为nil。
	//m的值中的元素数量有所减少，使得实际迭代的次数小于先前初始化的快照值的长度。
	//这样快照值的尾部存在若干个没有任何意义的值为nil的元素，
	//可以通过snapshot = snapshot[:actualLen]将无用的元素值从快照值中去掉。
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

func (set *HashSet) String() string {
	var buf bytes.Buffer //作为结果值的缓冲区
	buf.WriteString("HashSet{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

func (set *HashSet) IsSuperset(other *HashSet) bool {
	if other == nil { //如果other为nil，则other不是set的子集
		return false
	}
	setLen := set.Len()                    //获取set的元素值数量
	otherLen := other.Len()                //获取other的元素值数量
	if setLen == 0 || setLen == otherLen { //set的元素值数量等于0或者等于other的元素数量
		return false
	}
	if setLen > 0 && otherLen == 0 { //other为元素数量为0，set元素数量大于0，则set也是other的超集
		return true
	}
	for _, v := range other.Elements() {
		if !set.Contains(v) { //只要set中有一个包含other中的数据，就返回false
			return false
		}
	}
	return true
}

func (set *HashSet) Union(other *HashSet) *HashSet {
	if set == nil || other == nil { // set和other都为nil，则它们的并集为nil
		return nil
	}
	unionSet := NewHashSet()           //新创建一个HashSet类型值，它的长度为0，即元素数量为0
	for _, v := range set.Elements() { //将set中的元素添加到unionedSet中
		unionSet.Add(v)
	}
	if other.Len() == 0 {
		return unionSet
	}
	for _, v := range other.Elements() { //将other中的元素添加到unionedSet中，如果遇到相同，则不添加（在Add方法逻辑中体现）
		unionSet.Add(v)
	}
	return unionSet
}

func (set *HashSet) Intersect(other *HashSet) *HashSet {
	if set == nil || other == nil { // set和other都为nil，则它们的交集为nil
		return nil
	}
	intersectedSet := NewHashSet() //新创建一个HashSet类型值，它的长度为0，即元素数量为0
	if other.Len() == 0 {          //other的元素数量为0，直接返回intersectedSet
		return intersectedSet
	}
	if set.Len() < other.Len() { //set的元素数量少于other的元素数量
		for _, v := range set.Elements() { //遍历set
			if other.Contains(v) { //只要将set和other共有的添加到intersectedSet
				intersectedSet.Add(v)
			}
		}
	} else { //set的元素数量多于other的元素数量
		for _, v := range other.Elements() { //遍历other
			if set.Contains(v) { //只要将set和other共有的添加到intersectedSet
				intersectedSet.Add(v)
			}
		}
	}
	return intersectedSet
}

func (set *HashSet) Difference(other *HashSet) *HashSet {
	if set == nil || other == nil { // set和other都为nil，则它们的差集为nil
		return nil
	}
	diffSet := NewHashSet() //新创建一个HashSet类型值，它的长度为0，即元素数量为0
	if other.Len() == 0 {   // 如果other的元素数量为0
		for _, v := range set.Elements() { //遍历set，并将set中的元素v添加到differencedSet
			diffSet.Add(v)
		}
		return diffSet //直接返回differencedSet
	}
	for _, v := range set.Elements() { //other的元素数量不为0，遍历set
		if !other.Contains(v) { //如果other中不包含v，就将v添加到differencedSet中
			diffSet.Add(v)
		}
	}
	return diffSet
}

func (set *HashSet) SymmetricDifference(other *HashSet) *HashSet {
	if set == nil || other == nil { // set和other都为nil，则它们的对称差集为nil
		return nil
	}
	diffA := set.Difference(other) //生成集合 set 对集合 other 的差集
	if other.Len() == 0 {          //如果other的元素数量等于0，那么other对集合set的差集为空，则直接返回diffA
		return diffA
	}
	diffB := other.Difference(set) //生成集合 other 对集合 set 的差集
	return diffA.Union(diffB)      //返回集合 diffA 和集合 diffB 的并集
}

type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
}

func IsSuperset(one Set, other Set) bool {
	if one == nil || other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0 {
		return true
	}
	for _, v := range other.Elements() {
		if !one.Contains(v) {
			return false
		}
	}
	return true
}
