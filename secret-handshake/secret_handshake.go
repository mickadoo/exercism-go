package secret

var binaryVals = [4]uint{1, 2, 4, 8}
var codes = [4]string{"wink", "double blink", "close your eyes", "jump"}

const reverseOrder = 16

// Handshake returns the encoded input
func Handshake(input uint) (output []string) {
	for index, binaryVal := range binaryVals {
		if input&binaryVal > 0 {
			output = append(output, codes[index])
		}
	}

	if input&reverseOrder > 0 {
		return reverseSlice(output)
	}

	return output
}

func reverseSlice(input []string) []string {
	for left, right := 0, len(input)-1; left < right; left, right = left+1, right-1 {
		input[left], input[right] = input[right], input[left]
	}

	return input
}
