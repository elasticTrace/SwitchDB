package database

import "fmt"
import "log"
import "database/sql"
import "os"

import _ "github.com/mattn/go-sqlite3"


func Create(name string) bool {
    filename := Path(name)
    if Exists(name) == false {
        db, err := sql.Open("sqlite3", filename)
        if err != nil {
            log.Fatal(err)
        }
        defer db.Close()
        
        //create _collections
        sqlStmt := "create table _namespace (id integer not null primary key, name text, type text);"
        _, err = db.Exec(sqlStmt)
        if err != nil {
            log.Printf("%q: %s\n", err, sqlStmt)
            return false
        }
        return true
    }
    return false
}

func Open(name string) (*sql.DB, error) {
    if Exists(name) == true {
        return sql.Open("sqlite3", Path(name))
    }
    return sql.Open("sqlite3", Path(name))
}

func NamespaceExists(database, namespace string) bool {
    if Exists(database) == true {
        db, _ := Open(database)
        
        rows, err := db.Query("SELECT COUNT(id) as count FROM _namespace WHERE name='"+namespace+"' LIMIT 1;")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()
        for rows.Next() {
            var count int
            rows.Scan(&count)
            if count > 0 {
                return true
            }
        }
    }
    
    return false
}

func CreateNamespace(database, namespace, ns_type string) bool {
    db, err := Open(database)
    if err != null {
        return log.Fatal(err)
    }
    
    db.Exec("INSERT INTO ")
}

func Exists(name string) bool {
    if _, err := os.Stat(Path(name)); os.IsNotExist(err) {
        return false
    }
    return true
}

func Path(name string) string{
    return fmt.Sprintf("./dbs/%s.db", name)
}
