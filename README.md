# 🔧 Gel - A Git Clone Implementation in Go

Gel is a lightweight implementation of Git's core functionality written in Go. It provides essential version control operations with a clean, modular architecture inspired by Git's internal design.

## ✨ Features

- **Repository Initialization**: Create new Gel repositories
- **Object Storage**: Store and retrieve blobs, trees, and commits
- **File Tracking**: Add files to the staging area
- **Tree Operations**: Create and inspect directory trees
- **Commit Creation**: Generate commit objects with proper ancestry
- **Content Inspection**: Examine stored objects

## 🏗️ Architecture

Gel follows a layered architecture similar to Git:

```
gel/
├── cmd/gel/                    # CLI entry point
├── internal/
│   ├── core/                   # Core domain objects
│   │   ├── object/            # Git objects (blob, tree, commit)
│   │   └── repository/        # Repository abstraction
│   ├── plumbing/              # Low-level operations
│   │   ├── gel-path/          # Path utilities
│   │   └── storage/           # Storage abstraction
│   └── porcelain/             # High-level user commands
│       ├── add/               # File staging
│       ├── cat-file/          # Object inspection
│       ├── commit-tree/       # Commit creation
│       ├── hash-object/       # Object hashing
│       ├── init/              # Repository initialization
│       ├── ls-tree/           # Tree listing
│       └── write-tree/        # Tree creation
└── pkg/                       # Shared utilities
    ├── compression/           # Object compression
    ├── constant/              # Constants and messages
    └── hashing/               # SHA-1 hashing
```

## 🚀 Installation

```bash
# Clone the repository
git clone <repository-url>
cd gel

# Build the binary
go build -o gel cmd/gel/main.go

# Make it executable (optional)
chmod +x gel
```

## 📖 Usage

### Initialize a Repository
```bash
./gel init
```

### Add Files to Staging
```bash
./gel add <file>
```

### Create Object from File
```bash
./gel hash-object -w <file>
```

### Inspect Objects
```bash
./gel cat-file -p <hash>
```

### Create Tree from Current Directory
```bash
./gel write-tree
```

### List Tree Contents
```bash
./gel ls-tree <tree-hash>
```

### Create Commit
```bash
./gel commit-tree <tree-hash> -m "commit message"
./gel commit-tree <tree-hash> -p <parent-hash> -m "commit message"
```

## 🔧 Commands

| Command | Description | Status |
|---------|-------------|---------|
| `init` | Initialize a new Gel repository | ✅ |
| `add` | Add files to the staging area | ✅ |
| `hash-object -w` | Create and store object from file | ✅ |
| `cat-file -p` | Display object contents | ✅ |
| `write-tree` | Create tree object from current directory | ✅ |
| `ls-tree` | List contents of a tree object | ✅ |
| `commit-tree` | Create commit object | ✅ |

## 🗂️ Repository Structure

When you initialize a Gel repository, it creates:

```
.gel/
├── objects/           # Object storage (blobs, trees, commits)
│   ├── 00/           # Objects starting with 00
│   ├── 01/           # Objects starting with 01
│   └── ...           # More object directories
└── refs/             # References (branches, tags)
    └── heads/        # Branch references
```

## 🔍 Object Types

Gel supports three main object types:

### Blob Objects
Store file contents with compression and SHA-1 hashing.

### Tree Objects
Represent directory structures, containing references to blobs and other trees.

### Commit Objects
Contain:
- Tree SHA (snapshot of the repository)
- Parent commit SHA(s) (for history)
- Author and committer information
- Commit message

## 🏃‍♂️ Example Workflow

```bash
# Initialize repository
./gel init

# Create a file
echo "Hello, Gel!" > hello.txt

# Add to staging
./gel add hello.txt

# Create tree from current state
./gel write-tree
# Output: <tree-hash>

# Create initial commit
./gel commit-tree <tree-hash> -m "Initial commit"
# Output: <commit-hash>

# Make changes
echo "Hello, World!" > hello.txt
./gel add hello.txt

# Create new tree
./gel write-tree
# Output: <new-tree-hash>

# Create commit with parent
./gel commit-tree <new-tree-hash> -p <commit-hash> -m "Update greeting"
```

## 🔧 Development

### Prerequisites
- Go 1.24.0 or higher

### Building
```bash
go build -o gel cmd/gel/main.go
```

### Testing
```bash
go test ./...
```

### Project Structure
- **cmd/**: Command-line interface
- **internal/core/**: Domain objects and business logic
- **internal/plumbing/**: Low-level Git operations
- **internal/porcelain/**: High-level user-facing commands
- **pkg/**: Reusable utilities