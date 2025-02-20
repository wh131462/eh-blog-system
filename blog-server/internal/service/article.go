func InitDB(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // 自动迁移
    db.AutoMigrate(&Article{})

    return db, nil
}