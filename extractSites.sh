#!/bin/bash
## Script extracts column containing splice acceptor sites from bamReader.go output

wkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/
gapDIR=${wkDIR}/Split/gapInfoFiles
outDIR=${wkDIR}/Split/SpliceSites

# if [ -d $outDIR ]; then
#     echo "Folder $outDIR exists ..." 
# else
#     mkdir $outDIR
#     echo "Folder $outDIR does not exist"     
#     mkdir $outDIR
# fi

# cd ${wkDIR}

# for file in gapInRead*.txt; do

# 	makeName=$(echo $file | sed 's/_/	/g' | awk '{print $2}')

# 	cat ${file} | awk '{print $7 "\t" $8} '| sed -e 's/./&      /g' | awk '{print $1 $2 "\t" $3}' > "${outDIR}/${makeName}.txt"
		
# done


tableDIR=${outDIR}
plotDIR=${wkDIR}/plots/SpliceSitePlots


cd ${tableDIR}

for gapFile in *.txt ; do 
	echo $gapFile
	gapTable=${gapFile}
	makeFileName=$(echo $gapFile | sed 's/.txt//g' | sed s/-/\_/g)
	makeName=$(echo $gapFile | sed 's/.txt//g' | sed s/-/\ /g)
	pdfSplice5="${makeFileName}_5.pdf"
	pdfSplice3="${makeFileName}_3.pdf"
	



Rscript makeSplicePlot.R ${gapTable} ${plotDIR}/${pdfSplice3} ${makeFileName} ${plotDIR}/${pdfSplice5}


# args 1: table input
# args 2: pdf of table of 3' sites (mononucleotide)
# args 3: title on both plots: should be sample name
# args 4: pdf of table of 5' sites (dinucleotides)


done
 echo "complete"





