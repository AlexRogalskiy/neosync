package transformer_utils

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_GenerateRandomStringWithDefinedLength(t *testing.T) {
	val := int64(6)

	res, err := GenerateRandomStringWithDefinedLength(val)
	require.NoError(t, err)

	require.Equal(t, val, int64(len(res)), "The output string should be the same length as the input length")
}

func Test_GenerateRandomStringWithDefinedLengthError(t *testing.T) {
	val := int64(0)

	_, err := GenerateRandomStringWithDefinedLength(val)
	require.Error(t, err)
}

func Test_SliceStringEmptyString(t *testing.T) {
	res := SliceString("", 10)
	require.Empty(t, res, "Expected result to be an empty string")
}

func Test_SliceStringShortString(t *testing.T) {
	s := "short"
	res := SliceString(s, 10)
	require.Equal(t, s, res, "Expected result to be equal to the input string")
}

func Test_SliceStringValidSlice(t *testing.T) {
	s := "hello, world"
	length := 5
	expected := "hello"
	res := SliceString(s, length)
	require.Equal(t, expected, res, "Expected result to be a substring of the input string with the specified length")
}

func Test_GenerateRandomStringBounds(t *testing.T) {
	type testcase struct {
		min int64
		max int64
	}
	testcases := []testcase{
		{min: int64(2), max: int64(5)},
		{min: int64(2), max: int64(7)},
		{min: int64(23), max: int64(24)},
		{min: int64(4), max: int64(24)},
		{min: int64(2), max: int64(2)},
		{min: int64(2), max: int64(4)},
		{min: int64(1), max: int64(1)},
		{min: int64(0), max: int64(0)},
	}
	for _, tc := range testcases {
		name := fmt.Sprintf("%s_%d_%d", t.Name(), tc.min, tc.max)
		t.Run(name, func(t *testing.T) {
			output, err := GenerateRandomStringWithInclusiveBounds(tc.min, tc.max)
			require.NoError(t, err)
			length := int64(len(output))
			require.GreaterOrEqual(t, length, tc.min, "%d>=%d was not true. output should be greater than or equal to the min. output: %s", length, tc.min, output)
			require.LessOrEqual(t, length, tc.max, "%d<=%d was not true. output should be less than or equal to the max. output: %s", length, tc.max, output)
		})
	}
}

func Test_GenerateRandomStringError(t *testing.T) {
	min := int64(-2)
	max := int64(4)

	_, err := GenerateRandomStringWithInclusiveBounds(min, max)
	require.Error(t, err, "The min or max cannot be less than 0")
}

func Test_GenerateRandomStringErrorMinGreaterThanMax(t *testing.T) {
	min := int64(5)
	max := int64(4)

	_, err := GenerateRandomStringWithInclusiveBounds(min, max)
	require.Error(t, err, "The min cannot be greater than the max")
}

func Test_IsValidEmail(t *testing.T) {
	require.True(t, IsValidEmail("test@example.com"), "Email follows the valid email format")
	require.False(t, IsValidEmail("invalid"), "Email doesn't have a valid email format")
}

func Test_IsValidDomain(t *testing.T) {
	require.True(t, IsValidDomain("@example.com"), "Domain should have an @ sign and then a domain and top level domain")
	require.False(t, IsValidDomain("invalid"), "Domain doesn't contain an @ sign or a top level domain")
}

func Test_IsValidUsername(t *testing.T) {
	require.True(t, IsValidUsername("test"), "Username should be an alphanumeric value comprised of  a-z A-Z 0-9 . - _ and starting and ending in alphanumeric chars with a max length of 63")
	require.True(t, IsValidUsername("test-test"), "Username should be an alphanumeric value comprised of  a-z A-Z 0-9 . - _ and starting and ending in alphanumeric chars with a max length of 63")
	require.True(t, IsValidUsername("test-TEST"), "Username should be an alphanumeric value comprised of  a-z A-Z 0-9 . - _ and starting and ending in alphanumeric chars with a max length of 63")
	require.False(t, IsValidUsername("eger?45//"), "Username contains non-alphanumeric characters")
}

func Test_isValidCharTrue(t *testing.T) {
	val := "12wefg w1231"

	res := IsValidChar(val)

	require.True(t, res)
}

func Test_isValidCharFalse(t *testing.T) {
	val := "ij諏計"

	res := IsValidChar(val)

	require.False(t, res)
}

func Test_IsAllowedSpecialCharTrue(t *testing.T) {
	val := "$*#))"

	for _, r := range val {
		require.True(t, IsAllowedSpecialChar(r), "Expected true for rune: %v", r)
	}
}

func Test_IsAllowedSpecialCharFalse(t *testing.T) {
	val := "諏計飯利"

	for _, r := range val {
		require.False(t, IsAllowedSpecialChar(r), "Expected false for rune: %v", r)
	}
}

func Test_StringinSliceTrue(t *testing.T) {
	slice := []string{"hello", "world"}
	val := "hello"

	res := StringInSlice(val, slice)

	require.True(t, res)
}

