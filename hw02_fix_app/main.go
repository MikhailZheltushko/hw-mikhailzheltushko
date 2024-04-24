package main

import (
	"fmt"

	"github.com/MikhailZheltushko/hw-mikhailzheltushko/hw02_fix_app/printer"
	"github.com/MikhailZheltushko/hw-mikhailzheltushko/hw02_fix_app/reader"
	"github.com/MikhailZheltushko/hw-mikhailzheltushko/hw02_fix_app/types"
)

func main() {
	var path string

	fmt.Printf("Enter data file path: ")
	fmt.Scanln(&path)

	var err error
	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
