# parse-equations

`parse-equations` is a program to convert Mathematica or Maple style linear equations to the format for `go-finite-field.`

## Usage

To clone the source code to a local folder, run
```
> git clone https://github.com/andyytliu/parse-equations.git
> cd parse-equations
```

To start the program, run the following command with flags
```
> go run main.go -input=input_file.txt  -output=output_file.txt -vars=vars_file.txt
```

`input_file.txt` is a file with Mathematica or Maple style linear equations;
it should contain a list of linear combinations of variables (to be equated to zero), inside curly brackets `{ }` and separated by commas. For example
```
{ 3*x + 5/8*y - 18/7, 87/24*y - 5*z + 103/2, x + y + z }
```
Note that the multiplication sign `*` is essential and shouldn't be omitted; replacing it with an empty space would cause error when parsing equations.

`output_file.txt` is the file to write parsed equations to in the `go-finite-field` format.

`vars_file.txt` is a list of variables that appear in `input_file.txt`.
`go-finite-field` uses indices to represent variables, and `vars_file.txt` would act as translation from variables to indices;
for example, it could contain the following
```
{ x, y, z }
```
which would translate `x` to index `1`, `y` to `2`, and `z` to `3`.
