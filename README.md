# Gel - A Git Implementation in Go

Gel is a simplified Git implementation written in Go that provides core Git functionality including repository initialization, object storage, and basic Git commands.

## 🚀 Features

- **Repository Management**: Initialize Git repositories with `gel init`
- **Object Storage**: Store and retrieve Git objects (blobs, trees) with SHA-1 hashing
- **Core Commands**: 
  - `gel init` - Initialize a new repository
  - `gel hash-object -w <file>` - Store file as Git object and return hash
  - `gel cat-file -p <hash>` - Print Git object contents
  - `gel ls-tree --name-only <hash>` - List tree contents
  - `gel write-tree` - Create tree object from current directory

## 🏗️ Architecture

### Project Structure

```
gel/
├── cmd/gel/                    # Main application entry point
│   └── main.go                 # CLI command handling
├── internal/
│   ├── core/
│   │   ├── object/             # Git object implementations
│   │   │   ├── blob.go         # Blob object type
│   │   │   ├── object.go       # Base object interface
│   │   │   └── tree.go         # Tree object type
│   │   └── repository/         # Repository management
│   │       └── repository.go   # Repository operations
│   ├── plumbing/               # Low-level Git operations
│   │   ├── gitpath/            # Git directory path management
│   │   │   └── path.go         # Path discovery and management
│   │   └── storage/            # Object storage layer
│   │       ├── filesystem.go   # Filesystem-based storage
│   │       └── storage.go      # Storage interface
│   └── porcelain/              # High-level Git commands
│       ├── add/                # File staging operations  
│       │   └── add.go          # Add command implementation
│       ├── cat-file/           # Object inspection
│       │   └── cat_file.go     # Cat-file command
│       └── init/               # Repository initialization
│           └── init.go         # Init command implementation
├── pkg/                        # Shared utilities
│   ├── compression/            # Data compression utilities
│   │   └── compression.go      # Zlib compression/decompression
│   └── hashing/                # Cryptographic hashing
│       └── hashing.go          # SHA-1 hashing utilities
└── constant/                   # Application constants
    ├── constants.go            # Git object types and constants
    └── messages.go             # Error messages and strings
```

### Key Components

#### 1. Object Storage System
- **Hash-based storage**: Objects stored using SHA-1 hash in `.gel/objects/`
- **Object types**: Support for blob and tree objects
- **Compression**: Uses zlib compression for efficient storage

#### 2. Path Management
- **Repository discovery**: Automatically finds `.gel` directory by walking up directory tree
- **Lazy initialization**: Path discovery happens only when needed
- **Thread-safe**: Uses `sync.Once` for concurrent access safety

#### 3. Command System
- **Modular design**: Each command implemented as separate package
- **Error handling**: Comprehensive error handling and user feedback
- **Git compatibility**: Commands behave similarly to standard Git

## 🛠️ Installation

### Prerequisites
- Go 1.24.0 or higher

### Build from Source

```bash
# Clone the repository
git clone <repository-url>
cd gel

# Build the application
go build -o gel cmd/gel/main.go

# Or install globally
go install ./cmd/gel
```

## 📖 Usage

### Initialize a Repository

```bash
# Initialize a new gel repository in current directory
./gel init
```

This creates a `.gel` directory with the necessary structure:
```
.gel/
└── objects/    # Object storage directory
```

### Store Files as Objects

```bash
# Store a file as a Git object and get its hash
./gel hash-object -w filename.txt
# Output: a1b2c3d4e5f6... (SHA-1 hash)
```

### Inspect Objects

```bash
# View the contents of a Git object
./gel cat-file -p a1b2c3d4e5f6...
```

### Work with Trees

```bash
# Create a tree object from current directory
./gel write-tree
# Output: tree_hash...

# List contents of a tree object
./gel ls-tree --name-only tree_hash...
```

## 🔧 Development

### Running Tests

```bash
go test ./...
```

### Code Structure Guidelines

- **Internal packages**: Core functionality in `internal/` directory
- **Clean architecture**: Separation between plumbing (low-level) and porcelain (high-level) operations
- **Interface-driven design**: Use interfaces for testability and modularity
- **Error handling**: Always return and handle errors appropriately

### Adding New Commands

1. Create a new package under `internal/porcelain/`
2. Implement the command logic
3. Add command handling to `cmd/gel/main.go`
4. Update this README with command documentation

## 🧰 Technical Details

### Object Storage Format

Objects are stored in `.gel/objects/` using the following structure:
- Path: `.gel/objects/ab/cdef123...` (first 2 chars as directory, rest as filename)
- Content: zlib-compressed object data
- Format: `<type> <size>\0<content>`

### Supported Object Types

- **Blob**: File content storage
- **Tree**: Directory structure with file/directory entries

### Hash Algorithm

- Uses SHA-1 for object identification (same as Git)
- Content-addressable storage ensures data integrity

