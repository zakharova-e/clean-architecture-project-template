package command_test

import(
	"os"
	"io"
	"bytes"
	"github.com/zakharova-e/clean-architecture-project-template/internal/structureCreation/repository/command"
	"testing"
	"sync"
)

func TestCommandRepository_Execute(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		command string
		params  []string
		wantErr bool
		expectedOutput string
	}{
		{"test echo command","echo",[]string{"123"},false,"123\n"},
		{"test exit with error command","false",nil,true,""},
		{"test with non existent command","gsdjgkjkfs",nil,true,""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo:= command.NewCommandRepository()

			output,gotErr := captureStdout(repo.Execute,tt.command, tt.params)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Execute() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Execute() succeeded unexpectedly")
			}
			if output != tt.expectedOutput{
				t.Errorf("Execute() received wrong result: has %v, wanted %v", output,tt.expectedOutput)
			}
		})
	}
}


// Helper function to capture stdout
func captureStdout(f func(string,[]string) error, command string, params []string) (string,error) {
	oldStdout := os.Stdout 
	r, w, _ := os.Pipe()   
	os.Stdout = w          

	var wg sync.WaitGroup
	wg.Add(1)
	var buf bytes.Buffer
	go func() {
		defer wg.Done()
		io.Copy(&buf, r) 
	}()

	err := f(command,params) 

	w.Close()       
	wg.Wait()       
	os.Stdout = oldStdout 
	return buf.String(),err   
}
