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

	cat ${file} | awk '{print $7 "\t" $8} '| sed -e 's/./&      /g' | awk '{print $1 $2 "\t" $3}' > "${outDIR}/${makeName}.txt"
		
done

echo "complete"





