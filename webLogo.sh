#!/bin/bash
# Script makes webLOGOs



wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/ClusterFilter/findGaps/outputGapFinding
inTable=${wkDIR}

site="combination"

plotDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/ClusterFilter/findGaps/webLogo/${site}



if [ -d $plotDIR ]; then
	rm -r $plotDIR 
	mkdir $plotDIR
	echo "plot folder exists... replacing" 
else 
	echo "creating plot folder" 
	mkdir $plotDIR
fi 



cd ${inTable}
for table in *.fasta ; do 

	fileName="${table%.fasta}"

	#noGAP="$(echo $fileName | sed 's/gapInRead//g')"
	noGAP="genome sites"
	echo $noGAP


	weblogo --number-interval 2 --aspect-ratio 4 -y "" --size large --title "$noGAP" --annotate '-2,-1,+1,+2' -W 30  -c classic --format PDF -f ${table} -U probability -o ${plotDIR}/${fileName}.pdf 
done

echo "complete"
