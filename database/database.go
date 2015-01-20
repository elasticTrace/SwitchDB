package database

import "fmt"
import "log"
//import "database/sql"
import "os"
import "path/filepath"
import "strings"

import _ "github.com/mattn/go-sqlite3"
import "github.com/jinzhu/gorm"


func Create(name string) bool {
    if Exists(name) == false {
        fmt.Print(Path(name))
        db, _ := Open(name)
        
        //create _collections
        db.CreateTable(&Namespace{})
        return true
    }
    return false
}

func Open(name string) (gorm.DB, error) {
    
    db, err := gorm.Open("sqlite3", Path(name))
    //db.DB()
    db.LogMode(true)
    //db.SingularTable(true)
    
    
    return db, err
}

func NamespaceExists(database, ns_name string) bool {
    if Exists(database) == true {
        db, _ := Open(database)
        namespace := Namespace{}
        
        //db.Where(&Namespace{Name: namespace}).First(&Namespace)
        query := db.Where(&Namespace{Name: ns_name}).First(&namespace)
        //rows, err := db.Query("SELECT COUNT(id) as count FROM _namespace WHERE name='"+namespace+"' LIMIT 1;")
        
        if query.RecordNotFound() {
            return false
        }
        return true
    }
    
    return false
}

func NamespaceCreate(database, ns_name, ns_type string) bool {
    db, err := Open(database)
    if err != nil {
        log.Fatal(err)
        return false
    }
    
    //check if the name space exists
    if NamespaceExists(database, ns_name) == true {
        return false
    }
    
    namespace := Namespace{Name: ns_name, Type: ns_type}
    db.Create(&namespace)
    
    return true
}

func Exists(name string) bool {
    if _, err := os.Stat(Path(name)); os.IsNotExist(err) {
        return false
    }
    return true
}

func List(filter string) []string {
    files, _ := filepath.Glob("./dbs/*.db")
    dbs := make([]string, 0)
    var db []string
    
    for _, f := range files {
        db = strings.Split(f, "/")
        db = strings.Split(db[len(db)-1], ".")
        dbs = append(dbs, db[0])
    }
    return dbs
}

func NamespaceList(db_name, filter string) []string {
    db, _ := Open(db_name)
    var namespace_list Namespace
    query := db.Find(&namespace_list).Limit(10)
    fmt.Println(query)
    var aa []string
    return aa
}

func Path(name string) string{
    return fmt.Sprintf("./dbs/%s.db", name)
}
