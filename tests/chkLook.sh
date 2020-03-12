rm *.out #remove previous student output files
rm ex # remove previous student executable
go build -o ex $1 # build new student executable

echo "run all LOOK samples"
./ex look01.txt>look01.out
diff look01.base look01.out
./ex look20.txt>look20.out
diff look20.base look20.out
./ex lookPA.txt>lookPA.out
diff lookPA.base lookPA.out
