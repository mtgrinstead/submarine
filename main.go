package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    var forward_pos = 0
    var depth_pos = 0
    var aim = 0

    file, err := os.Open("submarine_kata_input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Fields(line)
        if len(parts) != 2 {
            log.Fatal("Invalid line format")
        }
        dir := parts[0]
        num, err := strconv.Atoi(parts[1])
        if err != nil {
            log.Fatalf("Invalid number format: %s", parts[1])
        }

        if dir == "forward" {
            forward_pos += num
            depth_pos += aim * num
        } else if dir == "up" {
            aim -= num
        } else if dir == "down" {
            aim += num
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("The forward position is: %d\nThe depth position is: %d\nThe forward position is: %d\n", forward_pos, depth_pos, forward_pos * depth_pos)
}
