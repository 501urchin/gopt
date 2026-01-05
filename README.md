# go-optimizations

A collection of **low-level, performance-oriented optimizations for Go** focused on reducing **allocations** or **CPU time** in hot paths.

This package intentionally leverages the `unsafe` package and other non-idiomatic techniques where they provide measurable performance benefits. As a result, **correctness, safety, and portability are the responsibility of the consumer**.

If you need safety, simplicity, or long-term stability, this package is likely **not appropriate**.


## Important Warning

This package **uses `unsafe` extensively**.

By using this library, you acknowledge that:

* APIs may **violate Go’s memory safety guarantees** if misused
* Incorrect usage can lead to **undefined behavior**, **memory corruption**, or **hard-to-debug crashes**
* Behavior may change across **Go versions**, **architectures**, or **GC implementations**

If you are unsure whether an optimization is safe for your use case, **do not use it**.


## Goals

* Reduce heap allocations
* Reduce CPU cycles in hot paths
* Avoid unnecessary copying
* Provide opt-in, explicit performance trade-offs
* Keep all unsafe behavior well-documented and isolated

This is **not** a general-purpose utility library. Each optimization targets a specific, measurable bottleneck.


## Versioning & Compatibility

* Go version compatibility is **best-effort**
* Minor Go releases may change runtime behavior and break assumptions
* No backward-compatibility guarantees are provided

Pin the module version and re-run benchmarks and tests when upgrading Go.

## Testing Recommendations

When using this package:

* Test under load
* Run with `-race`
* Add regression tests around every call site
* Validate behavior across Go versions
* Prefer benchmarks over intuition

If an optimization does not produce a measurable improvement in your workload, remove it.

