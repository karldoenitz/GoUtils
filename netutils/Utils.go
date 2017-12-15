package netutils

import (
	"strings"
	"strconv"
	"math"
)

// convert an IPv4 address string into an int number.
// if the input string is not an IPv4 address,
// this function will not throw error but return an Error value
func IpToNumber(ip string) (result int) {
	ipSegments := strings.Split(ip, ".")
	firstNumber, _ := strconv.Atoi(ipSegments[0])
	firstNumber = firstNumber * int(math.Pow(256, 3))
	secondNumber, _ := strconv.Atoi(ipSegments[1])
	secondNumber = secondNumber * int(math.Pow(256, 2))
	thirdNumber, _ := strconv.Atoi(ipSegments[2])
	thirdNumber = thirdNumber * int(math.Pow(256, 1))
	forthNumber, _ := strconv.Atoi(ipSegments[3])
	forthNumber = forthNumber * int(math.Pow(256, 0))
	result = firstNumber + secondNumber + thirdNumber + forthNumber
	return result
}

// convert an int number which return from IpToNumber into an IPv4 address string.
// if the input number is invalid,
// this function will not throw error but return an Invalid value.
func NumberToIp(number int) (result string) {
	firstIpSegmentNumber := (number / int(math.Pow(256, 3))) % 256
	firstIpSegment := strconv.Itoa(firstIpSegmentNumber)
	secondIpSegmentNumber := (number / int(math.Pow(256, 2))) % 256
	secondIpSegment := strconv.Itoa(secondIpSegmentNumber)
	thirdIpSegmentNumber := (number / int(math.Pow(256, 1))) % 256
	thirdIpSegment := strconv.Itoa(thirdIpSegmentNumber)
	forthIpSegmentNumber := (number / int(math.Pow(256, 0))) % 256
	forthIpSegment := strconv.Itoa(forthIpSegmentNumber)
	result = strings.Join([]string{firstIpSegment, secondIpSegment, thirdIpSegment, forthIpSegment}, ".")
	return result
}
