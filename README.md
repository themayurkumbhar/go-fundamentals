# Go - Funtamentals

This repository contains the go-funtamentals to understand the basic contructs of go-lang.

## Go features :bulb:

* Modern & clean
* System language
* Wide range of usecases
* Strongly typed
* Compiled
* Concurrency support
* Cross-platform support

----------


## Go Modules, Packages and Source-files Basics :spiral_notepad:

Installing go lang is very easy, just go to [go install](https://go.dev/doc/install) page and follow instructions.

* package structure:  

  ```go
  module
  |--- package
      |--- source-files
  ```
* Create directory to start using as `go module` which will contain all the srouce code related to that module.
* e.g. `mkdir go-fundamentals` this is my go module for developement.
* To init the above directory to be work as go module you will need to initialize with module initialization command.
* Use command `go mod init {optional/package/name}` this will create [go.mod](go.mod) file with module name and go versoin used to create the module.

* Create new directory to start using as `go package` which will logically group your source code files.
* e.g. `mkdir hello-world` this is my go package which will have code related to hello-world.
* Inside the `hello-world` folder you can have source files as part of hello-world package.  
* e.g. you can refer to [hello-world/hello.go](hello-world/hello.go) file is source code file containing the source code related to hello-world package.

## Build, Run, Install Basics :wrench:

Go provide different ways to bundle, run application. Depends upon the requirement you can use the any command to either `run` `build` or `install` your application.

* `go run {filename}` will allow you to run the go source file directly.
  * it will build the application at `temp location` and then `run` it and `clean` the file 
* `go build <optional|package>` will create the `executable file` depends on the which enviorment you are building the source code.
  * by default takes the current package directory name, you can aleaways specify the name of file
* `go install <optional|packagename>` will `build` and place the executable file in your `$HOME` directory. 

### 1. Hello World :smile:

* Creating your first `Hello World` go code, first follow the packaging instruction above.
* Now you have package with `hello.go` file created, to write you first code refer to [hello.go](hello-world/hello.go) file.
* To import any existing package use `import {package-name}` e.g. `import fmt` so we can use the fuctions present in the package.

### 2. Variables :unlock: and Constants :lock:

* Declaring variables
  * Any variable declared outside the function at package level must do it with `var` keyword. No non declartion statements at package level are allowed.
  * Varibale names should
    * **start** with `letter`
    * **not** be `Go Keyword`
    * **not** have any `spaces`
    * **not** have any `special character`
  * Good practices:
    * use meaningful names
    * keep short
    * can use `camelCase`
> Keep first character of variable `lowercase`, else making `capitalletter` will make that variable visible outside the package it defined.
  * Example: 
    ```go
    var (
    name    string
    course  string
    module, clip = 2, 1
    )
    ```
  * Scope of above declared vars is global in package.
  * Variable declaration inside the function
    * `var description string` this is if you dont have any default assignment to set variable
    * `description := "this is go fundamental course"` you can set the value directly to initialize.
    * operator ` := ` is short assignment operator used to assign the value to vars.

> You can declare varibale at `package` level and not use is allowed. but if you declare variable inside the `function` you have to use the variable.

  * Values and Pointers
    * Go passes arguments `by values` and `not by references`
    * When you pass the var directly to any function its pass by value ` len(title) ` 
      * To get memory address, you have to use `&` sign like ` fmt.Println(&title) ` will give something like `0xc000088250` 
    * Go provides `pointers` to point to other variables
      * Pointer variable e.g. `var ptr *string = &title` here `*` make the variable `ptr` a pointer var pointing to memory address of `title` var
      * `*ptr` will yeid value is hold by the variable pointed to which `var ptr` is pointing.
    * Pass by **Reference**
      * To pass the value as reference to function you need to use the `&variable` and in function you need to use the pointer variabel to reference the var like `variable *type`
      * then you can use `*variable = {new valuue}` to update the variable passed by reference which will actually update the value in original variable.
    * Constants in go
      * Contants can be created using keyword `const` in go
      * for example `const speedOfLight = 186000` 
      * Once create you can not change the value of constant.
    * Enviorment Variables
      * You can access the enviorment vars using package `os` in go
      * which will allow you to read env vars.


### 3. Functions in Go (basics)

* why do you functions? **code reusability!**
  * `input` -> `Function` -> `output`
  * Syntax of functions in Go
  ```go
  func functionName(params, type) returnType {
    <code>
  }
  ```
  * above code snippet says, `func` is keyword to declare function followed by `name` of fuction. which takes `parameters` with its `type` then followed by the `return type` of function.
  * Example:
  ```go
    func updateCourse(course string) course string {
      course = "new course"
      return course
    }
  ```
  * You can refer to [convert.go](functions/converter.go) full running source file.
* Funtion with **Varying** number of parameters
  * To create functoin which accepts varying params
    ```go
    func varyingfuncName(param ...type) returnType {
      <code>
    }
    ```
  * Important part is `paramName ...type` is says it accepts the varying number of parameters. you can refer [varying-params-fuction.go](functions/varying-params-function.go)


### 4. Conditionals in Go (basics)

* the `if` syntax evaluates the condition and executes code based on it.
  * conditional if else in go
  
  ```go
  if userAge > 18 {
    <code>
  }
  ```
  * conditional `else` with if
  
  ```go
  if userAge > 18 {
    <code>
  } else {
    <code>
  }
  ```
  * if else with nesting of `if-else` conditions
  
  ```go
  if userAge > 18 {
    <code>
  } else if userAge > 15 {
    <code>
  } else {
    <code>
  }
  ```
  * you can also declare the vars in if scope which only avaulable to scope of if else
  ```go
  if userAge, requiredAge := 15, 18; userAge > requiredAge {
    <code>
  } else {
    <code>
  }
  ```
  * You can refere to [if-example.go](conditionals/if-example.go) to running source code.
* Switch-case statements in go
  * Syntax of `switch` statement.
  ```go
  switch <simple-statment>; <expression> {
  case <value>: <code>
  case <value>: <code>
  default: <code>
  }
  ```
  * Go will not execute any statment after matching the case, so `only matched` case will be executed and if not match it will execute the `default` one.
  * Go will have default `break` for each case statement.
    * to override behaviour go provides `fallthrough` keyword.
    * which only applies to case statement its added, which will execute next case statememt only.
  * You can refer to runnning example of [switch-example.go](conditionals/switch-example.go).
  * Multiple conditional match in same `case` statement
    * go provides the way to match the multiple conditionals in same case so you can group behaviour and reduce multiple case statements.
    * example: [switch-multi-case.go](conditionals/switch-multi-case.go)
* If with Error handling
  * if is most common way to handle `error` in go
  * `error` is defined type is go lang.
  ```go
  _, err := os.Open("filename")
  if err != nil {
    <code_to_handle_error>
  }
  ```
  * You can refer to [if-with-errors.go](conditionals/if-with-errors.go) for running example


### 5. Loops in Go (Basics)

* GO has only one loop keyword `for` which has many way to configure and use as required.
* Syntax for loop
  ```go
  for <expression> {
    <code>
  }
  ```
  * here if `expression` is `true` that means it will loop until the expression evaluates to `true`
* Infinite `for` loop
  ```go
  for {
    <code>
  }
  ```
  * here if we dont specify any expression in for it will loop infinite times
* For loop with `boolean expression`
  * you can use `for {expression_evaluates_to|true|false}` to continue/break loop
  * example can be found at [for-loop.go](loops/for-loop.go)
  * You can also write the traditional for loop.
  ```go
  for i:=0; i<10; i++{
    <code>
  }
  ```
  * Example to run [for-loop.go](loops/for-loop.go)
* For Loop with `range expression`
  * You can iterate over the range of items using for loop
  ```go
  for index, item := range rangeVariable {
    <code>
  }
  ```
  * Execute the code for range loop [range-for-loop.go](loops/range-for-loop.go)
* Break and Continue statements in loop
  * `break` is used in loop to stop the execution of at given statement.
  * `continue` is used in loop to go back and continue to execute the loop statement.
  * Example for brak and continue [for-break-continue.go](loops/for-break-continue.go)

### 6. Arrays and Sclices

* for `lists` in go, all the data types stored in list must be of `same type`
* For `list of string` all the data should be `string` type, similarly for `numbers` it should br `int` type data
* **Arrays** : are like `lists`
  * they have `fix size`
* **Slices** : are like `arrays`
  * but they are `resizable` 
  * they are built top of the `array`
    * So basically no actual data is stored inside the slice, it is actually the `pointers` pointing to data stored inside the array
    * as they are pointers, dont store the data, its cheaper.
    * Initialization of empty slice
  ```go
  varName := make([]<type>, <length>, <capacity>)
  ```
    * You can have slice of slices which creates futher smaller pointers of slice
    ```go
    someSlice := []int{1, 2, 3, 4, 5, 6, 7}
    smallSlice := someSlice[2:7]
    ```
    * here we have to specify the `[start_index:end_index]` to create the slice, where `end_index` is `excluded` from the result
    * You can refer the more executable [arrays-slices](arrays-slices)
> this is very much similar concept of pythons list slicing. you can read more [here!](https://www.pythontutorial.net/advanced-python/python-slicing/)
  
  * How Slices are Dyamic (resizable)?
    * When the size of `slice` is completly filled, as it uses underlying data storage `array`, it will double the size of `array` and store the next element pointing to next position
    * You with `slice` you dont worry about the lenght of array that you need, you can use as resizable array data structure however in case of `array` you have to specify and create the size and have manually upgrade lenght do manipulations
    * Resizable better understanding execute [resizable.go](arrays-slices/resizable.go).
    
### 7. Maps in Go (basics)
  * `Map` is `key-value` type of data structure same as maps, hashTable in other languages. 
  * It stores data in `key -> value` format.
  * initialization of map
    * syntax
  ```go
  map[<key-type>]<value-type>
  ```
    * `<key-type>` must be **comaprable** data type, like it should work with `== and !=`
    * `key entry` must be **unique**.
  ```go
  someMapVar := make(map[string]int)
  ```
    * Here we have create empty map which will hold `string` keys and `int` values.
    * Execute [maps-basics.go](maps-in-go/maps-basics.go) for more examples.
    * Go maps Object will default return the data with `key ordered` way. like data is sorted based on `keys`
    * Retriving data from map using `for` loop will return data **randomly**.
    * Execute code [map-iteration.go](maps-in-go/map-iteration.go) to understand this.
  * add, update, delete and access value from the map
    * Adding new entry to map
      * Syntax `mapVar[<newKey>] = <newValue>`
    * Update entry in map
      * Syntax `mapVar[<existingKey>] = <newValue>`, if the key dose not exisits in `map` it will create new entry, else it will update the value.
    * Delete the entry in map
      * Syntax `delete(mapVar, <KeyToDelete>)` 
    * Access any value in map
      * Syntax `fmt.Println(mapVar[<mapKey>])`
  * Execute the [map-iteration.go](maps-in-go/map-iteration.go) to understand wokring
  * Map are `reference types` they will be passed by reference to `function`, so modifies actual data
    * so its cheap to work with as its only pointers are passed
  * they are not `thread safe`

### 8. Structs in Go (Basics)
  * `struct` is way of defining the custom data types in go.
  * For a basic `person` we can create struct like
  ```go
  type person struct {
    name string
    age int
    address string
  }
  ```
  * Object Oriented in GO
    * GO does not have `object` type 
    * Also it dose not have `class` keyword
    * No `inhearitance` support
  * Execute code [structs-basics.go](structs/structs-basics.go) to understand more.
  * You can access the `struct` fields using `.` with `var`
    * `person.name` or `person.age`
  * To update the structs fileds data
    * you can use `person.age = <new_age>` 

### 9. Concurrency in go (basics)
  * `concurrency` is multiple process executing *independently*
    * Lots of tasks *executing*, but only one *active*
> **Concurrency** is *dealing* with multiple things at once, **Parallelism** is *doing* lots of things at once
  * "Go does not use `threads` for concurrency :confused:	?
  * Go concurrency model
    * So go abstracts the threading model using the `goroutine` 
      * they are scheduled by the `go-rutime` not the `os`
      * `go-routines` way lighter than the `os` threads
      * Can grow and shrink as needed
      * GO manages all the `go-routines`
      * Dev dose not need to learn the os specific threads.
      * Fewer conetext switches
      * Faster startup time
    * Go uses design model called `Actor`  [Read More](https://www.oreilly.com/library/view/applied-akka-patterns/9781491934876/ch01.html)
    * Creating first `go-routine`
    ```go
    go func() {
      //come go-routine code
    }()
    ```
    * Here `go` keyword infornt of function will make this code block as `go-routine`
    * Go routine uses `channel` to share the data
      * two type of channels are present `buffered` & `unbuffered`
      * `unbuffered` channel is `synchronous` behaviour channel
        * creating channel is go
        ```go
        myChannel := make(chan int)
        ```
      * `buffered` channel is `asynchronous` behaviour channel
      ```go
        myChannel := make(chan int, 5) // this specifies that it can hold 5 ints in buffer
      ```
    * Execute your code related to go-routines [concurrency-go.go](concurrecy/concurrency-go.go)


----------

## Pull Requests :twisted_rightwards_arrows:

Feel Free to update the documentation and add more relavent contents. Thanks in advance. 🤟
