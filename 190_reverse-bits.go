package leetcode

func reverseBits(num uint32) uint32 {
	output := uint32(0)

	for i := 0; i < 32; i++ {
		bit := (num & 1)
		num = num >> 1

		output = (output << 1) + bit
	}

	return output
}

// beat 100% execute time
func reverseBits2(num uint32) uint32 {
	output := uint32(0)

	for i := 0; i < 32; i++ {
		num, output = num>>1, (output<<1)|(num&1)
	}

	return output
}
