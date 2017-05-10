##submits many bamReader.go scripts

bamDIR=/data/rc003/Brittany/Alignment/mouseBAM/normaliseRunTest
scriptDIR=/data/rc003/Brittany/Scripts

cd ${scriptDIR}
cp runBamReader.sh ${bamDIR}
cp bamReader.go ${bamDIR}
cd ${bamDIR}

bamList=$(ls *.bam.bai)

for bam in $bamList ; do

	echo "Running for $bam"
		scriptDIR=/data/rc003/Brittany/Scripts  bamDIR=/data/rc003/Brittany/Alignment/mouseBAM/normaliseRunTest  intDIR=/data/rc003/Brittany/Data/L1Location  refGenDIR=/data/rc003/Brittany/Data/genomes  outDIR=/data/rc003/Brittany/findGaps/testReadNormalise  sbatch runBamReader.sh ${bam}
done




#runRunBamReaderRun.sh