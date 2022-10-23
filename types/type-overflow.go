package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(`int has a platform-dependent size, 
	on a 32-bit system, it holds a 32 bit signed integer, 
	while on a 64-bit system, it holds a 64-bit signed integer`)

	var maxInt8 = math.MaxInt8
	var minInt8 = math.MinInt8
	fmt.Println("\nint8")
	fmt.Println("max number of int8: ", maxInt8)
	fmt.Println("min number of int8: ", minInt8)
	fmt.Println("type overflow int8(127 + 1): ", int8(maxInt8+1))
	fmt.Println("type overflow int8(127 + 2): ", int8(maxInt8+2))
	fmt.Println("type overflow int8(127 + 129): ", int8(maxInt8+129))
	fmt.Println("type overflow int8(127 + 130): ", int8(maxInt8+130))

	var maxInt16 = math.MaxInt16
	var minInt16 = math.MinInt16
	fmt.Println("\nint16")
	fmt.Println("max number of int16: ", maxInt16)
	fmt.Println("min number of int16: ", minInt16)
	fmt.Println("type overflow int16(32767 + 1): ", int16(maxInt16+1))
	fmt.Println("type overflow int16(32767 + 2): ", int16(maxInt16+2))
	fmt.Println("type overflow int16(32767 + 32767 + 2): ", int16(maxInt16+maxInt16+2))
	fmt.Println("type overflow int16(32767 + 32767 + 3): ", int16(maxInt16+maxInt16+3))

	var maxInt32 = math.MaxInt32
	var minInt32 = math.MinInt32
	fmt.Println("\nint32")
	fmt.Println("max number of int32: ", maxInt32)
	fmt.Println("min number of int32: ", minInt32)
	fmt.Println("type overflow int32(2147483647 + 1): ", int32(maxInt32+1))
	fmt.Println("type overflow int32(2147483647 + 2): ", int32(maxInt32+2))
	fmt.Println("type overflow int32(2147483647 + 2147483647 + 2): ", int32(maxInt32+maxInt32+2))
	fmt.Println("type overflow int32(2147483647 + 2147483647 + 3): ", int32(maxInt32+maxInt32+3))

	var maxInt64 = math.MaxInt64
	var minInt64 = math.MinInt64
	fmt.Println("\nint64")
	fmt.Println("max number of int64: ", maxInt64)
	fmt.Println("min number of int64: ", minInt64)
	fmt.Println("type overflow int64(9223372036854775807 + 1): ", int64(maxInt64+1))
	fmt.Println("type overflow int64(9223372036854775807 + 2): ", int64(maxInt64+2))
	fmt.Println("type overflow int64(9223372036854775807 + 9223372036854775807 + 2): ", int64(maxInt64+maxInt64+2))
	fmt.Println("type overflow int64(9223372036854775807 + 9223372036854775807 + 3): ", int64(maxInt64+maxInt64+3))

	fmt.Printf("\n")
	fmt.Println(`uint has a platform-dependent size, 
	on a 32-bit system, it holds a 32 bit unsigned integer, 
	while on a 64-bit system, it holds a 64-bit unsigned integer`)

	var maxUInt8 = math.MaxUint8
	var minUInt8 = 0
	fmt.Println("\nuint8")
	fmt.Println("max number of uint8: ", maxUInt8)
	fmt.Println("min number of uint8: ", minUInt8)
	fmt.Println("type overflow uint8(255 + 1): ", uint8(maxUInt8+1))
	fmt.Println("type overflow uint8(255 + 2): ", uint8(maxUInt8+2))

	var maxUInt16 = math.MaxUint16
	var minUInt16 = 0
	fmt.Println("\nuint16")
	fmt.Println("max number of uint16: ", maxUInt16)
	fmt.Println("min number of uint16: ", minUInt16)
	fmt.Println("type overflow uint16(65535 + 1): ", uint16(maxUInt16+1))
	fmt.Println("type overflow uint16(65535 + 2): ", uint16(maxUInt16+2))

	var maxUInt32 = math.MaxUint32
	var minUInt32 = 0
	fmt.Println("\nuint32")
	fmt.Println("max number of uint32: ", maxUInt32)
	fmt.Println("min number of uint32: ", minUInt32)
	fmt.Println("type overflow uint32(4294967295 + 1): ", uint32(maxUInt32+1))
	fmt.Println("type overflow uint32(4294967295 + 2): ", uint32(maxUInt32+2))

	var maxUInt64 uint64 = math.MaxUint64
	var minUInt64 = 0
	fmt.Println("\nuint64")
	fmt.Println("max number of uint64: ", maxUInt64)
	fmt.Println("min number of uint64: ", minUInt64)
	fmt.Println("type overflow uint64(18446744073709551615 + 1): ", uint64(maxUInt64+1))
	fmt.Println("type overflow uint64(18446744073709551615 + 2): ", uint64(maxUInt64+2))

	var minFloat32 float32 = math.SmallestNonzeroFloat32
	var maxFloat32 = math.MaxFloat32
	fmt.Println("\nfloat32")
	fmt.Printf("max number of float32: %.50e\n", maxFloat32)
	fmt.Printf("min number of float32: %.50e\n", minFloat32)
	fmt.Printf("type overflow float32(MaxFloat32 + 1): %.50e\n", float32(maxFloat32+1))
	fmt.Printf("type overflow float32(MaxFloat32 + 20): %.50e\n", float32(maxFloat32+20))

	var minFloat64 = math.SmallestNonzeroFloat64
	var maxFloat64 = math.MaxFloat64
	fmt.Println("\nfloat64")
	fmt.Printf("min number of float64: %.50e\n", minFloat64)
	fmt.Printf("max number of float64: %.50e\n", maxFloat64)
	fmt.Printf("type overflow float64(MaxFloat64 + 1): %.50e\n", float64(maxFloat64+1))
	fmt.Printf("type overflow float64(MaxFloat64 + 20): %.50e\n", float64(maxFloat64+20))
}
