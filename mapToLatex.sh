#!/bin/bash
## Some commands to make latex tables out of the map file from CENSOR

# place amphisands between each column: 
cat  humanUniqueActiveL1.fasta.map | awk '{print $1 " & " $2 " & " $3 " & " $5 " & " $6 " & " $7 " & " $8 " & " $9 " & " $10 " & " $11 " & " $12}'

# Pretty simple, if you want, write a script around it to echo it all to a file. 
