# runtime_utils

Golang related runtime utilities.

## Caller utilities

The caller functions provide helpers to get information about the function and package that actually called the
functions.

  - GetCallerInfo - Parse caller stack.
  - GetCallerFunctionName - Return the caller of that function with it's stack information.
  - GetCallerStack - Provides A list of the entire execution stack, with possibilities to skip callers and limit the
    amount of results.


## License

The current package is released under [MPL v2](https://www.mozilla.org/en-US/MPL/2.0/) license.

