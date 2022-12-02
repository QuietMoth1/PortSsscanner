build:
	go build -ldflags "-w -s" -o portssscanner portssscanner.go

run:
	go run portssscanner.go

compile:
	# 32-Bit Systems
	# FreeBDS
	GOOS=freebsd GOARCH=386 go build -ldflags "-w -s" -o portssscanner portssscanner.go
	# MacOS
	GOOS=darwin GOARCH=386 go build -ldflags "-w -s" -o portssscanner portssscanner.go
	# Linux
	GOOS=linux GOARCH=386 go build -ldflags "-w -s" -o portssscanner portssscanner.go
	# Windows
	GOOS=windows GOARCH=386 go build -ldflags "-w -s" -o portssscanner.exe portssscanner.go
	    # 64-Bit
	# FreeBDS
	GOOS=freebsd GOARCH=amd64 go build -ldflags "-w -s" -o portssscanner portssscanner.go
	# MacOS
	GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o portssscanner portssscanner.go
	# Linux
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o portssscanner portssscanner.go
	# Windows
	GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o portssscanner.exe portssscanner.go

