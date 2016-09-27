#!/bin/bash
## Script extracts column containing splice acceptor sites from bamReader.go output

wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/limitedGap
outDIR=${wkDIR}/SpliceSites

if [ -d $outDIR ]; then
    echo "Folder $outDIR exists ..." 
else
    mkdir $outDIR
    echo "Folder $outDIR does not exist"     
    mkdir $outDIR
fi

cd ${wkDIR}

for file in gapInRead*.txt; do
	makeName=$(echo $file | sed 's/_/	/g' | awk '{print $2}')
	awk '{print $7}' $file > "${outDIR}/${makeName}_5.txt"
	awk '{print $8}' $file > "${outDIR}/${makeName}_3.txt"	
done

echo "complete"