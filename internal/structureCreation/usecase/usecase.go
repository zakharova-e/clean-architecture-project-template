package usecase

import (
	"embed"
	"errors"
	"fmt"
	"strings"
	"text/template"
	"regexp"

	"golang.org/x/sync/errgroup"

	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
	creationModule "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation"
)

var fileTemplates *template.Template

//go:embed templates/*
var fs embed.FS

func init(){
	var err error
	fileTemplates, err = template.ParseFS(fs, "templates/*.tmpl")
    if err != nil {
        panic(err)
    }
}

type StructureCreationUseCase struct{
	filesRepo creationModule.IFilesAndFolderRepository
	commandsRepo creationModule.ICommandRepository
}

func NewStructureCreationUseCase(filesRepo creationModule.IFilesAndFolderRepository,commandsRepo creationModule.ICommandRepository) creationModule.IStructureCreationUseCase{
	return &StructureCreationUseCase{filesRepo: filesRepo,commandsRepo: commandsRepo}
}

func(uc *StructureCreationUseCase) GoModuleInitialisation(projectName string) error{
	if projectName == ""{
		return errors.New("invalid project name")
	}
	validErr := ValidateProjectName(projectName)
	if validErr!= nil{
		return validErr
	}
	return uc.commandsRepo.Execute("go",[]string{"mod","init",projectName})
}

func(uc *StructureCreationUseCase) GitRepositoryInitialisation() error{
	return uc.commandsRepo.Execute("git",[]string{"init"})
}

func(uc *StructureCreationUseCase) CreateBaseStructure(options creationModels.Options) error{
	eg := new(errgroup.Group)
	eg.Go(func() error {
		return uc.filesRepo.CreateFolder("cmd")
	})
	eg.Go(func() error {
		err := uc.filesRepo.CreateFolder("internal")
		if err!=nil{
			return err
		}
		err = uc.filesRepo.CreateFolder("internal/config")
		if err!=nil{
			return err
		}
		err = uc.filesRepo.CreateFile("internal/config/app.go",[]byte{})
		return err
	})
	eg.Go(func() error {
		content,errR := getReadmeContent(options)
		if errR!=nil{
			return errR
		}
		return uc.filesRepo.CreateFile("Readme.md",content)
	})
	eg.Go(func() error {
		content,errR := getGitignoreContent(options)
		if errR!=nil{
			return errR
		}
		return uc.filesRepo.CreateFile(".gitignore",content)
	})
	eg.Go(func() error {
		content,errR := getDockerFileContent(options)
		if errR!=nil{
			return errR
		}
		return uc.filesRepo.CreateFile("Dockerfile",content)
	})
	eg.Go(func() error {
		content,errR := getDockerignoreContent(options)
		if errR!=nil{
			return errR
		}
		return uc.filesRepo.CreateFile(".dockerignore",content)
	})
	eg.Go(func() error {
		return uc.filesRepo.CreateFile(".env",[]byte{})
	})
	return eg.Wait(); 
}

func(uc *StructureCreationUseCase) CreateModules(moduleNames []string) error{
	eg := new(errgroup.Group)
	for _,module := range moduleNames{
		if module == ""{
			continue
		}
		moduleFolder := "internal/"+module
		eg.Go(func() error {
			err := uc.filesRepo.CreateFolder(moduleFolder)
			if err!=nil{
				return err
			}
			err =  uc.filesRepo.CreateFile(moduleFolder+"/repository.go",[]byte{})
			if err!=nil{
				return err
			}
			err =  uc.filesRepo.CreateFile(moduleFolder+"/usecase.go",[]byte{})
			if err!=nil{
				return err
			}
			err =  uc.filesRepo.CreateFolder(moduleFolder+"/delivery")
			if err!=nil{
				return err
			}
			err =  uc.filesRepo.CreateFolder(moduleFolder+"/repository")
			if err!=nil{
				return err
			}
			err = uc.filesRepo.CreateFolder(moduleFolder+"/usecase")
			if err!=nil{
				return err
			}
			err =  uc.filesRepo.CreateFile(moduleFolder+"/usecase/usecase.go",[]byte{})
			return err
		})

	}
	return eg.Wait(); 
}


func ValidateProjectName(projectName string) error{
	projectNameTrimmed := strings.TrimSpace(projectName)
	validPattern := `^[a-zA-Z0-9._\-]+(/[a-zA-Z0-9._\-]+)*$`
	matched, _ := regexp.MatchString(validPattern, projectNameTrimmed)
	if !matched {
		return fmt.Errorf("invalid module name %q",projectName)
	}
	return nil
}


	
	
	