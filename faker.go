package faker

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// Faker is Faker instance
type Faker struct {
	cfg *Config
}

// New returns a new Faker instance with options
//
//    faker.New(
//        faker.WithRand(rand.New(rand.NewSource(time.Now().Unix()))),
//        faker.WithColors("", ""),
//        faker.WithHexSymbols(""),
//        faker.WithWeekdays("", ""),
//        faker.WithMonths("", ""),
//        faker.WithPasswordMin(0),
//        faker.WithPasswordMax(0),
//        faker.WithPasswordChars(""),
//        faker.WithGenericTopLevelDomains("", ""),
//        faker.WithFirstNames("", ""),
//        faker.WithLastNames("", ""),
//    )
//
func New(opts ...Option) *Faker {
	cfg := newConfig(opts...)
	return &Faker{cfg: cfg}
}

type (
	// Config -.
	Config struct {
		rand *rand.Rand
		// Unicode characters as integers
		asciifyUnicodeDecimals []int
		// Address
		postCodeFormats []string
		// Auth
		passwordMin   int
		passwordMax   int
		passwordChars string
		// Color
		colors     []string
		hexSymbols string
		// Date
		weekdays []string
		months   []string
		// Internet
		genericTopLevelDomains []string
		httpMethods            []string
		httpStatusCodes        []int
		// Person
		firstNames []string
		lastNames  []string
	}

	// Option -.
	Option func(opts *Config)
)

// WithRand -.
func WithRand(r *rand.Rand) Option {
	return func(cfg *Config) {
		cfg.rand = r
	}
}

// WithAsciifyUnicodeDecimals -.
func WithAsciifyUnicodeDecimals(asciifyUnicodeDecimals ...int) Option {
	return func(cfg *Config) {
		cfg.asciifyUnicodeDecimals = asciifyUnicodeDecimals
	}
}

// WithPostCodeFormats -.
func WithPostCodeFormats(postCodeFormats ...string) Option {
	return func(cfg *Config) {
		cfg.postCodeFormats = postCodeFormats
	}
}

// WithColors -.
func WithColors(colors ...string) Option {
	return func(cfg *Config) {
		cfg.colors = colors
	}
}

// WithHexSymbols -.
func WithHexSymbols(hexSymbols string) Option {
	return func(cfg *Config) {
		cfg.hexSymbols = hexSymbols
	}
}

// WithWeekdays -.
func WithWeekdays(weekdays ...string) Option {
	return func(cfg *Config) {
		cfg.weekdays = weekdays
	}
}

// WithMonths -.
func WithMonths(months ...string) Option {
	return func(cfg *Config) {
		cfg.months = months
	}
}

// WithPasswordMin -.
func WithPasswordMin(passwordMin int) Option {
	return func(cfg *Config) {
		cfg.passwordMin = passwordMin
	}
}

// WithPasswordMax -.
func WithPasswordMax(passwordMax int) Option {
	return func(cfg *Config) {
		cfg.passwordMax = passwordMax
	}
}

// WithPasswordChars -.
func WithPasswordChars(passwordChars string) Option {
	return func(cfg *Config) {
		cfg.passwordChars = passwordChars
	}
}

// WithGenericTopLevelDomains -.
func WithGenericTopLevelDomains(genericTopLevelDomains ...string) Option {
	return func(cfg *Config) {
		cfg.genericTopLevelDomains = genericTopLevelDomains
	}
}

// WithHTTPMethods -.
func WithHTTPMethods(httpMethods ...string) Option {
	return func(cfg *Config) {
		cfg.httpMethods = httpMethods
	}
}

// WithHTTPStatusCodes -.
func WithHTTPStatusCodes(httpStatusCodes ...int) Option {
	return func(cfg *Config) {
		cfg.httpStatusCodes = httpStatusCodes
	}
}

// WithFirstNames -.
func WithFirstNames(firstNames ...string) Option {
	return func(cfg *Config) {
		cfg.firstNames = firstNames
	}
}

// WithLastNames -.
func WithLastNames(lastNames ...string) Option {
	return func(cfg *Config) {
		cfg.lastNames = lastNames
	}
}

func newConfig(opts ...Option) *Config {
	var cfg Config
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.rand == nil {
		var seed int64
		if err := binary.Read(crand.Reader, binary.BigEndian, &seed); err != nil {
			panic(err)
		}
		WithRand(rand.New(rand.NewSource(seed).(rand.Source64)))(&cfg)
	}
	return &cfg
}

// Integer returns a integer between a given minimum and maximum values
func (f *Faker) Integer(min, max int) int {
	return Integer(min, max, WithRand(f.cfg.rand))
}

// Integer returns a int between a given minimum and maximum values
func Integer(min, max int, opts ...Option) int {
	diff := max - min
	if diff == 0 {
		return min
	}
	cfg := newConfig(opts...)
	return cfg.rand.Intn(diff) + min
}

// Number returns a float64 between a given minimum and maximum values
func (f *Faker) Number(min, max float64) float64 {
	return Number(min, max, WithRand(f.cfg.rand))
}

// Number returns a float64 between a given minimum and maximum values
func Number(min, max float64, opts ...Option) float64 {
	cfg := newConfig(opts...)
	numbers := make([]float64, 10)
	for i := range numbers {
		numbers[i] = min + cfg.rand.Float64()*(max-min)
	}
	return RandomElement(numbers)
}

// RandomElement returns a random element from a given slice
func RandomElement[T any](slice []T, opts ...Option) T {
	i := Integer(0, len(slice)-1, opts...)
	return slice[i]
}

// Numerify returns a string that replace all "#" characters with numbers from 0 to 9 as string
func (f *Faker) Numerify(in string) string {
	return Numerify(in, WithRand(f.cfg.rand))
}

// Numerify returns a string that replace all "#" characters with numbers from 0 to 9 as string
func Numerify(in string, opts ...Option) string {
	var out strings.Builder
	out.Grow(len(in))
	for i := range in {
		if in[i] == '#' {
			out.WriteString(strconv.Itoa(Integer(0, 9, opts...)))
		} else {
			out.WriteByte(in[i])
		}
	}
	return out.String()
}

// Asciify returns a string that replace all "*" characters with latin letters
func (f *Faker) Asciify(in string) string {
	return Numerify(
		in,
		WithRand(f.cfg.rand),
		WithAsciifyUnicodeDecimals(f.cfg.asciifyUnicodeDecimals...),
	)
}

// Asciify returns a string that replace all "*" characters with unicode characters
func Asciify(in string, opts ...Option) string {
	cfg := newConfig(opts...)
	if len(cfg.asciifyUnicodeDecimals) == 0 {
		latinLetters := make([]int, 0, 50)
		// Capital letters
		latinLetters = append(latinLetters, intArrRange(65, 90)...)
		// Small letters
		latinLetters = append(latinLetters, intArrRange(97, 122)...)
		WithAsciifyUnicodeDecimals(latinLetters...)(cfg)
	}
	var out strings.Builder
	out.Grow(len(in))
	for i := range in {
		if in[i] == '*' {
			out.WriteString(fmt.Sprintf("%c", RandomElement(cfg.asciifyUnicodeDecimals, opts...)))
		} else {
			out.WriteByte(in[i])
		}
	}
	return out.String()
}

func intArrRange(min, max int) []int {
	arr := make([]int, max-min+1)
	for i := range arr {
		arr[i] = min + i
	}
	return arr
}
