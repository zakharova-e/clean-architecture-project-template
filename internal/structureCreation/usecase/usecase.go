package usecase

import (
	"errors"

	creationModule "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation"
)

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

func(uc *StructureCreationUseCase) CreateBaseStructure() error{
	var err error
	err = uc.filesRepo.CreateFolder("cmd")
	err = uc.filesRepo.CreateFolder("internal")
	err = uc.filesRepo.CreateFolder("internal/config")
	err = uc.filesRepo.CreateFile("internal/config/app.go",[]byte{})
	err = uc.filesRepo.CreateFile("Readme.md",[]byte{})
	err = uc.filesRepo.CreateFile(".gitignore",[]byte{})
	err = uc.filesRepo.CreateFile("Dockerfile",[]byte{})
	err = uc.filesRepo.CreateFile(".dockerignore",[]byte{})
	err = uc.filesRepo.CreateFile(".env",[]byte{})
	return err
}

func(uc *StructureCreationUseCase) CreateModules(moduleNames []string) error{
	var err error
	for _,module := range moduleNames{
		moduleFolder := "internal/"+module
		err = uc.filesRepo.CreateFolder(moduleFolder)
		err = uc.filesRepo.CreateFile(moduleFolder+"/repository.go",[]byte{})
		err = uc.filesRepo.CreateFile(moduleFolder+"/usecase.go",[]byte{})
		err = uc.filesRepo.CreateFolder(moduleFolder+"/delivery")
		err = uc.filesRepo.CreateFolder(moduleFolder+"/repository")
		err = uc.filesRepo.CreateFolder(moduleFolder+"/usecase")
		err = uc.filesRepo.CreateFile(moduleFolder+"/usecase/usecase.go",[]byte{})
	}
	return err
}



	
	
	