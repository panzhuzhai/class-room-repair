package logger

// Config represents the setting for zap logger.
//type Config struct {
//	ZapConfig zap.Config        `json:"zap_config" yaml:"zap_config"`
//	LogRotate lumberjack.Logger `json:"log_rotate" yaml:"log_rotate"`
//}

//
//// Logger is an alternative implementation of *gorm.Logger
//type Logger interface {
//	GetZapLogger() *zap.SugaredLogger
//	LogMode(level gormLogger.LogLevel) gormLogger.Interface
//	Info(ctx context.Context, msg string, data ...interface{})
//	Warn(ctx context.Context, msg string, data ...interface{})
//	Error(ctx context.Context, msg string, data ...interface{})
//	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
//}
//
//type logger struct {
//	Zap *zap.SugaredLogger
//}

// NewLogger is constructor for logger
//func NewLogger(sugar *zap.SugaredLogger) Logger {
//	return &logger{Zap: sugar}
//}

// InitLogger create logger object for *gorm.DB from *echo.Logger
//func InitLogger(yamlFile embed.FS) {
//	configYaml, err := yamlFile.ReadFile(constants.LoggerConfigPath)
//	if err != nil {
//		fmt.Printf("Failed to read logger configuration: %s", err)
//		os.Exit(constants.ErrExitStatus)
//	}
//	var myConfig *Config
//	if err = yaml.Unmarshal(configYaml, &myConfig); err != nil {
//		fmt.Printf("Failed to read zap logger configuration: %s", err)
//		os.Exit(constants.ErrExitStatus)
//	}
//	var zap *zap.Logger
//	zap, err = build(myConfig)
//	if err != nil {
//		fmt.Printf("Failed to compose zap logger : %s", err)
//		os.Exit(constants.ErrExitStatus)
//	}
//	sugar := zap.Sugar()
//	// set package varriable logger.
//	GormLogger = NewLogger(sugar)
//	ZapLogger = sugar
//	//log.GetZapLogger().Infof("Success to read zap logger configuration: zaplogger.yml")
//	_ = zap.Sync()
//	//return log
//}
//
//// GetZapLogger returns zapSugaredLogger
//func (log *logger) GetZapLogger() *zap.SugaredLogger {
//	return log.Zap
//}
