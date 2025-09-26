package main

import (
	"flag"
	"fmt"
	"log"

	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
	creationModuleHandler "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation/delivery/cmd"
	creationUsecase "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation/usecase"
    creationCommandRepo "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation/repository/command"
    creationFileRepo  "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation/repository/file"
)


var options creationModels.Options

func init(){
    flag.StringVar(&options.ProjectName,"name","","Project name")
}

func main() {
    flag.Parse()
    args := flag.Args()
    options.Modules = make([]string, len(args))
    for _,v := range args{
        if v!=""{
            options.Modules = append(options.Modules, v)
        }
    }
    filesRepo := creationFileRepo.NewFilesAndFolderRepository()
    commandRepo:= creationCommandRepo.NewCommandRepository()
    usecase := creationUsecase.NewStructureCreationUseCase(filesRepo,commandRepo)
    handler := creationModuleHandler.NewCmdHandler(usecase)
    err := handler.CreateProjectTemplate(options)
    if err!=nil{
        log.Fatal(err.Error())
    }
    fmt.Println("Project created successfully")
} 