@echo off

rem Check if the project was built
if not exist bin\myproject.exe goto :build

rem Run the project
bin\myproject.exe %*