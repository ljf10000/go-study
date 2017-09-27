package main

import (
	"fmt"
	"time"
)

func main() {
	input := "00:11:22:33:44:55"
	mac := [6]byte{}
	sep := byte(0)
	
	fmt.Sscanf(input, "%d%c%d%c%d%c%d%c%d%c%d",
		&mac[0], &sep,
		&mac[1], &sep,
		&mac[2], &sep,
		&mac[3], &sep,
		&mac[4], &sep,
		&mac[5])
	
	fmt.Println(mac)
	
	fmt.Println("AABBCCDDEEFF" + time.Now().Format("20060102150405"))
}
