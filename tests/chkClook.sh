rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all C-LOOK samples"
./ex c-look01.txt>c-look01.out
diff c-look01.base c-look01.out
./ex c-look20.txt>c-look20.out
diff c-look20.base c-look20.out
./ex c-lookPA.txt>c-lookPA.out
diff c-lookPA.base c-lookPA.out
