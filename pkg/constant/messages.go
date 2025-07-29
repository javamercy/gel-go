package constant

const (
	ERR_USAGE_MESSAGE      = "usage: gel <command> [<args>...]\n"
	ERR_UNKNOWN_COMMAND    = "Unknown command %s\n"
	ERR_GEL_NOT_REPOSITORY = "not a gel repository (or any of the parent directories): .gel"
)

const (
	ERR_CREATING_DIRECTORY = "Error creating directory: %s\n"
	ERR_WRITING_FILE       = "Error writing file: %s\n"
	ERR_WRITING_TREE       = "Error writing tree: %s\n"
)

const (
	ERR_GEL_PATH_NOT_SET = "gel path not set"
	ERR_INVALID_HASH     = "invalid hashing"
)

const (
	MSG_INITIALIZED_GEL_DIRECTORY = "Initialized gel directory"
	MSG_WRITING_TREE_FOR          = "Writing tree for %s\n"
)
