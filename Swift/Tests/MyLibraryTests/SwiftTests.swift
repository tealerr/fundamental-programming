import XCTest
@testable import MyLibrary

class SwiftTests: XCTestCase {
    func test_WhenInputScore_80_ShouldReturn_A() {
        let grade = gradeCal(score: 85)
        let result = showGrade(grade: grade)
        XCTAssertEqual("Your grade is A", result)
    }

    func test_WhenInputScore_75_ShouldReturn_B() {
        let grade = gradeCal(score: 75)
        let result = showGrade(grade: grade)
        XCTAssertEqual("Your grade is B", result)
    }

    func test_WhenInputScore_60_ShouldReturn_F() {
        let grade = gradeCal(score: 60)
        let result = showGrade(grade: grade)
        XCTAssertEqual("Your grade is F", result)
        }
}
