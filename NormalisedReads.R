# load required libraries
library(IRanges)
library(ggplot2)
require(magrittr); require(tidyr) # to unite cols
library(reshape) # for melt
library(scales) # to access break formatting functions


setwd("~/Documents/University/Honours_2016/Project/Data/readData/humC/")


# Plot all reads
  reads <- read.table(file = "intervals.txt")
  reads <- read.table(file = "INTNOT154529005.txt")
  colnames(reads) <- c("chr", "LStart", "LEnd", "mapStart", "mapEnd")
  
  start <- reads$mapStart
  end <- reads$mapEnd
  
  # intervals stored as an IRanges object
    intervals <- IRanges(start = start, end = end)
    cov <- coverage(intervals)
  
  # ggplot - stacked bar plot
    bins <- disjointBins(IRanges(start(intervals), end(intervals) + 1))
    dat <- cbind(as.data.frame(intervals), bin = bins)
    
    ggplot(dat) + 
      geom_rect(aes(xmin = start, xmax = end,
                    ymin = bin, ymax = bin + 0.9 )) 
  
  # line coverage plot
    plot(cov, type = "l", main = "All read coverage - human", xlab = "L1 coordinate", ylab = "number of reads")
    
    
# Read Summary Section - How many reads align to each L1 and how many are spliced.
    readSummaryRaw <- read.table(file = "gapsIn24h_5FU_R1_readSummary.txt")
    colnames(readSummaryRaw) <- c("Chr", "lStart", "lEnd", "numReads", "splReads", "nonSplReads", "propSplice")
    readSummary <- readSummaryRaw[order(-readSummaryRaw$numReads, -readSummaryRaw$splReads),] # have to order it, so that levels=unique works
    
    # paste coordinates into the one column:
    readSummary %<>%
    unite( L1, Chr, lStart, lEnd, remove = TRUE) # makes Chr, Lstart and Lend into one 'L1' col
 
    # makes separate rows based on numReads and splReads - allows for grouped columns
    df <- melt(readSummary[,c('L1','numReads','splReads')],id.vars = 1) 
    df$L1 <- factor(df$L1, levels=unique(as.character(df$L1))) # plots IN ORDER - so, it has to be sorted.
    
    ggplot(df,aes(x = L1 ,y = value+1)) + # value is + 1 because otherwise (due to log scale) 0 is negative.
      geom_bar(aes(fill = variable),stat = "identity",position = "dodge") + 
      xlab("L1 chr and interval") +
      ylab("reads aligned on L1") +
      scale_y_continuous(trans = "log10", breaks = trans_breaks("log10", function(x) 10^x)) +
      theme(axis.text.x=element_text(angle=50,hjust=1), plot.margin=unit(c(1,1,1.5,1.2),"cm")) # top, right, bottom, left
    

    