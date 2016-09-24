#!/bin/bash
# Runs bamReader.go

# go run bamReader.go -index=sample5Change.bam.bai -bam=sample5Change.bam -intervalsBed=tiny.bed -outPath=./Split/ -outName=smallSplit.bed

WkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/BAMs
index="Mut-F2-Rep1_CGTACG_L007.STAR.10.45.Aligned.sortedByCoord.out.bam.bai"
bam="Mut-F2-Rep1_CGTACG_L007.STAR.10.45.Aligned.sortedByCoord.out.bam"

intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/L1Location
# intervalsBed="ClusActiveL1s.bed"
intervalsBed="ClusAllL1s.bed"

outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/BAMs/Split/
outFile="gapInReadCigar.bed"

# WkDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts
# index="sample5Change.bam.bai"
# bam="sample5Change.bam"

# intDIR=${WkDIR}
# intervalsBed="tiny.bed"

# outDIR=${WkDIR}/Split/
# outFile="splitShort.bed"


go run bamReader.go -index=${WkDIR}/${index} -bam=${WkDIR}/${bam} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outFile}
