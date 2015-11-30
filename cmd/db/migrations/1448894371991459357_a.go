
package migrations
import (
  "database/sql"
  "github.com/wawandco/transporter/transporter"
)

func init(){
  migration := transporter.Migration{
    Identifier: 1448894371991459357,
    Up: func(tx *sql.Tx){
      tx.Exec("Create table down_table (a varchar(255) );")
    },
    Down: func(tx *sql.Tx){
      tx.Exec("Drop table down_table;")
    },
  }

  //Register the migration to run up or down acordingly.
  transporter.Register(migration)
}
