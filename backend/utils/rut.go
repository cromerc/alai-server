package utils

import (
	"errors"
	"strconv"
	"strings"
)

type RutType int8

const (
	Run RutType = iota
	Rut
)

func (r RutType) String() string {
	switch r {
	case Run:
		return "RUN"
	case Rut:
		return "RUT"
	default:
		return "unknown"
	}
}

func CleanRut(rut *string) {
	*rut = strings.ToUpper(*rut)
	*rut = strings.TrimSpace(*rut)
	*rut = strings.Replace(*rut, ".", "", -1)
	*rut = strings.Replace(*rut, "-", "", -1)
}

func PrettyRut(rut *string) {
	tempRut := *rut
	verifier := strings.ToUpper(tempRut[len(tempRut)-1:])
	tempRut = tempRut[:len(tempRut)-1]
	tempRut = Reverse(tempRut)
	tempRut = InsertNth(tempRut, 3, '.')
	tempRut = Reverse(tempRut)
	tempRut = tempRut + "-" + verifier
	*rut = tempRut
}

func IsValidRut(rut string) (bool, error) {
	// rut should be 8 or 9 characters
	if len(rut) != 8 && len(rut) != 9 {
		return false, errors.New("incorrect RUT length")
	}

	verifier := strings.ToUpper(rut[len(rut)-1:])
	tempRut := rut[:len(rut)-1]

	_, err := strconv.Atoi(verifier)
	if err != nil && verifier != "K" {
		return false, errors.New("invalid RUT identifier")
	}

	generatedVerifier, err := generateVerifier(tempRut)
	if err != nil {
		return false, err
	}
	if verifier != generatedVerifier {
		return false, errors.New("incorrect RUT verifier")
	}
	return true, nil
}

func GetRutType(rut string) (RutType, error) {
	tempRut := rut[:len(rut)-1]
	numericRut, err := strconv.Atoi(tempRut)
	if err != nil {
		return Run, errors.New("invalid RUN/RUT")
	}

	if numericRut < 100000000 && numericRut > 50000000 {
		return Rut, nil
	} else {
		return Run, nil
	}
}

func generateVerifier(rut string) (string, error) {
	if _, err := strconv.Atoi(rut); err != nil {
		return "", errors.New("invalid RUT")
	}

	var multiplier = 2
	var sum = 0
	var remainder int
	var division int
	var rutLength = len(rut)

	for i := rutLength - 1; i >= 0; i-- {
		sum = sum + toInt(rut[i:i+1])*multiplier
		multiplier++
		if multiplier == 8 {
			multiplier = 2
		}
	}

	division = sum / 11
	division = division * 11.0
	remainder = sum - int(division)

	if remainder != 0 {
		remainder = 11 - remainder
	}

	if remainder == 10 {
		return "K", nil
	} else {
		return strconv.Itoa(remainder), nil
	}
}
