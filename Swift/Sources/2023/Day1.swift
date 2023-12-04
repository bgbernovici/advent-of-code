import Foundation

/*
    https://adventofcode.com/2023/day/1
    Bogdan Bernovici
*/

extension String {
    func parseToInt() -> String? {
        return self.components(separatedBy: CharacterSet.decimalDigits.inverted).joined()
    }
}

struct SumGenerator {
    var _sum = 0

    mutating func sumNext(_ element: Int) -> Void {
        _sum += element 
    }

    var sum: Int {
        return _sum
    }
}

@main
struct Day1 {
    
    public static func main() {
        let inputURL = URL(fileURLWithPath: "Sources/../../Inputs/2023/Day1_.txt")
        print("Part 1: ", sumCalibrations(inputURL, false))
        print("Part 2: ", sumCalibrations(inputURL, true))
    }

    public static func sumCalibrations(_ inputURL: URL, _ needsRecalibration: Bool) -> Int {
        var summator = SumGenerator()
        do {
            let input = try String(contentsOf: inputURL, encoding: .utf8)
            let elements = calibrate(input, needsRecalibration)
            for element in elements {
                summator.sumNext(element)
            }
            return summator.sum
        } catch {
            print("Error reading input file: \(error)")
        }
        return 0
    }

    public static func calibrate(_ input: String, _ needsRecalibration: Bool) -> [Int] {
        let rows = input.components(separatedBy: CharacterSet.newlines).filter { !$0.isEmpty }
        var calibratedValues = [Int]()
        for var row in rows {
            if needsRecalibration {
                recalibrate(line: &row)
            }
            let optionalValue = row.parseToInt()
            guard let value = optionalValue else {
                continue
            }
            if value.count == 1 {
                guard let v = Int(String(repeating: value, count: 2)) else {
                    continue
                }
                calibratedValues.append(v)
            } else {
                guard let f = value.first, let l = value.last else {
                    continue
                }
                guard let v = Int(String([f, l])) else {
                    continue
                }
                calibratedValues.append(v)
            }
        }
        return calibratedValues
    }

    static let domain = [
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
    ]

    public static func recalibrate(line: inout String) {
        var newLine = ""

        for (k, v) in domain {
            line = line.replacingOccurrences(of: String(v), with: k)
        }
    
        var i = 0
        while i < line.count {
            for (k, v) in domain {
                let spanning = i + k.count
                let clamped_spanning = min(spanning, line.count)
                let startIndex = line.index(line.startIndex, offsetBy: i)
                let endIndex = line.index(line.startIndex, offsetBy: clamped_spanning)
                let subString = line[startIndex..<endIndex]
                if subString.contains(k) {
                    newLine += String(v)
                    break;
                }
            }
            i += 1
        }
        line = newLine
    }
}