
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: 1448643678407987006,
    Up: func(tx *sql.Tx){
      tx.Exec("Create table other_table (a varchar(255) );")
    },
    Down: func(tx *sql.Tx){
      tx.Exec("Drop Table other_table;")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
