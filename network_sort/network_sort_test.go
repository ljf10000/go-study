package array

import (
	"fmt"
	"encoding/binary"
	"testing"
)

func TestBigEndianGet(t *testing.T){
	b := []byte{1,2,3,4}
	fmt.Println("binary is ", b)
	
	var c uint32 = binary.BigEndian.Uint32(b)
	fmt.Printf("u32 is 0x%x", c)
}

func TestBigEndianPut(t *testing.T){
	a := [4]byte{}
	b := a[:]
	
	binary.BigEndian.PutUint32(b, 0x01020304)
	fmt.Println("binary is ", b)
}

func TestLittleEndianGet(t *testing.T){
	b := []byte{1,2,3,4}
	fmt.Println("binary is ", b)
	
	var c uint32 = binary.LittleEndian.Uint32(b)
	fmt.Printf("u32 is 0x%x", c)
}

func TestLittleEndianPut(t *testing.T){
	a := [4]byte{}
	b := a[:]
	
	binary.LittleEndian.PutUint32(b, 0x01020304)
	fmt.Println("binary is ", b)
}