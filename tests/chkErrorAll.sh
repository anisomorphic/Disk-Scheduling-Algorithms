rm ex # remove previous student executable
go build -o ex $1 # build new student executable
echo "error 01 req > upper or req < lower - produces ERROR 15"
./ex error01.txt >error01.out
diff error01.base error01.out
echo "error 02 upper < lower - produces ABORT 13" 
./ex error02.txt>error02.out
diff error02.base error02.out
echo "error 03 initial < lower - produces ABORT 12"
./ex error03.txt>error03.out
diff error03.base error03.out
echo "error 04 initial > upper - produces ABORT 11"
./ex error04.txt>error04.out
diff error04.base error04.out
