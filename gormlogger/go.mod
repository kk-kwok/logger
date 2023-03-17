module github.com/kk-kwok/logger/gormlogger

go 1.20

replace github.com/kk-kwok/logger => ../

require (
	github.com/kk-kwok/logger v1.0.0
	gorm.io/gorm v1.23.8
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	go.opentelemetry.io/otel v1.8.0 // indirect
	go.opentelemetry.io/otel/trace v1.8.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
)
