package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"go-rss-reader-service/db_utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

var logger = zap.L()

func GetFieldErrorsArray(err error) map[string]interface{} {
	var ve validator.ValidationErrors
	fieldErrorsArray := make(map[string]interface{})
	if errors.As(err, &ve) {
		for _, fe := range ve {
			fieldErrorsArray[fe.Field()] = fe.Error()
		}
		return fieldErrorsArray
	}
	return nil
}

func ValidateUrl(address string) (*http.Response, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		logger.Error(fmt.Sprintf("error in callUrl while making request for %v", address), zap.Error(err))
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("error in callUrl while calling the request for %v", address), zap.Error(err))
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errorOccurred := errors.New(fmt.Sprintf("response code is %v", resp.StatusCode))
		logger.Error(fmt.Sprintf("error in callUrl having non-success response for %v", address), zap.Error(errorOccurred))
		return nil, err
	}
	return resp, nil
}

func CalculateTotalPagesForPagination(totalCount int64, limit int64) int64 {
	var totalPages float64
	totalPages = float64(totalCount) / float64(limit)
	if totalPages < 1 {
		totalPages = 1
	}
	decimalPoints, _ := strconv.ParseFloat(strings.Split(fmt.Sprintf("%.2f", totalPages), ".")[1], 64)
	if decimalPoints > 0 {
		totalPages = totalPages + 1
	}
	return int64(totalPages)
}

func CheckIfUrlExistInDb(url string) (bool, error) {
	rssUrlDto, err := db_utils.GetRssUrlByUrl(url)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("error in AddRssUrl", zap.Error(err))
		return false, err
	}
	if rssUrlDto.Id > 0 {
		return true, nil
	}
	return false, nil
}
