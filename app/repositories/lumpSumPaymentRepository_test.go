package repositories

import (
	"time"
	"testing"

	entity "github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type lumpSumPaymentRepositoryTestSuite struct {
	suite.Suite
	ctx *gin.Context
	db  *gorm.DB
}

func (suite *lumpSumPaymentRepositoryTestSuite) SetupLumpSumPaymentRepositoryTest() {
	suite.db, _ = config.ConnectDB()
}

func (suite *lumpSumPaymentRepositoryTestSuite) TestBuildNewLumpSumPaymentRepository() {
	repoTest := NewLumpSumPaymentRepository()
	var dummyImpl *LumpSumPaymentRepository
	assert.NotNil(suite.T(), repoTest)
	assert.Implements(suite.T(), dummyImpl, repoTest)
}

func (suite *lumpSumPaymentRepositoryTestSuite) CreateLumSumPayment() {
	repoTest := NewLumpSumPaymentRepository()
	dummyLumpSumPayment := entity.LumpSumPayment{
		ID					: 1,
		FeeId				: "1234567",
		CollectionId		: "1234567",
		ProposalNumber		: "30012341234122",
		PolicyNumber		: "30012341234123",
		FirstEffectiveDate	: time.Now(),
		EffectiveDate		: time.Now(),
		SettledDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		BankName			: "BCA",
	}
	repoTest.CreateLumSumPayment(dummyLumpSumPayment)
}

func (suite *lumpSumPaymentRepositoryTestSuite) UpdateLumSumPayment() {
	repoTest := NewLumpSumPaymentRepository()
	dummyLumpSumPayment := entity.LumpSumPayment{
		ID					: 1,
		FeeId				: "1234568",
		CollectionId		: "1234568",
		ProposalNumber		: "30012341234124",
		PolicyNumber		: "30012341234121",
		FirstEffectiveDate	: time.Now(),
		EffectiveDate		: time.Now(),
		SettledDate			: time.Now(),
		PaymentMethod		: "Indomaret",
		TotalAmount			: 10000,
		BankName			: "BCA",
	}
	repoTest.UpdateLumSumPayment(dummyLumpSumPayment)
}

func (suite *lumpSumPaymentRepositoryTestSuite) GetAllLatestGroupLumpSumPayments() {
	repoTest := NewLumpSumPaymentRepository()
	lumpSumPaymentDummy := repoTest.GetAllLatestGroupLumpSumPayments()
	assert.NotNil(suite.T(), lumpSumPaymentDummy)
}

func (suite *lumpSumPaymentRepositoryTestSuite) GetAllLumpSumPayments() {
	repoTest := NewLumpSumPaymentRepository()
	lumpSumPaymentDummy := repoTest.GetAllLumpSumPayments()
	assert.NotNil(suite.T(), lumpSumPaymentDummy)
}
func (suite *lumpSumPaymentRepositoryTestSuite) GetLumpSumPayment() {
	repoTest := NewLumpSumPaymentRepository()
	lumpSumPaymentDummy := repoTest.GetLumpSumPayment("30012341234121")
	assert.NotNil(suite.T(), lumpSumPaymentDummy)
}

func (suite *lumpSumPaymentRepositoryTestSuite) RemoveLumSumPayment() {
	repoTest := NewLumpSumPaymentRepository()
	dummyLumpSumPayment := entity.LumpSumPayment{
		ID: 1,
	}
	repoTest.DeleteLumSumPayment(dummyLumpSumPayment)
}

func TestLumpSumPaymentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(lumpSumPaymentRepositoryTestSuite))
}