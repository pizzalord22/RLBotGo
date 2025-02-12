RLBotGo
===========
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/Trey2k/RLBotGo)


This repository holds a library for making Rocket League bots in Go!

It provides:

  * An easy to use interface for writing bots
  * An [example bot](https://github.com/Trey2k/RLBotGoExmaple/blob/main/main.go) using this library

Table of Contents:

  * [About](#about)
  * [Todo](#todo)
  * [Usage](#usage)
  * [Compiling](#compiling)
  * [Contributing](#contributing)
  * [License](#license)

About
-----

This project was made to make it easy to write [RL Bots](https://rlbot.org/) in Go. Instead of using flatbuffer's datatypes, this library converts everything to Go types for ease of use.

Todo
-----

Here are some things that could use some work in this repository's current state:

  * ~~Add support for render groups~~
  * ~~Add support for desired game state~~
  * Add game message support
  * Add some (potentially) useful math functions
  * Get #Go channel in [RLBot Discord](https://discord.com/invite/yc643yyd)

Usage
------------

The suggested starting point for using this library is using the [RLBotGoExample](https://github.com/Trey2k/RLBotGoExmaple) repository as a template for your bot.

If you don't start with the example repository, start out with a connection to RLBot:
```Go
	rlBot, err := RLBot.Connect(23234)
	if err != nil {
		panic(err)
	}
```
After that, send your bot's ready message:
```Go
	err = rlBot.SendReadyMessage(true, true, true)
	if err != nil {
		panic(err)
	}

```
Call SetGetInput with the name of your desired callback function:
```Go
    rlBot.SetGetInput(tick)
```
Finally, write a function to return the player input every tick:
```Go
// getInput takes in a GameState which contains the gameTickPacket, ballPredidctions, fieldInfo and matchSettings
// it also takes in the RLBot object. And returns a PlayerInput
func getInput(gameState *RLBot.GameState, rlBot *RLBot.RLBot) *RLBot.ControllerState {
	PlayerInput := &RLBot.ControllerState{}

	// Count ball touches up to 10 and on 11 clear the messages and jump
	wasjustTouched := false
	if gameState.GameTick.Ball.LatestTouch.GameSeconds != 0 && lastTouch != gameState.GameTick.Ball.LatestTouch.GameSeconds {
		totalTouches++
		lastTouch = gameState.GameTick.Ball.LatestTouch.GameSeconds
		wasjustTouched = true
	}

	if wasjustTouched && totalTouches <= 10 {
    // DebugMessage is a helper function to let you quickly get debug text on screen. it will automaticaly place it so text will not overlap
		rlBot.DebugMessageAdd(fmt.Sprintf("The ball was touched %d times", totalTouches))
		PlayerInput.Jump = false
	} else if wasjustTouched && totalTouches > 10 {
		rlBot.DebugMessageClear()
		totalTouches = 0
		PlayerInput.Jump = true
	}
	return PlayerInput

}
```

After that, you should have a functional bot!

Some other useful things:
```go
// Sending a quick chat
// (QuickChatSelection, teamOnly) refer to the godocs or RLBot documentation for all QuickChatSelection types
rlBot.SendQuickChat(RLBot.QuickChat_Custom_Toxic_404NoSkill, false)

// Sending a desired game state
// view https://pkg.go.dev/github.com/Trey2k/RLBotGo#DesiredGameState for more info
// Most fields are optional
desiredState := &RLBot.DesiredGameState{}
desiredState.BallState.Physics.Velocity = RLBot.Vector3{X: 0, Y: 0, Z: 1000}
rlBot.SendDesiredGameState(desiredState)

// Getting ball predictions
// This will be in the gameState sturct that you recive when the getInput callback is called
func getInput(gameState *RLBot.GameState, rlBot *RLBot.RLBot) *RLBot.ControllerState {
	// Loop through all the predictions we have and print the position and predicted time.
	// There should be a total of 6 * 60 predictions. 60 for every secound and a total of 6 secounds
	for i := 0; i < len(gameState.BallPrediction.Slices); i++ {
		prediction := gameState.BallPrediction.Slices[i]
		fmt.Printf("The ball will be at pos (%f, %f, %f) at %f game time", prediction.Physics.Location.X,
			prediction.Physics.Location.Y, prediction.Physics.Location.Z, prediction.GameSeconds)
	}
	return nil
}


```

Compiling
------------
In order to use this library, you'll need to install and configure the following:

  * [Go](https://golang.org) installed and [configured](https://golang.org/doc/install) atleast go 1.15
  * [Setup](https://www.youtube.com/watch?v=oXkbizklI2U) [RLBot](https://rlbot.org/)
  * A copy of [Rocket League](https://www.rocketleague.com/) installed
  * Port 23234 availible on your local machine for [RLBot](https://rlbot.org/)
  * A little patience :)

To compile your bot the first thing you will want to do is take a look at the bot folder in the [Example Repo](https://github.com/Trey2k/RLBotGoExample/tree/main/bot). Modify the config files to your liking and make sure you point to the correct executable file. After that you can simply use `go build ./` and your bot should be built.

To add it to RLBot simply click the +Add button in RL Bot GUI and select the folder that contains the bot folder.

If sending your bot for a tournament it will probably be easiest to place the exe in the bot/src/ folder. Give the bot folder a more unique(Your bots name) name and zip that folder. Make sure to change the path to the exe in the bot.cfg file as well!

Contributing
------------

Contributions are always welcome. If you're interested in contributing feel free to submit a PR.

License
-------

This project is currently licensed under the permissive MIT license. Please refer to the [license](/LICENSE) file for more information.
