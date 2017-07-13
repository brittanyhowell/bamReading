#!/bin/bash
##Script finds all families from censor map file, then counts how many of each family are in the familiy. 

## A required file is a list of families. To find this file, use the following:
## It has been removed because I generated the file from all of the mice L1s (superset) and then wanted to count those families in the spliced subset.

# if [ -f allFam.txt ]; then
# 	rm allFam.txt
# 	echo "allFam deleted"
# fi
#
# echo "extracting families"
# awk '{print $4}' ${InFile} | sort | uniq > allFam.txt


wkDIR=/data/rc003/Brittany/censor/Human/Active

InFile="humanUniqueActiveL1_noBreak.fasta"

countFAM="numFamHumanSplice.txt"
# labelFAM="countHumanAll.txt"

cd ${wkDIR}

if [ -f ${countFAM} ]; then
	rm ${countFAM}
	touch ${countFAM}
    echo "${countFAM} exists, replacing"
else
    touch ${countFAM}
    echo "${countFAM} did not exist, now does" 
fi



# if [ -f ${labelFAM} ]; then
# 	rm ${labelFAM}
# 	echo "${labelFAM} deleted"
# fi




echo "counting families"
for fam in $(cat allFam.txt); do
	echo "finding ${fam}"
	egrep -c "$fam " ${InFile} >> ${countFAM}
done

# paste allFam.txt ${countFAM} > ${labelFAM}

## To create a table with uniq and all, repeat the script in both directories, then paste the files together.


## Calculate the number of samples in the columns
# awk '{print $2}' compareMouseFam.txt | paste -sd+ - | bc

### Mouse
## Then, use awk to calculate the percentage:
# awk '{printf("%s \t %s \t %.2f \t %s \t %.2f \n", $1,$2,$2/37.86,$3,$3/.92)}' compareMouseFam.txt > compareMouseFamP.txt
# awk '{print $1 " & " $2 " & " $3 " & " $4 " & " $5 "\\\\"}' compareMouseFamP.txt | sed s/\_/-/g > compareMouseFamP.tex
# cp compareMouseFamP.tex ../../../../../../Thesis/appendices/


### Human
## Then, use awk to calculate the percentage:
# awk '{printf("%s \t %s \t %.2f \t %s \t %.2f \n", $1,$2,$2/6.2,$3,$3/.12)}' compareHumFam.txt > compareHumFamP.txt
# awk '{print $1 " & " $2 " & " $3 " & " $4 " & " $5 "\\\\"}' compareHumFamP.txt | sed s/\_/-/g > compareHumFamP.tex
# cp compareHumFamP.tex ../../../../../../Thesis/appendices/
