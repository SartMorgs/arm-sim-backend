package controller

import "strconv"

type ExecuteUnit struct{
	config map[string]int
}

func (eu *ExecuteUnit) ConfigForLdi(reg string, value string){
	x, _ := strconv.Atoi(reg)
	eu.config["register"] = x

	y, _ := strconv.Atoi(value)
	eu.config["value"] = y
}