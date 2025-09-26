package command

import(
	"os"
	"os/exec"
	creationModule "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation"
)

type CommandRepository struct{
	
}

func NewCommandRepository() creationModule.ICommandRepository{
	return &CommandRepository{}
}

func(repo *CommandRepository) Execute(command string, params []string) error{
	cmd := exec.Command(command, params...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    return  cmd.Run()
}