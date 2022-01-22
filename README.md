# Go - Funtamentals

This repository contains the go-funtamentals to understand the basic contructs of go-lang.

----------

## Go features

* Modern & clean
* System language
* Wide range of usecases
* Strongly typed
* Compiled
* Concurrency support
* Cross-platform support

----------


## Go Modules, Packages and Source-files Basics

Installing go lang is very easy, just go to [go install](https://go.dev/doc/install) page and follow instructions.

* package structure:
```
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

## Build, Run, Install Basics

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

### 2. Variables and Constants

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
```
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


## Pull Requests
I only accept PR's to fix mistakes. Thanks in advance.

