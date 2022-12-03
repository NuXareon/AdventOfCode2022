package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    score := int64(0)

    for scanner.Scan(){
        opponentHand := scanner.Text()
        scanner.Scan() // YOLO: We assume the inputs is correct 
        matchOutcome := scanner.Text()
        myHand := calculateHand(opponentHand, matchOutcome)
        handScore := evaluateHand(myHand)
        matchScore := evaluateMatch(opponentHand, myHand)
        score += handScore + matchScore
    }

    fmt.Println(score)
}

func calculateHand(opponentHand string, matchOutcome string) string {
    if (matchOutcome == "Y") {
        return opponentHand // Draw
    }

    if (opponentHand == "A") { // Rock
        if (matchOutcome == "X") {
            return "C" // Lose
        }
        return "B" // Win
    } else if (opponentHand == "B") { // Paper
        if (matchOutcome == "X") {
            return "A" // Lose
        }
        return "C" // Win
    } else if (opponentHand == "C") { // Scissor
        if (matchOutcome == "X") {
            return "B" // Lose
        }
        return "A" // Win
    }

    return opponentHand // Draw (should never happen unless our oppenent is cheating)
}

func evaluateHand(hand string) int64{
    switch hand {
    case "A": // Rock
        return 1
    case "B": // Paper
        return 2
    case "C": // Scissors
        return 3
    }
    return 0
}

func evaluateMatch(opponentHand string, myHand string) int64 {
    if (opponentHand == myHand) {
        return 3 // Draw
    }

    if (opponentHand == "A") { // Rock
        if (myHand == "B") {
            return 6 // Win
        }
    } else if (opponentHand == "B") { // Paper
        if (myHand == "C") {
            return 6 // Win
        }
    } else if (opponentHand == "C") { // Scissor
        if (myHand == "A") {
            return 6 // Win
        }
    }

    return 0 // Lose
}