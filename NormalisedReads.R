

setwd("~/Documents/University/Honours_2016/Project/Data/readData/testing")
# load required libraries
library(IRanges)
library(ggplot2)


### Summary reads section

  # summary reads table
  reads <- read.table(file = "gapsIn24h_1mg_R2.STAR.5.25.bam_readSummary.txt")
  colnames(reads) <- c("chromosome", "cStart", "cEnd", "numReads", "numSpliced", "numNonSpliced", "pSpliced") 
  
  percentSpliced <- reads$pSpliced
  pSpliceTrunc <- pSpliceAll [pSpliceAll>0 & pSpliceAll<100]
  pSplice <- na.omit(pSpliceTrunc)
  
  numberReads <- reads$numReads
  numberSplicedReads <- reads$numSpliced
  
  ## The plot of percent spliced and whatnot
  ggplot(data=reads, aes(x= numberSplicedReads, y= reads$numNonSpliced))+#, colour = percentSpliced))+
    geom_point(shape=5) +
    # geom_smooth(method=lm) +
    coord_cartesian(xlim= c(0,100), ylim= c(0,100)) +
    # scale_colour_gradientn(colours=rainbow(4))
  
  plot(x = reads$numReads, y = reads$numSpliced, xlim = c(0,1000), ylim= c(0,1000))



### Plotting all of the reads
  
  # all reads
  allReads <- read.table("gapsIn24h_1mg_R2.STAR.5.25.bam_reads.txt")
  colnames(allReads) <- c("Chr", "L1PosStart", "L1PosEnd", "lStart", "lEnd", "cStart", "cEnd", "cigar" )
  
  start = allReads$lStart
  end = allReads$lEnd
  
  # intervals stored as an IRanges object
  intervals <- IRanges(start = start, end = end)
  cov <- coverage(intervals)
  
  ## Function : Plots boxy box plots
  plotRanges <- function(x, xlim = x, main = "boxy reads",
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
  


   
   # cov <- as.vector(cov)
   # mat <- cbind(seq_along(cov)-0.5, cov)
   # 
   # d <- diff(cov) != 0
   # mat <- rbind(cbind(mat[d,1]+1, mat[d,2]), mat)
   # mat <- mat[order(mat[,1]),]
   # lines(mat, col="red", lwd=4)
   # axis(2)
  
  # Plot coverage - all reads
  plot(cov, type = "l")
  # ggplot(data = intervals) # THIS DOES NOT DO ANYTHING :/
  
  
  
  ## Bringin' in the split reads
  
  splitReads <- read.table("./gapsIn24h_1mg_R2.STAR.5.25.bam.txt")
  colnames(splitReads) <- c("ID", "chr", "cStart", "cEnd", "lStart", "lEnd", "5nucs", "3nucs", "m5nucs", "m3nucs", "gapLen", "flags")
  
  splitStart <- splitReads$lStart
  splitEnd <- splitReads$lEnd
  
  splitIntervals  <- IRanges(start = splitStart, end = splitEnd)
  splitCov <- coverage(splitIntervals)
  
  plotRanges(splitIntervals) # Oh so boxy
  
  # Coverage plot with both
  par(mfrow=c(1,1), mar=c(5,5,4,4))
  plot(splitCov, type = "l", col = "blue", ylim = c(0,13000), xlab = "Coordinate on L1", ylab = "read coverage")
  lines(cov, type = "l", col = "red")
  
  
  # Combining plots (multiPanel)
  par(mfrow=c(2,1), mar=c(1,1,1,1))
  plot(cov, type = "l", xlim = c(0,10000), xaxt = 'n', xlab = '')
  par(mar = c(4, 2, 1, 1))
  plot(splitCov, type = "l", xlim = c(0,10000))
  

  