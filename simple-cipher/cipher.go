package cipher

import (
	"errors"
	"strings"
	"unicode"
)

// Cipher is an interface for encoding and decoding strings
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

func moveCharacter(char rune, distance int) (rune, error) {
	if distance < 0 {
		distance = 26 + distance
	}

	val := int(char - 'a')
	if val > 26 || val < 0 {
		return 0, errors.New("Invalid rune")
	}
	val = ((val + distance) % 26) + 'a'

	return rune(val), nil
}

func encodeWithShift(input string, distance int) string {
	result := ""
	input = strings.ToLower(input)
	for _, char := range input {
		newChar, err := moveCharacter(char, distance)
		if err != nil {
			continue
		}
		result += string(newChar)
	}

	return result
}

// Caesar is a basic cipher
type Caesar struct{}

// Encode encodes the input using the caesar method
func (c Caesar) Encode(input string) string {
	return encodeWithShift(input, 3)
}

// Decode decodes the input using the caesar method
func (c Caesar) Decode(input string) string {
	return encodeWithShift(input, -3)
}

// NewCaesar returns a new Caesar cipher
func NewCaesar() Cipher {
	return Caesar{}
}

// Shift is a basic cipher
type Shift struct {
	distance int
}

// Encode encodes the input using the shift method
func (s Shift) Encode(input string) string {
	return encodeWithShift(input, s.distance)
}

// Decode decodes the input using the shift method
func (s Shift) Decode(input string) string {
	return encodeWithShift(input, -s.distance)
}

// NewShift returns a new Shift cipher
func NewShift(distance int) Cipher {
	if distance > 25 || distance < -25 || distance == 0 {
		return nil
	}

	return Shift{distance}
}

// Vigenere is a basic cipher
type Vigenere struct {
	key string
}

// Encode encodes the input using the vigenere method
func (v Vigenere) Encode(input string) string {
	input = strings.ToLower(input)
	result := ""
	pos := 0
	for _, char := range input {
		distance := int(v.key[pos%len(v.key)] - 'a')
		newChar, err := moveCharacter(char, distance)
		if err != nil {
			continue
		}
		pos++
		result += string(newChar)
	}

	return result
}

// Decode decodes the input using the vigenere method
func (v Vigenere) Decode(input string) string {
	input = strings.ToLower(input)
	result := ""
	pos := 0
	for _, char := range input {
		distance := int(v.key[pos%len(v.key)] - 'a')
		distance = 26 - distance
		newChar, err := moveCharacter(char, distance)
		if err != nil {
			continue
		}
		pos++
		result += string(newChar)
	}

	return result
}

// NewVigenere returns a new Vigenere cipher
func NewVigenere(key string) Cipher {
	if key == "" {
		return nil
	}
	isAllAs := true
	for _, char := range key {
		if !unicode.IsLetter(char) || !unicode.IsLower(char) {
			return nil
		}
		if char != 'a' {
			isAllAs = false
		}
	}
	if isAllAs {
		return nil
	}

	return Vigenere{key}
}
