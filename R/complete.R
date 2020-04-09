# Write a function that reads a directory full of files and reports the number of completely observed cases in each data file. The function should return a data frame where the first column is the name of the file and the second column is the number of complete cases. 

complete <- function(directory,id=1:332){
  path <- "C:\\Users\\ashwi\\Documents\\R\\rprog_data_specdata\\"
  y <- data.frame()

  
  for ( i in id ){
    data <- read.csv(paste(path,directory,"\\",sprintf("%03d", i),".csv",sep=""))
    y <- rbind(y,c(i,sum(complete.cases(data))))
  }
  
  colnames(y) <- c("id","nobs")
  print(y)
}