### Transporter

Transporter is a database migration tool (CLI), it helps you keep your databases in order by creating migration files, those migration files are Go files that run SQL commands on your database to keep multiple environments databases in sync on its structure.

#### Installation

To install Transporter please run `go get -u github.com/wawandco/transporter`, with that, the transporter command should be available.

#### Commands

- transporter init

  `init` checks for `db` folder and creates this one and the `migrations` folder inside if these doesn't exist, also generates a `config.yml` inside the `db` folder.

- transporter create

  `generate` creates a migration inside the `db/migrations` folder.

- transporter up
- transporter down

#### Migrations

Migrations are simple `.go` files generated by transporter, these define a struct that implements 2 functions `Up_XXXXXXXXXXXXXX()` and `Down_XXXXXXXXXXXXXX()` which will be called according to the command you're calling from the CLI.

This is an example of a transporter Migration.

```go
package migrations
import "github.com/wawandco/transporter/transporter"

func init(){
  transporter.Register(&Migration{
    Up:         Up_20151119112619,
    Down:       Down_20151119112619,
    Identifier: 20151119112619,
  })
}

func Up_20151119112619(txn *transporter.Tx) {
  //Here you run your logic for the migration.
}

func Down_20151119112619(txn *transporter.Tx) {
  //Here you run your logic to rollback the migration.
}
```

#### Copyright
Fako is Copyright © 2008-2015 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.
