JSON validator
==============

*A simple JSON validator written in go*

### How to run the tests

Move inside the root directory of the project and run `go test` from the terminal

---

### How to use the tool

Move inside the root directory of the project and run the following command:

* `go run go-json [filename]` from the terminal

### Arguments

#### filename(mandatory):

* the name of the file to process, it can be a relative path or an absolute one

---

### Technical details

* The semantic definition of an object or array or value is compliant to the standard: https://www.json.org/json-en.html
    * ![object.png](assets/object.png)
    * ![array.png](assets/array.png)
    * ![value.png](assets/value.png)

* The semantic definition of a whitespace or number or a string is not compliant to the standard because of the lack of
  value from the learning point of view
    * a whitespace is a `" "` character, excluding linefeeds, carriage returns and horizontal tabs;
    * a number is a sequence of one or more digits belonging to the set `{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}` excluding the
      scientific notation;
    * a string is a sequence of one or more unicode character excluding quotations marks, reverse solidus, solidus,
      backspaces, formfeeds, linefeeds, carriage returns, horizontal tabs and hex values
