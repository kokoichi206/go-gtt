package generate

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

// Return cli command about generate.
func Command() *cli.Command {
	return &cli.Command{
		Name:  "generate",
		Usage: "Clone/sync with a remote repository",
		Description: `
	Generate a template module directory. If the folder name already
	exists, nothing will happen unless '-f' ('--force') flag is supplied.
	'force' option will override existing files`,
		Action: doGenerate,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "force", Aliases: []string{"f"}},
		},
	}
}

// Main action
func doGenerate(c *cli.Context) error {
	var (
		f = c.Bool("force")
	)
	var slice = c.Args().Slice()
	fmt.Println(slice)
	if len(slice) != 1 {
		return errors.New("only one input file name should be followd after command.")
	}

	var module = slice[0]
	// If the folder already exists, do nothing.
	if _, err := os.Stat(module); !os.IsNotExist(err) {
		if !f {
			fmt.Printf("module %s already exists.\n", module)
			return fmt.Errorf("module %s already exists.\n", module)
		}
		// Continue if force option exists.
	} else if err := os.Mkdir(module, 0755); err != nil {
		fmt.Println(err)
		return err
	}
	os.Create(filepath.Join(module, "outputs.tf"))
	os.Create(filepath.Join(module, "variables.tf"))
	os.Create(filepath.Join(module, "main.tf"))

	var readmePath = filepath.Join(module, "README.md")
	err := os.WriteFile(readmePath, []byte(readmeContents(module)), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
		return err
	}

	return nil
}

// This is template for README.md
func readmeContents(module string) string {
	return "# " + module + `

## Usage

### [variables](./variables.tf)

| variable | description |
| -------- | ----------- |

### [outputs](./outputs.tf)

| output | description |
| ------ | ----------- |
`
}
