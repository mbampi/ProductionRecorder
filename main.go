package main

import (
	"log"
)

func main() {
	fieldsFazenda := []Field{
		{Name: "1", Area: 200, Productions: []int{200, 450, 120, 30}},
		{Name: "2", Area: 50, Productions: []int{500, 210, 720}},
	}
	farm := &Farm{Name: "Fazenda", Fields: fieldsFazenda}
	log.Printf("Talhão 0: %.2f", farm.Fields[0].Productivity())
	log.Printf("Talhão 1: %.2f", farm.Fields[1].Productivity())
	log.Printf("Fazenda: %.2f", farm.Productivity())
}
