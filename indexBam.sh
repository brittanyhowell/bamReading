#!/bin/bash
## Script indexes bams
## Date: 22-9-2016

# Invoked by: sbatch indexBam.sh

#SBATCH -p batch
#SBATCH -N 1 
#SBATCH -n 8 
#SBATCH --time=0-04:00
#SBATCH --mem=10GB 

# Notification configuration 
#SBATCH --mail-type=END                                         
#SBATCH --mail-type=FAIL           
#SBATCH --mail-user=brittany.howell@student.adelaide.edu.au      


# Load the necessary modules
module load SAMtools/1.3.1-GCC-5.3.0-binutils-2.25

cd /data/rc003/Brittany/Alignment/bamMouse

samtools index -b AllMut.STAR.Aligned.sortedByCoord.out.bam AllMut.STAR.Aligned.sortedByCoord.out.bam.bai

# samtools index -b AllWT.STAR.Aligned.sortedByCoord.out.bam AllWT.STAR.Aligned.sortedByCoord.out.bam.bai



