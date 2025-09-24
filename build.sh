#!/bin/bash
echo =================================
echo   GoLog Analyzer - Build Script
echo =================================
echo.

echo.
echo Compilation de l'application...
go build -o loganalyzer.exe .

echo L'application est prete a utiliser:
echo   .\loganalyzer.exe analyze --config config.json --output report.json
echo.
pause