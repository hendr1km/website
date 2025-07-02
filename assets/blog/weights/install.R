install.packages("renv")
renv::init()
renv::install("ggdist")
renv::install("doParallel")
renv::install("dvmisc")
renv::install("dplyr")
renv::snapshot()

