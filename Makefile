bkname = ./backup/thitcho-$(shell date +"%Y-%m-%d-%H-%M")
pid = $(shell cat .pid +"%T")
logfile = log/$(pid).log

default: build  server

build:
	@clear
	@mkdir -p bin
	@echo "STEP : BUILD"
	@go build -o bin/thitcho
	@echo "Build successfully."

install::
	@clear
	@echo "STEP : INSTALL"
	@go install
	@echo "Install successfully."



back-up:
	@echo "STEP : BACK UP"
	@mkdir -p ./backup
	@cp ./bin/thitcho $(bkname)
	@echo "Done"

run:
	@echo "STEP : RUNNING"
	@./bin/thitcho --help

server:
	@clear
	@echo "STEP : RUN SERVER"
	@./bin/thitcho server

monitor:
	@./bin/thitcho monitor

stop:
	@clear
	@echo "Kill PID = $(pid)"
	@kill $(pid)
	@echo "=> done.."
