package interfaces

import "danglingmind.com/ddd/Test/mock"

var (
	userAppMock    mock.UserAppInterface
	blogAppMock    mock.BlogAppInterface
	tagAppMock     mock.TagAppInterface
	blogTagAppMock mock.BlogTagAppInterface
	tagServiceMock mock.TagServiceInterface
	tokenMock      mock.TokenInterface
	authMock       mock.AuthInterface

	us = NewUser(&userAppMock)
	b  = NewBlog(&blogAppMock, &userAppMock, &tagAppMock, &blogTagAppMock, &tagServiceMock, &tokenMock, &authMock)
	au = NewAuthenticator(&userAppMock, &authMock,&tokenMock)
)
