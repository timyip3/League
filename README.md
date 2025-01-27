To run the code, please execute it with the command "go run ." to run the main method on main.go under the directory League.

To run the functions, please send the request(s) with:
/echo:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

/invert:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"

/flatten:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"

/sum:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"
        
/multiply:
        curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"