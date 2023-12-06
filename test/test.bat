@echo off
cmake . -G "Unix Makefiles"
make
.\bin\test\test_program.exe
del .\bin\test\test_program.exe
