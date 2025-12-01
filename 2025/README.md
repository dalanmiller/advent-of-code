# Advent of Code 2025 (Go)

All 2025 solutions live in a single Go package `year2025`, so a plain `go test` from this directory (or the repo root via `go.work`) runs every day's tests together.

## Workflow
- Generate a new day: `go run ./cmd/newday` (optionally pass a day number). The script creates `dayXX.go`, `dayXX_test.go`, and downloads `inputs/dayXX.txt` when `AOC_SESSION` (or `session`) is present in your environment or `.env`.
- Use the helpers in `aoc2025.go` for inputs and assertions, e.g. `ReaderForInput(3)` or `AssertEqual(t, Day03PartOne, reader, expected)`.
- Keep all solution files in this folder with package `year2025` so `go test` continues to fan out across the whole year.
