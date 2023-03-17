v0.6.0 (2022-07-28)
-----------
Add: WithTraceID(), Ctx() method to interface, allow to integrate with tracing
Fix: Fix NewLogger() created custom logger stacktrace caller incorrect problem

v0.5.0 (2022-07-21)
-----------
Change: merge v2 code back into v1, and we will remove v2 soon

v2.0.1 (2022-07-07)
-----------
Removed: exported var ZapLogger has been removed

v2.0.0 (2022-07-07)
-----------
Refactor: upgrade mod version to v2, split gorm and echo logger to standalone pkg


v0.4.4 (2022-02-21)
-----------
Feature: interface expose `SetLevel()`


v0.4.2 (2022-02-16)
-----------
Change: `With()` signature changed to allow varadic params
Feature: With now allow chain call like `.With().With().With()`


v0.4.0 (2021-6-15)
-----------
Feature: level color
Feature: short time


v0.3.4 (2020-12-04)
-----------
Change: echo log ignore /metrics


v0.3.3 (2020-11-26)
-----------
Change: echo log ignore /healthz


v0.3.2 (2020-11-10)
-----------
Feature: colorized


v0.3.0 (2020-9-08)
-----------
Feature: logger interface


v0.2.2 (2020-9-03)
-----------
Change: field pair


v0.2.1 (2020-8-26)
-----------
Change: set CallerSkip


v0.2.0 (2020-7-31)
-----------
Feature: echo logger


v0.1.0 (2020-7-30)
-----------
* The first public version
