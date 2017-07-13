args = commandArgs(TRUE)

# test if there is at least one argument: if not, return an error
if (length(args)==0) {
  stop("At least one argument must be supplied (input file).n", call.=FALSE)
} else if (length(args)==1) {
  # default output file
  args[2] = "out.txt"
}





splices <- read.table(args[1])
colnames(splices) <- c("five", "three" )
monocolours <- c("darkgrey", "darkgrey", "cornflowerblue", "darkgrey")
dicolours <- c("darkgrey","darkgrey", "cornflowerblue", "darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey","darkgrey")

num <- nrow(splices)

pdf(args[2], width = 10, height = 6)
  plot(splices$three, main = args[3], col = monocolours)
  abline(h=num/4)
graphics.off()

pdf(args[4], width = 10, height = 6)
  plot(splices$five, main = args[3], col = dicolours)
  abline(h = num/16, col = 1)
graphics.off()

#p = sum(width(intervals))/(length(coverage(intervals))*n)

# p <- (1/16)
# monoP <- (1/4)
# abline(h= num/4 - (sqrt((p*(1-p))/num)), lty = 2) # mean - 2 sd
# 
# library(stringr)
# numAA <- sum(str_count(splices$site, "AA"))
# numAC <- sum(str_count(splices$site, "AC"))
# numAG <- sum(str_count(splices$site, "AG"))
# numAT <- sum(str_count(splices$site, "AT"))
# numCA <- sum(str_count(splices$site, "CA"))
# numCC <- sum(str_count(splices$site, "CC"))
# numCG <- sum(str_count(splices$site, "CG"))
# numCT <- sum(str_count(splices$site, "CT"))
# numGT <- sum(str_count(splices$site, "GT"))
# numTG <- sum(str_count(splices$site, "TG"))
# numTT <- sum(str_count(splices$site, "TT"))
# prop.test(numAA,num,1/16)
# prop.test(numAC,num,1/16)
# prop.test(numAG,num,1/16)
# prop.test(numAT,num,1/16)
# prop.test(numCA,num,1/16)
# prop.test(numCC,num,1/16)
# prop.test(numCG,num,1/16)
# prop.test(numCT,num,1/16)
# prop.test(numGT,num,1/16)
# prop.test(numTG,num,1/16)
# prop.test(numTT,num,1/16)

