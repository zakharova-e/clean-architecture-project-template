# Project template creator

This is a simple tool for creating all necessary files and folders to start a new project

## Input parameters

#### Project name 

Use --name flag to specify the project name. This name will be used for golang module creation.
Be careful, use only accepted names.

`projectTemplate --name "amazing-parser" `


#### Modules

Use position arguments to specify module names, that you want to create.
Module is a logically separated part of the application. Each module is a folder with "layered" structure, which comprises delivery, usecase and repository layers.

`projectTemplate --name "amazing-parser" module1 module2 module3`


## Compilation and using

Compile to executable file

`go build -o projectTemplate cmd/project-template/main.go`

Put it in $GOPATH/bin folder
