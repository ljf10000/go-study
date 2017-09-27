package libipip

import (
	"encoding/binary"
	"net"
)

/******************************************************************************/

type U32_16_16 uint32

func (me U32_16_16) A() uint16 {
	return uint16(me >> 16)
}

func (me U32_16_16) B() uint16 {
	return uint16(me & 0xffff)
}

func MakeU32_16_16(a, b uint16) U32_16_16 {
	return (U32_16_16(a) << 16) | (U32_16_16(b) & 0xffff)
}

/******************************************************************************/

type U32_22_10 uint32

func (me U32_22_10) A() uint32 {
	return uint32(me >> 10)
}

func (me U32_22_10) B() uint32 {
	return uint32(me & 0x3ff)
}

func MakeU32_22_10(a, b uint32) U32_22_10 {
	return U32_22_10((a << 10) | (b & 0x3ff))
}

/******************************************************************************/

type U32_26_6 uint32

func (me U32_26_6) A() uint32 {
	return uint32(me >> 6)
}

func (me U32_26_6) B() uint32 {
	return uint32(me & 0x3f)
}

func MakeU32_26_6(a, b uint32) U32_26_6 {
	return U32_26_6((a << 6) | (b & 0x3f))
}

/******************************************************************************/

type U32_27_5 uint32

func (me U32_27_5) A() uint32 {
	return uint32(me >> 5)
}

func (me U32_27_5) B() uint32 {
	return uint32(me & 0x1f)
}

func MakeU32_27_5(a, b uint32) U32_27_5 {
	return U32_27_5((a << 5) | (b & 0x1f))
}

/******************************************************************************/

type U64_32_32 uint64

func (me U64_32_32) A() uint32 {
	return uint32(me >> 32)
}

func (me U64_32_32) B() uint32 {
	return uint32(me & 0xffffffff)
}

func MakeU64_32_32(a, b uint32) U64_32_32 {
	return (U64_32_32(a) << 32) | (U64_32_32(b) & 0xffffffff)
}

/******************************************************************************/

type U64_16_16_32 uint64

func (me U64_16_16_32) A() uint16 {
	return uint16(me >> 48)
}

func (me U64_16_16_32) B() uint16 {
	return uint16(me << 16 >> 48)
}

func (me U64_16_16_32) C() uint32 {
	return uint32(me & 0xffffffff)
}

func MakeU64_16_16_32(a, b uint16, c uint32) U64_16_16_32 {
	return (U64_16_16_32(a) << 48) |
		(U64_16_16_32(b) << 32) |
		(U64_16_16_32(c) & 0xffffffff)
}

/******************************************************************************/
// ip as a.b.x.x
type IpB = U32_16_16

func (me IpB) Prefix() uint16 {
	return me.A()
}

func (me IpB) Suffix() uint16 {
	return me.B()
}

func (me IpB) String() string {
	bin := [4]byte{}

	binary.BigEndian.PutUint32(bin[:], uint32(me))

	return net.IP(bin[:]).String()
}

func MakeIpB(prefix, suffix uint16) IpB {
	return MakeU32_16_16(prefix, suffix)
}

func NewIpB(ip string) IpB {
	return IpB(binary.BigEndian.Uint32(net.ParseIP(ip).To4()))
}

/******************************************************************************/

type IndexDesc = U64_32_32

func (me IndexDesc) Offset() uint32 {
	return me.A()
}

func (me IndexDesc) Count() uint32 {
	return me.B()
}

func MakeIndexDesc(offset, count uint32) IndexDesc {
	return MakeU64_32_32(offset, count)
}

/******************************************************************************/

type SymbolDesc = U32_26_6

func (me SymbolDesc) Offset() uint32 {
	return me.A()
}

func (me SymbolDesc) Size() uint32 {
	return me.B()
}

func MakeSymbolDesc(offset, size uint32) SymbolDesc {
	return MakeU32_26_6(offset, size)
}

/******************************************************************************/

type IpIndex = U64_16_16_32

func (me IpIndex) MinSuffix() uint16 {
	return me.A()
}

func (me IpIndex) MaxSuffix() uint16 {
	return me.B()
}

func (me IpIndex) IdxEntry() uint32 {
	return me.C()
}

func MakeIpIndex(min, max uint16, iEntry uint32) IpIndex {
	return MakeU64_16_16_32(min, max, iEntry)
}

/******************************************************************************/