func Test_StringinSliceFalse(t *testing.T) {
	slice := []string{"hello", "world"}
	val := "hellomeow"

	res := StringInSlice(val, slice)

	require.False(t, res)
}

func Test_TrimStringIfExceeds(t *testing.T) {
	type testcase struct {
		input    string
		maxl     int64
		expected string
	}

	testcases := []testcase{
		{"foo", 3, "foo"},
		{"foo", 2, "fo"},
		{"foo", 0, ""},
		{"", 1, ""},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			actual := TrimStringIfExceeds(tc.input, tc.maxl)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func Test_GetSmallerOrEqualNumbers(t *testing.T) {
	type testcase struct {
		input    []int64
		val      int64
		expected []int64
	}

	testcases := []testcase{
		{[]int64{1, 2, 3}, 2, []int64{1, 2}},
		{[]int64{1, 2, 3}, 0, []int64{}},
		{[]int64{1, 2, 3}, 3, []int64{1, 2, 3}},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			actual := GetSmallerOrEqualNumbers(tc.input, tc.val)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func Test_ToSet(t *testing.T) {
	actual := ToSet([]int64{1, 1, 1, 2, 3, 4, 5, 10})
	require.Len(t, actual, 6)
}

func Test_WithoutCharacters(t *testing.T) {
	type testcase struct {
		input    string
		invalid  []rune
		expected string
	}

	testcases := []testcase{
		{"foobar", []rune{'r'}, "fooba"},
		{"foobar", []rune{'r', 'o'}, "fba"},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			actual := WithoutCharacters(tc.input, tc.invalid)
			require.Equal(t, tc.expected, actual)
		})
	}
}

func Test_GetRandomCharacterString(t *testing.T) {
	actual := GetRandomCharacterString(rand.New(rand.NewSource(1)), 100)
	assert.Len(t, actual, 100)
}

func Test_GenerateStringFromCorpus(t *testing.T) {
	randomizer := rand.New(rand.NewSource(1))
	values := []string{"aa", "bb", "cc", "dd"}
	stringMap := map[int64][2]int{2: {0, 3}}
	sizeIndices := []int64{2}

	output, err := GenerateStringFromCorpus(
		randomizer,
		values,
		stringMap,
		sizeIndices,
		nil,
		2,
		nil,
	)
	require.NoError(t, err)
	require.NotEmpty(t, output)
}

func Test_GenerateStringFromCorpus_No_Candidates(t *testing.T) {
	randomizer := rand.New(rand.NewSource(1))
	values := []string{"aa", "bb", "cc", "dd"}
	stringMap := map[int64][2]int{2: {0, 3}}
	sizeIndices := []int64{2}

	minLength := int64(3)
	output, err := GenerateStringFromCorpus(
		randomizer,
		values,
		stringMap,
		sizeIndices,
		&minLength,
		4,
		nil,
	)
	require.Error(t, err)
	require.Empty(t, output)
}

func Test_GenerateStringFromCorpus_Mismatched_MapAndIndices(t *testing.T) {
	randomizer := rand.New(rand.NewSource(1))
	// the index has a key of 3, but it is not present in the map
	values := []string{"aa", "bb", "cc", "dd"}
	stringMap := map[int64][2]int{2: {0, 3}}
	sizeIndices := []int64{3}

	output, err := GenerateStringFromCorpus(
		randomizer,
		values,
		stringMap,
		sizeIndices,
		nil,
		4,
		nil,
	)
	require.Error(t, err)
	require.Empty(t, output)
}
func Test_GenerateStringFromCorpus_NoDice(t *testing.T) {
	randomizer := rand.New(rand.NewSource(1))
	// the index has a key of 3, but it is not present in the map
	values := []string{"aa", "bb"}
	stringMap := map[int64][2]int{2: {0, 1}}
	sizeIndices := []int64{2}

	output, err := GenerateStringFromCorpus(
		randomizer,
		values,
		stringMap,
		sizeIndices,
		nil,
		4,
		[]string{"aa", "bb"},
	)
	require.Error(t, err)
	require.Empty(t, output)
}

func Test_getRangeFromCandidates(t *testing.T) {
	type testcase struct {
		candidates []int64
		lengthMap  map[int64][2]int
		expected   [2]int64
	}
	testcases := []testcase{
		{[]int64{}, map[int64][2]int{}, [2]int64{-1, -1}},
		{[]int64{2, 3, 4, 5}, map[int64][2]int{2: {0, 3}, 5: {10, 20}}, [2]int64{0, 20}},
		{[]int64{2}, map[int64][2]int{2: {0, 3}}, [2]int64{0, 3}},
		{[]int64{2, 5}, map[int64][2]int{2: {0, 3}}, [2]int64{0, 3}},
		{[]int64{2, 5}, map[int64][2]int{5: {0, 3}}, [2]int64{0, 3}},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			actual := getRangeFromCandidates(tc.candidates, tc.lengthMap)
			require.Equal(t, tc.expected, actual)
		})
	}
}
