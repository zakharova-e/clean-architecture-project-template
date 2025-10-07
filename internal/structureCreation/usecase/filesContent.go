package usecase

import(
	"runtime"
	"log"
	"bytes"
	creationModels "github.com/zakharova-e/clean-architecture-project-template/internal/models"
)

func getDockerFileContent(options creationModels.Options) []byte{
	var buf bytes.Buffer
	data := map[string]string{"Version": runtime.Version(),"Name":options.ProjectName}
    if err := fileTemplates.ExecuteTemplate(&buf, "dockerfile.tmpl", data); err != nil {
        log.Fatal(err)
    }

	return buf.Bytes()
}

func getDockerignoreContent(options creationModels.Options) []byte{
	var buf bytes.Buffer
    if err := fileTemplates.ExecuteTemplate(&buf, "dockerignore.tmpl", nil); err != nil {
        log.Fatal(err)
    }

	return buf.Bytes()
}

func getGitignoreContent(options creationModels.Options) []byte{
	var buf bytes.Buffer
    if err := fileTemplates.ExecuteTemplate(&buf, "gitignore.tmpl",nil); err != nil {
        log.Fatal(err)
    }

	return buf.Bytes()
}

func getReadmeContent(options creationModels.Options) []byte{
	var buf bytes.Buffer
    if err := fileTemplates.ExecuteTemplate(&buf, "readme.tmpl", nil); err != nil {
        log.Fatal(err)
    }

	return buf.Bytes()
}