package main

type Farm struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name        string
	Area        int
	Productions []int
}

func (f *Field) Productivity() float64 {
	totalProduction := 0
	for _, production := range f.Productions {
		totalProduction += production
	}
	productivity := float64(totalProduction) / float64(f.Area)
	return productivity
}

func (f *Farm) Productivity() float64 {
	productivity := 0.0
	totalArea := 0
	for _, field := range f.Fields {
		productivity += float64(field.Productivity()) * float64(field.Area)
		totalArea += field.Area
	}
	productivity = productivity / float64(totalArea)
	return productivity
}
