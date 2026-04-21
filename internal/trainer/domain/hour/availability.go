package hour

// Availability is enum.
//
// Using struct instead of `type Availability string` for enums allows us to ensure,
// that we have full control of what values are possible.
// With `type Availability string` you are able to create `Availability("i_can_put_anything_here")`
type Availability struct {
	a string
}

// Every type in Go have zero value. In that case it's `Availability{}`.
// It's always a good idea to check if provided value is not zero!
func (h Availability) IsZero() bool {
	return h == Availability{}
}

var (
	Available         = Availability{"available"}
	NotAvailable      = Availability{"not_available"}
	TrainingScheduled = Availability{"training_scheduled"}
)
