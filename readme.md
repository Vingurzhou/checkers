# cosmosProject

基于Cosmos SDK和Tendermint的区块链游戏.

## 开始

```shell
checkersd start
```

### 配置

```shell
vim config.yml
```

### 开始

```shell
checkersd start
```
### 查看系统信息
```shell
checkersd query checkers show-system-info
```
### 创建玩家

```shell
checkersd keys add zwz
```

### 创建棋局

```shell
export alice=$(checkersd keys show alice -a) 
export bob=$(checkersd keys show bob -a)
checkersd tx checkers create-game $alice $bob --from $alice 
```
### 查看棋局
```shell
checkersd query checkers show-stored-game 1
```
### 查看棋盘

```shell
checkersd query checkers show-stored-game 1 --output json | jq ".storedGame.board" | sed 's/"//g' | sed 's/|/\n/g'
```

### 移动棋子

```shell
checkersd tx checkers play-move 1 1 2 2 3 --from $alice
```

### 查看所有棋局

```shell
checkersd query checkers list-stored-game
```
