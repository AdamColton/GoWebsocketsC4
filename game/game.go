package game

import(
  "../c4"
  "../c4mc"
  "code.google.com/p/go.net/websocket"
  "strconv"
)

func Game(ws *websocket.Conn){
  game := c4.NewGame()
  ai := c4mc.NewMonteCarlo_P(50000)

  for turn := game.Turn(); turn != 0; turn = game.Turn() {
    if turn == 1 {
      websocket.Message.Send(ws, game.String() )
      var moveString []byte
      websocket.Message.Receive(ws, &moveString)
      move, _ := strconv.Atoi( string(moveString) )
      game.Move(move)
    } else {
      ai.Move(&game)
    }
  }
  websocket.Message.Send(ws, game.String() )
}