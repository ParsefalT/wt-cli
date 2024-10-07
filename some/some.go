package some

import (
	"fmt"

	"github.com/fatih/color"
)

func Success(val string) {
	color.Green(val)
}
func Exit(val string) {
	color.Red(val)
}
func Info(val map[int]string) {
	if len(val) == 0 {
		fmt.Println("not tasks")
	} else {
		for key, value := range val {
			fmt.Printf("%d - %s\n", key,value)
		}
	}
}
func Delete(val string) {
	color.Yellow(val)
}