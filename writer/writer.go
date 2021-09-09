package writer

import (
	"encoding/json"
	"fmt"
	. "github.com/flydevtools/goas/openApi3Schema"
	log "github.com/sirupsen/logrus"
	"os"
	"gopkg.in/yaml.v3"
)

type Writer interface {
	Write(OpenAPIObject, string) error
}

type fileWriter struct{
	format string
}

func NewFileWriter(format string) *fileWriter {
	return &fileWriter{format}
}

func (w *fileWriter) Write(openApiObject OpenAPIObject, path string) error {
	log.Info("Writing to open api object file ...")
	fd, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Can not create the file %s: %v", path, err)
	}
	defer fd.Close()

	var output []byte

	if w.format == "yaml" {
		output, err = yaml.Marshal(openApiObject)
	} else {
		output, err = json.MarshalIndent(openApiObject, "", "  ")
	}
	if err != nil {
		return err
	}
	_, err = fd.WriteString(string(output))
	return err
}
