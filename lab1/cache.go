//Shaurya Sharma Lab1
package main

import (
	"container/list"
	//"fmt"
)
// DO NOT CHANGE THIS CACHE SIZE VALUE
const CACHE_SIZE int = 3

var l *list.List = list.New()
var keyMap = make(map[int]*list.Element)

func Get(key int) int {
	v := keyMap[key]
	if v != nil {
		l.MoveToFront(v)
		return v.Value.(int)
	}
	return -1
}

func Set(key int, value int) {
	v := keyMap[key]
	if v == nil {
		if len(keyMap) >= CACHE_SIZE {
			for v := range keyMap {
				if keyMap[v].Value == l.Back().Value {
					delete(keyMap, v)
					//fmt.Println(l)
					l.Remove(l.Back())
					break
				}
			}
		}
		l.PushFront(value)
		keyMap[key] = l.Front()
	} else {
		l.MoveToFront(v)
	}
}
