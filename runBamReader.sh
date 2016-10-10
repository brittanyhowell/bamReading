#!/bin/bash
# Runs bamReader.go
# Updated version can be found at "PhoenixRunBamReader.sh" because running this in a for loop is really not advisable. Don't do it.

scriptDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/
bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/BAMs/transfer/
# # index="Mut-F5-Rep1_ACAGTG_L008.STAR.10.45.bam.bai"
# # bam="Mut-F5-Rep1_ACAGTG_L008.STAR.10.45.bam"


intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/L1Location/
intervalsBed="ClusActiveL1s.bed"
# intervalsBed="ClusAllL1s.bed"

refGenDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/genomeAssembliess
refGen="mm10.fa"

outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/mouse/
outPrefix=gapInRead${bamRecord%.STAR.10.45.bam}



# testDIRs - for when a quick run is needed
# bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# intervalsBed="tiny.bed"
# outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams/
# outPrefix=gapInRead${bamRecord%.STAR.10.45.bam}
# refGenDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/genomeAssemblies
# refGen="mm10.small"

cp bamReader.go ${bamDIR}

cd ${bamDIR}
bamFiles=$(ls *.bam.bai)

for bam in $bamFiles ; do
	bamRecord=${bam%.bai}

 	go run bamReader.go -index=${bamDIR}/${bam} -bam=${bamDIR}/${bamRecord} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outPrefix}.txt -seqOutName=${outPrefix}_FullIntron.txt -refGen=${refGenDIR}/${refGen} -logo5Name=${outPrefix}_5SJ.txt -logo3Name=${outPrefix}_3SJ.txt
 
done 
rm bamReader.go

echo "I finished making tables"
