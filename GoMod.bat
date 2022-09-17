@echo off

set /p module=Input Module Name:

go mod init %module%

go mod tidy

