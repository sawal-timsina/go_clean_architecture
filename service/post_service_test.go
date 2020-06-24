package service

import (
	"testing"
	"prototype2/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error){
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1
	
	post := entity.Post{ID: identifier, Title: "Apple", Text: "Ball"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Apple", result[0].Title)
	assert.Equal(t, "Ball", result[0].Text)
}


func TestCreate(t *testing.T){
	mockRepo := new(MockRepository)
	post := entity.Post{Title: "Test Title", Text: "Test Text"}

	//Setting up the expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "Test Title", result.Title)
	assert.Equal(t, "Test Text", result.Text)
	assert.Nil(t, err)
}