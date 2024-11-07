package main

func ReverseBits(oct byte) byte {
	var resultByte byte = 0x00
	for i := 0; i < 8; i++ {
		octByte := (oct >> i) & 1
		resultByte |= octByte << (7 - i)
	}
	return resultByte
}
