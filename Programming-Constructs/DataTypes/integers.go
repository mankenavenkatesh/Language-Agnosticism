package main

import "fmt"

var (
	eightBitU uint8 = uint8(255)
	eightBitS int8  = int8(-128)

	sixteenBitU uint16 = uint16(65535)
	sixteenBitS int16  = int16(-32768)

	thirtyTwoBitU uint32 = uint32(4294967295)
	thirtyTwoBitS int32  = int32(-2147483648)

	sixtyFourBitU uint64 = uint64(18446744073709551615)
	sixtyFourBitS int64  = int64(-9223372036854775808)
)

func main() {
	fmt.Println("Eight bit unsigned", eightBitU)
	fmt.Println("Eight bit signed", eightBitS)

	fmt.Println("Sixteen bit unsigned", sixteenBitU)
	fmt.Println("Sixteen bit signed", sixteenBitS)

	fmt.Println("thirtyTwo bit unsigned", thirtyTwoBitU)
	fmt.Println("thirtyTwo bit signed", thirtyTwoBitS)

	fmt.Println("sixtyFour bit unsigned", sixtyFourBitU)
	fmt.Println("sixtyFour bit signed", sixtyFourBitS)
}
