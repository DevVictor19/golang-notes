package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros (int8, int16, int32/int, int64)
	var integer int = -32000
	fmt.Println("Literal inteiro é", reflect.TypeOf(integer)) // int

	// ** Sem sinal positivos
	// o "u" vem de unsigned (sem sinal)
	// uint8 uint16 uint32 unit64 -> (1 byte, 2 bytes, 4 bytes, 8 bytes)
	// uint8 uint16 uint32 unit64 -> (byte, short, int, long)

	// 2⁸ - 1 = 255
	var b byte = 255               // ou uint8
	fmt.Println(reflect.TypeOf(b)) // uint8

	// 2¹⁶ - 1 = 65535
	var s uint16 = 65535
	fmt.Println(reflect.TypeOf(s))

	// 2³² - 1 = 4294967295
	var i uint32 = 4294967295
	fmt.Println(reflect.TypeOf(i))

	// 2⁶⁴ - 1 = 1,844674407×10¹⁹
	var l uint64 = math.MaxUint64
	fmt.Println(reflect.TypeOf(l))

	var i2 rune = 'a'               // representa um mapeamento da tabela Unicode (int32)
	fmt.Println(reflect.TypeOf(i2)) // int32

	// ** Números reais (float32, float64)
	// por padrão numeros de ponto flutuante vem como float64
	var f32 float32 = 49.9932
	fmt.Println(reflect.TypeOf(f32))

	var f64 float64 = 3235.0343
	fmt.Println(reflect.TypeOf(f64))

	// ** Booleans
	b1 := true
	b2 := false
	fmt.Println(reflect.TypeOf(b1), reflect.TypeOf(b2))

	// ** Strings
	s1 := "This is a string"
	fmt.Println(reflect.TypeOf(s1))

	// ** Strings com multiplas linhas
	s2 := `This is a 
	multiline string`
	fmt.Println(reflect.TypeOf(s2))

	// ** Char??
	char := 'a'
	fmt.Println(reflect.TypeOf(char)) //int32
}
