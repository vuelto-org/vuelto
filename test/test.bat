@echo off
cmake .
make
.\bin\test\test_program
del .\bin\test\test_program
