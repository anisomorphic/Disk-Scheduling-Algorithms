rm ex # remove previous student executable
go build -o ex $1 # build new student executable
echo "error 01 req > upper or req < lower - produces ERROR 15"
./ex error01.txt
echo "error 02 upper < lower - produces ABORT 13" 
./ex error02.txt
echo "error 03 int < lower - produces ABORT 12"
./ex error03.txt
echo "error 04 upper < lower - produces ABORT 11"
./ex error04.txt
