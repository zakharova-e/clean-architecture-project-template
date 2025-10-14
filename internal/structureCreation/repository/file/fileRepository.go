package file

import(
	"bufio"
	"os"
	creationModule "github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation"
)

type FilesAndFolderRepository struct{
	
}

func NewFilesAndFolderRepository() creationModule.IFilesAndFolderRepository{
	return &FilesAndFolderRepository{}
}

func(repo *FilesAndFolderRepository) CreateFolder(folderName string) error{
	return os.Mkdir(folderName, 0755)
}

func(repo *FilesAndFolderRepository) CreateFile(fileName string, content []byte) error{
	f, err := os.Create(fileName)
    if err != nil {
        return err
    }
    defer f.Close()
	bf := bufio.NewWriter(f)
	defer bf.Flush()

	_,err = bf.Write(content)
	return err
}