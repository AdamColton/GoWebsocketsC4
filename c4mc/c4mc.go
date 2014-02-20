package c4mc

import (
  "../c4"
  "math/rand"
  "runtime"
)

/*
Carlo the Monkey
    __
w c(..)o    (
\__(-)    __)
    /\   (
  /(_)___)
 w /|
  | \
  m m
*/

func randomMove(game *c4.Game) int {
  moves := game.Moves()
  l := len(moves)
  return moves[ rand.Intn(l) ]
}

func monkeyCarlo(game *c4.Game){
  game.Move( randomMove(game) )
}

func fullMonkeyCarloGame(game *c4.Game) int{
  positions := 0
  for game.Turn() != 0 {
    monkeyCarlo(game)
    positions++
  }
  return positions
}

func moveValue(startingGame, endingPosition *c4.Game) int {
  if (endingPosition.Winner() == startingGame.Turn()) {
    return 1
  }
  return 0
}

type MonkeyCarlo struct{}
func (self *MonkeyCarlo) Move(game *c4.Game) {
  monkeyCarlo(game)
}

type MonteCarlo struct{
  Positions int
}
func NewMonteCarlo(positions int) MonteCarlo {
  return MonteCarlo{positions}
}
func (self *MonteCarlo) Move(game *c4.Game) {
  moveRecord := [7][2]int{}
  for i := 0; i<self.Positions; {
    move := randomMove(game)
    sim := game.Copy()
    moveRecord[move][0]++
    sim.Move(move)
    i += fullMonkeyCarloGame(&sim)
    moveRecord[move][1] += moveValue(game, &sim)
  }
  game.Move( bestMove(&moveRecord) )
}

func bestMove(moves *[7][2]int) int {
  bestMove := 0
  bestScore := 0.0
  for i := 0; i<7; i++ {
    if (moves[i][0] > 0){
      if score := float64(moves[i][1]) / float64(moves[i][0]); score > bestScore {
        bestScore = score
        bestMove = i
      }
    }
  }
  return bestMove
}

type MonteCarlo_P struct{
  positions int
  threads int
}
func NewMonteCarlo_P(simulations int) MonteCarlo_P {
  threads := runtime.NumCPU()
  return MonteCarlo_P{simulations, threads}
}
type moveData struct{
  col int
  val int
  positions int
}
func (self *MonteCarlo_P) Move(game *c4.Game) {
  moveRecord := [7][2]int{}
  ch := make(chan moveData)
  moveChannel := make(chan int, self.threads)
  for i:=0; i<self.threads; i++ {
    moveChannel <- randomMove(game)
    go gameThread(ch, moveChannel, game)
  }
  for i, positions := 0, 0 ; i<self.threads;{
    move := <-ch
    moveRecord[move.col][0]++
    moveRecord[move.col][1] += move.val
    positions += move.positions
    if positions > self.positions {
      moveChannel <- -1
      i++
    } else {
      moveChannel <- randomMove(game)
    }
  }
  game.Move( bestMove(&moveRecord) )
}

func gameThread(out chan moveData, moveChannel chan int, game *c4.Game){
  var move int
  var positions int
  var sim c4.Game
  for {
    move = <- moveChannel
    if move == -1 {
      return
    }
    sim = game.Copy()
    sim.Move(move)
    positions = fullMonkeyCarloGame(&sim)
    out <- moveData{move, moveValue(game, &sim), positions}
  }
  
}