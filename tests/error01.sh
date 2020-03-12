rm ex # remove previous student executable
go build -o ex $1 # build new student executable

./ex error01.txt
