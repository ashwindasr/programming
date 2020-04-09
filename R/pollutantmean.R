# Write a function named 'pollutantmean' that calculates the mean of a pollutant (sulfate or nitrate) across a specified list of monitors. The function 'pollutantmean' takes three arguments: 'directory', 'pollutant', and 'id'. Given a vector monitor ID numbers, 'pollutantmean' reads that monitors' particulate matter data from the directory specified in the 'directory' argument and returns the mean of the pollutant across all of the monitors, ignoring any missing values coded as NA.


pollutantmean <- function(directory, pollutant, id=1:332){
  path <- "C:\\Users\\ashwi\\Documents\\R\\rprog_data_specdata\\"
  totalsum <- 0
  len <- 0
  
  for (i in id){
    data <- read.csv(paste(path,directory,"\\",sprintf("%03d", i),".csv",sep=""))
    totalsum <- totalsum + sum(data[[pollutant]], na.rm = TRUE)
    len <- len + sum(!is.na(data[[pollutant]]))
  }
  print(totalsum/len)

}