# ====================================================
# 2025-08-27   WED
# 
# NOT WORKING in WSL2 or git bash.
# In WSL2 I need to install the latest golang.
# In git bash I need to install make.
# 
# This would probably work on the raspberry pi
# but it would take too long to compile on the pi.
# ====================================================

BINARY_NAME=myapp

all:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME}_linux_amd64 .
	GOOS=linux GOARCH=arm64 go build -o ${BINARY_NAME}_linux_arm64 .
	GOOS=linux GOARCH=arm GOARM=7 go build -o ${BINARY_NAME}_linux_armv7 .
	GOOS=linux GOARCH=arm GOARM=6 go build -o ${BINARY_NAME}_linux_armv6 .
	GOOS=windows GOARCH=amd64 go build -o ${BINARY_NAME}_windows_amd64.exe .

clean:
	rm -f ${BINARY_NAME}_*
