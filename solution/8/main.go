/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

type Data struct {
	number int64
	pos    int
	value  int
}

func main() {
	data, err := ParseDate()
	if err != nil {
		log.Fatal(err)
	}
	data.PrintBinary()
	data.SetBit()
	data.PrintBinary()
}

func ParseDate() (Data, error) {
	number := flag.Int64("n", 1234567, "variable")
	pos := flag.Int("i", 0, "bit number")
	value := flag.Int("v", 0, "bit value")
	flag.Parse()

	data := Data{*number, *pos, *value}

	if *pos > 63 || *pos < 0 || *value < 0 || *value > 1 {
		return data, errors.New("Out of range")
	}
	return data, nil
}

func (d *Data) SetBit() {
	mask := int64(1) << d.pos // создаем маску на ту позицую бита, которую нужно изменить
	if d.value == 0 {
		d.number = d.number &^ mask // и исключающее или
	} else if d.value == 1 {
		d.number = d.number | mask // или
	}

}

func (d Data) PrintBinary() {
	fmt.Printf("Binary representation of %d:\t", d.number)
	for i := 63; i >= 0; i-- {
		bit := (d.number >> i) & 1
		fmt.Printf("%d", bit)
	}
	fmt.Println()
}
