package utils

const MigrationTemplate = `
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: {{.Identifier}},
    Up: func(tx *sql.Tx){
      tx.Exec("{{.UpCommand}}")
    },
    Down: func(tx *sql.Tx){
      tx.Exec("{{.DownCommand}}")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
`
