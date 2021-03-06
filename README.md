### Transporter

[![Circle CI](https://circleci.com/gh/wawandco/transporter.svg?style=svg&circle-token=93794e8b2f6b9b594822f00b72284f4928d21056)](https://circleci.com/gh/wawandco/transporter)
[![docs](https://img.shields.io/badge/godoc-docs-blue.svg)](https://godoc.org/github.com/wawandco/transporter)


Transporter is a database migration tool (CLI), it helps you keep your databases in order by creating migration files, those migration files are Go files that run SQL commands on your database to keep multiple environments databases in sync on its structure.

Features:

- Allows you to keep your database environments sync'ed
- You can run your migrations up and down as needed
- Supports multiple DBMS
- Provides functions for common database operations
- You can run your own SQL

#### Installation

To install Transporter please run `go get -u github.com/wawandco/transporter`, with that, the transporter command should be available.

#### Supported DBMS

Transporter supports the following DBMS's

- postgresql

  ```yml

  development:
    driver: postgres
    url: "user=username dbname=my_db_development sslmode=disable"

  ```

- mysql

  ```yml

  development:
    driver: mysql
    url: "transporter:password@tcp(mysql.local:3306)/ttest"

  ```

- mariadb

  ```yml

  development:
    driver: mariadb
    url: "transporter:password@tcp(mariadb.local:3306)/ttest"

  ```



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
import transporter "github.com/wawandco/transporter/core"

func init(){
  transporter.Register(&Migration{
    Identifier: 20151119112619,
    Up:         func (txn *transporter.Tx) {
      txn.Exec("ALTER TABLE my_table ADD COLUMN new_column varchar(255);")
    },
    Down:       func (txn *transporter.Tx) {
      txn.Exec("ALTER TABLE my_table DROP COLUMN new_column;")
    },
  })
}
```

#### Common operation functions

As you may have noticed, there are some database operations we do frequently when creating and maintaining database migrations, we've added some functions to make your life easier :), you can use these as the following example shows:

```go

package migrations
import transporter "github.com/wawandco/transporter/core"

func init(){
  transporter.Register(&Migration{
    Identifier: 20151119112619,
    Up:         func (txn *transporter.Tx) {

      //Instead of doing it this way:
      txn.Exec("ALTER TABLE my_table ADD COLUMN new_column varchar(255);")

      //You could simply do:
      tx.AddColumn("my_table", "new_column", "varchar(255)") // And is equivalent

    },
    Down:       func (txn *transporter.Tx) {
      txn.Exec("ALTER TABLE my_table DROP COLUMN new_column;")
    },
  })
}
```

We have provided the following other functions you could use the same manner

```go
tx.CreateTable("my_table", transporter.Table{
   "old_column": "varchar(12)",
}) // We will need to add this transporter.Table struct.

tx.DropTable("my_table")
tx.AddColumn("my_table", "new_column", "varchar(12)")
tx.DropColumn("my_table", "new_column")
tx.ChangeColumnType("my_table", "new_column", "varchar(12)")
tx.RenameColumn("my_table", "new_column", "varchar(12)")
tx.RenameTable("my_table", "other_table_name")

```

As usual, we're open to suggestions on common database operations, please let us know!.

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

### Flags Support

In some scenarios you won't want to persist credentials to your database inside the source code, in that case you can specify the database url and driver by using parameters for the transporter command.

```bash
transporter --database="user=username dbname=my_db_development sslmode=disable" --driver=postgres up 
```

#### Developer Setup

To work on this we've set a docker-compose file, you will need to instal docker toolkit in order to run our tests against the docker images.

1. Install docker, docker-compose
2. run `docker-compose build`
3. run `./dbsetup.sh` to setup testing databases.
4. run `docker-compose run lib go test ./... -v` to run the tests.

Ensure you have all the tests passing on the docker-compose machines before pushing into master.

#### Copyright
Transporter is Copyright © 2008-2015 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.
