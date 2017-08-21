# load required libraries
library(IRanges)
library(ggplot2)
require(magrittr); require(tidyr) # to unite cols
library(reshape) # for melt
library(scales) # to access break formatting functions


setwd("~/Documents/University/Honours_2016/Project/Data/readData/humC/")


# covONLY section
  reads <- read.table(file = "intervals.txt")
  reads <- read.table(file = "INTNOT154529005.txt")
  colnames(reads) <- c("chr", "LStart", "LEnd", "mapStart", "mapEnd")
  
  start <- reads$mapStart
  end <- reads$mapEnd
  
  # intervals stored as an IRanges object
    intervals <- IRanges(start = start, end = end)
    cov <- coverage(intervals)
  
  # ggplot - iRanges
    bins <- disjointBins(IRanges(start(intervals), end(intervals) + 1))
    dat <- cbind(as.data.frame(intervals), bin = bins)
    
    ggplot(dat) + 
      geom_rect(aes(xmin = start, xmax = end,
                    ymin = bin, ymax = bin + 0.9 )) 
  
  # Line coverage plot
    plot(cov, type = "l", main = "All read coverage - human", xlab = "L1 coordinate", ylab = "number of reads")
    
    
# Read Summary Section
    readSummaryRaw <- read.table(file = "gapsIn24h_5FU_R1_readSummary.txt")
    colnames(readSummaryRaw) <- c("Chr", "lStart", "lEnd", "numReads", "splReads", "nonSplReads", "propSplice")
    readSummary <- readSummaryRaw[order(-readSummaryRaw$numReads, -readSummary$splReads),]
    
    # paste coordinates into the one column:
    readSummary %<>%
    unite( L1, lStart, lEnd, remove = TRUE)
 
    
    x <- readSummary$L1
    y <- readSummary$numReads
    z <- readSummary$splReads
    
    df <- melt(readSummary[,c('L1','numReads','splReads')],id.vars = 1)
    df$L1 <- factor(df$L1, levels=unique(as.character(df$L1)))
    
    ggplot(dfr,aes(x = L1 ,y = value+1)) + # value is + 1 because otherwise (due to log scale) 0 is negative.
      geom_bar(aes(fill = variable),stat = "identity",position = "dodge") + 
      scale_y_continuous(trans = "log10", breaks = trans_breaks("log10", function(x) 10^x)) +
      theme(axis.text.x=element_text(angle=60,hjust=1)) 
    
    
    