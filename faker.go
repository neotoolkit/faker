package faker

import (
	"math/rand"
	"time"
)

// Faker -.
type Faker struct {
	options *Options
}

// New -.
func New(opts ...Option) *Faker {
	options := setOptions(opts...)
	return &Faker{options: options}
}

type (
	// Option -.
	Option func(opts *Options)

	// Options -.
	Options struct {
		rand *rand.Rand
		// Color
		colors     []string
		hexLetters []string
		// Date
		weekdays []string
		months   []string
		// Internet
		passwordMin            int
		passwordMax            int
		passwordChars          string
		genericTopLevelDomains []string
		// Person
		firstNames []string
		lastNames  []string
	}
)

// SetRand -.
func (o *Options) SetRand(r *rand.Rand) {
	o.rand = r
}

// SetColors -.
func (o *Options) SetColors(colors ...string) {
	o.colors = colors
}

// SetHexLetters -.
func (o *Options) SetHexLetters(hexLetters ...string) {
	o.hexLetters = hexLetters
}

// SetWeekdays -.
func (o *Options) SetWeekdays(weekdays ...string) {
	o.weekdays = weekdays
}

// SetMonths -.
func (o *Options) SetMonths(months ...string) {
	o.months = months
}

// SetPasswordMin -.
func (o *Options) SetPasswordMin(passwordMin int) {
	o.passwordMin = passwordMin
}

// SetPasswordMax -.
func (o *Options) SetPasswordMax(passwordMax int) {
	o.passwordMax = passwordMax
}

// SetPasswordChars -.
func (o *Options) SetPasswordChars(passwordChars string) {
	o.passwordChars = passwordChars
}

// SetGenericTopLevelDomains -.
func (o *Options) SetGenericTopLevelDomains(genericTopLevelDomains ...string) {
	o.genericTopLevelDomains = genericTopLevelDomains
}

// SetFirstNames -.
func (o *Options) SetFirstNames(firstNames ...string) {
	o.firstNames = firstNames
}

// SetLastNames -.
func (o *Options) SetLastNames(lastNames ...string) {
	o.lastNames = lastNames
}

// SetRand -.
func SetRand(r *rand.Rand) Option {
	return func(opts *Options) {
		opts.rand = r
	}
}

// SetColors -.
func SetColors(colors ...string) Option {
	return func(opts *Options) {
		opts.colors = colors
	}
}

// SetHexLetters -.
func SetHexLetters(hexLetters ...string) Option {
	return func(opts *Options) {
		opts.hexLetters = hexLetters
	}
}

// SetWeekdays -.
func SetWeekdays(weekdays ...string) Option {
	return func(opts *Options) {
		opts.weekdays = weekdays
	}
}

// SetMonths -.
func SetMonths(months ...string) Option {
	return func(opts *Options) {
		opts.months = months
	}
}

// SetPasswordMin -.
func SetPasswordMin(passwordMin int) Option {
	return func(opts *Options) {
		opts.passwordMin = passwordMin
	}
}

// SetPasswordMax -.
func SetPasswordMax(passwordMax int) Option {
	return func(opts *Options) {
		opts.passwordMax = passwordMax
	}
}

// SetPasswordChars -.
func SetPasswordChars(passwordChars string) Option {
	return func(opts *Options) {
		opts.passwordChars = passwordChars
	}
}

// SetGenericTopLevelDomains -.
func SetGenericTopLevelDomains(genericTopLevelDomains ...string) Option {
	return func(opts *Options) {
		opts.genericTopLevelDomains = genericTopLevelDomains
	}
}

// SetFirstNames -.
func SetFirstNames(firstNames ...string) Option {
	return func(opts *Options) {
		opts.firstNames = firstNames
	}
}

// SetLastNames -.
func SetLastNames(lastNames ...string) Option {
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
		options.SetRand(rand.New(rand.NewSource(time.Now().Unix())))
	}
	return &options
}

// Integer returns a integer between a given minimum and maximum values
func (f *Faker) Integer(min, max int) int {
	return Integer(min, max, SetRand(f.options.rand))
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
	return Number(min, max, SetRand(f.options.rand))
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
