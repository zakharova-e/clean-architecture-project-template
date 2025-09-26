package structurecreation

type IStructureCreationUseCase interface{
	GoModuleInitialisation(name string) error
	GitRepositoryInitialisation() error
	CreateBaseStructure() error
	CreateModules(moduleNames []string) error
}