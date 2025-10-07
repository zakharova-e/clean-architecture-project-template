package structurecreation

import (
	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
)
type IStructureCreationUseCase interface{
	GoModuleInitialisation(name string) error
	GitRepositoryInitialisation() error
	CreateBaseStructure(creationModels.Options) error
	CreateModules(moduleNames []string) error
}