package main

import (
	"fmt"

	"github.com/jskelcy/interview/LRU/lru"
)

func main() {
	cache, _ := lru.NewLRU(3)

	cache.Insert("jake", "skelcy")
	cache.Insert("matt", "disipio")
	cache.Insert("shawn", "heneghan")
	cache.Read("jake")
	cache.Insert("sam", "corbett")

	for _, name := range []string{"jake", "matt", "shawn", "sam"} {
		data, err := cache.Read(name)
		if err != nil {
			fmt.Printf("read error: %s\n", err.Error())
			continue
		}
		fmt.Printf("last name: %s\n", data.(string))
	}
}
