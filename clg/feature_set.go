package clg

import (
	"github.com/xh3b4sd/anna/feature-set"
)

func (i *clgIndex) GetFeaturesFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	features := fs.GetFeatures()

	return []interface{}{features}, nil
}

func (i *clgIndex) GetFeaturesByCountFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	count, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	features := fs.GetFeaturesByCount(count)

	return []interface{}{features}, nil
}

func (i *clgIndex) GetFeaturesByLengthFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	min, err := ArgToInt(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	max, err := ArgToInt(args, 2)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 3 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 3 got %d", len(args))
	}

	features := fs.GetFeaturesByLength(min, max)

	return []interface{}{features}, nil
}

func (i *clgIndex) GetFeaturesBySequenceFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	sequence, err := ArgToString(args, 1)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 2 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 2 got %d", len(args))
	}

	features := fs.GetFeaturesBySequence(sequence)

	return []interface{}{features}, nil
}

func (i *clgIndex) GetMaxLengthFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	maxLength := fs.GetMaxLength()

	return []interface{}{maxLength}, nil
}

func (i *clgIndex) GetMinLengthFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	minLength := fs.GetMinLength()

	return []interface{}{minLength}, nil
}

func (i *clgIndex) GetMinCountFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	minCount := fs.GetMinCount()

	return []interface{}{minCount}, nil
}

func (i *clgIndex) GetNewFeatureSet(args ...interface{}) ([]interface{}, error) {
	var err error
	newConfig := featureset.DefaultFeatureSetConfig()

	newConfig.Sequences, err = ArgToStringSlice(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.MaxLength, err = ArgToInt(args, 1, newConfig.MaxLength)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.MinLength, err = ArgToInt(args, 2, newConfig.MinLength)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.MinCount, err = ArgToInt(args, 3, newConfig.MinCount)
	if err != nil {
		return nil, maskAny(err)
	}
	newConfig.Separator, err = ArgToString(args, 4, newConfig.Separator)
	if err != nil {
		return nil, maskAny(err)
	}

	if len(args) > 5 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 5 got %d", len(args))
	}

	newFeatureSet, err := featureset.NewFeatureSet(newConfig)
	if err != nil {
		return nil, maskAny(err)
	}

	err = newFeatureSet.Scan()
	if err != nil {
		return nil, maskAny(err)
	}

	return []interface{}{newFeatureSet}, nil
}

func (i *clgIndex) GetSeparatorFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	separator := fs.GetSeparator()

	return []interface{}{separator}, nil
}

func (i *clgIndex) GetSequencesFeatureSet(args ...interface{}) ([]interface{}, error) {
	fs, err := ArgToFeatureSet(args, 0)
	if err != nil {
		return nil, maskAny(err)
	}
	if len(args) > 1 {
		return nil, maskAnyf(tooManyArgumentsError, "expected 1 got %d", len(args))
	}

	sequences := fs.GetSequences()

	return []interface{}{sequences}, nil
}
