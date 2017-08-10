

setwd("~/Documents/University/Honours_2016/Project/Data/readData/humanSJnucTest/")
setwd("~/Documents/University/Honours_2016/Project/Data/readData/mouse/")
setwd("~/Documents/University/Honours_2016/Project/Data/readData/mouseSJTest/")
# load required libraries
library(IRanges)
library(ggplot2)


# covONLY section
reads <- read.table(file = "Ints.txt")
colnames(reads) <- c("chr", "LStart", "LEnd", "mapStart", "mapEnd")

start <- reads$mapStart
end <- reads$mapEnd
# intervals stored as an IRanges object
intervals <- IRanges(start = start, end = end)
cov <- coverage(intervals)

## ggplot - iRanges

bins <- disjointBins(IRanges(start(intervals), end(intervals) + 1))
dat <- cbind(as.data.frame(intervals), bin = bins)

ggplot(dat) + 
  geom_rect(aes(xmin = start, xmax = end,
                ymin = bin, ymax = bin + 0.9 )) 



### Summary reads section
chrX <- read.table(file = "chrXread.txt")
colnames(chrX) <- 

  # summary reads table
  reads <- read.table(file = "gapsIn24h_1mg_R2_readSummary.txt")
  reads <- read.table(file = "gapsInMut-F2-Rep1_CGTACG_L007_readSummary.txt")
  colnames(reads) <- c("chromosome", "cStart", "cEnd", "numReads", "numSpliced", "numNonSpliced", "pSpliced") 

  percentSpliced <- reads$pSpliced
  pSpliceTrunc <- pSpliceAll [pSpliceAll>0 & pSpliceAll<100]
  pSplice <- na.omit(pSpliceTrunc)
  
  numberReads <- reads$numReads
  numberSplicedReads <- reads$numSpliced
  
  ## The plot of percent spliced and whatnot
  ggplot(data=reads, aes(x= reads$numReads, y= reads$numSpliced))+#, colour = percentSpliced))+
    geom_point(shape=5) +
    # geom_smooth(method=lm) +
    coord_cartesian(xlim= c(0,100), ylim= c(0,100)) # +
    # scale_colour_gradientn(colours=rainbow(4))
  
  plot(x = reads$numReads, y = reads$numSpliced, xlim = c(0,1000), ylim= c(0,1000))


  

### Plotting all of the reads
  chrX <- read.table(file = "chrXread.txt")
  colnames(chrX) <-  c("Chr", "L1PosStart", "L1PosEnd", "lStart", "lEnd", "cStart", "cEnd", "cigar" )
  start = chrX$lStart
  end = chrX$lEnd 
   
  # all reads
  sallReads <- read.table("../humanSJnucTest/gapsIn24h_1mg_R2_reads.txt")
  sallReads <- read.table("../humanSJnucTest/gapsIn24h_1mg_R2_splitReads.txt")
  sallReads <- read.table("gapsInMut-F2-Rep1_CGTACG_L007_reads.txt")
  splitMouse <- read.table("gapsInMut-F2-Rep1_CGTACG_L007_splitReads.txt")
  colnames(splitMouse) <-c("ID", "chr", "cStart", "cEnd", "lStart", "lEnd", "class5", "class3","Nclass5", "Nclass3", "gapLen", "cigar", "flags")
  colnames(sallReads) <- c("Chr", "L1PosStart", "L1PosEnd", "lStart", "lEnd", "cStart", "cEnd", "cigar" )
  colnames(sallReads) <- c("ID", "chr", "cStart", "cEnd", "lStart", "lEnd", "class5", "class3", "gapLen", "cigar", "flags")
  
  allReads <- sallReads[order(sallReads$lStart),]
 
  start = sallReads$lStart
  end = sallReads$lEnd
  
  start = splitMouse$lStart
  end = splitMouse$lEnd
  
  # intervals stored as an IRanges object
  intervals <- IRanges(start = start, end = end)
  cov <- coverage(intervals)
  
  ## ggplot - iRanges
  
  bins <- disjointBins(IRanges(start(intervals), end(intervals) + 1))
  dat <- cbind(as.data.frame(intervals), bin = bins)
  
  ggplot(dat) + 
    geom_rect(aes(xmin = start, xmax = end,
                  ymin = bin, ymax = bin + 0.9 )) 
  
  ## Function : Plots boxy box plots
  plotRanges <- function(x, xlim = x, main = "",
                         col = "black", sep = .1, ...)
  {
    height <- .2
    if (is(xlim, "Ranges"))
      xlim <- c(min(start(xlim)), 10000)
    bins <- disjointBins(IRanges(start(x), end(x) + 1))
    plot.new()
    plot.window(xlim, c(0, max(bins)*(height + sep)))
    ybottom <- bins * (sep + height) - height
    rect(start(x)-0.5, ybottom, end(x)+0.5, ybottom + height, col = col, ...)
    title(main)
    axis(1)
  }
  
  plotRanges(intervals) 
  
