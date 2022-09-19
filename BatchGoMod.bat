@echo off
@echo.
@echo MIT License
@echo.
@echo Copyright (c) 2022 yjhi
@echo.
@echo Permission is hereby granted, free of charge, to any person obtaining a copy
@echo of this software and associated documentation files (the "Software"), to deal
@echo in the Software without restriction, including without limitation the rights
@echo to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
@echo copies of the Software, and to permit persons to whom the Software is
@echo furnished to do so, subject to the following conditions:
@echo.
@echo The above copyright notice and this permission notice shall be included in all
@echo copies or substantial portions of the Software.
@echo.
@echo THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
@echo IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
@echo FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
@echo AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
@echo LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
@echo OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
@echo SOFTWARE.



set /p host=Input Host

echo ---------------------------------------------------------
cd jconfig
del go.mod
del go.sum
echo "%host%/jconfig" | ..\GoMod.bat
cd ..



echo ---------------------------------------------------------
cd jerrors
del go.mod
del go.sum
echo "%host%/jerrors" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------
cd jtime
del go.mod
del go.sum
echo "%host%/jtime" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------
cd jlog
del go.mod
del go.sum
echo "%host%/jlog" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------
cd jlog2
del go.mod
del go.sum
echo "%host%/jlog2" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------
cd jsql
del go.mod
del go.sum
echo "%host%/jsql" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------

cd jutils
del go.mod
del go.sum
echo "%host%/jutils" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------

cd jhttp
del go.mod
del go.sum
echo "%host%/jhttp" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------
cd test
del go.mod
del go.sum
echo "%host%/test" | ..\GoMod.bat
cd ..


echo ---------------------------------------------------------
del go.mod
del go.sum
echo "%host%" | GoMod.bat



echo Finished GoMod

pause
