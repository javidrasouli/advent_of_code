package main

import (
	"fmt"
	"strings"
)

func main() {
	// part 1
	StartOfPacket := false
	index := 0
	for !StartOfPacket {
		// part 1
		//part := Buffer[index : index+4]
		// part 2
		part := Buffer[index : index+14]
		PartList := strings.Split(part, "")
		mp := make(map[string]int)
		for _, s := range PartList {
			mp[s] = mp[s] + 1
		}
		IsRepet := false
		for _, i := range mp {
			if i > 1 {
				IsRepet = true
			}
		}
		if !IsRepet {
			StartOfPacket = true
		} else {
			index = index + 1
		}
	}
	fmt.Println("count of character :", index+14)
}
