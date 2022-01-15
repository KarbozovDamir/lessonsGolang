package main

// func sum(x, y int) int {
// 	return x + y
// }
// func partialSum(x int) func(int) int {
// 	return func(y int) int {
// 		return sum(x, y)
// 	}
// }
// func main() {
// 	partial := partialSum(3)
// 	fmt.Println(partial(7))
// }

//LESSON6 Closures Functions
// func main() {
// 	l := 3
// 	k := 4
// 	func() {
// 		ares := l + k
// 		fmt.Println(ares)
// 	}()
// }

// func intSeq() func() int {
// 	i := 1
// 	return func() int {
// 		i++
// 		return i
// 	}
// }
// func main() {
// 	next := intSeq()
// 	fmt.Println(next()) // 2
// 	fmt.Println(&next)  //0xc00000e028
// 	// second exempl
// 	next1 := intSeq()
// 	fmt.Println(next1()) // 2
// 	fmt.Println(&next)   //0xc00000e028
// }

//LESSON5 -How to iterate over an Array using for loop?
// func main() {
// 	intArray := [5]int{10, 20, 30, 40, 50}

// 	fmt.Println("\n---------------Example 1--------------------\n")
// 	for i := 0; i < len(intArray); i++ {
// 		fmt.Println(intArray[i])
// 	}

// 	fmt.Println("\n---------------Example 2--------------------\n")
// 	for index, element := range intArray {
// 		fmt.Println(index, "=>", element)

// 	}

// 	fmt.Println("\n---------------Example 3--------------------\n")
// 	for _, value := range intArray {
// 		fmt.Println(value)
// 	}

// 	j := 0
// 	fmt.Println("\n---------------Example 4--------------------\n")
// 	for range intArray {
// 		fmt.Println(intArray[j])
// 		j++
// 	}
// }

// LESSON4 - Map
// func main() {
// 	var employee = make(map[string]int)
// 	employee["Mark"] = 10
// 	employee["Sandy"] = 20
// 	employee["Rocky"] = 30
// 	employee["Josef"] = 40
// 	fmt.Println(employee)

//How delet of map
//delete(employee, "Mark")

// Empty Map
// 	employeeList := make(map[string]int)
// 	fmt.Println(len(employee))
// 	fmt.Println(len(employeeList))
// }

//LESSON3 - Anonymous
// func sayHi(mssg string) {
// 	fmt.Println(mssg)
// }
// func main() {
// 	// regular function
// 	// sayHi("Hello There")

// 	//anonymous function
// 	func(mssg string) {
// 		fmt.Println(mssg)
// 	}("ANONYMOUS")
// }

// func main() {
// 	func(l int, b int) {
// 		fmt.Println(l * b)
// 	}(20, 30)
// }

// LESSON 2 Methods

// type Person struct {
// 	firstname string // name string
// 	lastname  string
// 	age       int
// 	hobbyone  hobby // hobbies []hobby
// 	hobbytwo  hobby
// 	hobby     //наследование from type hobby all atributes and methods
// }

// func (p Person) Speak() { // var p Person
// 	fmt.Printf("Hello, my name is %v\n", p.firstname)
// }

// func (p *Person) SetName(name string) { // var p *Person = &Damir
// 	p.firstname = name
// }

// func (p Person) SetWrongName(name string) { // var p Person = Damir
// 	p.firstname = name
// }

// type hobby struct {
// 	name string
// 	like string
// }

// func main() {
// 	// var p Person
// 	var damir Person
// 	damir.name = "football"
// 	damir.hobbyone.name = "valeyball"
// 	damir.hobbytwo.name = "tennis"
// 	damir.SetWrongName("damir")
// 	damir.SetName("damir")
// 	damir.Speak()
// 	// Alika new() => var alika *Person= &Person{}
// 	alika := new(Person)
// 	alika.SetName("alika")
// 	alika.Speak()

