@echo off

set PROJECT_ROOT=.
set SRC=%PROJECT_ROOT%\pb\proto
set TAR=%PROJECT_ROOT%\pb

protoc --proto_path=%SRC%  --go_out=%TAR% --go_opt=paths=import %SRC%\*.proto
pause