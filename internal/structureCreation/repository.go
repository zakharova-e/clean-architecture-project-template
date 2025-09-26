package structurecreation

type IFilesAndFolderRepository interface{
	CreateFolder(name string) error
	CreateFile(name string, content []byte) error
}

type ICommandRepository interface{
	Execute(command string, params []string) error
}