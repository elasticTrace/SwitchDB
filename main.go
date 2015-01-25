package main

//import "log"
//import "database/sql"
//import "os"
//import "fmt"

import "./database"

import "github.com/gin-gonic/gin"
//import _ "github.com/mattn/go-sqlite3"



func main() {
    r := gin.Default()
    r.PUT("/database/:db_name/create", func(c *gin.Context) {
        
        db_name  := c.Params.ByName("db_name")
        if database.Exists(db_name) == false{
            
            database.Create(db_name)
            c.JSON(201, gin.H{"db": c.Params.ByName("db_name")})
            return
        }
        
        c.JSON(409, gin.H{"db": c.Params.ByName("db_name"), "error": "Database already exists"})
    })
    
    r.PUT("/database/:db_name/namespace/:ns_name/create/:ns_type", func(c *gin.Context) {
        
        db_name := c.Params.ByName("db_name")
        ns_name := c.Params.ByName("ns_name")
        ns_type := c.Params.ByName("ns_type")
        
        if database.Exists(db_name) == false {
            c.JSON(404, gin.H{"db": c.Params.ByName("db_name"), "error": "Database does not exists"})
            return
        }
        
        if database.NamespaceExists(db_name, ns_name) == false {
            if database.NamespaceCreate(db_name, ns_name, ns_type) == true {
                c.JSON(201, gin.H{"db": db_name, "ns": ns_name})
                return
            }
        }
        c.JSON(409, gin.H{"db": db_name, "ns": ns_name, "error": "Name space already exists"})
    })
    
    r.PUT("/insert/:db_name/:ns_name", func(c *gin.Context) {
        db_name := c.Params.ByName("db_name")
        ns_name := c.Params.ByName("ns_name")
        
        if database.Exists(db_name) == false {
            c.JSON(404, gin.H{"db": c.Params.ByName("db_name"), "error": "Database does not exists"})
            return
        }
        
        if database.NamespaceExists(db_name, ns_name) {
            c.JSON(404, gin.H{"db": db_name, "ns", ns_name, "error": "Namespace does nto exist in database"})
            return
        }
        
        
    })
    
    
    r.GET("/database/list/:db_name", func(c *gin.Context) {
        db_name := c.Params.ByName("db_name")
        list := database.NamespaceList(db_name, "")
        c.JSON(200, gin.H{"db": db_name, "ns": list})
        return
    })
    
    r.GET("/database/list", func(c *gin.Context) {
        list := database.List("")
        c.JSON(200, gin.H{"dbs": list})
    })
    
    /*r.PUT("/:db_name/:ns_name/insert", func(c *gin.Context) {
        
    })*/

    // Listen and server on 0.0.0.0:8080
    r.Run(":8089")
}
