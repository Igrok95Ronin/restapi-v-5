package main

import (
	"fmt"
	"log"
	"strconv"
)

type Calculator struct {
	fn, sn int
}

func (f *Calculator) Add(fn, sn int) int {
	f.fn = fn
	f.sn = sn
	return f.fn + f.sn
}

func (s *Calculator) Subtract(fn, sn int) int {
	s.fn = fn
	s.sn = sn
	return s.fn - s.sn
}
func (m *Calculator) Multiply(fn, sn int) int {
	m.fn = fn
	m.sn = sn
	return m.fn * m.sn
}

func (d *Calculator) Divide(fn, sn int) int {
	d.fn = fn
	d.sn = sn
	return d.fn / d.sn
}

type iCalculator interface {
	Add(fn, sn int) int
	Subtract(fn, sn int) int
	Multiply(fn, sn int) int
	Divide(fn, sn int) int
}

func main() {
	var (
		n1, n2, sign string
		i            iCalculator
	)

	f := Calculator{}
	i = &f

	s := Calculator{}
	i = &s

	m := Calculator{}
	i = &m

	d := Calculator{}
	i = &d

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&n1)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&n2)

	fn1, err := strconv.Atoi(n1)
	if err != nil {
		log.Println("Только целочисленные значения")
		log.Fatal(err)
	}
	sn2, err := strconv.Atoi(n2)
	if err != nil {
		log.Println("Только целочисленные значения")
		log.Fatal(err)
	}

	fmt.Print("Введите оператор (+,-,*,/): ")
	fmt.Scanln(&sign)

	switch sign {
	case "+":
		result := i.Add(fn1, sn2)
		fmt.Println(result)
	case "-":
		result := i.Subtract(fn1, sn2)
		fmt.Println(result)
	case "*":
		result := i.Multiply(fn1, sn2)
		fmt.Println(result)
	case "/":
		if fn1 == 0 || sn2 == 0 {
			log.Println("Делить на ноль нельзя")
			return
		} else {
			result := i.Divide(fn1, sn2)
			fmt.Println(result)
		}
	default:
		log.Println("Допустим только один из следующих оператор (+,-,*,/)")
		return
	}
}
