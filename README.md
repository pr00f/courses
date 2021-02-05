# courses

**Strings** + **Test System**

    Tests result: all tests were passed
    Mark: 4 bytes
    Duration: ~1 hour
    Run test: go test ./tasks/strings -count 1 -v

**Lucky Tickets**

    Tests result: all tests were passed
    Mark: 3 bytes
    Duration: ~1.5 hour
    Run test: go test ./tasks/tickets -count 1 -v

**Spells**

    Tests result: 12 tests were passed
    Mark: 3 bytes
    Duration: ~1 hour
    Answers destination: /data/spells/answers
    Run test: go test ./tasks/spells -count 1 -v

**Power**

    Tests result: 10 tests were passed
    Mark: 2 bytes
    Duration: ~1 hour
    Run test: go test ./tasks/power -count 1 -v
    Сomparison table: 

Name  | 2^1.000.000.000 | 2^10.000.000.000 | 2^100.000.000.000
------------- | ------------- | ------------- | -------------
SolutionSimple  | 1.22s | 12.47s | 130.61s
SolutionPowerOfTwoWithMultiplication  | 0.58s | 1.80s | 40.32s
SolutionBinaryExpansionOfExponent  | 0.00s | 0.00s | 0.00s

**Fibo**

    Tests result: 7 tests were passed
    Mark: ?
    Reason: very difficult to work with big numbers (do not have pow for 'big number' library in language)
    Duration: ~2 hour
    Run test: go test ./tasks/fibo -count 1 -v

