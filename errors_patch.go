package tun

// isErrNotPollable returns true if the error is internal/poll.ErrNotPollable.
func isErrNotPollable(err error) bool {
        if err == nil {
                return false
        }
        if err.Error() == "not pollable" {
                return true
        }

        switch err := err.(type) {
        case interface{ Unwrap() error }:
                return isErrNotPollable(err.Unwrap())
        case interface{ Unwrap() []error }:
                for _, e := range err.Unwrap() {
                        if isErrNotPollable(e) {
                                return true
                        }
                }
        }

        return false
}
