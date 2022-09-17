@echo off


set host="gitee.com/yjhi/go"

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
