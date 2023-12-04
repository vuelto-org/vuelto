@echo off
cmake .
make
.\bin\test\test_program.exe
del .\bin\test\test_program.exe
