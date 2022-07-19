package update

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

// Return cli command about generate.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update README according to the tf files",
		Description: `
	Update README according to the tf files.
	If target file is not found, it will return error.`,
		Action: doUpdate,
	}
}

type variable struct {
	name        string
	description string
}

type output struct {
	name        string
	description string
}

// Main action
func doUpdate(c *cli.Context) error {

	// If the folder already exists, do nothing.
	if _, err := os.Stat("variables.tf"); os.IsNotExist(err) {
		return errors.New("variables.tf should exist.")
	}
	if _, err := os.Stat("outputs.tf"); os.IsNotExist(err) {
		return errors.New("variables.tf should exist.")
	}
	if _, err := os.Stat("README.md"); os.IsNotExist(err) {
		return errors.New("variables.tf should exist.")
	}

	// Get variables from 'variables.tf' file.
	bytes, err := ioutil.ReadFile("variables.tf")
	if err != nil {
		return err
	}
	variablesStr := string(bytes)
	r := regexp.MustCompile(`(?s)variable "([^{]*)" {.*description[ ]*= "(.*)"`)
	slice := strings.Split(variablesStr, "}\n")
	variables := []variable{}
	// MUST: description
	for _, s := range slice {
		res := r.FindAllStringSubmatch(s, -1)

		for _, v := range res {
			if len(v) < 3 {
				continue
			}
			variables = append(variables, variable{
				name:        v[1],
				description: v[2],
			})
		}
	}

	// Get outputs from 'outputs.tf' file.
	bytes, err = ioutil.ReadFile("outputs.tf")
	if err != nil {
		return err
	}
	outputsStr := string(bytes)
	r = regexp.MustCompile(`(?s)output "([^{]*)" {.*description[ ]*= "(.*)"`)
	slice = strings.Split(outputsStr, "}\n")
	outputs := []output{}
	// MUST: description
	for _, s := range slice {
		res := r.FindAllStringSubmatch(s, -1)

		for _, v := range res {
			if len(v) < 3 {
				continue
			}
			outputs = append(outputs, output{
				name:        v[1],
				description: v[2],
			})
		}
	}

	rbytes, err := ioutil.ReadFile("README.md")
	if err != nil {
		return err
	}
	rms := string(rbytes)

	// Update variables.
	vr := regexp.MustCompile(`(?s)\| variable ([^#]*)\n\n`)
	rms = vr.ReplaceAllString(rms, createVariableContents(variables))

	// Update outputs.
	or := regexp.MustCompile(`(?s)\| output ([^#]*)\n\n`)
	rms = or.ReplaceAllString(rms, createOutputContents(outputs))

	// FIXME: not smart...
	// Update outputs for last line.
	orEnd := regexp.MustCompile(`(?s)\| output ([^#]*)\n$`)
	rms = orEnd.ReplaceAllString(rms, createOutputContents(outputs))

	// Override the README file.
	f, err := os.Create("README.md")
	if err != nil {
		return err
	}
	f.Write([]byte(rms))

	return nil
}

var sz8k = 8 * 1024

func createVariableContents(variables []variable) string {
	b := make([]byte, 0, sz8k)
	for _, vari := range variables {
		b = append(b, fmt.Sprintf("| %s | %s |", vari.name, vari.description)...)
		b = append(b, "\n"...)
	}
	b = append(b, "\n"...)
	return `| variable | description |
| -------- | ----------- |
` + string(b)
}

func createOutputContents(outputs []output) string {
	b := make([]byte, 0, sz8k)
	for _, out := range outputs {
		b = append(b, fmt.Sprintf("| %s | %s |", out.name, out.description)...)
		b = append(b, "\n"...)
	}
	b = append(b, "\n"...)
	return `| output | description |
| ------ | ----------- |
` + string(b)
}
