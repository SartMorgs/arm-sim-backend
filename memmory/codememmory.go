package memmory

import (
	"strconv"
	"encoding/json"
	"strings"
	//"reflect"
)

type CodeMemmory struct{
	romList map[string]*Memmory
	functionMap map[string]map[string]map[string]string
	inputCode map[string]interface{}
	directiveCode map[string]interface{}
}

func NewCodeMemmory() *CodeMemmory{
	codememmory := new(CodeMemmory)
	codememmory.romList = make(map[string]*Memmory)

	var addressMemmoryCode string

	// Interruption Area
	for countInt := 1; countInt < 17; countInt++{
		for count := ((countInt - 1) * 256); count < (countInt * 256); count++{
			addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
			codememmory.romList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
			codememmory.romList[addressMemmoryCode].SetAliasField("Int" + strconv.Itoa(countInt))
		}
	}

	for count := (4 * 1024); count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		codememmory.romList[addressMemmoryCode] = NewMemmory(addressMemmoryCode)
	}

	codememmory.functionMap = nil 
	codememmory.inputCode = nil
	codememmory.directiveCode = nil

	return codememmory
}

func (cdmem *CodeMemmory) AddAliasMemmoryField(addr string, alias string){
	cdmem.romList[addr].SetAliasField(alias)
}

func (cdmem *CodeMemmory) AddInstructionField(addr string, inst string){
	cdmem.romList[addr].SetFullBinValue(inst)
}

func (cdmem *CodeMemmory) ResetCodeMemmory(){
	var addressMemmoryCode string
	for count := 0; count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		cdmem.romList[addressMemmoryCode].SetFullBinValue("00000000000000000000000000000000")
	}
}

func (cdmem *CodeMemmory) GetRomList() map[string]*Memmory{
	return cdmem.romList
}

func (cdmem *CodeMemmory) GetCodeMemmoryJson() string{
	code_memmory := make(map[int]string)

	var addressMemmoryCode string
	for count := 0; count < (12 * 1024); count++{
		addressMemmoryCode = "0x" + strconv.FormatInt(int64(count), 16)
		str_memmory := map[string]string{
			"decimal_value": strconv.FormatInt(int64(cdmem.romList[addressMemmoryCode].GetDecValue()), 10),
			"hexadecimal_value": cdmem.romList[addressMemmoryCode].GetHexValue(),
			"binary_value": cdmem.romList[addressMemmoryCode].GetBinValue(),
			"full_binary": cdmem.romList[addressMemmoryCode].GetFullBinValue(),
			"config_type": cdmem.romList[addressMemmoryCode].GetConfigType(),
			"memmory_address": cdmem.romList[addressMemmoryCode].GetAddress(),
			"alias_field": cdmem.romList[addressMemmoryCode].GetAliasField()}

		jmem, _ := json.MarshalIndent(str_memmory, "", "")

		code_memmory[count] = string(jmem)
	}

	jcode_memmory, _ := json.MarshalIndent(code_memmory, "", "")

	return string(jcode_memmory)
}

func (cdmem *CodeMemmory) SetInputCode(code map[string]interface{}){
	cdmem.inputCode = code
}

func (cdmem *CodeMemmory) GetInputCode() map[string]interface{}{
	return cdmem.inputCode
}

func (cdmem *CodeMemmory) SetDirectiveCode(directive map[string]interface{}){
	cdmem.directiveCode = directive
}

func (cdmem *CodeMemmory) GetDirectiveCode() map[string]interface{}{
	return cdmem.directiveCode
}

func (cdmem *CodeMemmory) BuildFunctionMap() map[string]map[string]map[string]string{
	cdmem.functionMap = make(map[string]map[string]map[string]string)

	start_address := 0
	end_address := 0

	code_block := make([]string, 0, len(cdmem.inputCode))
	for k := range cdmem.inputCode {
		code_block = append(code_block, k)
	}

	cdmem.functionMap["interruptions"] = make(map[string]map[string]string)
	cdmem.functionMap["functions"] = make(map[string]map[string]string)

	for item := 0; item < len(code_block); item++{
		if strings.Contains(code_block[item], "INT") && strings.Contains(code_block[item], "_Handler"){
			size_interruption_block := len(cdmem.inputCode[code_block[item]].([]interface{})) - 1

			cdmem.functionMap["interruptions"][code_block[item]] = make(map[string]string)
			cdmem.functionMap["interruptions"][code_block[item]]["start"] = "0x0"
			cdmem.functionMap["interruptions"][code_block[item]]["end"] = "0x" + string(strconv.FormatInt(int64(size_interruption_block), 16))
		} else if code_block[item] == "main" {
			size_main_block := len(cdmem.inputCode[code_block[item]].([]interface{})) - 1

			cdmem.functionMap["functions"][code_block[item]] = make(map[string]string)
			cdmem.functionMap["functions"][code_block[item]]["start"] = "0x0"
			cdmem.functionMap["functions"][code_block[item]]["end"] = "0x" + string(strconv.FormatInt(int64(size_main_block), 16))
		
			end_address = end_address + size_main_block
		} else{
			size_function_block := len(cdmem.inputCode[code_block[item]].([]interface{}))
			start_address = end_address + 1
			end_address = end_address + size_function_block

			cdmem.functionMap["functions"][code_block[item]] = make(map[string]string)
			cdmem.functionMap["functions"][code_block[item]]["start"] = "0x" + string(strconv.FormatInt(int64(start_address), 16))
			cdmem.functionMap["functions"][code_block[item]]["end"] = "0x" + string(strconv.FormatInt(int64(end_address), 16))
		}
	}

	return cdmem.functionMap
}

