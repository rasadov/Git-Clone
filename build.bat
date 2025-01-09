@echo off

rem check the project directory
set /p projectDir=Enter the project directory or leave empty for current path (e.g. C:\Projects\MyProject):
if "%projectDir%"=="" set projectDir=%cd%
cd %projectDir%

rem check the build directory
if not exist %projectDir%\bin mkdir %projectDir%\bin

rem build the project
cd src
go.exe build -o %projectDir%\bin\myproject.exe %projectDir%\src\main.go
if %errorlevel% neq 0 goto :error
