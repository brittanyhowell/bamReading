#!/bin/bash
# Script finds unique lines

wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Mouse/sorted/either
uniqueDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Mouse/sorted/eitherUnique

cd ${wkDIR}
for file in *.txt ; do
	name="${file%.txt}"
  awk '{print $2 "\t" $3 "\t" $4}' ${file} | sort | uniq > ${uniqueDIR}/${name}.bed


cd ${wkDIR}
done

cd ${uniqueDIR}
cat *.bed | sort | uniq > Unique-either.bed

echo "complete"