
## R file variables
tableDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/runWithGenomeSpliceSites/Human/sorted/stringent
plotDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/plots/Human_Stringent
scriptDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts


cp ${scriptDIR}/coverageSplitReads.R ${tableDIR}

cd ${tableDIR}

for gap in *.txt; do 

	gapTable=${gap}
	noGAP="$(echo $gap | sed 's/gapInRead//g')"
	makeName="${noGAP%.txt}"
	pdfGAP="${makeName}.pdf"
	plotName="human_${makeName}"
	## Run the R script for plotting

# echo $noGAP
# echo $makeName
# echo $plotName
# echo ""

# echo "table: ${tableDIR}/${gapTable}"
# echo "plot: ${plotDIR}/${pdfGAP} "
# echo "name on plot: ${plotName}" 
# echo ""

# echo ${gapTable}
# echo ${makeName}
# echo ${pdfGAP}
# echo ${plotName}
# echo 
	Rscript coverageSplitReads.R ${tableDIR}/${gapTable} ${plotDIR}/${pdfGAP} ${plotName}
	# args 1: table name
	# args 2: name of PDF
	# args 3: title on plot

done

rm coverageSplitReads.R