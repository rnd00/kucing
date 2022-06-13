# kucing

an attempt to recreate [**kochengs**](https://github.com/rnd00/kochengs) with better understanding in concurrency;

in short -- an *'attempt on learning'* go-concurrency and channels.

per [**kochengs**](https://github.com/rnd00/kochengs), is still;
  > Will generate cat image, based from [thiscatdoesnotexist](https://thiscatdoesnotexist.com)

## plan

three modules,
  - `download`
  - `compare`
  - `cache`

and a `helpers`,

- will try as many goroutine as needed for `download` and `compare`, both communicating through channels,
  - `download` will be a simple GET operation, passing the retrieved data to channels toward `compare`
  - `compare` will get the first data inside channels, take the parts of data as `[]byte`, use it as `string`  and then check it to `cache`
    - `map[string]bool` => `{"SomeKindOfKeyHereFromData": true}`
    - if exist then return
    - if nonexist (returning `false`) then save to `cache` as `true` 
- `cache` will be in-memory cache, holding parts of `[]byte` from the retrieved img, using it as `string` key in `map[string]bool` type

for logging;
- spin another routine for printing logs from `helpers`

!! !! **if everything works, make a `cli` on top of it** !! !!