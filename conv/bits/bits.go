package bits

import "math"


const (
	kb float64 = 1024
)

var (
	mb float64 = math.Pow(kb, 2)
	gb float64 = math.Pow(kb, 3)
)

func byteConv(bytes uint64, format float64) float64 {
	return math.Round((float64(bytes) / format) * 100 ) / 100
}


// ToGB function will convert bytes to gigabytes (GB)
func ToGB(bytes uint64) float64 {
	return byteConv(bytes, gb)
}

// ToMB function will convert bytes to megabytes (MB)
func ToMB(bytes uint64) float64 {
	return byteConv(bytes, mb)
}

// ToKB function will convert bytes to kilobytes (KB)
func ToKB(bytes uint64) float64 {
	return byteConv(bytes, kb)
}

// DoubleDecimal function will return a float with two 
// decimal numbers only
func DoubleDecimal(raw float64) float64 {
	return math.Round(raw * 100) / 100
}

// ShiftLoads function will divide the loads values by 
// 2^16; or, n / 65536; or, n / (1 << 16)
func ShiftLoads(raw [3]uint64) [3]float64 {
	return [3]float64{
		DoubleDecimal(float64(raw[0]) / (1 << 16)),
		DoubleDecimal(float64(raw[1]) / (1 << 16)),
		DoubleDecimal(float64(raw[2]) / (1 << 16)),
	}
}