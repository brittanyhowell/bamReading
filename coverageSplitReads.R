args = commandArgs(TRUE)

# test if there is at least one argument: if not, return an error
if (length(args)==0) {
  stop("At least one argument must be supplied (input file).n", call.=FALSE)
} else if (length(args)==1) {
  # default output file
  args[2] = "out.txt"
}


#setwd("~/Documents/University/Honours_2016/Project/Data/BAMs/Split/")
# load IRanges library
library(IRanges)


split <- read.table(file = args[1])
#split <- read.table("./gapInReadCigar.bed")
colnames(split) <- c("read name", "chromosome", "startL1", "endL1", "startGap", "endGap", "length", "cigar", "flag")

start = split$startGap
end = split$endGap

# intervals stored as an IRanges object
intervals <- IRanges(start = start, end = end)

plotRanges <- function(x, xlim = x, main = args[3],
                          col = "black", sep = .1, ...)
   {
  height <- .2
  if (is(xlim, "Ranges"))
    xlim <- c(min(start(xlim)), max(end(xlim)))
    bins <- disjointBins(IRanges(start(x), end(x) + 1))
    plot.new()
    plot.window(xlim, c(0, max(bins)*(height + sep)))
    ybottom <- bins * (sep + height) - height
    rect(start(x)-0.5, ybottom, end(x)+0.5, ybottom + height, col = col, ...)
    title(main)
    axis(1)
}

pdf(args[2], width = 10, height = 6)
plotRanges(intervals)
graphics.off()





