package seed

import (
    "fmt"
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID        uint   `gorm:"primaryKey"`
    Name      string
    Email     string
    CreatedAt time.Time
}

type Order struct {
    ID        uint      `gorm:"primaryKey"`
    UserID    uint      `gorm:"not null"`
    TotalPrice float64  `gorm:"not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func SeedDB(db *gorm.DB) {
   
    users := []User{
        {Name: "John Doe", Email: "john@example.com", CreatedAt: time.Now()},
        {Name: "Jane Smith", Email: "jane@example.com", CreatedAt: time.Now()},
    }

    for _, user := range users {
        db.FirstOrCreate(&user)
    }

    orders := []Order{
        {UserID: 1, TotalPrice: 100.50, CreatedAt: time.Now(), UpdatedAt: time.Now()},
        {UserID: 2, TotalPrice: 200.75, CreatedAt: time.Now(), UpdatedAt: time.Now()},
    }

    for _, order := range orders {
        db.FirstOrCreate(&order)
    }

    fmt.Println("Database seeded successfully")
}

