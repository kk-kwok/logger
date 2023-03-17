module code.hellotalk.com/infra/logger/gormlogger/v2

go 1.17

replace code.hellotalk.com/infra/logger/v2 => ../

require (
	code.hellotalk.com/infra/logger/v2 v2.0.7
	gorm.io/gorm v1.23.8
)

require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
)
