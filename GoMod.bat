@echo off

set /p module=Input Module Name:
@echo.
go mod init %module%
go mod tidy

