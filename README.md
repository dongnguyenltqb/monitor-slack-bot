## MONITOR SLACK BOT

- Clone repository

  ```
  #clone this repository
  git clone https://github.com/dongnguyenltqb/monitor-slack-bot.git
  
  #jump to folder monitor-slack-bot
  cd monitor-slack-bot
  
  #download dependency
  go mod download
  ```

- Install

  ```
  #install to $GOBIN
  make install
  ```

- Help

  ```
  #get help
  thitcho -h
  
  #get list help command
  thitcho list -h
  
  #get add help command
  thitcho add -h
  ```

- Edit config.yaml, slack bot token, channel id

  ```
  #use sample config
  cp config.yaml.example config.yaml
  
  ```

- Start monitor

  ```
  #run with nohup
  nohup thitcho monitor --config config.yaml > monitor.log &
  
  #get log
  sudo tail -f monitor.log
  
  #handle slack command 
  # url =. /bot/slack/sysinfo
  nohup thitcho server --config config.yaml > monitor-server.log &
  ```

- Stop 

  ```
  sudo pkill thitcho
  ```

  