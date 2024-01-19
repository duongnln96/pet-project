package open_api3

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/duongnln96/blog-realworld/third_party/OpenAPI/user"
	"github.com/ghodss/yaml"
	"github.com/urfave/cli/v2"
)

func Run(cliCtx *cli.Context) error {

	var output string

	if cliCtx.NumFlags() == 0 {
		log.Fatalf("Path is required")
	}

	output = cliCtx.Path("output")

	swagger := user.NewOpenAPI3()

	// openapi3.json
	data, err := json.Marshal(&swagger)
	if err != nil {
		return fmt.Errorf("Couldn't marshal json: %s", err)
	}
	if err := os.WriteFile(path.Join(output, "swagger.json"), data, 0600); err != nil {
		return fmt.Errorf("Couldn't write json: %s", err)
	}

	// openapi3.yaml
	data, err = yaml.Marshal(&swagger)
	if err != nil {
		return fmt.Errorf("Couldn't marshal yaml: %s", err)
	}
	if err := os.WriteFile(path.Join(output, "swagger.yaml"), data, 0600); err != nil {
		return fmt.Errorf("Couldn't write yaml: %s", err)
	}

	return nil
}
