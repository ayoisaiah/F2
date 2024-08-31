package config

import "github.com/ayoisaiah/f2/internal/apperr"

var (
	errInvalidArgument = &apperr.Error{
		Message: "requires one of: -f, -r, --csv, or -u. Run f2 --help for usage",
	}

	errInvalidSimpleModeArgs = &apperr.Error{
		Message: "at least one argument must be specified in simple mode",
	}

	errParsingFixConflictsPattern = &apperr.Error{
		Message: "the provided --fix-conflicts-pattern '%s' is invalid",
	}

	errInvalidSort = &apperr.Error{
		Message: "the provided sort '%s' is invalid",
	}
)