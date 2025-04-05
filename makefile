MODULE = TikTok-rpc

DIR = $(shell pwd)
IDL_PATH = $(DIR)/idl

.PHONY:target
target:
		sh build.sh
		sh output/bootstrap.sh

.PHONY:newHz
newHz:
	hz new -module $(MODULE)

.PHONY: hz-%
hz-%:
	hz update -idl ${IDL_PATH}/api/$*.thrift

.PHONY: genKt
genKt:
	kitex idl/model.thrift
	kitex idl/user.thrift
	kitex idl/video.thrift
	kitex idl/interact.thrift
	kitex idl/socialize.thrift
	mkdir rpc
	mkdir rpc/user
	mkdir rpc/video
	mkdir rpc/interact
	mkdir rpc/socialize
	cd rpc/user && kitex -module $(MODULE) -thrift no_default_serdes -service user ${IDL_PATH}/user.thrift && \
	cd ../video && kitex -module $(MODULE) -thrift no_default_serdes -service video ${IDL_PATH}/video.thrift && \
	cd ../interact && kitex -module $(MODULE) -thrift no_default_serdes -service interact ${IDL_PATH}/interact.thrift && \
	cd ../socialize && kitex -module $(MODULE) -thrift no_default_serdes -service socialize ${IDL_PATH}/socialize.thrift

.PHONY: kit
kit:
		kitex -thrift no_default_serdes idl/model.thrift
		kitex -thrift no_default_serdes idl/user.thrift
		kitex idl/video.thrift
		kitex idl/interact.thrift
		kitex idl/socialize.thrift

