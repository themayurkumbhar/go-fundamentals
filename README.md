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




## Pull Requests
I only accept PR's to fix mistakes. Thanks in advance.

