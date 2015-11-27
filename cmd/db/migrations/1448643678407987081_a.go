
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: 1448643678407987081,
    Up: func(tx *sql.Tx){
      tx.Exec("Alter table shshshshs other_table add column o varchar(12);")
    },
    Down: func(tx *sql.Tx){
      tx.Exec("")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
