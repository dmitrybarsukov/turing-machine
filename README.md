# Solver for board game "Turing machine"

https://turingmachine.info/

---

## What it does?

* This program analyzes current state in the game and gives recommendations about the next move;
* Uses a yaml file as an input info about game state.

---

## 1. Build it

Just run `make` or `go build cmd/turing-machine/main.go`

OR

run `go install github.com/dmitrybarsukov/turing-machine/cmd/turing-machine@latest`

## 2. Prepare game state file

`cp game_state_sample.yaml game_state.yaml`

### 2.1. Fill validators in `game_state.yaml`

Majority of known validators are already defined in beginning of `game_state_sample.yaml` file,
but you can also define your own. 
Example:

```yaml
validators:
  A: *countOfNumber1
  B: *majorParity
  C: *parityOfAllSum
  D: *sumOfTriangleAndSquareComparedTo6
  E: *hasPair
  F: *hasOrder
```
If you do not have validators E and F or only F, just remove them from this file

### 2.2. Fill checked tests (if any) in `game_state.yaml`

```yaml
tests:
  124:        # Input number 124
    D: true   # Testing validator D gave true
    B: false  # Testing validator D gave false
    E: true   # Testing validator E gave true
  542:        # Input number 542
    A: false  # Testing validator A gave false
```

## 3. Run it

* `./bin/turing-machine game_state.yaml` - if built via `make`
* `./turing-machine game_state.yaml` - if built via `go build ...`
* `turing-machine game_state.yaml` - if built via `go install ...`

## 4. Do what app says

1. `Try code 123` -> you should try input code `123`
2. `Check validator B` -> you should check validator `B` result
3. Return to step **2**

## 5. Win game!

If app says `Found solution: 531`, so the solution for game is `531`

---

## Version history

### v1.0
* Initial version with ~70% of all known validators

### v1.1
* Improved validator recommendation
