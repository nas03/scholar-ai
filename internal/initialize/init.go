package initialize

// Bootstrap initializes all required services for the application
func Bootstrap() error {
	// Load configuration first
	LoadConfig()

	// Initialize logger first (errors are logged internally)
	InitLogger()
	defer SyncLogger()

	// Initialize services (errors are logged internally)
	InitGorm()
	InitMailClient()
	InitRedis()

	return nil
}
