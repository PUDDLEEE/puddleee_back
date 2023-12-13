//mockery --dir=pkg/interfaces --name=IUserRepository --filename=IUserRepository.go --output=pkg/interfaces/mocks --outpkg=mocks

mockery = mockery
dir = pkg/interfaces

userRepo = IUserRepository
userService = IUserService

roomRepo = IRoomRepository
roomService = IRoomService

jwtAuthService = IJwtAuthService

authRepo = IAuthRepository
authService = IAuthService

outputDir = pkg/interfaces/mocks
outputPkg = mocks

mainDir = ../../main.go
userDir = internal/user
roomDir = internal/room
viewDir = internal/view
messageDir = internal/message

createUserRepoMock :
	${mockery} --dir=${dir} --name=${userRepo} --filename=${userRepo}.go --output=${outputDir} --outpkg=${outputPkg}

createUserServiceMock :
	${mockery} --dir=${dir} --name=${userService} --filename=${userService}.go --output=${outputDir} --outpkg=${outputPkg}

createRoomRepoMock :
	${mockery} --dir=${dir} --name=${roomRepo} --filename=${roomRepo}.go --output=${outputDir} --outpkg=${outputPkg}

createRoomServiceMock :
	${mockery} --dir=${dir} --name=${roomService} --filename=${roomService}.go --output=${outputDir} --outpkg=${outputPkg}

createJwtAuthServiceMock :
	${mockery} --dir=${dir} --name=${jwtAuthService} --filename=${jwtAuthService}.go --output=${outputDir} --outpkg=${outputPkg}

createAuthServiceMock :
	${mockery} --dir=${dir} --name=${authService} --filename=${authService}.go --output=${outputDir} --outpkg=${outputPkg}

createAuthRepositoryMock :
	${mockery} --dir=${dir} --name=${authRepo} --filename=${authRepo}.go --output=${outputDir} --outpkg=${outputPkg}

createDocs :
	swag init -d ${userDir},${roomDir} -g ${mainDir} --parseDependency