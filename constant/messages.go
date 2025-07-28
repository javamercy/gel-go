package constant

const (
	ERR_USAGE_MESSAGE      = "usage: mygit <command> [<args>...]\n"
	ERR_UNKNOWN_COMMAND    = "Unknown command %s\n"
	ERR_GIT_NOT_REPOSITORY = "not a git repository (or any of the parent directories): .git"
)

const (
	ERR_CREATING_DIRECTORY = "Error creating directory: %s\n"
	ERR_WRITING_FILE       = "Error writing file: %s\n"
	ERR_WRITING_TREE       = "Error writing tree: %s\n"
)

const (
	ERR_GIT_PATH_NOT_SET = "git gitpath not set"
	ERR_INVALID_HASH     = "invalid hashing"
)

const (
	MSG_INITIALIZED_GIT_DIRECTORY = "Initialized git directory"
	MSG_WRITING_TREE_FOR          = "Writing tree for %s\n"
)
