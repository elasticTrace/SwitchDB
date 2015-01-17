package database

import "fmt"
import "log"
import "database/sql"
import "os"

import _ "github.com/mattn/go-sqlite3"
import "github.com/jinzhu/gorm"


func Create(name string) bool {
    filename := Path(name)
    if Exists(name) == false {
        db = Open(Path(name))
        
        //create _collections
        db.CreateTable(&Namespace{})
        return true
    }
    return false
}

func Open(name string) (*sql.DB, error) {
    if Exists(name) == true {
        db, err := gorm.Open("sqlite3", Path(name))
    }
    
    db, err := gorm.Open("sqlite3", Path(name))
    db.SingularTable(true)
    return db, err
}

func NamespaceExists(database, namespace string) bool {
    if Exists(database) == true {
        db, _ := Open(database)
        
        db.Where(&Namespace{Name: namespace}).First(&namespace)
        
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
