package database

import "time"

type Namespace struct {
    Id           int64
    Name         string  `sql:"size:255"`
    Type         string  `sql:"size:255"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    //DeletedAt    time.Time
}
