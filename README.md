# 🎬 Video Renamer

A fast and reliable CLI tool written in Go that renames video files based on their actual capture date extracted from metadata using ExifTool.

Instead of relying on file names or modification dates, Video Renamer reads metadata directly from video files, sorts them chronologically, and renames them in a clean sequential format.

Example:

```text
VID_8452.MP4  ->  0001.mp4
VID_8453.MP4  ->  0002.mp4
VID_8454.MP4  ->  0003.mp4
```

---

## ✨ Features

- Extracts recording timestamps using ExifTool
- Sorts videos by actual capture time
- Sequential renaming (`0001.mp4`, `0002.mp4`, ...)
- Supports mixed video formats
- Dry-run mode for safe previews
- Two-phase rename process to avoid filename conflicts
- Cross-platform (Linux, macOS, Windows)
- Fast batch metadata extraction
- No cloud services or external APIs

---

## 📦 Requirements

### ExifTool

Video Renamer requires ExifTool to be installed and available in your PATH.

Check installation:

```bash
exiftool -ver
```

---

## 🐧 Linux

### Ubuntu / Debian

```bash
sudo apt update
sudo apt install libimage-exiftool-perl
```

### Arch Linux

```bash
sudo pacman -S perl-image-exiftool
```

### Fedora

```bash
sudo dnf install perl-Image-ExifTool
```

Verify:

```bash
exiftool -ver
```

---

## 🍎 macOS

Using Homebrew:

```bash
brew install exiftool
```

Verify:

```bash
exiftool -ver
```

---

## 🪟 Windows

Download ExifTool from:

https://exiftool.org

Add the executable directory to your PATH.

Verify:

```powershell
exiftool -ver
```

---

# 🚀 Installation

## Download Release

Download the latest binary from the GitHub Releases page.

Linux:

```bash
chmod +x video-renamer-linux
./video-renamer-linux
```

Windows:

```powershell
video-renamer.exe
```

---

## Build From Source

### Clone

```bash
git clone https://github.com/isnot00/video-renamer.git
cd video-renamer
```

### Build

```bash
go build -o video-renamer ./cmd
```

### Run

```bash
./video-renamer
```

---

# 🛠 Build Targets

Using Makefile:

Build current platform:

```bash
make build
```

Build Windows binary:

```bash
make windows
```

Build Linux binary:

```bash
make linux
```

Build all releases:

```bash
make release
```

Run tests:

```bash
make test
```

Clean build artifacts:

```bash
make clean
```

---

# 📖 Usage

## Dry Run

Preview changes without modifying files.

```bash
video-renamer --dir ./videos --dry-run
```

Example output:

```text
VID_1001.MP4 -> 0001.mp4
VID_1002.MP4 -> 0002.mp4
VID_1003.MP4 -> 0003.mp4
```

No files are modified.

---

## Execute Rename

Apply changes to files.

```bash
video-renamer --dir ./videos
```

or

```bash
video-renamer --dir ./videos --execute
```

The tool will display a preview before renaming.

---

# 🔄 How It Works

1. Scan the target directory
2. Collect video files
3. Extract metadata using ExifTool
4. Read capture timestamps
5. Sort videos chronologically
6. Generate sequential names
7. Preview results
8. Rename files safely

---

# 📂 Example

Before:

```text
Holiday_003.mp4
VID_8452.mp4
Vacation.mp4
```

Capture dates:

```text
2022-06-10 12:05:21
2022-06-10 12:08:44
2022-06-10 12:15:02
```

After:

```text
0001.mp4
0002.mp4
0003.mp4
```

---

# ⚠ Safety

Video Renamer uses a two-phase rename strategy:

```text
original -> temporary
temporary -> final
```

This prevents filename collisions during the renaming process.

Always run a dry run first:

```bash
video-renamer --dir ./videos --dry-run
```

---

# 🧪 Testing

Run all tests:

```bash
go test ./...
```

Run verbose tests:

```bash
go test -v ./...
```

---

# 🏗 Project Structure

```text
cmd/
internal/
├── exif/
├── model/
├── renamer/
├── scanner/
└── ui/

Makefile
README.md
```

---

# 🤝 Contributing

Pull requests, issues, and suggestions are welcome.

If you find a bug or have an idea for an improvement, please open an issue.

---

# 📄 License

MIT License

---

# 🙏 Acknowledgements

- ExifTool by Phil Harvey
- Go Programming Language