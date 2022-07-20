package faker

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strconv"
	"strings"
)

// Faker is Faker instance
type Faker struct {
	options *Options
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
	options := setOptions(opts...)
	return &Faker{options: options}
}

type (
	// Options -.
	Options struct {
		rand *rand.Rand
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
		// Person
		firstNames []string
		lastNames  []string
	}

	// Option -.
	Option func(opts *Options)
)

// SetRand sets Rand instance
func (o *Options) SetRand(r *rand.Rand) {
	o.rand = r
}

// SetPostCodeFormats sets post code formats
func (o *Options) SetPostCodeFormats(postCodeFormats ...string) {
	o.postCodeFormats = postCodeFormats
}

// SetColors sets colors
func (o *Options) SetColors(colors ...string) {
	o.colors = colors
}

// SetHexSymbols sets hex symbols
func (o *Options) SetHexSymbols(hexSymbols string) {
	o.hexSymbols = hexSymbols
}

// SetWeekdays sets weekdays
func (o *Options) SetWeekdays(weekdays ...string) {
	o.weekdays = weekdays
}

// SetMonths sets months
func (o *Options) SetMonths(months ...string) {
	o.months = months
}

// SetPasswordMin sets minimal length of password
func (o *Options) SetPasswordMin(passwordMin int) {
	o.passwordMin = passwordMin
}

// SetPasswordMax sets maximum length of password
func (o *Options) SetPasswordMax(passwordMax int) {
	o.passwordMax = passwordMax
}

// SetPasswordChars sets password chars
func (o *Options) SetPasswordChars(passwordChars string) {
	o.passwordChars = passwordChars
}

// SetGenericTopLevelDomains sets generic top-level domains
func (o *Options) SetGenericTopLevelDomains(genericTopLevelDomains ...string) {
	o.genericTopLevelDomains = genericTopLevelDomains
}

// SetHTTPMethods sets HTTP methods
func (o *Options) SetHTTPMethods(httpMethods ...string) {
	o.httpMethods = httpMethods
}

// SetFirstNames sets first names
func (o *Options) SetFirstNames(firstNames ...string) {
	o.firstNames = firstNames
}

// SetLastNames sets last names
func (o *Options) SetLastNames(lastNames ...string) {
	o.lastNames = lastNames
}

// WithRand -.
func WithRand(r *rand.Rand) Option {
	return func(opts *Options) {
		opts.rand = r
	}
}

// WithPostCodeFormats -.
func WithPostCodeFormats(postCodeFormats ...string) Option {
	return func(opts *Options) {
		opts.postCodeFormats = postCodeFormats
	}
}

// WithColors -.
func WithColors(colors ...string) Option {
	return func(opts *Options) {
		opts.colors = colors
	}
}

// WithHexSymbols -.
func WithHexSymbols(hexSymbols string) Option {
	return func(opts *Options) {
		opts.hexSymbols = hexSymbols
	}
}

// WithWeekdays -.
func WithWeekdays(weekdays ...string) Option {
	return func(opts *Options) {
		opts.weekdays = weekdays
	}
}

// WithMonths -.
func WithMonths(months ...string) Option {
	return func(opts *Options) {
		opts.months = months
	}
}

// WithPasswordMin -.
func WithPasswordMin(passwordMin int) Option {
	return func(opts *Options) {
		opts.passwordMin = passwordMin
	}
}

// WithPasswordMax -.
func WithPasswordMax(passwordMax int) Option {
	return func(opts *Options) {
		opts.passwordMax = passwordMax
	}
}

// WithPasswordChars -.
func WithPasswordChars(passwordChars string) Option {
	return func(opts *Options) {
		opts.passwordChars = passwordChars
	}
}

// WithGenericTopLevelDomains -.
func WithGenericTopLevelDomains(genericTopLevelDomains ...string) Option {
	return func(opts *Options) {
		opts.genericTopLevelDomains = genericTopLevelDomains
	}
}

// WithHTTPMethods -.
func WithHTTPMethods(httpMethods ...string) Option {
	return func(opts *Options) {
		opts.httpMethods = httpMethods
	}
}

// WithFirstNames -.
func WithFirstNames(firstNames ...string) Option {
	return func(opts *Options) {
		opts.firstNames = firstNames
	}
}

// WithLastNames -.
func WithLastNames(lastNames ...string) Option {
	return func(opts *Options) {
		opts.lastNames = lastNames
	}
}

func setOptions(opts ...Option) *Options {
	var options Options
	for _, opt := range opts {
		opt(&options)
	}
	if options.rand == nil {
		var seed int64
		if err := binary.Read(crand.Reader, binary.BigEndian, &seed); err != nil {
			panic(err)
		}
		options.SetRand(rand.New(rand.NewSource(seed).(rand.Source64)))
	}
	return &options
}

// Integer returns a integer between a given minimum and maximum values
func (f *Faker) Integer(min, max int) int {
	return Integer(min, max, WithRand(f.options.rand))
}

// Integer returns a int between a given minimum and maximum values
func Integer(min, max int, opts ...Option) int {
	diff := max - min
	if diff == 0 {
		return min
	}
	options := setOptions(opts...)
	return options.rand.Intn(diff) + min
}

// Number returns a float64 between a given minimum and maximum values
func (f *Faker) Number(min, max float64) float64 {
	return Number(min, max, WithRand(f.options.rand))
}

// Number returns a float64 between a given minimum and maximum values
func Number(min, max float64, opts ...Option) float64 {
	options := setOptions(opts...)
	numbers := make([]float64, 10)
	for i := range numbers {
		numbers[i] = min + options.rand.Float64()*(max-min)
	}
	return RandomElement(numbers)
}

// RandomElement returns a random element from a given slice
func RandomElement[T any](slice []T, opts ...Option) T {
	i := Integer(0, len(slice)-1, opts...)
	return slice[i]
}

// Numerify returns a string that replace all "*" characters with numbers from 0 to 9 as string
func Numerify(in string, opts ...Option) string {
	var out strings.Builder
	for i := range in {
		if in[i] == '*' {
			out.WriteString(strconv.Itoa(Integer(0, 9, opts...)))
		} else {
			out.WriteByte(in[i])
		}
	}
	return out.String()
}
