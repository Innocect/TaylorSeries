package main

import (
	"fmt"
	"os"
	"strconv"
	"taylor/series"
)

func main() {

	ts := series.InitTaylorSeries()

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Print(fmt.Errorf("enter a valid value of "))
	}
	x, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Print(fmt.Errorf("enter a valid value of "))
	}

	ans := ts.Calculate(uint64(n), uint64(x), nil)

	fmt.Println(ans)
}
