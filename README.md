# 💻 RMLang

## 📄 Introduction

This project is an implementation of an interpreter written in [Go](https://go.dev/) for a fictional programming language called "RMLang". The project is completely dependency-free, using only native Go packages.

## 🚀 Features

- Basic syntax
  - Variables binding
  - Arithmetic expressions 
- Common data types support
  - Integer
  - Boolean
  - String
  - Array
  - Hash map
- Operators
  - Arithmetic operators: +, -, *, /
  - Comparison operators: ==, !=, <, >, <=, >=
  - Logical operators: ! (not)
- Control structures
  - If statements: Basic conditional statements
- Functions
  - First class citizens
  - High order functions
  - Anonimous functions
- Built-in functions
  - **len**: Accepts an array or string as unique argument and returns its size or length
  - **first**: Accepts an array as unique argument and returns its first element
  - **last**: Accepts an array as unique argument and returns its last element
  - **rest**: Accepts an array as unique argument and returns its elements except the first one
  - **push**: Accepts an array as first argument and a expression as second argument, creates a copy of the array adding the element at the last position and returns it
  - **puts**: Prints the arguments to the STDOUT
- REPL

## 🏃 Running the project

To run the project execute the following steps:

1. Install Go language if it is not yet. Installation guide: https://go.dev/doc/install
2. Clone the repository using the following command:
```bash
$ git clone git@github.com:RafaLopesMelo/rmlang.git
```
3. Run the project using the following command:
```bash
$ go run cmd/main.go
```
4. After the previous step the REPL will be started and you can start testing:

![image](https://github.com/user-attachments/assets/28a63311-9b75-45de-be4a-7ae98e867f2e)

## 📁 Project Structure
```go
root/
├── cmd/
|   └── main.go        -> Application entry point, runs the REPL
├── internal/          -> Application code
|   ├── ast/           -> RMLang AST nodes, are evaluated by the Evaluator
|   ├── evaluator/     -> Responsible for actually run the language code
|   ├── lexer/         -> Responsible to transform source code into tokens
|   ├── object/        -> Are generated from the AST by the Evaluator and then interpreted
|   ├── parser/        -> Responsible for create the code AST from tokens generated by the Lexer
|   ├── repl/          -> REPL implementation
|   └── token/         -> Tokens generated from the source code by the Lexer
├── go.mod             -> Go module file
└── README.md          -> Project README
```
## ▶️ Demonstration
### Basic operations
![image](https://github.com/user-attachments/assets/4d5e525e-69ba-4633-9857-25aa49b38561)
### Conditional
![image](https://github.com/user-attachments/assets/22231140-428d-4a5a-847b-a18fc797ab36)
### Arrays
![image](https://github.com/user-attachments/assets/b1f24805-3cb8-4479-b17e-8f74035483cb)
### Hash maps
![image](https://github.com/user-attachments/assets/a8d85f3e-a22d-487e-83bb-7913f800e435)
### High order function
![image](https://github.com/user-attachments/assets/572aa157-c206-42c1-9fcd-eacb48f1182f)

## License

This project is licensed under the MIT License

## Acknowledgments
- Thorsten Ball for his amazing book, [Writing an Interpreter in Go](https://interpreterbook.com/)
