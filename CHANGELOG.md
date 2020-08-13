## 0.4.0 (2020/08/13)

INCOMPATIBILITIES AND NOTES:

* Terraform v0.13 support ([#21](https://github.com/minamijoyo/tfschema/pull/21))

## 0.3.0 (2019/05/23)

INCOMPATIBILITIES AND NOTES:

* Terraform v0.12 support ([#14](https://github.com/minamijoyo/tfschema/pull/14))

You can use both Terraform v0.11/v0.12 supported providers.

* Change type notation to HCL2 type annotation ([#16](https://github.com/minamijoyo/tfschema/pull/16))

For most tfschema users, this appears as if the type notation had just been changed to lowercase.
It was originally capitalized because cty's Go type was capitalized.
I know we can still use capitalized letters for maximum compatibility, but I believe that it will be easier to be consistent to use HCL2 style when representing complex data types.

## 0.2.0 (2018/08/31)

INCOMPATIBILITIES AND NOTES:

* Change JSON output format for easy parsing ([#6](https://github.com/minamijoyo/tfschema/pull/6))

## 0.1.2 (2018/07/27)

BUG FIXES:

* Use newest plugin when multiple versions are found ([#2](https://github.com/minamijoyo/tfschema/pull/2))
