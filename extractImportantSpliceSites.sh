#!/bin/bash
# Script extracts splice junctions

outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Mouse/added
wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Mouse/FullTable
newDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Mouse/withSplice
sortedDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Mouse/sorted

cd ${wkDIR}
for file in *.txt ; do

	# extract sites to check, one with the four intronic sequences, one with the added bits
	stringentSplice="${file%.txt}_stringentSplice.txt"
	cat ${file} | awk '{print $7 "\t" $8}' | sed -e 's/./&      /g'  | awk '{print $2$3$4$5 "\t" $8$9$10}' > ${outDIR}/${stringentSplice}

	partSplice="${file%.txt}_PartSplice.txt"
 	cat ${file} | awk '{print $7 "\t" $8}' | sed -e 's/./&      /g'  | awk '{print $4$5 "\t" $8$9}' > ${outDIR}/${partSplice}

 	cd ${outDIR}

 	# Paste them together 
 	fullSplice="${file%.txt}_FullSplice.txt"
 	paste ${stringentSplice} ${partSplice} > ${fullSplice}

 	rm ${partSplice}
 	rm ${stringentSplice}
 	mv ${fullSplice} ${wkDIR}

 	cd ${wkDIR}

 	# Append them to the original file
 	fullFile="${file%.txt}_WithSplice.txt"
 	paste ${file} ${fullSplice} > ${newDIR}/${fullFile}

	rm ${fullSplice}

  	cd ${newDIR}
	
	echo "checking for canonical SJ"
 # 	# Check if the splice site is canonical
 # 	awk '{if ( $14 == "AGGT" && $15 == "AGG") print $0}' ${fullFile} > ${sortedDIR}/"${file%.txt}_bothStringent.txt"
	# awk '{if ( $16 == "GT" && $17 == "AG") print $0}' ${fullFile} > ${sortedDIR}/"${file%.txt}_bothIntron.txt"
 # 	awk '{if ( $16 == "GT" || $17 == "AG") print $0}' ${fullFile} > ${sortedDIR}/"${file%.txt}_eitherIntron.txt"

 	awk '{if ( $14 == "aggt" && $15 == "agg") print $0}' ${fullFile} > ${sortedDIR}/"${file%.txt}_bothStringent.txt"
	awk '{if ( $16 == "gt" && $17 == "ag") print $0}' ${fullFile} > ${sortedDIR}/"${file%.txt}_bothIntron.txt"
 	awk '{if ( $16 == "gt" || $17 == "ag") print $0}' ${fullFile} > ${sortedDIR}/"${file%.txt}_eitherIntron.txt"
 # 	# 14: four base 5' splice site: AGGT
 # 	# 15: three base 3' splice site: AGG
 # 	# 16: two base 5' intronic splice site: GT
 # 	# 17: two base 3' intronic splice site: AG


cd ${wkDIR}
done

cd ${sortedDIR}
mv *_bothIntron.txt ./bothIntron
mv *_bothStringent.txt ./stringent 
mv *_eitherIntron.txt ./either 

echo "complete"





