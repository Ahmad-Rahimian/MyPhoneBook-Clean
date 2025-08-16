project-root/
│
├── cmd/                  # نقطه ورود برنامه
│   └── app/
│       └── main.go
│
├── internal/             # منطق اصلی برنامه (core)
│   ├── domain/           # مدل‌ها و اینترفیس‌های اصلی
│   │   └── contact.go
│   │
│   ├── repository/       # دسترسی به دیتابیس
│   │   └── contact_repository.go
│   │
│   ├── service/          # منطق بیزینسی
│   │   └── contact_service.go
│   │
│   ├── handler/          # کنترلرها / HTTP handlers
│   │   └── contact_handler.go
│   │
│   └── router/           # تعریف و رجیستر Routeها
│       └── router.go
│
│
|
├── pkg/ 
|   ├── db/        # پکیج‌های عمومی (مثل ابزارک‌ها)
│       └── db.go
|
├── go.mod
└── go.sum