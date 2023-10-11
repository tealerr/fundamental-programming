cmake_minimum_required(VERSION 3.0.0)
project(myCmakeCpp VERSION 0.1.0 LANGUAGES C CXX)

include(CTest)
enable_testing()

# Add source files from subdirectories
add_subdirectory(loop)
add_subdirectory(operator)

# Add main executable
add_executable(myCmakeCpp main.cpp condition.cpp)

# Link libraries or targets from subdirectories
target_link_libraries(myCmakeCpp loop_lib operator_lib)

# Add a target for building
add_custom_target(build-all
    COMMAND ${CMAKE_COMMAND} --build ${CMAKE_BINARY_DIR}
    COMMENT "Building all targets"
)

set(CPACK_PROJECT_NAME ${PROJECT_NAME})
set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
include(CPack)

# Custom clean target
add_custom_target(clean-all
    COMMAND ${CMAKE_COMMAND} --build ${CMAKE_BINARY_DIR} --target clean
    COMMAND ${CMAKE_COMMAND} -E remove_directory ${CMAKE_BINARY_DIR}/*
)
