# kucing

`kucing` is an indonesian word for `cat`

## what

an attempt to recreate [**kochengs**](https://github.com/rnd00/kochengs) with better understanding in concurrency;

in short -- an _'attempt on learning'_ go-concurrency and channels.

per [**kochengs**](https://github.com/rnd00/kochengs), is still;

> Will generate cat image, based from [thiscatdoesnotexist](https://thiscatdoesnotexist.com)

or more actually kind of _will **GET** cat image from thiscatdoesnotexist_

## plan

three modules,

- `download`
- `compare`
- `cache`

and a `helpers`,

- will try as many goroutine as needed for `download` and `compare`, both communicating through channels,
  - `download` will be a simple GET operation, passing the retrieved data to channels toward `compare`
  - `compare` will get the first data inside channels, take the parts of data as `[]byte`, use it as `string` and then check it to `cache`
    - `map[string]bool` => `{"SomeKindOfKeyHereFromData": true}`
    - if exist then return
    - if nonexist (returning `false`) then save to `cache` as `true`
- `cache` will be in-memory cache, holding parts of `[]byte` from the retrieved img, using it as `string` key in `map[string]bool` type

for logging;

- spin another routine for printing logs from `helpers`

!! !! **if everything works, make a `cli` on top of it** !! !!

## variables being used here

note: these variables are being used as is for now, the part where we can change it when running this app has not yet been implemented.

| key                  | value  | desc                                                                                                      |
| -------------------- | ------ | --------------------------------------------------------------------------------------------------------- |
| `workersAmt`         | int    | amount of `workers` being used here. more `workers` = more cpu used                                       |
| `jobsAmt`            | int    | amount of `jobs` being ordered to the `workers`.                                                          |
| `severityLevelPrint` | int    | not yet implemented, will only print equal or more than inputted. `Debug` is 0, `Info` is 2, `Fatal` is 3 |
| `fileCodeLength`     | int    | filename will be `kucing[xxxx].jpg`, this will decide how long `xxxx` will be                             |
| `fileCodeSeed`       | string | characters written here will be used as the base for `fileCode` in `xxxx`                                 |
