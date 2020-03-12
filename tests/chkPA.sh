rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

./ex fcfsPA.txt>fcfsPA.out
diff fcfsPA.base fcfsPA.out
./ex sstfPA.txt>sstfPA.out
diff sstfPA.base sstfPA.out
./ex scanPA.txt>scanPA.out
diff scanPA.base scanPA.out
./ex c-scanPA.txt>c-scanPA.out
diff c-scanPA.out c-scanPA.base
./ex lookPA.txt>lookPA.out
diff lookPA.base lookPA.out
./ex c-lookPA.txt>c-lookPA.out
diff c-lookPA.out c-lookPA.base
