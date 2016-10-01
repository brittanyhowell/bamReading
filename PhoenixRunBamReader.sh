#!/bin/bash
# Runs bamReader.go
## runBamReader.sh
## invoked by: scriptDIR=/data/rc003/Brittany/Scripts bamDIR=/data/rc003/Brittany/Alignment/mouseBAM intDIR=/data/rc003/Brittany/Data/L1Location refGenDIR=/data/rc003/Brittany/genomes outDIR=/data/rc003/Brittany/findGaps/outputGapFinding/ sbatch test.sh Mut-F2-Rep1_CGTACG_L007.STAR.10.45.bam.bai

#SBATCH -p batch
#SBATCH -N 1 
#SBATCH -n 8 
#SBATCH --time=0-04:00
#SBATCH --mem=50GB 

# Notification configuration 
#SBATCH --mail-type=END                                         
#SBATCH --mail-type=FAIL           
#SBATCH --mail-user=brittany.howell@student.adelaide.edu.au    




# scriptDIR=/data/rc003/Brittany/Scripts
# bamDIR=/data/rc003/Brittany/Alignment/mouseBAM

# intDIR=/data/rc003/Brittany/Data/L1Location
intervalsBed="ClusActiveL1s.bed"
# intervalsBed="ClusAllL1s.bed"

# refGenDIR=/data/rc003/Brittany/genomes
refGen="mm10.fa"

# outDIR=/data/rc003/Brittany/findGaps/outputGapFinding/











# testDIRs
# bamDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# intDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams
# intervalsBed="tiny.bed"
# outDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/bamReading/Scripts/sampleBams/
# outPrefix=gapInRead${bamRecord%.STAR.10.45.bam}
# refGenDIR=/Users/brittanyhowell/Documents/University/Honours_2016/Project/Data/genomeAssemblies
# refGen="mm10.small"



cd ${bamDIR}

	bamRecord=${1%.bai}
	bam=${1}

outPrefix="gapInRead${bamRecord%.STAR.10.45.bam}"

 	go run bamReader.go -index=${bamDIR}/${bam} -bam=${bamDIR}/${bamRecord} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outPrefix}.txt -seqOutName=${outPrefix}_FullIntron.txt -refGen=${refGenDIR}/${refGen} -logo5Name=${outPrefix}_5SJ.txt -logo3Name=${outPrefix}_3SJ.txt
 




echo "Complete table"
