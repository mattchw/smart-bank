package util

import (
	"github.com/stretchr/testify/suite"
)

type PasswordTestSuite struct {
	suite.Suite
	password string
}

func (suite *PasswordTestSuite) SetupSuite() {
	suite.password = RandomString(10)
}

func (suite *PasswordTestSuite) TestHashPassword() string {
	hash, err := HashPassword(suite.password)
	suite.Nil(err)
	suite.NotEmpty(hash)

	return hash
}

func (suite *PasswordTestSuite) TestCheckPassword_Success() {
	hash := suite.TestHashPassword()
	err := CheckPassword(suite.password, hash)
	suite.Nil(err)
}

func (suite *PasswordTestSuite) TestCheckPassword_Fail() {
	hash := suite.TestHashPassword()
	err := CheckPassword("wrong password", hash)
	suite.NotNil(err)
}
