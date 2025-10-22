package usecase_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
	"github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation/usecase"
)

func TestStructureCreationUseCase_GoModuleInitialisation(t *testing.T) {
	executionError := errors.New("execution error")
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		projectName    string
		initError error
		executionError error
	}{
		{"test with empty project name","",errors.New("invalid project name"),nil},
		{"test with invalid project name","github.com/name/project/",fmt.Errorf("invalid module name %q","github.com/name/project/"),nil},
		{"test with correct project name","github.com/name/project",nil,nil},
		{"test with error during the command execution","github.com/name/project",executionError,executionError},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//mockFileRepo := new(usecase.MockFilesAndFolderRepository)
			mockCommandRepo := new(usecase.MockCommandRepository)
			uc := usecase.NewStructureCreationUseCase(nil,mockCommandRepo)
			mockCommandRepo.On("Execute","go",[]string{"mod","init",tt.projectName}).Return(tt.executionError).Once()
			gotErr := uc.GoModuleInitialisation(tt.projectName)
			if gotErr != nil {
				if tt.initError == nil {
					t.Errorf("GoModuleInitialisation() failed: %v", gotErr)
				}
				return
			}
			if tt.initError != nil {
				t.Fatal("GoModuleInitialisation() succeeded unexpectedly")
			}
		})
	}
}

func TestStructureCreationUseCase_GitRepositoryInitialisation(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		executionErr error
	}{
		{"test with successful execution",nil},
		{"test with unsuccessful execution",errors.New("some execution error")},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCommandRepo := new(usecase.MockCommandRepository)
			uc := usecase.NewStructureCreationUseCase(nil,mockCommandRepo)
			mockCommandRepo.On("Execute","git",[]string{"init"}).Return(tt.executionErr).Once()
			gotErr := uc.GitRepositoryInitialisation()
			if gotErr != tt.executionErr {
				t.Errorf("GitRepositoryInitialisation() failed: want %v have %v", tt.executionErr,gotErr)
			}
		})
	}
}

func TestStructureCreationUseCase_CreateBaseStructure(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		options creationModels.Options
		wantErr bool
	}{
		{ "test with successful creation",creationModels.Options{ProjectName: "",Modules: []string{}},false},
		{ "test with error during creation",creationModels.Options{ProjectName: "",Modules: []string{}},true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFileRepo := new(usecase.MockFilesAndFolderRepository)
			uc := usecase.NewStructureCreationUseCase(mockFileRepo,nil)
			mockFileRepo.On("CreateFile",mock.Anything,mock.Anything).Return(nil)
			if tt.wantErr == true{
				mockFileRepo.On("CreateFolder","internal").Return(errors.New("error folder creating")).Once()
			}
			mockFileRepo.On("CreateFolder",mock.Anything).Return(nil)
			gotErr := uc.CreateBaseStructure(tt.options)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateBaseStructure() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateBaseStructure() succeeded unexpectedly")
			}
		})
	}
}

func TestStructureCreationUseCase_CreateModules(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		moduleNames []string
		wantErr     bool
	}{
		{"test with successful creation",[]string{"cart","catalog"},false},
		{"test with error during creation",[]string{"cart","catalog"},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFileRepo := new(usecase.MockFilesAndFolderRepository)
			uc := usecase.NewStructureCreationUseCase(mockFileRepo,nil)
			if tt.wantErr == true{
				mockFileRepo.On("CreateFile","internal/"+tt.moduleNames[0]+"/repository.go",mock.Anything).Return(errors.New("error file creation")).Once()
			}
			mockFileRepo.On("CreateFile",mock.Anything,mock.Anything).Return(nil)
			mockFileRepo.On("CreateFolder",mock.Anything).Return(nil)
			gotErr := uc.CreateModules(tt.moduleNames)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateModules() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateModules() succeeded unexpectedly")
			}
		})
	}
}

func TestValidateProjectName(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		projectName string
		wantErr     bool
	}{
		{"test name with spaces"," github.com/name/project ",false},
		{"test name with slash in the end","github.com/name/project/",true},
		{"test name with symbols","github.com/!name/project@",true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := usecase.ValidateProjectName(tt.projectName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ValidateProjectName() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ValidateProjectName() succeeded unexpectedly")
			}
		})
	}
}
