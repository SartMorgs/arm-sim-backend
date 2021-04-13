package program

import(
	"encoding/json"
)

type Step struct{
	step string
	code_memmory string
	data_memmory string
	device_memmory string
	register_bank string
}

type Simulator struct{
	prog Program
	steps map[int]Step
}

// Realizar simulação