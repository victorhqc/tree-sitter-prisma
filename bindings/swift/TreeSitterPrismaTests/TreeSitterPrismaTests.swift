import XCTest
import SwiftTreeSitter
import TreeSitterPrisma

final class TreeSitterPrismaTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_prisma())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Prisma grammar")
    }
}
