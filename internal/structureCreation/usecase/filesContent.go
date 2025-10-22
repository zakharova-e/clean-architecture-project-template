package usecase

import(
	"runtime"
	"bytes"
	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
)

func getDockerFileContent(options creationModels.Options) ([]byte,error){
	var buf bytes.Buffer
	data := map[string]string{"Version": runtime.Version(),"Name":options.ProjectName}
    if err := fileTemplates.ExecuteTemplate(&buf, "dockerfile.tmpl", data); err != nil {
        return nil,err
    }

	return buf.Bytes(),nil
}

func getDockerignoreContent(options creationModels.Options) ([]byte,error){
	var buf bytes.Buffer
    if err := fileTemplates.ExecuteTemplate(&buf, "dockerignore.tmpl", nil); err != nil {
        return nil,err
    }

	return buf.Bytes(),nil
}

func getGitignoreContent(options creationModels.Options) ([]byte,error){
	var buf bytes.Buffer
    if err := fileTemplates.ExecuteTemplate(&buf, "gitignore.tmpl",nil); err != nil {
        return nil,err
    }

	return buf.Bytes(),nil
}

func getReadmeContent(options creationModels.Options) ([]byte,error){
	var buf bytes.Buffer
    if err := fileTemplates.ExecuteTemplate(&buf, "readme.tmpl", nil); err != nil {
        return nil,err
    }
	return buf.Bytes(),nil
}