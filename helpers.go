package goLittleHelpers

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return err
}

func round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

type CovertItem int

const (
	Speed CovertItem = iota
	Size
	Memory
)

// ConvertBytes get size in bytes(type float64) and return rounded,
// converted size as formatted string
func ConvertBytes(in float64, i CovertItem) string {

	if in <= 0.0 {
		return "0 B"
	}

	var suffixes [5]string

	switch i {
	case Speed:
		suffixes = [5]string{"b/s", "kB/s", "MB/s", "GB/s", "TB/s"}
	case Size:
		suffixes = [5]string{"B", "kB", "MB", "GB", "TB"}
	case Memory:
		suffixes = [5]string{"B", "KiB", "MiB", "GiB", "TiB"}
	default:
		return "0 B"
	}

	base := math.Log(in) / math.Log(1024)
	getSize := round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	s := int(math.Floor(base))
	if s <= 5 {
		getSuffix := suffixes[s]
		return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + getSuffix
	} else {
		return "0 B"
	}
}

// Split []int to [](parts * []int)
func SplitArray(item []int, parts int) [][]int {

	var result [][]int
	length := len(item)
	sliceLength := 0

	// Checks ==========
	if length <= parts {
		return append(result, item)
	}

	if length%parts == 0 {
		sliceLength = length / parts
	} else {
		if length/parts == 1 {
			return append(result, item)
		}
		sliceLength = length / parts
	}
	if sliceLength == 1 {
		return append(result, item)
	}

	// Split ==========
	start := 0
	stop := sliceLength
	for i := 0; i <= parts; i++ {
		if i != parts {
			result = append(result, item[start:stop])
		} else {
			result = append(result, item[start:])
			break
		}
		start = stop
		stop = start + sliceLength
	}

	return result
}
