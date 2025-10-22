package usecase

import (
	"strings"
	"testing"

	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
)

func Test_getDockerFileContent(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		options creationModels.Options
		expectedSubstrings []string
		wantErr bool
	}{
		{"test successful rendering",creationModels.Options{ProjectName: "testProject", Modules: []string{}},[]string{"FROM","RUN","WORKDIR","COPY"},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got,gotErr := getDockerFileContent(tt.options)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getDockerFileContent() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getDockerFileContent() succeeded unexpectedly")
			}
			gotString := string(got)
			for _,sub := range tt.expectedSubstrings{
				if !strings.Contains(gotString,sub){
					t.Errorf("getDockerFileContent() failed: output doesn't contain substing %s, got %v", sub, gotString)
				}
			}
			
		})
	}
}

func Test_getDockerignoreContent(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		options creationModels.Options
		expectedSubstrings []string
		wantErr bool
	}{
		{"test successful rendering",creationModels.Options{ProjectName: "testProject", Modules: []string{}},[]string{".gitignore",".git"},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got,gotErr := getDockerignoreContent(tt.options)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getDockerignoreContent() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getDockerignoreContent() succeeded unexpectedly")
			}
			gotString := string(got)
			for _,sub := range tt.expectedSubstrings{
				if !strings.Contains(gotString,sub){
					t.Errorf("getDockerignoreContent() failed: output doesn't contain substing %s, got %v", sub, gotString)
				}
			}
			
		})
	}
}

func Test_getGitignoreContent(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		options creationModels.Options
		expectedSubstrings []string
		wantErr bool
	}{
		{"test successful rendering",creationModels.Options{ProjectName: "testProject", Modules: []string{}},[]string{".env","*.log"},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got,gotErr := getGitignoreContent(tt.options)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getGitignoreContent() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getGitignoreContent() succeeded unexpectedly")
			}
			gotString := string(got)
			for _,sub := range tt.expectedSubstrings{
				if !strings.Contains(gotString,sub){
					t.Errorf("getGitignoreContent() failed: output doesn't contain substing %s, got %v", sub, gotString)
				}
			}
			
		})
	}
}

func Test_getReadmeContent(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		options creationModels.Options
		expectedSubstrings []string
		wantErr bool
	}{
		{"test successful rendering",creationModels.Options{ProjectName: "testProject", Modules: []string{}},[]string{},false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got,gotErr := getReadmeContent(tt.options)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("getReadmeContent() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("getReadmeContent() succeeded unexpectedly")
			}
			gotString := string(got)
			for _,sub := range tt.expectedSubstrings{
				if !strings.Contains(gotString,sub){
					t.Errorf("getReadmeContent() failed: output doesn't contain substing %s, got %v", sub, gotString)
				}
			}
			
		})
	}
}
