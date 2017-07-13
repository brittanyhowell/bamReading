args = commandArgs(TRUE)

# test if there is at least one argument: if not, return an error
if (length(args)==0) {
  stop("At least one argument must be supplied (input file).n", call.=FALSE)
} else if (length(args)==1) {
  # default output file
  args[2] = "out.txt"
}



setwd("~/Documents/University/Honours_2016/Project/bamReading/Split/Mouse/CondensedTable/")
# load IRanges library
library(IRanges)

# Reading in the final stringent tables
split <- read.table("./stringent/UniqueMouse_stringent.txt")
split <- read.table("./stringent/UniqueMouse_stringent_relative.txt")
colnames(split) <- c("lStart", "lEnd")
colnames(split) <- c("chromosome", "gStart", "gEnd", "lStart", "lEnd")

start = split$lStart
end = split$lEnd

# intervals stored as an IRanges object
intervals <- IRanges(start = start, end = end)

plotRanges <- function(x, xlim = x, main = "",
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


plotRanges(intervals)






