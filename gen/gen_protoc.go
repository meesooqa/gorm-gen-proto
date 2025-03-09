package gen

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"text/template"
)

const (
	tplDir       = `pb/proto/{{.Dir}}`
	tplCmdProtoc = `protoc -I. -I../ --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative {{.File}}.proto`
)

type ProtocExecutor struct{}

func NewProtocExecutor() *ProtocExecutor {
	return &ProtocExecutor{}
}

/*
	```bash
	cd ./pb/proto/{{.Package}}
	protoc -I. -I../ \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative \
	{{.Package}}.proto
	```
*/
// Run generates Go files using `protoc`
func (o *ProtocExecutor) Run(dirName, fileStem string) error {
	origDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(origDir)

	dir, err := o.getDir(dirName)
	if err != nil {
		return err
	}
	cmdProtoc, err := o.getCmdProtoc(fileStem)
	if err != nil {
		return err
	}

	err = os.Chdir(dir)
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(context.TODO(), "sh", "-c", cmdProtoc)
	err = cmd.Run()
	if err != nil {
		return err
	}
	err = os.Chdir(origDir)
	if err != nil {
		return err
	}
	cmd = exec.CommandContext(context.TODO(), "sh", "-c", "go mod tidy")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func (o *ProtocExecutor) getCmdProtoc(fileStem string) (string, error) {
	b := bytes.Buffer{}
	err := template.Must(template.New("protoc").Parse(tplCmdProtoc)).Execute(&b, struct {
		File string
	}{
		File: fileStem,
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

func (o *ProtocExecutor) getDir(dirName string) (string, error) {
	b := bytes.Buffer{}
	err := template.Must(template.New("dir").Parse(tplDir)).Execute(&b, struct {
		Dir string
	}{
		Dir: dirName,
	})
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
