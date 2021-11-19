package handlers

import (
	"errors"
	"fmt"

	"github.com/latonaio/salesforce-data-models"
)

func HandlePriceRecord(metadata map[string]interface{}) error {
	priceMasters, err := models.MetadataToPriceRecords(metadata)
	if err != nil {
		return fmt.Errorf("failed to convert metadata to models: %v", err)
	}
	if err := models.RegisterPriceRecordsAndCacheClear(priceMasters); err != nil {
		return errors.New("failed to cache clear and register table: " +  err.Error())
	}
	return nil
}

func HandlePriceRecordSeriesNumber(metadata map[string]interface{}) error {
	prsn, err := models.MetadataToPriceRecordSeriesNumbers(metadata)
	if err != nil {
		return fmt.Errorf("failed to convert metadata to models: %v", err)
	}
	if err := models.RegisterPriceRecordSeriesNumbersAndCacheClear(prsn); err != nil {
		return errors.New("failed to cache clear and register table: " +  err.Error())
	}
	return nil
}

func HandlePriceType(metadata map[string]interface{}) error {
	priceTypes, err := models.MetadataToPriceTypes(metadata)
	if err != nil {
		return fmt.Errorf("failed to convert metadata to models: %v", err)
	}
	if err := models.RegisterPriceTypesAndCacheClear(priceTypes); err != nil {
		return errors.New("failed to cache clear and register table: " +  err.Error())
	}
	return nil
}