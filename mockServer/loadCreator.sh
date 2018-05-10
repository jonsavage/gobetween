#!/bin/bash

echo "./loadCreator.sh"
# frequency=.04
# totalDuration=5
# numberOfIterations=$totalDuration/$frequency

for i in {1..725}
do
     curl localhost:3000
     sleep .04
done
