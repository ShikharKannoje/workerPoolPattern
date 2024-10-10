# The Worker Pool Pattern
This is a pattern which is used to handle multiple tasks asynchronously in an efficient way. It allows you to utilise the system's resources to an optimum level. This pattern is useful when you have a queue of tasks and you want to execute them independently and in parallel.

The basic idea of this pattern is to create a set of worker goroutines which would pick tasks from a shared channel and execute them. This approach limits the program to grow beyond a certain limit and keeps the number of concurrently running goroutines, preventing resource exhaustion. 
