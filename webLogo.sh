#!/bin/bash
# Script makes webLOGOs



wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Human/
inTable=${wkDIR}/Five

site="Human-Five"

plotDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/plots/webLogo/${site}


if [ -d $plotDIR ]; then
	rm -r $plotDIR 
	mkdir $plotDIR
	echo "plot folder exists... replacing" 
else 
	echo "creating plot folder" 
	mkdir $plotDIR
fi 



cd ${inTable}
for table in *.txt ; do 

	fileName="${table%.txt}"

	noGAP="$(echo $fileName | sed 's/gapInRead//g')"
	echo $noGAP


	weblogo --number-interval 2 --aspect-ratio 4 -y "" --size large --title "$noGAP" --annotate '-3,-2,-1,+1,+2,+3' -W 30  -c classic --format PDF -f ${table} -U probability -o ${plotDIR}/${fileName}.pdf 
done

echo "complete"
