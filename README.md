# go-gtt

Generate Terraform Template by Go

## COMMANDS

### _generate_

Generate a module template.
Sample files are found [here](./cmd/generate/_examples).

```sh
gtt generate [-f] <module name>
```

### _update_

Update a module's README according to the `variables.tf` and `outputs.tf`.

- Fill the tables in `README.md`.

```sh
gtt update
```

## INSTALLATION

Built binaries are available from GitHub Releases. https://github.com/kokoichi206/go-gtt/releases

## LICENSE

"go-gtt" is unde [MIT License](./LICENSE).
