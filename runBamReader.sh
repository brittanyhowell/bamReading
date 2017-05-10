#!/bin/bash
# Runs bamReader.go
# Updated version can be found at "PhoenixRunBamReader.sh" because running this in a for loop is really not advisable. Don't do it.

# scriptDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/
# bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/BAMs/mutSplice/
# # # index="Mut-F5-Rep1_ACAGTG_L008.STAR.10.45.bam.bai"
# # # bam="Mut-F5-Rep1_ACAGTG_L008.STAR.10.45.bam"


# intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/L1Location/
# intervalsBed="ClusActiveL1s.bed"
# # intervalsBed="ClusAllL1s.bed"

# refGenDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/genomeAssemblies
# refGen="mm10.fa"

# outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/mouse/
# outPrefix=gapInRead${bamRecord%.STAR.10.45.bam}



# testDIRs - for when a quick run is needed
# bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# intervalsBed="tiny.bed"
# outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams/
# outPrefix=gapInRead${bamRecord%.STAR.10.45.bam}
# refGenDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/genomeAssemblies
# refGen="mm10.small"


# testDIRs - for when a quick run is needed
bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams/OneTest
intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
intervalsBed="tiny.bed"
outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams/sampleOutput/
outPrefix=gapInRead${bamRecord}
refGenDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/genomes
refGen="mm10Baby.fa"

bam=test5.bam.bai

	bamRecord=${bam%.bai}

 	go run bamReader.go -index=${bamDIR}/${bam} -bam=${bamDIR}/${bamRecord} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outPrefix}.txt -seqOutName=${outPrefix}_FullIntron.txt -refGen=${refGenDIR}/${refGen} -logo5Name=${outPrefix}_5SJ.txt -logo3Name=${outPrefix}_3SJ.txt
 
# done 
# rm bamReader.go

# echo "I finished making tables"


# go run bamReader.go -index=BH090415.Aligned.sortedByCoord.out.bam.bai -bam=BH090415.Aligned.sortedByCoord.out.bam -intervalsBed=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/humanL1Location/convert/genbankBoth.bed -outPath=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/BAMs/mutSplice/out/ -outName=spliceMut.txt -seqOutName=spliceMut_FullIntron.txt  -logo5Name=spliceMut_5SJ.txt -logo3Name=spliceMut_3SJ.txt 


# go run bamReader.go -index=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/testSetBAM/smallBAM.bam.bai -bam=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/testSetBAM/smallBAM.bam -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outPrefix}.txt -seqOutName=${outPrefix}_FullIntron.txt -refGen=${refGenDIR}/${refGen} -logo5Name=${outPrefix}_5SJ.txt -logo3Name=${outPrefix}_3SJ.txt