#!/bin/bash
## some criteria to count numbers of split reads according to different criteria.

species="Mouse"
wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/${species}
outDIR=${wkDIR}/CondensedTable/countNumReads
appDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Thesis/appendices/

FullDIR=${wkDIR}/FullTable
bothDIR=${wkDIR}/findActive/bothIntron
eitherDIR=${wkDIR}/findActive/either
stringDIR=${wkDIR}/findActive/stringent



if [ -d $outDIR ]; then
    echo "Folder $outDIR exists ... replacing"
    rm -r ${outDIR} 
    mkdir ${outDIR}
else
    mkdir $outDIR
    echo "Folder $outDIR does not exist... creating"     
    mkdir ${outDIR}
fi


# Count all lines
cd ${FullDIR}
	wc -l *.txt > ${outDIR}/Full.txt

cd ${bothDIR}
	wc -l *.txt > ${outDIR}/both.txt

cd ${eitherDIR}
	wc -l *.txt > ${outDIR}/either.txt	

cd ${stringDIR}
	wc -l *.txt > ${outDIR}/string.txt

cd ${outDIR}


# Extract just numbers
for file in *.txt ; do 
	awk '{print $1}' ${file} > "num${file}"
done

# Print sample names
cat Full.txt | awk '{print $2}' | sed s/gapInRead//g | sed s/.txt//g > samples.txt

paste samples.txt numFull.txt numeither.txt numboth.txt numstring.txt > reduceNum.txt

cat reduceNum.txt | awk '{print $1 " & " $2 " & " $3 " & " $4 " & " $5 " \\\\"}' | sed s/\_/-/g > reduce-${species}Num.tex
cp reduce-${species}Num.tex ${appDIR}

echo "complete"



