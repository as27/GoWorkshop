package main

type Sensor struct {
	ID     int     `json:"id"`
	Name   string  `json:"name,omitempty"`
	Values []Value `json:"values"`
}

type Value struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Timestamp int     `json:"timestamp"`
}

func main() {
	s := Sensor{
		1,
		"Sensor1",
		[]Value{
			Value{"Temperatur", 24, 201706010800},
			Value{"Temperatur", 28, 201706011200},
		},
	}
}