func (cdmem *CodeMemmory) FixFunctionAddressIntoMainCode(){

	code_block := make([]string, 0, len(cdmem.inputCode))
	for k := range cdmem.inputCode {
		code_block = append(code_block, k)
	}

	//print(cdmem.inputCode[code_block[0]].([]interface{})[0])

	for item := 0; item < len(code_block); item++{
		for instruction_position := 0; instruction_position < len(cdmem.inputCode[code_block[item]].([]interface{})); instruction_position++{
			instruction := cdmem.inputCode[code_block[item]].([]interface{})[instruction_position].(string)
			if strings.Contains(instruction, "F"){
				instruction_splited_list := strings.Split(instruction, "F")
				//print(instruction_splited_list[0] + " " + instruction_splited_list[1] + " " + instruction_splited_list[2])
				func_reference_integer, _ := strconv.ParseInt(instruction_splited_list[1], 10, 64)
				function_start_addr_hex := cdmem.functionMap["functions"][code_block[func_reference_integer]]["start"]
				function_start_addr_bin, _ := strconv.ParseInt(function_start_addr_hex[2:], 16, 64)
				function_bin := string(strconv.FormatInt(function_start_addr_bin, 2))

				for iterator := 0; len(function_bin) < 26; iterator++{
					function_bin = "0" + function_bin
				}

				instruction = instruction_splited_list[0] + function_bin
				cdmem.inputCode[code_block[item]].([]interface{})[instruction_position] = instruction
			}
		}
	}
}

func (cdmem *CodeMemmory) AddCodeIntoMemmory(){
	code_block := make([]string, 0, len(cdmem.inputCode))
	for k := range cdmem.inputCode {
		code_block = append(code_block, k)
	}

	var addressMemmoryCode string
	var interruption_number int64
	var start_address int64 
	var int_addr int64
	//var end_address int64
	for item := 0; item < len(code_block); item++{
		if code_block[item] == "main"{
			int_addr, _ = strconv.ParseInt(cdmem.functionMap["functions"]["main"]["start"][2:], 10, 64)
			start_address = (4 * 1024) + int_addr
			//end_address, _ = strconv.ParseInt(cdmem.functionMap["functions"]["main"]["end"][2:], 10, 64)
		} else if strings.Contains(code_block[item], "INT") && strings.Contains(code_block[item], "_Handler"){
			if len(code_block[item]) <= 11{
				interruption_number, _ = strconv.ParseInt(string(code_block[item][3]), 10, 64)
			} else{
				interruption_number, _ = strconv.ParseInt(string(code_block[item][3:4]), 10, 64)
			}

			start_address = int64(interruption_number * 256)
			//end_address = int64(((interruption_number + 1) * 256) - 1)
		} else{
			int_addr, _ = strconv.ParseInt(cdmem.functionMap["functions"][code_block[item]]["start"][2:], 10, 64)
			start_address = (4 * 1024) + int_addr
			//end_address, _ = strconv.ParseInt(cdmem.functionMap["functions"][code_block[item]]["end"][2:], 10, 64)
		}

		for instruction_position := 0; instruction_position < len(cdmem.inputCode[code_block[item]].([]interface{})); instruction_position++{
			instruction := cdmem.inputCode[code_block[item]].([]interface{})[instruction_position].(string)
			addressMemmoryCode = "0x" + strconv.FormatInt(int64(instruction_position) + start_address, 16)

			cdmem.AddInstructionField(addressMemmoryCode, instruction)
		}
	}
}

func (cdmem *CodeMemmory) AddDirectivesAliasIntoMemmory(){
	directive_amount := len(cdmem.directiveCode)

	var directive_string_item string
	for directive_item := 0; directive_item < directive_amount; directive_item++{
		directive_string_item = strconv.FormatInt(int64(directive_item+1), 10)
		if cdmem.directiveCode[directive_string_item].([]interface{})[1].(string) == "EQU"{
			addr := cdmem.directiveCode[directive_string_item].([]interface{})[2].(string)
			alias := cdmem.directiveCode[directive_string_item].([]interface{})[0].(string)
			cdmem.AddAliasMemmoryField(addr, alias)
		}
	}
}