# SQLBoiler Expose Relations Generator

If you are a fan of [SQLBoiler](https://github.com/volatiletech/sqlboiler) you know that the relations are not exposed but place in a struct called `.R`. This works fine if you don't need expose the relations with json. This package is made, so you don't have to manually make custom functions to expose them as stated here: [https://github.com/volatiletech/sqlboiler/issues/272](https://github.com/volatiletech/sqlboiler/issues/272).

## Installation

First you have to install the code generator binaries. Keep in mind that you still have to install `sqlboiler` separately.


```bash
go install github.com/expanse-agency/sqlboiler-erg
```

## Usage

### Config options
You can also pass these options as flags.

| Name                | Defaults  | Description |
| ------------------- | --------- | ----------- |
| pkgname             | "erg"  | The name you wish to assign to your generated package |
| output              | "erg_models"  | The name of the folder to output to |
| output-ts |     | The name of the ts file (models.ts), no file is no typescript generation. Can also be ../../models.ts or dir/other_dir/filename.ts |
| wipe                | false     | Delete the output folder before generation to ensure sanity |
| blacklist          | []        | Can be table or column. Valid: "table", "table.column", "*.column" |

### Example

```toml
...
// your sqlboiler config
...
[erg]
output = "models/erg"
output-ts = "models/models.ts"
pkgname = "erg"
wipe    = true
blacklist = ["*.password", "table.token", "*.secret_column", "table"]
```

### Initial Generation

After generating your models with sqlboiler, you can run the following command to generate the exposed models.

```text
SQL Boiler generates a Go ORM from template files, tailored to your database schema.
Complete documentation is available at http://github.com/volatiletech/sqlboiler

Usage:
  sqlboiler-erg [flags]

Examples:
sqlboiler-erg -c sqlboiler-erg.toml 

Flags:
  -c, --config string              Filename of config file to override default lookup
  -h, --help                       help for sqlboiler-erg
```

### Reasoning of the blacklist

In the original question it was proposed to just add the model like this:
```go
type User struct {
    *sqlboilerModel.User
    Books BookSlice `json:"books,omitempty"`
}
```

I found however that sqlboiler does not omit the json fields and disabling fields is possible with `tag-ignore` but you have to do this for every field. I found it easier to just blacklist the fields I don't want to expose and not include them in the struct. You can also easily blacklist a whole table if you don't want it to show up as a relation anywhere.