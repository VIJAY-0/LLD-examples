package lru

import (
	"fmt"
	"testing"
)

func Test_lru(t *testing.T) {
	cache := NewLruCache(3)

	cache.put(1, 100)
	cache.printCahe()
	cache.put(2, 200)
	cache.printCahe()
	cache.put(3, 300)
	cache.printCahe()

	if elem, found := cache.get(2); found {
		fmt.Printf("get 2 :%v \n", elem)
	} else {
		fmt.Printf("get 2 : cache miss \n")
	}

	cache.printCahe()

	if elem, found := cache.get(5); found {
		fmt.Printf("get 5 :%v \n", elem)
	} else {
		fmt.Printf("get 5 : cache miss \n")
	}

	cache.printCahe()
	cache.put(5, 5500)
	cache.printCahe()
}
