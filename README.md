# connect-4
Connect 4 multiplayer terminal game

You need to have the game engine before you run the game server.

You need to have the game server running before you run the game client.

Open them separately, each in its own terminal.

How to run the game engine:
```sh
$ make -s build SYSTEM=ge
$ ./bin/ge --port=50051
```

How to run the game server:
```sh
$ make -s build SYSTEM=gs
$ ./bin/gs --port=50052
```

How to run the game client:
```sh
$ make -s build SYSTEM=gc
$ ./bin/gc
```
