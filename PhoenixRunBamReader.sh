#!/bin/bash
# Runs bamReader.go
## runBamReader.sh
## Script is the batch script which is called by "runPhoenixRunBamReader" and hence the filepath variables need to be declared not here, but in the call. They are presented here still only for ease. 
## invoked by: scriptDIR=/data/rc003/Brittany/Scripts bamDIR=/data/rc003/Brittany/Alignment/mouseBAM intDIR=/data/rc003/Brittany/Data/L1Location refGenDIR=/data/rc003/Brittany/genomes outDIR=/data/rc003/Brittany/findGaps/outputGapFinding/ sbatch test.sh Mut-F2-Rep1_CGTACG_L007.STAR.10.45.bam.bai

#SBATCH -p batch
#SBATCH -N 1 
#SBATCH -n 8 
#SBATCH --time=0-08:00
#SBATCH --mem=10GB 

# Notification configuration 
#SBATCH --mail-type=END                                         
#SBATCH --mail-type=FAIL           
#SBATCH --mail-user=brittany.howell@student.adelaide.edu.au    


## Variables
	# Variables for mouse:

		# Filepath variables:
		scriptDIR=/data/rc003/Brittany/Scripts/ 
		bamDIR=/data/rc003/Brittany/Alignment/mouseBAM/normaliseRunTest
		intDIR=/data/rc003/Brittany/Data/L1Location/ 
		dataDIR=/data/rc003/Brittany/Data/sjMaps/
		refGenDIR=/data/rc003/Brittany/Data/genomes/
		outDIR=~/musNorm/

		# Non-filepath variables
		intervalsBed="L1_Mouse_merge_sort_ORF2only-bothORF.bed"
		refGen="mm10.fa"
		SJMap5="SJMap5.txt"
		SJMap3="SJMap3.txt"

	

	# Variables for human:

		# # Filepath variables:
		# scriptDIR=/data/rc003/Brittany/Scripts/ 
		# bamDIR=/data/rc003/Brittany/humanAlignment
		# intDIR=/data/rc003/Brittany/Data/L1Location/ 
		# dataDIR=/data/rc003/Brittany/Data/sjMaps/
		# refGenDIR=/data/rc003/Brittany/Data/genomes/
		# outDIR=~/humNormC/

		# # Non-filepath variables
		# # intervalsBed="human_L1_bothORF.bed"
		# intervalsBed="human_L1_ORF2_bothORF.bed"
		# refGen="hg38.fa"
		# SJMap5="SJMap5.txt"
		# SJMap3="SJMap3.txt"



cd ${bamDIR}

	bamRecord=${1%.bai}
	bam=${1}

outPrefix="gapsIn${bamRecord%.STAR.10.45.bam}"
echo "Running for ${outPrefix}"
 	go run kamala.go -index=${bamDIR}/${bam} -bam=${bamDIR}/${bamRecord} -intervalsBed=${intDIR}/${intervalsBed} -outPath=${outDIR} -outName=${outPrefix}_splitReads.txt -seqOutName=${outPrefix}_FullIntron.txt -refGen=${refGenDIR}/${refGen}  -readName=${outPrefix}_reads.txt -readSumName=${outPrefix}_readSummary.txt -SJMap5=${dataDIR}/${SJMap5} -SJMap3=${dataDIR}/${SJMap3} -report=${outPrefix}_report.txt


echo "Complete table"