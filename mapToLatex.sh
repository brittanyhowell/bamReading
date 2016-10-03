#!/bin/bash
## Some commands to make latex tables out of the map file from CENSOR

# place amphisands between each column, place \\ at end, and remove _s
cat  humanUniqueActiveL1.fasta.map | awk '{print "{\\footnotesize " $1 "} & " $2 " & " $3 " & " $4 " & "  $5 " & " $6 " & " $7 " & " $8 " & " $9 " & " $10 " & " $11 " & " $12 " \\\\"}' | sed 's/\_/ /g' > latexHumanActiveFamilies.map 

cat mouseUniqueL1_noBreak.map | awk '{print "{\\footnotesize " $1 "} & " $2 " & " $3 " & " $4 " & "  $5 " & " $6 " & " $7 " & " $8 " & " $9 " & " $10 " & " $11 " & " $12 " \\\\"}' | sed 's/\_/ /g' > latexMouseActiveFamilies.map

# Pretty simple, if you want, write a script around it to echo it all to a file. 
