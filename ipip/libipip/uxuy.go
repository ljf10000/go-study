package libipip

import (
	"encoding/binary"
	"net"
)

/******************************************************************************/
// ip as a.b.x.x
type IpB uint32

func (me IpB) Prefix() uint16 {
	return uint16(me >> 16)
}

func (me IpB) Suffix() uint16 {
	return uint16(me & 0xffff)
}

func (me IpB) String() string {
	bin := [4]byte{}

	binary.BigEndian.PutUint32(bin[:], uint32(me))

	return net.IP(bin[:]).String()
}

func MakeIpB(prefix, suffix uint16) IpB {
	return (IpB(prefix) << 16) | (IpB(suffix) & 0xffff)
}

func NewIpB(ip string) IpB {
	return IpB(binary.BigEndian.Uint32(net.ParseIP(ip).To4()))
}

/******************************************************************************/
type IndexDesc uint64

func (me IndexDesc) Offset() uint32 {
	return uint32(me >> 32)
}

func (me IndexDesc) Count() uint32 {
	return uint32(me & 0xffffffff)
}

func MakeIndexDesc(offset, count uint32) IndexDesc {
	return (IndexDesc(offset) << 32) | (IndexDesc(count) & 0xffffffff)
}

/******************************************************************************/
type SymbolDesc uint32

func (me SymbolDesc) Offset() uint32 {
	return uint32(me >> 6)
}

func (me SymbolDesc) Size() uint32 {
	return uint32(me & 0x3f)
}

func MakeSymbolDesc(offset, size uint32) SymbolDesc {
	return SymbolDesc((offset << 6) | (size & 0x3f))
}

/******************************************************************************/
type IpIndex uint64

func (me IpIndex) MinSuffix() uint16 {
	return uint16(me >> 48)
}

func (me IpIndex) MaxSuffix() uint16 {
	return uint16(me << 16 >> 48)
}

func (me IpIndex) IdxEntry() uint32 {
	return uint32(me & 0xffffffff)
}

func MakeIpIndex(min, max uint16, iEntry uint32) IpIndex {
	return (IpIndex(min) << 48) |
		(IpIndex(max) << 32) |
		(IpIndex(iEntry) & 0xffffffff)
}
