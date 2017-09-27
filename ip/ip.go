package main

import (
	"encoding/binary"
	"net"

	. "asdf"
)

/*
func Ipv4IntToString(n uint32) string {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, n)
	b := buf.Bytes()
	var b [4]byte

	binary.LittleEndian.PutUint32()
	return net.IPv4(b[0], b[1], b[2], b[3]).String()
}
*/

func main() {
	var bin [4]byte

	s1 := "192.168.0.1"
	s2 := "192.168.0.2"
	b1 := net.ParseIP(s1).To4()
	b2 := net.ParseIP(s2).To4()

	ip1 := binary.BigEndian.Uint32(b1)
	ip2 := binary.BigEndian.Uint32(b2)
	Log.Info("ip1=%d, ip2=%d", ip1, ip2)

	prefix1 := ip1 >> 16
	prefix2 := ip2 >> 16
	Log.Info("prefix1=%d, prefix2=%d", prefix1, prefix2)

	suffix1 := ip1 & 0xffff
	suffix2 := ip2 & 0xffff
	Log.Info("suffix1=%d, suffix2=%d", suffix1, suffix2)

	binary.BigEndian.PutUint32(bin[:], ip1)

	Log.Info("%s==>%v:%x==>%d.%d.%d.%d", s1, b1, ip1, bin[0], bin[1], bin[2], bin[3])
}
