package main

import (
	// "reflect"
	"fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
)

// Doctor declaration
type Doctor struct { 
	number int;
	actorName string;
	companions []string;
}

// Animal declaration
type Animal struct {
	// Name string `required max:"100"`;
	Origin string;
}

// Bird declarationb
type Bird struct {
	Animal;
	SpeedKPH float32;
	canFly bool;
}

func main() {
	// statePopulations := map[string]int{}
	// statePopulations = map[string]int {
	// 	"California": 12374994,
	// 	"Texas": 8743258,
	// 	"Florida": 8332024,
	// 	"New York": 53824462,
	// 	"Pennsylvania": 74580325,
	// 	"Illinois": 73294634,
	// 	"Ohio": 897657234,
	// }

	// if pop, ok := statePopulations["Florida"]; ok {
	// 	fmt.Println(pop)
	// }
	// fmt.Println(statePopulations);
	// statePopulations["Gorgia"] = 7467838
	// delete(statePopulations, "Ohio")
	// fmt.Println(statePopulations);

	// aDoctor := Doctor{
	// 	number: 3,
	// 	actorName: "John Wick",
	// 	companions: []string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }

	// myDoctor := struct{name string; age int}{name: "Oluwatobi", age: 31}
	// fmt.Println(myDoctor)

	// b := Bird{
	// 	Animal: Animal{Name: "Emu", Origin: "Lagos"},
	// 	SpeedKPH: 30.5,
	// 	canFly: false,
	// }
	// b.Name = "Emu";
	// b.Origin = "West-East";
	// b.SpeedKPH = 20.56;
	// b.canFly = true;

	// t := reflect.TypeOf(Animal{})
	// field, _ := t.FieldByName("Name")

	// fmt.Println(field.Tag)

	// number := 50
	// guess := 50

	// if number > guess {
	// 	fmt.Println("Too low")
	// }

	// if number < guess {
	// 	fmt.Println("Too High")
	// }

	// if number == guess {
	// 	fmt.Println("You got it!")
	// }

	// if false {
	// 	fmt.Println("I am a warrior")
	// }

	// s := []int{1, 2, 3}
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }

	// // control statement
	// // defer, panic, recover

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")


	// res, err := http.Get("https://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// robots, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // Pointer
	// var a int = 42;
	// var b *int = &a

	// a = 48

	// fmt.Println(a, *b)

	// var ms *myStruct
	// ms = new(myStruct)
	// ms.foo = 42
	// fmt.Println(ms.foo)

	greeting := "Hello";
	name := "Stacey";

	sayGreeting(&greeting, &name);
	fmt.Println(name)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", *s)

}

type myStruct struct {
	foo int
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result +=v
	}
	return &result
}