package usecase

import(
	"github.com/stretchr/testify/mock"
)

type MockFilesAndFolderRepository struct{
	mock.Mock
}

func(m *MockFilesAndFolderRepository) CreateFolder(name string) error{
	args := m.Called(name)
	return args.Error(0)
}

func(m *MockFilesAndFolderRepository) CreateFile(name string, content []byte) error{
	args := m.Called(name,content)
	return args.Error(0)
}

type MockCommandRepository struct{
	mock.Mock
}


func(m *MockCommandRepository) Execute(command string, params []string) error{
	args := m.Called(command,params)
	return args.Error(0)
}