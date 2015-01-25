package database

import "time"

type Namespace struct {
    Id           int64
    Name         string  `sql:"size:255;unique"`
    Type         string  `sql:"size:255"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    time.Time
}

type DocMeta struct {
    Id          string
    CreatedAt   time.Time
    UpdateAt    time.Time
}
