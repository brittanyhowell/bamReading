#!/bin/bash
# Runs bamReader.go


bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/BAMs/transfer
# index="Mut-F5-Rep1_ACAGTG_L008.STAR.10.45.bam.bai"
# bam="Mut-F5-Rep1_ACAGTG_L008.STAR.10.45.bam"

intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/L1Location
# intervalsBed="ClusActiveL1s.bed"
intervalsBed="ClusAllL1s.bed"

outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split/limitedGap/
# outFile="gapInRead_Mut_F5_Rep1_10_45.txt"

# # testDIRs
# bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams

# # index="sample5Change.bam.bai"
# # bam="sample5Change.bam"

# intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# # intervalsBed="ClusActiveL1s.bed"
# intervalsBed="tiny.bed"

# outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# # outFile="gapInRead_testing.txt"


cd ${bamDIR}
bamFiles=$(ls *.bam.bai)

for bam in $bamFiles ; do
	bamRecord=${bam%.bai}
# The go code
# go run bamReader.go -index=${bamDIR}/${index} -bam=${bamDIR}/${bam} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outFile}


 go run bamReader.go -index=${bamDIR}/${bam} -bam=${bamDIR}/${bamRecord} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=gapInRead_${bamRecord%.STAR.10.45.bam}.txt
done 

echo "finished making tables"

## R file variables
tableDIR=${outDIR}
plotDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/plots/full



cd ${outDIR}

for gap in *.txt; do 

	gapTable=${gap}
	makeName=$(echo $gap | sed 's/_/	/g' | awk '{print $2}')
	pdfGAP="${makeName}.pdf"
	plotName="STAR-${makeName}"
	## Run the R script for plotting

# echo ${gapTable}
# echo ${makeName}
# echo ${pdfGAP}
# echo ${plotName}
# echo .
	Rscript coverageSplitReads.R ${tableDIR}/${gapTable} ${plotDIR}/${pdfGAP} ${plotName}
	# args 1: table name
	# args 2: name of PDF
	# args 3: title on plot

done