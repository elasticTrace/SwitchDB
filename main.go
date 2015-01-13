package main

//import "log"
//import "database/sql"
//import "os"
import "fmt"

import "./database"

import "github.com/gin-gonic/gin"
//import _ "github.com/mattn/go-sqlite3"



func main() {
    r := gin.Default()
    r.PUT("/database/create/:name", func(c *gin.Context) {
        
        db_name  := c.Params.ByName("name")
        if database.Exists(db_name) == false{
            
            database.Create(db_name)
            c.JSON(200, gin.H{"name": c.Params.ByName("name")})
            return
        }
        
        c.JSON(409, gin.H{"name": c.Params.ByName("name"), "error": "Database already exists"})
    })
    
    r.PUT("/collection/create/:db/:col", func(c *gin.Context) {
        
        db_name  := c.Params.ByName("db")
        col_name := c.Params.ByName("col")
        
        if database.NamespaceExists(db_name, col_name) == false {
            
        }
        fmt.Println(database.NamespaceExists(db_name, col_name))
    })

    // Listen and server on 0.0.0.0:8080
    r.Run(":8089")
}
