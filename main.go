package main

import (
	"fmt"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	form := tview.NewForm().
		AddDropDown("Start Search from The Year \n", []string{"2017", "2018", "2019", "2020", "2021", "2022"}, 0, nil).
		AddDropDown("Search from The Month", []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}, 0, nil).
		AddDropDown("Search from The Day of the Month", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}, 0, nil).
		AddDropDown("Search From The Add Hour", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24"}, 0, nil).
		AddDropDown("Search Till The Year", []string{"2017", "2018", "2019", "2020", "2021", "2022"}, 0, nil).
		AddDropDown(" Search Till The Month", []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}, 0, nil).
		AddDropDown("Search Till The Day of the Month", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}, 0, nil).
		AddDropDown("Search Till The Add Hour", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24"}, 0, nil).
		//	AddButton("Save", nil).
		AddButton("Save", func() {
			app.Stop()
		})
	form.SetBorder(true).SetTitle("Enter Date to Search").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
	_, a := form.GetFormItem(0).(*tview.DropDown).GetCurrentOption()
	//fmt.Printf(b)
	_, b := form.GetFormItem(1).(*tview.DropDown).GetCurrentOption()
	var k, l int
	switch {
	case b == "January":
		k = 01
	case b == "February":
		k = 02
	case b == "March":
		k = 03
	case b == "April":
		k = 4
	case b == "May":
		k = 5
	case b == "June":
		k = 6
	case b == "July":
		k = 7
	case b == "August":
		k = 8
	case b == "September":
		k = 9
	case b == "October":
		k = 10
	case b == "November":
		k = 11
	case b == "December":
		k = 12
	}

	_, c := form.GetFormItem(2).(*tview.DropDown).GetCurrentOption()
	//fmt.Printf(c)

	_, d := form.GetFormItem(3).(*tview.DropDown).GetCurrentOption()
	//fmt.Printf(d)

	_, e := form.GetFormItem(4).(*tview.DropDown).GetCurrentOption()
	//fmt.Printf(e)

	_, f := form.GetFormItem(5).(*tview.DropDown).GetCurrentOption()
	//fmt.Printf(f)

	switch {
	case f == "January":
		l = 1
	case f == "February":
		l = 2
	case f == "March":
		l = 3
	case f == "April":
		l = 4
	case f == "May":
		l = 5
	case f == "June":
		l = 6
	case f == "July":
		l = 7
	case f == "August":
		l = 8
		l = 9
	case f == "October":
		l = 10
	case f == "November":
		l = 11
	case f == "December":
		l = 12
	}

	_, g := form.GetFormItem(6).(*tview.DropDown).GetCurrentOption()
	_, h := form.GetFormItem(6).(*tview.DropDown).GetCurrentOption()
	//fmt.Printf(g)

	fmt.Println(a, "-", k, "-", c, "-", d, "to", e, "-", l, "-", g, "-", h)

}
