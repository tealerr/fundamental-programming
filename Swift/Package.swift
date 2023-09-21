// swift-tools-version:5.7.1
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "Swift",
    products: [
        .library(
            name: "Swift",
            targets: ["Swift"]),
    ],
    targets: [
        // Targets are the basic building blocks of a package, defining a module or a test suite.
        // Targets can depend on other targets in this package and products from dependencies.
        .target(
            name: "Swift"),
        .testTarget(
            name: "SwiftTests",
            dependencies: ["Swift"]),
    ]
)