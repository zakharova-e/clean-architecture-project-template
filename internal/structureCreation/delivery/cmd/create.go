package cmd

import(
	creationModule "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation"
	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
)

type CmdHandler struct{
	usecase creationModule.IStructureCreationUseCase
}

func NewCmdHandler(usecase creationModule.IStructureCreationUseCase) *CmdHandler{
	return &CmdHandler{usecase: usecase}
}

func(handler *CmdHandler) CreateProjectTemplate(options creationModels.Options) error{
	var err error

	err = handler.usecase.GoModuleInitialisation(options.ProjectName)
	if err!=nil{
		return err
	}

	err = handler.usecase.GitRepositoryInitialisation()
	if err!=nil{
		return err
	}

	err = handler.usecase.CreateBaseStructure(options)
	if err!=nil{
		return err
	}

	err = handler.usecase.CreateModules(options.Modules)
	return err
}