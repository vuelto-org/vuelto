@echo off
cmake . -G "Unix Makefiles"
make -j8
.\bin\test\test_program.exe
del .\bin\test\test_program.exe
