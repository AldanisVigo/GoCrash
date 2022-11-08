package main

import(
	"fmt" 
	// "bufio" 
	// "os"
	// "log"
	// "reflect"
	// "strconv"
	// "strings"
	// "unicode/utf8"
	// "time"
	// "math/rand"
)

var pl = fmt.Println

//Comments
/*
	Multiline comments
*/

// func main(){
// 	pl("What is your name?")
// 	reader := bufio.NewReader(os.Stdin)
// 	name, err := reader.ReadString('\n')
// 	if err == nil {
// 		pl("Hello", name)
// 	}else{
// 		log.Fatal(err)
// 	}
// }


func main() {
	//Variables
	// var vName string = "Aldanis Vigo"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "hello"
	// v4 := 2.4
	// v4 = 5.4


	//Data Types
	// pl(reflect.TypeOf(25))
	// pl(reflect.TypeOf(3.14))
	// pl(reflect.TypeOf(true))
	// pl(reflect.TypeOf(25))
	// pl(reflect.TypeOf("Hello"))
	// pl(reflect.TypeOf("😃"))

	//Casting
	// cv1 := 1.5
	// cv2 := int(cv1)
	// pl(cv2)

	// //string to int
	// cv3 := "50000000"
	// cv4, err := strconv.Atoi(cv3)
	// pl(cv4, err, reflect.TypeOf(cv4))
	
	// //int to string
	// cv5 := 500000000
	// cv6 := strconv.Itoa(cv5)
	// pl(cv6,err, reflect.TypeOf(cv6))


	// //string to float
	// cv7 := "3.14"
	// if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {
	// 	pl(cv8)
	// }

	// //float to string
	// cv9 := fmt.Sprintf("%f",3.14)
	// pl(cv9, reflect.TypeOf(cv9))

	// //The if conditional
	// iAge := 24
	// if(iAge >= 1) && (iAge <= 18) {
	// 	pl("Important birthday")
	// }else if (iAge == 21) || (iAge == 50) {
	// 	pl("Important birthday")
	// }else if (iAge >= 65) {
	// 	pl("Important birthday")
	// } else {
	// 	pl("Not an important birthday")
	// }



	
	// //Strings
	// sv10 := "A word"
	// //Arrays of bytes

	// //replacing parts of string
	// replacer := strings.NewReplacer("A","Another")
	// sv11 := replacer.Replace(sv10)
	// pl(sv11)

	// //printing the length of a string
	// pl("Length : ", len(sv11))
	
	// //Check if a string contains another string
	// pl("Contains Another :", strings.Contains(sv11, "Another"))

	// //Get the first index match for letter "o"
	// pl("o index:", strings.Index(sv11,"o"))

	// //Replace all matches of the letter "o" with a 0
	// pl("Replace: ", strings.Replace(sv11,"o","0",-1))

	// //Remove white spaces
	// sv12 := "\nSome Words\n" 
	// sv12 = strings.TrimSpace(sv12) //Get rid of whitespaces

	// //Splitting by delimeter
	// pl("Split :", strings.Split("a-b-c-d", "-"))
	
	// //Convert string to lowercase
	// pl("Lower: ", strings.ToLower(sv12))

	// //Convert string to uppercase
	// pl("Upper: ", strings.ToUpper(sv12))

	// //Check if string starts with another 
	// pl("Prefix : ", strings.HasPrefix("tacocat","taco"))

	// //Check if a string ends with another
	// pl("Suffix : ", strings.HasSuffix("tacocat","taco"))

	//----------------------------------------------------------
	//Runes - Unicodes that represent characters
	// rStr := "abcdefg"
	// pl("Rune count: ", utf8.RuneCountInString(rStr)) 
	// for i, runeVal := range rStr {
	// 	fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	// }

	//Working with times

	// //current time
	// now := time.Now()
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())

	//Arithmetic operations
	// pl("5 + 4 =", 5 + 4)
	// pl("5 - 4 =", 5 - 4)
	// pl("5 * 4 =", 5 * 4)
	// pl("5 / 4 =", 5 / 4)
	// pl("5 % 4 =", 5 % 4)
	
	//shorthand increment
	// mInt := 1
	// mInt = mInt + 1
	// mInt += 1
	// mInt++
	// pl(mInt)


	//Floating value percision
	// pl("Float Precision = ", 0.11111111111111111111 + 0.11111111111111111111)

	//Creating random values
	//Get a random seed for random num generator
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) //Up to 50 but not including 50
	// pl("Random Number : ", randNum)


	//Math functions
	// pl("Abs(-10) = ", math.Abs(-10))
	// pl("Pow(4, 2) = ", math.Pow(4,2))
	// pl("Sqrt(16) = ", math.Sqrt(16))
	// pl("Cbrt(8) = ", math.Cbrt(8))
	// pl("Ceil(4.4) = ", math.Ceil(4.4))
	// pl("Floor(4.4) = ", math.Floor(4.4))
	// pl("Round(4.4) = ", math.Round(4.4))
	// pl("Log2(8) = ", math.Log2(8))
	// pl("Log10(100) = ", math.Log10(100))
	// pl("Log(7.389) = ", math.Log(math.Exp(2)))
	// pl("Max(5,4) = ", math.Max(5,4))
	// pl("Min(5,4) = ", math.Min(5,4))

	//Convert degrees to radians
	// r90 := 90 * math.Pi / 180
	// d90 := r90 * (180 / math.Pi)
	// fmt.Printf("%.2f radians = %.2f degrees\n", r90, d90)
	// pl("Sin(90) =", math.Sin(r90))

	//Formatted print

	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value
	
	// fmt.Printf("%s %d %c %f %t %o %x\n","Stuff", 1, 'A', 3.14, true, 1, 1)
	// fmt.Printf("%9f\n", 3.14)
	// fmt.Printf("%.2f\n", 3.141592)
	// fmt.Printf("%9.f\n", 3.141592)

	// sp1 := fmt.Sprintf("%9.f\n",3.141592)
	// pl(sp1)


	//For loops
	//for initialization; condition;
	//postStatement { BODY }
	// for x := 5; x >= 1; x-- {
	// 	//x only exists in the scope of this for loop
	// 	pl(x)
	// }

	// //while loops
	// fX := 0
	// for fX < 5 {
	// 	pl(fX)
	// 	fX++
	// }

	//Guessing game
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1
	// for true {
	// 	fmt.Print("Guess a number between 0 and 50 : ")
	// 	pl("Randnum is ", randNum)
	// 	reader := bufio.NewReader(os.Stdin)
	// 	guess, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	guess = strings.TrimSpace(guess)
	// 	iGuess, err := strconv.Atoi(guess)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if iGuess > randNum {
	// 		pl("Pick a lower value")
	// 	}else if iGuess < randNum {
	// 		pl("Pick a higher value")
	// 	}else {
	// 		pl("You guessed it.")
	// 		break
	// 	}

	// }


	//Cycling through an array with range
	// aNums := []int{1,2,3}
	// for index, num := range aNums {
	// 	pl(index,num)
	// }


	//Arrays
	//collection of values with the same data type
	//Size of an array is fixed
	// var arr1 [5]int
	// arr1[0] = 1
	// arr2 := [5]int{1,2,3,4,5}
	// pl("Index 0 : ", arr2[0])
	// pl("Array Length : ", len(arr2))

	// //Iterating through an array
	// for i := 0; i < len(arr2); i++ {
	// 	pl(arr2[i])
	// }

	// for i, v := range arr2 {
	// 	fmt.Printf("%d : %d", i, v)
	// }

	// //Multidimensional arrays
	// arr3 := [2][2]int {
	// 	{1,2},
	// 	{3,4},
	// }

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		pl(arr3[i][j])
	// 	}
	// }
}
