# Advent of Code 2024

Practising my acquired knowledge of `go`.

*⚠️ NOTE ⚠️* Most of these solutions are uploaded without any sort of cleanup, so
lower expectations on code readibility and abstraction.

The goal is to solve each day's problems while learning a bit of `go`, not to
come up with the optimal, cleanest solutions.

## Running

For each day, copy the `template/` directory and create a `dayXX/` directory.
Make sure to modify the package name, and in `main.go` replace the previous
day's import.

```fish
# Include debug logs
go run debug main.go a 01

# Release mode
go run r main.go a 01
```

## Testing

```fish
go test ./dayXX
```

## When will I get bored?

And some notes of what I learned

- [x] Day1
- [x] Day2 - `append` modifies the original slice
- [x] Day3
- [x] Day4 - make sure to copy slices as there are hacky ways in which my
slices can be modified due to reallocations
- [x] Day5
- [x] Day6 - oh so we cannot place obstacles in positions that the guard has
  already checked but is facing away from?
- [x] Day7
- [x] Day8
- [x] Day9
- [x] Day10
- [x] Day11
- [x] Day12
- [x] Day13
- [x] Day14
- [x] Day15
- [x] Day16
- [ ] Day17
- [ ] Day18
- [ ] Day19
- [ ] Day20
- [ ] Day21
- [ ] Day22
- [ ] Day23
- [ ] Day24
