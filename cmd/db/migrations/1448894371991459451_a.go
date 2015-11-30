
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: 1448894371991459451,
    Up: func(tx *sql.Tx){
      tx.Exec("Alter table down_table add column o varchar(12);")
    },
    Down: func(tx *sql.Tx){
      tx.Exec("Alter table down_table drop column o;")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
