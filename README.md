# tproccess

## Install the fyne

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
```

## Windows App build using Macbook M2 Pro

- Make sure there is a icon called `Icon.png` inside the cmd folder or root folder where the main function located.

**Install the dependencies:**

```bash
brew install mingw-w64
```

** Build EXE file using fyne**

```bash
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 fyne package -os windows -exe myapp.exe -src ./cmd
```

** Build EXE file using go build**

```bash
CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 CGO_ENABLED=1 go build -o myapp.exe ./cmd
```