# max x coord: max(end(xlim))

   
   # cov <- as.vector(cov)
   # mat <- cbind(seq_along(cov)-0.5, cov)
   # 
   # d <- diff(cov) != 0
   # mat <- rbind(cbind(mat[d,1]+1, mat[d,2]), mat)
   # mat <- mat[order(mat[,1]),]
   # lines(mat, col="red", lwd=4)
   # axis(2)
  
  # Plot coverage - all reads
  plot(cov, type = "l", main = "mouse coverage", xlab = "L1 coordinate", ylab = "mapped section reads")
  # ggplot(data = intervals) # THIS DOES NOT DO ANYTHING :/
  
  
  
  ## Bringin' in the split reads
  
  splitReads <- read.table("./gapsIn24h_1mg_R2_splitReads.txt")
  splitReads <- read.table("gapsInMut-F2-Rep1_CGTACG_L007_splitReads.txt")
  colnames(splitReads) <- c("ID", "chr", "cStart", "cEnd", "lStart", "lEnd", "class5", "class3", "gapLen", "cigar", "flags")
  colnames(splitReads) <- c("ID", "chr", "cStart", "cEnd", "lStart", "lEnd", "class5", "class3","Nclass5", "Nclass3", "gapLen", "cigar", "flags")
  
  # Making frequency table

  freqSplits <- splitReads[,c(5:11)]
  dfAll <- as.data.frame(freqSplits)
  df <- dfAll[dfAll$gapLen > 0,]
  df <- dfAll[dfAll$class5 >16 &  dfAll$class5 < 33,]
  library(plyr)
  counts <- ddply(df, .(df$lStart, df$lEnd, df$gapLen, df$class5, df$class3, df$Nclass5, df$Nclass3), nrow)
  names(counts) <- c("start", "end", "len","5'SJ", "3'SJ","n5'SJ", "n3'SJ","freq")

  cols <- c("firebrick1", "brown1", "chocolate", "orange1", "gold1", "darkolivegreen1", "darkolivegreen4", "olivedrab1", "lightseagreen", "mediumseagreen", "paleturquoise1", "powderblue", "royalblue1", "royalblue4", "slateblue3", "slateblue1")
  SJ3 <- c("AA", "AG", "AC", "AT", "GA", "GG", "GT", "CA", "CG", "CC", "CT", "TA", "TG", "TC", "TT")
  
ggplot(counts) +
    geom_segment(aes(x = counts$start, y = counts$freq, xend = counts$end, yend = counts$freq, color = factor(counts$`3'SJ`)), size = 2, alpha=.5, data = counts) +
    # scale_y_continuous(trans = "log10", minor_breaks = seq(0, 100000, 100)) +
     # scale_colour_brewer(palette = "Set1") +
    labs(title = "mouse 10:45", x = "L1 coordinate", y = "Reads supporting gap", fill="Cat") +
    # coord_cartesian(xlim = NULL, ylim = c(10, 2500)) +
    # scale_color_manual(values = cols ,labels=SJ3)+
    theme(legend.title=element_blank()) 



## ggplot for all reads

freqReads <- allReads[,c(4:5)]
dfRAll <- as.data.frame(freqReads)
dfR <- dfRAll[dfRAll$gapLen > 200,]
dfR <- dfRAll[dfRAll$class5 >16 &  dfAll$class5 < 33,]
library(plyr)
countsR <- ddply(dfRAll, .(dfRAll$lStart, dfRAll$lEnd ), nrow)
names(countsR) <- c("start", "end", "freq")

cols <- c("firebrick1", "brown1", "chocolate", "orange1", "gold1", "darkolivegreen1", "darkolivegreen4", "olivedrab1", "lightseagreen", "mediumseagreen", "paleturquoise1", "powderblue", "royalblue1", "royalblue4", "slateblue3", "slateblue1")
SJ3 <- c("AA", "AG", "AC", "AT", "GA", "GG", "GT", "CA", "CG", "CC", "CT", "TA", "TG", "TC", "TT")

