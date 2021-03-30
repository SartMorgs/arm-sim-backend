package controller

import "strconv"

type ExecuteUnit struct{
	config map[string]int64
}

func NewExecuteUnit() *ExecuteUnit{
	executeunit := new(ExecuteUnit)
	executeunit.config = make(map[string]int64)

	return executeunit
}

func (eu *ExecuteUnit) ConfigForLdi(reg string, value string){
	x, _ := strconv.ParseInt(reg, 2, 64)
	eu.config["register"] = x

	y, _ := strconv.ParseInt(value, 2, 64)
	eu.config["value"] = y
}

func (eu *ExecuteUnit) ConfigForLdr(reg_rt string, reg_rn string){
	x, _ := strconv.ParseInt(reg_rt, 2, 64)
	eu.config["register_rt"] = x

	y, _ := strconv.ParseInt(reg_rn, 2, 64)
	eu.config["register_rn"] = y
}

func (eu *ExecuteUnit) ConfigForStr(reg_rt string, reg_rn string, ram_addr string){
	x, _ := strconv.ParseInt(reg_rt, 2, 64)
	eu.config["register_rt"] = x

	y, _ := strconv.ParseInt(reg_rn, 2, 64)
	eu.config["register_rn"] = y

	z, _ := strconv.ParseInt(ram_addr, 2, 64)
	eu.config["ram_address"] = z
}

func (eu *ExecuteUnit) ConfigForAdds(reg_rd string, reg_rn string, reg_rm string, imm string){
	x, _ := strconv.ParseInt(reg_rd, 2, 64)
	eu.config["register_rd"] = x

	y, _ := strconv.ParseInt(reg_rn, 2, 64)
	eu.config["register_rn"] = y

	z, _ := strconv.ParseInt(reg_rm, 2, 64)
	eu.config["register_rm"] = z

	m, _ := strconv.ParseInt(imm, 2, 64)
	eu.config["immediate"] = m
}

func (eu *ExecuteUnit) ConfigForSubs(reg_rd string, reg_rn string, reg_rm string, imm string){
	x, _ := strconv.ParseInt(reg_rd, 2, 64)
	eu.config["register_rd"] = x

	y, _ := strconv.ParseInt(reg_rn, 2, 64)
	eu.config["register_rn"] = y

	z, _ := strconv.ParseInt(reg_rm, 2, 64)
	eu.config["register_rm"] = z

	m, _ := strconv.ParseInt(imm, 2, 64)
	eu.config["immediate"] = m
}

func (eu *ExecuteUnit) ConfigForAnds(reg_rd string, reg_rn string, reg_rm string){
	x, _ := strconv.ParseInt(reg_rd, 2, 64)
	eu.config["register_rd"] = x

	y, _ := strconv.ParseInt(reg_rn, 2, 64)
	eu.config["register_rn"] = y

	z, _ := strconv.ParseInt(reg_rm, 2, 64)
	eu.config["register_rm"] = z
}

func (eu *ExecuteUnit) ConfigForLdm(reg_rn string, reg_list []string){
	x, _ := strconv.ParseInt(reg_rn, 2, 64)
	eu.config["register_rn"] = x

	var registernumber string
	for count := 1; count <= len(reg_list); count++{
		registernumber = strconv.FormatInt(int64(count), 10)

		y, _ := strconv.ParseInt(reg_list[count - 1], 2, 64)

		eu.config["register_" + registernumber] = y
	}
}
