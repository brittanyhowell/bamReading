# make smaller bam

 sed -n 1,30p sample100.sam > sample5.sam

 # Convert sam to bam
samtools view -b sample5.sam > sample5.bam

# Convert bam to sam
samtools view -h sample5.bam > sample5.sam

# Index bam
samtools index -b sample5.bam sample5.bam.bai