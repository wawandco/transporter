
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: 1448643678407987101,
    Up: func(tx *sql.Tx){
      tx.Exec("Alter table other_table drop column a;")
    },
    Down: func(tx *sql.Tx){
      tx.Exec("")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
