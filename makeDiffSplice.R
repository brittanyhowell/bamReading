

setwd("~/Documents/University/Honours_2016/Project/bamReading/Split/Human/Compile")


lEvents <- read.table("landscapeNumEvents.txt")
events <- read.table("AllComp.txt")
subset <- t(data.frame(lEvents[,0:27]))

colnames(events) <- c("All", "Either", "Both", "Stringent")

colours <- c("salmon1" ,"salmon1" ,"salmon1" ,"palevioletred" ,"palevioletred" ,"palevioletred" ,"orchid4","orchid4","orchid4", "slateblue" ,"slateblue" ,"slateblue" ,"cornflowerblue", "cornflowerblue", "cornflowerblue", "aquamarine", "aquamarine", "aquamarine", "darkolivegreen1", "darkolivegreen1", "darkolivegreen1", "chartreuse4", "chartreuse4", "chartreuse4", "darkgreen", "darkgreen", "darkgreen")
barplot(subset, beside=T, col = colours, ylab = "Number of candidate splice events", names.arg = colnames(events))



































