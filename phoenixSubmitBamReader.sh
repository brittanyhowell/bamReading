##submits many bamReader.go scripts

bamDIR=/data/rc003/Brittany/Alignment/mouseBAM/normaliseRunTest
scriptDIR=/data/rc003/Brittany/Scripts

cd ${scriptDIR}
cp runBamReader.sh ${bamDIR}
cp kamala.go ${bamDIR}
cd ${bamDIR}

bamList=$(ls *.bam.bai)

for bam in $bamList ; do

	echo "Running for $bam"
				# scriptDIR=/data/rc003/Brittany/Scripts  bamDIR=/data/rc003/Brittany/humanStrictAlignment/testNormalise  intDIR=/data/rc003/Brittany/Data/L1Location  refGenDIR=/data/rc003/Brittany/Data/genomes outDIR=~/humNorm2/ dataDIR=/data/rc003/Brittany/Data/ sbatch runBamReader.sh ${bam}
				scriptDIR=/data/rc003/Brittany/Scripts/  bamDIR=/data/rc003/Brittany/Alignment/mouseBAM/normaliseRunTest intDIR=/data/rc003/Brittany/Data/L1Location/  dataDIR=/data/rc003/Brittany/Data/sjMaps/ refGenDIR=/data/rc003/Brittany/Data/genomes/ outDIR=~/musNorm/ sbatch runBamReader.sh ${bam}
done