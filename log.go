package log

// A global variable so that log functions can be directly accessed
var log Logger

const (
	//Debug has verbose message
	Debug = "debug"
	//Info is default log level
	Info = "info"
	//Warn is for logging messages about possible issues
	Warn = "warn"
	//Error is for logging errors
	Error = "error"
	//Fatal is for logging fatal messages. The sytem shutsdown after
	//logging the message.
	Fatal = "fatal"
)

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

//Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})

	Infof(format string, args ...interface{})

	Warnf(format string, args ...interface{})

	Errorf(format string, args ...interface{})

	Fatalf(format string, args ...interface{})

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger

	Named(name string) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers,
// for such the level of Console is picked by default
type Configuration struct {
	ConsoleJSONFormat bool
	ConsoleLevel      string
}

func init() {
	c := &Configuration{}
	c.ConsoleLevel = Info
	log, _ = newZapLogger(c)
}

// Setup applies given configuration to the future instances of logger
func Setup(config *Configuration) error {
	l, err := newZapLogger(config)
	if err != nil {
		return err
	}
	log = l
	return nil
}

//New returns an instance of logger
func New(config *Configuration) (Logger, error) {
	logger, err := newZapLogger(config)
	if err != nil {
		return nil, err
	}
	log = logger
	return log, nil
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}

func Named(name string) Logger {
	return log.Named(name)
}
