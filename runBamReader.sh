#!/bin/bash
# Runs bamReader.go


bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/BAMs
# index="Mut-F2-Rep1_CGTACG_L007.STAR.10.45.Aligned.sortedByCoord.out.bam.bai"
index="Mut-F2-Rep1_CGTACG_L007.tophat2_pe.mm10.bam.bai"
# bam="Mut-F2-Rep1_CGTACG_L007.STAR.10.45.Aligned.sortedByCoord.out.bam"
bam="Mut-F2-Rep1_CGTACG_L007.tophat2_pe.mm10.bam"

intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/L1Location
# intervalsBed="ClusActiveL1s.bed"
intervalsBed="ClusAllL1s.bed"

outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split
outFile="TopHatgapInReadCigar.bed"

## R file variables
tableDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Split
gapTable="TopHatgapInReadCigar.bed"
outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/plots
pdfGAP="Tophat.pdf"

# The go code
# go run bamReader.go -index=${bamDIR}/${index} -bam=${bamDIR}/${bam} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outFile}




## Run the R script for plotting

Rscript coverageSplitReads.R ${tableDIR}/${gapTable} ${outDIR}/${pdfGAP}
# args 1: table name
# args 2: name of PDF