ggplot(countsR) +
  geom_segment(aes(x = countsR$start, y = countsR$freq, xend = countsR$end, yend = countsR$freq), size = 1, alpha=.5, data = countsR) +
  scale_y_continuous(trans = "log10", minor_breaks = seq(0, 100000, 100)) +
  # scale_colour_brewer(palette = "Set1") +
  labs(title = "human MCF7 - all reads", x = "L1 coordinate", y = "Reads", fill="Cat") +
  coord_cartesian(xlim = NULL, ylim = c(10, 2500)) +
  # scale_color_manual(values = cols ,labels=SJ3)+
  theme(legend.title=element_blank()) 



  
 
  
  # subsetting data
  newdata <- splitReads[ which(splitReads$class5=='249'), ]
  newdata <- subset(splitReads, class5==127)
   newdata <- splitReads[splitReads$class5  > 19 & splitReads$class5 <33,]
  new3data <- splitReads[splitReads$class3 == 8,]
   
   
  ## Subsetted
   splitStart <- new3data$lStart
   splitEnd <- new3data$lEnd  
   
  ## Not subsetted
  splitStart <- splitReads$lStart
  splitEnd <- splitReads$lEnd
  
  splitIntervals  <- IRanges(start = splitStart, end = splitEnd)
  splitCov <- coverage(splitIntervals)
  
  plotRanges(splitIntervals) # Oh so boxy
  
  # ggplot of reads
  df <- as.data.frame(splitReads)
  x <- df$lStart
  xEnd <- df$lEnd
  
  ggplot(df) +
    geom_segment(aes(x = x, y = 1, xend = xEnd, yend = 100, colour = df$class3), data = df)
    
  ## ggplot - iRanges
  
  bins <- disjointBins(IRanges(start(splitIntervals), end(splitIntervals) + 1))
  dat <- cbind(as.data.frame(splitIntervals), bin = bins)
  
  ggplot(dat) + 
    geom_rect(aes(xmin = start, xmax = end,
                  ymin = bin, ymax = bin + 0.9 )) 
  
  # Coverage plot with both
  par(mfrow=c(1,1), mar=c(5,5,4,4))
  plot(splitCov, type = "l", col = "blue", ylim = c(0,13000), xlab = "Coordinate on L1", ylab = "read coverage")
  lines(cov, type = "l", col = "red")
  
  lines(splitCov, type = "l", col = "blue", ylim = c(0,13000), xlab = "Coordinate on L1", ylab = "read coverage")
  
  
  # Combining plots (multiPanel)
  par(mfrow=c(2,1), mar=c(1,4,1,1))
 plot(cov, type = "l", xlim = c(0,10000), xaxt = 'n', xlab = '') # All reads - coverage
  par(mar = c(4, 4, 1, 1))
  # plot(splitCov, type = "l", xlim = c(0,10000)) # Coverage
  plotRanges(splitIntervals) # Oh so boxy

  
  ### Matrix things
  
  
mat <- matrix(sample(c(0,1), size = 1000, replace = TRUE), 
                nrow = 100)
image(t(mat[nrow(mat):1,]), col = c("black", "white"))

mat[1:10][2:3][1]

mat[1:length(mat)]

mat[11]



HCLUST <- hclust(dist(mat))


par(mfrow=c(1,1), mar=c(5,5,4,4))

## test things
# set.seed(1)
# N <- 100
# library(GenomicRanges)
# ## GRanges
# gr <- GRanges(seqnames = 
#                 sample(c("chr1", "chr2", "chr3"),
#                        size = N, replace = TRUE),
#               IRanges(
#                 start = sample(1:300, size = N, replace = TRUE),
#                 width = sample(70:75, size = N,replace = TRUE)),
#               strand = sample(c("+", "-", "*"), size = N, 
#                               replace = TRUE),
#               value = rnorm(N, 10, 3), score = rnorm(N, 100, 30),
#               sample = sample(c("Normal", "Tumor"), 
#                               size = N, replace = TRUE),
#               pair = sample(letters, size = N, 
#                             replace = TRUE))
# ## automatically facetting and assign y
# ## this must mean geom_rect support GRanges object
# ggplot(gr) + geom_rect()
# 

viewReads <- function(reads){
  # sort by start
  subset <- splitReads[splitReads$class3 < 5,]
  sorted <- subset[order(subset$lStart),];
  
  #---
  # In the first iteration we work out the y-axis
  # positions that segments should be plotted on
  # segments should be plotted on the next availible
  # y position without merging with another segment
  #---
  yread <- c(); #keeps track of the x space that is used up by segments 
  
  # get x axis limits
  minstart <- min(sorted$lStart);
  maxend <- max(sorted$lEnd);
  
  # initialise yread
  yread[1] <- minstart - 1;
  ypos <- c(); #holds the y pos of the ith segment
  
  # for each read
  for (r in 1:nrow(sorted)){
    read <- sorted[r,];
    start <- read$lStart;
    placed <- FALSE;
    
    # iterate through yread to find the next availible
    # y pos at this x pos (start)
    y <- 1;
    while(!placed){
      
      if(yread[y] < start){
        ypos[r] <- y;
        yread[y] <- read$lEnd;
        placed <- TRUE;
      } 
      
      # current y pos is used by another segment, increment
      y <- y + 1;
      # initialize another y pos if we're at the end of the list
      if(y > length(yread)){
        yread[y] <- minstart-1;
      }
    }
  } 
  
  # find the maximum y pos that is used to size up the plot
  maxy <- length(yread);
  sorted$ypos <- ypos;
  
  # Now we have all the information, start the plot
  plot.new();
  plot.window(xlim=c(minstart, maxend+((maxend-minstart)/10)), ylim=c(1,maxy));
  axis(3);
  
  #---
  # This second iteration plots the segments using the found y pos and 
  # the start and end values
  #---
  for (r in 1:nrow(sorted)){
    read <- sorted[r,];
    # colour dependent on strand type
    if(read$class3 > 3){
      color = 'yellow'
    }else{
      color = 'red'
    }
    #plot this segment!
    segments(read$lStart, maxy-read$ypos, read$lEnd, maxy-read$ypos, col=color);
  }
}
