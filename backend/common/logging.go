package common

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
)

func LogFatal(tag string, message string) {
	color.Red(fmt.Sprintf("[%s] : %s", tag, message))
	os.Exit(1)
}

func LogPrintln(tag string, message string) {
	log.Println(fmt.Sprintf("[%s] : %s", tag, message))
}

func Banner() {
	var banner = `
$$\   $$\            $$\                                   
$$$\  $$ |           $$ |                                  
$$$$\ $$ | $$$$$$\ $$$$$$\    $$$$$$\  $$$$$$$\   $$$$$$$\ 
$$ $$\$$ |$$  __$$\\_$$  _|   \____$$\ $$  __$$\ $$  _____|
$$ \$$$$ |$$ /  $$ | $$ |     $$$$$$$ |$$ |  $$ |\$$$$$$\  
$$ |\$$$ |$$ |  $$ | $$ |$$\ $$  __$$ |$$ |  $$ | \____$$\ 
$$ | \$$ |\$$$$$$  | \$$$$  |\$$$$$$$ |$$ |  $$ |$$$$$$$  |
\__|  \__| \______/   \____/  \_______|\__|  \__|\_______/  
`
	fmt.Println(fmt.Sprintf("%s\nv.Alpha", banner))
}
