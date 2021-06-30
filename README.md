# Let's Go Bowling

## Overview
Create a 10-pin bowling scoring application that uses traditional scoring rules to score
each frame.
Put your project in Github and send us the repo when you are done.
Please note, that while the overall assessment is time boxed to approximately 48 hours,
the work required should only take a couple of hours. If you find that you are bumping
up against that time, you should look to try to simplify your approach.
Scoring rules
A good overview of how to score the bowling game can be found here:
https://www.thoughtco.com/bowling-scoring-420895

## Derived requirements
* Store code in Github
* Timebox to a few hours
* Score a hypothetical bowling game using [the "traditional scoring rules" described here](https://www.thoughtco.com/bowling-scoring-420895)
* Input: Accept user input representing frames via CLI. Input format is comma separated values [0-9,X,/]
* Output: Print the recently completed frame number and current and/or final score to CLI

## Example scenarios:
### Scenario 1:
Given a score of a strike or spare, and it is the first frame,
When I provide "X" or "7,/", respectively, via the command line,
Then the recently completed frame number is printed to the command line along with
the current score, which is empty.
### Scenario 2:
Given an open frame score, with the first attempting being 2 and the second being 3,
When I provide "2,3", via the command line,
Then the recently completed frame number is printed to the command line along with
the current score that is printed out which is the accumulative score of the previous
frames and the newly entered score.
### Scenario 3:
Given the current frame is the 10th frame, and 3 attempts are awarded due to at least 1
strike or spare in the frame,
When I provide "3,/,X", via the command line,
Then the recently completed frame number is printed to the command line along with
the final score after which the application exits.

## Requirements
Project requirements are outlined in the root folder.  See `requirements.md`.