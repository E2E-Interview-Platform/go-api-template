package environment

// This files lists all the environment variables used in init and set at shell level
const (
	// value: json or empty
	// Used in logger for deciding format
	LOG_FORMAT_KEY = "LOG_FORMAT"

	// value: dir name or empty
	// specifies the dir to log
	LOG_DIR_KEY = "LOG_DIR"
)
