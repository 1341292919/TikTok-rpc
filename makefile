MODULE = TikTok-rpc


.PHONY:target
target:
		sh build.sh
		sh output/bootstrap.sh

.PHONY:newHz
newHz:
	hz new -module $(MODULE)

.PHONY: genHz
genHz:
	hz update -idl ./idl/api/interact.thrift
	hz update -idl ./idl/api/socialize.thrift
	hz update -idl ./idl/api/user.thrift
	hz update -idl ./idl/api/video.thrift
	hz update -idl ./idl/api/model.thrift

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
	cd rpc/user
	kitex -module $(MODULE) -service user ../../idl/user.thrift
	cd ..
	cd video
	kitex -module $(MODULE) -service video ../../idl/video.thrift
	cd ..
	cd interact
	kitex -module $(MODULE) -service interact ../../idl/interact.thrift
	cd ..
	cd socialize
	kitex -module $(MODULE) -service socialize ../../idl/socialize.thrift