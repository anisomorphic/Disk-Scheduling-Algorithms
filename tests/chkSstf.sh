rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all SSTF samples"
./ex sstf01.txt>sstf01.out
diff sstf01.base sstf01.out
./ex sstf20.txt>sstf20.out
diff sstf20.base sstf20.out
./ex sstfPA.txt>sstfPA.out
diff sstfPA.base sstfPA.out
