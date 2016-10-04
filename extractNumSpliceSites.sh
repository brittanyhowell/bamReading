#!/bin/bash
## Some commands to see what splice sites occur in samples

# First, make files from all of the different sample groups:
	cat *Ctrl* > control.txt
	cat *1mg* > OneMg.txt
	cat *2mg* > TwoMg.txt
	cat *5FU* > FiveFU.txt


# Next, extract just the unique splice sites found in those samples:	  
	awk '{print $5 "\t" $6}' FiveFU.txt | sort | uniq > uniqFiveFU.txt
	awk '{print $5 "\t" $6}' OneMg.txt | sort | uniq > uniqOneMg.txt
	awk '{print $5 "\t" $6}' TwoMg.txt | sort | uniq > uniqTwoMg.txt
	awk '{print $5 "\t" $6}' control.txt | sort |uniq > uniqControl.txt


==> uniqControl.txt <==
1068	2619
2272	3118
2768	3197
3757	4315
4651	5071

==> uniqFiveFU.txt <==
1068	2619
2272	3118
3757	4315
4651	5071

==> uniqOneMg.txt <==
1068	2271
1068	2619
2272	3118
3757	4315
4651	5071

==> uniqTwoMg.txt <==
1068	2619
2272	3118
3757	4315
4378	5421




#Check the splice sites that occur in any sample: 

	awk '{print $5 "\t" $6}' allGap.txt | sort | uniq > uniqAllGap.txt

==> uniqAllGap.txt <==
nm 	start	stop	Total
A	1068	2271	5 	
B	1068	2619	1081
C	2272 	3118	1915
D	2768	3197	14
E	3757	4315	1839
F	4378	5421	3
G	4651	5071	31


egrep -c "Insert interval" allGap.txt


# site1=""
# site2=""

# for file in gapInRead*.txt; do 
# egrep -c "${site1} \t ${site2}" allGap.txt 
# done


# Now, make files from all of the different time:
	cat *Ctrl* > control.txt
	cat *1mg* > OneMg.txt
	cat *2mg* > TwoMg.txt
	cat *5FU* > FiveFU.txt

# Next, extract just the unique splice sites found in those samples:	  
	awk '{print $5 "\t" $6}' FiveFU.txt | sort | uniq > uniqFiveFU.txt
	awk '{print $5 "\t" $6}' OneMg.txt | sort | uniq > uniqOneMg.txt
	awk '{print $5 "\t" $6}' TwoMg.txt | sort | uniq > uniqTwoMg.txt
	awk '{print $5 "\t" $6}' control.txt | sort |uniq > uniqControl.txt


	for file in gapInRead*.txt ; do 

		awk '{print $5 "\t" $6}' $file | sort | uniq > "uniq${file}"
	done












