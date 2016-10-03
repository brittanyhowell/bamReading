#!/bin/bash
## Some commands to get just the unique censor results from the map file


# First, find the L1s that are broken up and found in multiple sequences: 
awk '{print $1}' mouseUniqueL1.fasta.map | sort | uniq -c | sort
awk '{print $1}' humanUniqueActiveL1.fasta.map | sort | uniq -c | sort

# Identify those >1, and determine which bits to keep.
# All of the breaks of hits were less than 555bp. To confirm them, print only those smaller than 555bp
awk '{if ($3-$2 < 555) print $0}' mouseUniqueL1.fasta.map
awk '{if ($3-$2 < 2000) print $0}' humanUniqueActiveL1.fasta.map

# If they are the correct ones, flip the sign to a ">" sign and it will print just the good ones. 
awk '{if ($3-$2 > 555) print $0}' mouseUniqueL1.fasta.map > mouseUniqueL1_noBreak.map
awk '{if ($3-$2 > 2000) print $0}' humanUniqueActiveL1.fasta.map