
setwd("~/Documents/University/Honours_2016/Project/bamReading/Split/Human/FullTable/")

library(ggplot2)
library(reshape2)

dat <- read.table("gapInRead0h_Ctrl_R2.txt")

colnames(dat) <- c("name", "chr", "cStart", "cEnd", "lStart", "lEnd", "nuc", "nuc2", "nuc3", "nuc4", "len", "cigar", "flags")


head(dat)

all <- c(dat$lStart, dat$lEnd)

mAll <- append(dat$lStart, dat$lEnd)

hist(mAll, breaks = 80, xlim = c(0, 8000))
hist(dat$lStart, breaks = 80, xlim = c(0, 8000))
hist(dat$lEnd, breaks = 80, xlim = c(0, 8000))



require(gridExtra)


plot1 <- qplot(mAll, geom = "histogram", binwidth = 100, xlim = c(0,7000), xlab = "All regions")
plot2 <- qplot(dat$lStart, geom = "histogram", binwidth = 100, xlim = c(0,7000), xlab = "Start of splices")
plot3 <- qplot(dat$lEnd, geom = "histogram", binwidth = 100, xlim = c(0,7000), xlab = "End of splices")

grid.arrange(plot1, plot2, plot3, nrow=3)

