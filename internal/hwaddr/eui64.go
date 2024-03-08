package hwaddr

func EUIFrom6(addr [6]byte) EUI64 {
	eui64 := [8]byte{}

	copy(eui64[:3], addr[:3])
	eui64[0] ^= 0x02

	eui64[3], eui64[4] = 0xFF, 0xFE

	copy(eui64[5:], addr[3:])

	return eui64
}

type EUI64 [8]byte

func (e EUI64) AppendToPrefixAddr(prefix [16]byte) [16]byte {
	copy(prefix[8:], e[:])
	return prefix
}
