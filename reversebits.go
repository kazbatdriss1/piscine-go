func ReverseBits(oct byte) byte {
	var result byte

	for i := 0; i < 8; i++ {
		bit := (oct >> i) & 1
		result |= bit << (7 - i)
	}

	return result
}
