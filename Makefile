include Makefile.variables

.PHONY: format test check prepare
## prefix before other make targets to run in your local dev environment

format:
	${DOCKRUN} bash ./scripts/format.sh
check: format
	${DOCKRUN} bash ./scripts/check.sh
test: check
	${DOCKTEST} bash ./scripts/test.sh
local: | quiet
	@$(eval DOCKRUN= )
	@mkdir -p tmp
	@touch tmp/dev_image_id
quiet: # this is silly but shuts up 'Nothing to be done for `local`'
	@:
prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id
db_start:
	@docker run --name task-scheduler-mysql-db -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3306:3306 -d mysql:5.6
	@docker run --name task-scheduler-mongo-db -p 27015-27017:27015-27017 -d mongo:4.2.0
db_prepare: db_start
	@docker cp task_scheduler.sql task-scheduler-mysql-db:task_scheduler.sql
	@echo "Executing databases...wait for 15 seconds"
	@sleep 15
	@docker exec -i task-scheduler-mysql-db sh -c 'mysql -uroot < task_scheduler.sql'
