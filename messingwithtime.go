package main

import(
	"fmt"
	"bufio"
	"os"
	"time"
	// "strings"
	// "strconv"
)

func main(){
	fmt.Println("Please enter a date (MM/DD/YYYY): ")
	reader := bufio.NewReader(os.Stdin)
	date,_ := reader.ReadString('\n') 
	
	//Old Way
	// fmt.Println("You typed in " + date)
	// month,_ := strconv.Atoi(strings.Split(date,"/")[0])
	// day,_ := strconv.Atoi(strings.Split(date,"/")[1])
	// year,_ := strconv.Atoi(strings.Split(date,"/")[2])
	
	// newDate := time.Date(year,time.Month(month),day,0,0,0,0,time.UTC)
	// fmt.Println(newDate.Weekday())

	//New Way
	myDay, _ := time.Parse(date,"")
	fmt.Println(myDay.Weekday())

}