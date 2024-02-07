# SQLBoiler Expose Relations Generator

If you are a fan of [SQLBoiler](https://github.com/volatiletech/sqlboiler) you know that the relations are not exposed but place in a struct called `.R`. This works fine if you don't need expose the relations with json. This package is made, so you don't have to manually make custom functions to expose them as stated here: [https://github.com/volatiletech/sqlboiler/issues/272](https://github.com/volatiletech/sqlboiler/issues/272).

## Installation

First you have to install the code generator binaries. Keep in mind that you still have to install `sqlboiler` separately.


```bash
go install github.com/expanse-agency/sqlboiler-erg@latest
```

## Usage

```go

import (
  "github.com/yourpackage/models/dm" // Database Models (from sqlboiler)
  "github.com/yourpackage/models/am" // Api Models (from sqlboiler-erg)
)

func (r *Router) GetUsers(c *framework.Context) {
  users, err := dm.Users().All(r.db, c)
  if err != nil {
    c.JSON(http.StatusInternalServerError, err)
    return
  }
  c.JSON(http.StatusOK, am.ToUsers(users, nil)) // Convert SQLBoiler model to API model
}

```

### Config options
You can also pass these options as flags.

| Name                | Defaults  | Description |
| ------------------- | --------- | ----------- |
| pkgname             | "erg"  | The name you wish to assign to your generated package |
| output              | "erg_models"  | The name of the folder to output to |
| output-ts |     | The name of the ts file (models.ts), no file is no typescript generation. Can also be ../../models.ts or dir/other_dir/filename.ts |
| wipe                | false     | Delete the output folder before generation to ensure sanity |
| inline             | false     | Whether to inline the slqboiler structs or not |
| blacklist          | []        | Can be table or column. Valid: "table", "table.column", "*.column" |

### Example

```toml
// ... existing sqlboiler.toml config
[erg]
output = "models/am"
output-ts = "models/models.ts"
pkgname = "am"
wipe    = true
inline = false
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

### Reasoning of the blacklist/inline options

In the original question it was proposed to just add the model like this:
```go
type User struct {
    *sqlboilerModel.User
    Books BookSlice `json:"books,omitempty"`
}
```

However sqlboiler does not omit empty fields and the `tag-ignore` option is only for specific fields. I wanted to disable certain fields like this: `*.password`. If you want to have the structs inlined you can use the `inline` option. 
```go
type User struct {
    ID:               a.ID,
    // ... other fields
    CreatedAt:        ConvertTime(a.CreatedAt),
    UpdatedAt:        ConvertTime(a.UpdatedAt),
    CreatedAt:        ConvertNullTime(a.UpdatedAt),
    Books BookSlice `json:"books,omitempty"`
}
```