// 	// Ahat
// 	var ahat *Person = &Person{}
// 	ahat.SetName("ahat")
// 	ahat.Speak()

// 	//Almas
// 	almas := &Person{}
// 	almas.SetName("almas")
// 	almas.Speak()

// 	//Faiz
// 	faiz := Person{}
// 	faiz.SetName("faiz")
// 	faiz.Speak()
// 	fmt.Println(*alika, damir, *ahat, *almas, faiz) // alika = 0x01              ahat = 0x02
// }

// LESSON 1

// type book struct {
// 	name    string
// 	zhanra  string
// 	date    int
// 	opened  bool     // []string                a = append(a, "asdasd") authors = append(authors, author{})
// 	authors []author // 0
// }

// type author struct {
// 	name    string // [] {damir, ali}
// 	age     int
// 	like    []string // [] {football, chess, valeyball}
// 	dislike []string
// }

// func main() {
// 	// var books book
// 	// book1 := book{name: "Pak", zhanra: "comedy", date: 1986, opened: true}

// 	var book2 book

// 	book2.name = "harry potter"
// 	book2.zhanra = "comedy"
// 	book2.date = 1758
// 	book2.opened = true

// 	// author1
// 	var author1 author // var a string
// 	// author1 := author{} 2 var obyavlen
// 	author1.name = "ahat" // a = "hello"

// 	// ...
// 	book2.authors = append(book2.authors, author1) // var b []string b = append(b, a)

// 	// author2
// 	book2.authors = append(book2.authors, author{ // b = append(b, "world")
// 		name: "ali",
// 	})

// 	// // author3
// 	a := author{
// 		name:    "ali3",
// 		age:     21,
// 		like:    []string{"football", "golang"},
// 		dislike: []string{"PHP", "JS"},
// 	}
// 	book2.authors = append(book2.authors, a)

// 	fmt.Println(book2)
// }

// Example with Student

// type Student struct {
// 	name        string
// 	age         string
// 	personality []personality
// 	speciality  []speciality
// }

// type personality struct {
// 	lastName string
// 	score    string
// 	gpa      string
// 	degree   string
// }
// type speciality struct {
// 	name    []string
// 	subject []string
// 	hours   string
// }

// func main() {
// 	var person1 Student
// 	person1.name = "Name: Alina,"
// 	person1.age = "Age 21"

// 	id1 := personality{
// 		lastName: "lastName: Mao,",
// 		score:    "score: 95,",
// 		gpa:      "GPA: 3.9,",
// 		degree:   "Degree: master",
// 	}

// 	id2 := speciality{
// 		name:    []string{"Speciality: programmer,", "translater"},
// 		subject: []string{"Subjects: math,", "lang, sport"},
// 		hours:   "hours 36",
// 	}

// 	person1.personality = append(person1.personality, id1)
// 	person1.speciality = append(person1.speciality, id2)
// 	fmt.Println(person1)
// }

// type people struct {
// 	name    string
// 	age     int
// 	gender  string
// 	hobbies hobby
// }

// type hobby struct {
// 	like    []string
// 	dislike []string
// }

// func main() {
// 	var person1 people
// 	person2 := people{name: "ali", age: 26, gender: "M"}
// 	person2.hobbies.like = append(person2.hobbies.like, "C/C++")
// 	a := []string{"python"}
// 	person2.hobbies.dislike = a

// 	person1.name = "damir"
// 	person1.age = 32
// 	person1.gender = "M"
// 	person1.hobbies.like = []string{"football", "golang"}
// 	person1.hobbies.dislike = []string{"java", "c++"}

// 	fmt.Println(person1)
// 	fmt.Println(person2)
// }

// }

// type myInt int

// var ali myInt

// // int, string, rune, byte, float, complex
// func main() {
// 	ali = 3
// 	var a int = 2
// 	res := myInt(a) + ali
// 	fmt.Println(res)
// }
