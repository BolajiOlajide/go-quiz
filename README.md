# GO QUIZ

Quiz solution from Gophercises.com

To run the app, run the command

```bash
make build
```

from your terminal and run the generated build file using the command

```bash
./quiz
```

this defaults to using the `problems.csv` file, if you wish to pass in your own csv use the -csv flag

```bash
./quiz -csv="new csv.csv"
```

you can also set a limit with the flag `-limit`, the default is 30seconds and the unit to be used is seconds.
An example is shown below:

```bash
./quiz -limit=10
```