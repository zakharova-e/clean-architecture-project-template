package file_test

import (
	"testing"
	"path/filepath"
	"os"
	"github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation/repository/file"
)

func TestFilesAndFolderRepository_CreateFolder(t *testing.T) {
	tmp := t.TempDir()
	fileName := "filename"
	//create a file to third test
	os.WriteFile(filepath.Join(tmp,fileName), []byte("data"), 0644)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		folderName    string
		wantErr bool
	}{
		{"test folder creation","newfolder",false},
		{"test non-existing path folder creation","a/b/c/newfolder",true},
		{"test create folder with already existing filename",fileName,true},
	}
	
	repo := file.NewFilesAndFolderRepository()
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folderName := filepath.Join(tmp,tt.folderName)
			gotErr := repo.CreateFolder(folderName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateFolder() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateFolder() succeeded unexpectedly")
			}
		})
	}
}

func TestFilesAndFolderRepository_CreateFile(t *testing.T) {
	tmp := t.TempDir()
	existingFilename := "existingfilewithsometext.txt"
	os.WriteFile(filepath.Join(tmp,existingFilename), []byte("any content"), 0644)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		fileName    string
		content []byte
		wantErr bool
	}{
		{"test file creation","newfile.txt",[]byte("Hello, world!"),false},
		{"test file creation in non-existing folder","a/b/c/newfile.txt",[]byte("Hello, world!"),true},
		{"test file rewrite",existingFilename,[]byte("Hello, world!"),false},
	}
	repo := file.NewFilesAndFolderRepository()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := filepath.Join(tmp,tt.fileName)
			gotErr := repo.CreateFile(path, tt.content)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateFile() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateFile() succeeded unexpectedly")
			}
			data, err := os.ReadFile(path)
			if err!=nil{
				t.Errorf("CreateFile() failed: %v", err)
			}
			if string(data) != string(tt.content){
				t.Errorf("CreateFile() failed: wrong file content. Have: %s, want: %s", string(data),string(tt.content))
			}
		})
	}
}
