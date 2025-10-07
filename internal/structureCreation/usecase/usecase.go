package usecase

import (
	"embed"
	"errors"
	"text/template"

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

func(uc *StructureCreationUseCase) GoModuleInitialisation(name string) error{
	if name == ""{
		return errors.New("invalid project name")
	}
	//todo check if all symbols are allowed
	return uc.commandsRepo.Execute("go",[]string{"mod","init",name})
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
		return uc.filesRepo.CreateFile("Readme.md",getReadmeContent(options))
	})
	eg.Go(func() error {
		return uc.filesRepo.CreateFile(".gitignore",getGitignoreContent(options))
	})
	eg.Go(func() error {
		return uc.filesRepo.CreateFile("Dockerfile",getDockerFileContent(options))
	})
	eg.Go(func() error {
		return uc.filesRepo.CreateFile(".dockerignore",getDockerignoreContent(options))
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





	
	
	