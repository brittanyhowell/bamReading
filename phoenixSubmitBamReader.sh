##submits many kamala.go scripts


# Source folders
bamDIR=/fast/users/a1646948/data/bam/human1045
scriptDIR=/fast/users/a1646948/scripts

# Out folder (make it the same name as the bam for consistency)
NAME="human1045"
outDIR=/fast/users/a1646948/gapTables/${NAME}

# Check BED folder exists, if not, make it
	if [ -d $outDIR ]; then
		rm -r $outDIR 
		mkdir $outDIR
		echo "OUT folder exists... replacing"  
	else 
		echo "creating OUT folder"  
		mkdir $outDIR
	fi 

cd ${scriptDIR}
# cp runKamala.sh ${bamDIR}
cp hrunKamala.sh ${bamDIR}
cp kamala.go ${bamDIR}
cd ${bamDIR}

bamList=$(ls *.bam.bai)

for bam in $bamList ; do

	echo "Running for $bam"
						scriptDIR=/fast/users/a1646948/scripts bamDIR=/fast/users/a1646948/data/bam/human1045 intDIR=/fast/users/a1646948/data/L1Location dataDIR=/fast/users/a1646948/data/SJMaps refGenDIR=/fast/users/a1646948/data/genomes outDIR=/fast/users/a1646948/gapTables/human1045/ sbatch hrunKamala.sh ${bam}
				# scriptDIR=/fast/users/a1646948/scripts bamDIR=/fast/users/a1646948/data/bam/mouse1045 intDIR=/fast/users/a1646948/data/L1Location dataDIR=/fast/users/a1646948/data/SJMaps refGenDIR=/fast/users/a1646948/data/genomes outDIR=/fast/users/a1646948/gapTables/mouse1045/ sbatch runKamala.sh ${bam}
done