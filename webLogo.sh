#!/bin/bash
# Script makes webLOGOs



wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Mouse/
plotDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/plots/webLogo
inTable=${wkDIR}/Three


cd ${wkDIR}/Five
for table in *.txt ; do 

fileName="${table%.txt}"
echo "making for ${fileName}"

	weblogo --number-interval 2 --aspect-ratio 4 -y "" -W 20 -c classic --format PDF -f ${table} -o ${plotDIR}/${fileName}.pdf
done

echo "complete"
