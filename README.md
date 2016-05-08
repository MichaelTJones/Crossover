# Crossover
Example of determining the performance crossover point for serial/parallel evaluation. Written to answer a question posed on the go-nuts mailing list.

For the particular computation under discussion, summing an array of ints, the input size where parallel evaluation becomes faster is approximately 28728 on a 4-core Apple MacBookPro. In a more substantial computation where overhead would not be so dominant, the value determined would be much lower. Despite this, the machinery of TestCrossover() is what is needed in every such case. 
