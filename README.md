### Transporter

[![Circle CI](https://circleci.com/gh/wawandco/transporter.svg?style=svg&circle-token=93794e8b2f6b9b594822f00b72284f4928d21056)](https://circleci.com/gh/wawandco/transporter)

Transporter is a database migration tool (CLI), it helps you keep your databases in order by creating migration files, those migration files are Go files that run SQL commands on your database to keep multiple environments databases in sync on its structure.

#### Installation

To install Transporter please run `go get -u github.com/wawandco/transporter`, with that, the transporter command should be available.

#### Supported DBMS

Transporter supports the following DBMS's

- postgresql
- mysql [comming soon]

#### Commands

- transporter init

  `init` checks for `db` folder and creates this one and the `migrations` folder inside if these doesn't exist, also generates a `config.yml` file inside `db`.

- transporter create

  `generate` creates a migration inside the `db/migrations` folder.

- transporter up

  `up` runs all pending migrations inside the `db/migrations` folder.

- transporter down

  `down` runs last migration down.

#### Migrations

Migrations are simple `.go` files generated by transporter, these define a struct that implements 2 functions `Up` and `Down` which will be called according to the command you're calling from the CLI.

This is an example of a transporter Migration.

```go
package migrations
import "github.com/wawandco/transporter/transporter"

func init(){
  transporter.Register(&Migration{
    Identifier: 20151119112619,
    Up:         func (txn *sql.Tx) {
      txn.Exec("ALTER TABLE my_table ADD COLUMN new_column varchar(255);")
    },
    Down:       func (txn *sql.Tx) {
      txn.Exec("ALTER TABLE my_table DROP COLUMN new_column;")
    },
  })
}
```


#### Multiple Environment support

Transporter supports you to have multiple environments on the same application, these could be used to run tests against different databases.

By default is only ships with the `development` environment, which will run if you invoke the `up` and `down` commands without any other argument.

if you need to have multiple environments you could change your `db/config.yml` file as in this example:


```yml

development:
  driver: postgres
  url: "user=username dbname=my_db_development sslmode=disable"

staging:
  driver: postgres
  url: "user=username dbname=my_db_staging"

production:
  driver: postgres
  url: "user=username dbname=my_db_production"

```

And call `transporter up staging` or `transporter down staging` or any other environment's database you would like to run the migrations against.


#### Copyright
Transporter is Copyright © 2008-2015 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.
