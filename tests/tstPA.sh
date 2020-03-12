rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

./ex fcfsPA.txt>check
./ex sstfPA.txt>>check
./ex scanPA.txt>>check
./ex c-scanPA.txt>>check
./ex lookPA.txt>>check
./ex c-lookPA.txt>>check

rm check
