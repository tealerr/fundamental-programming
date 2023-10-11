# C++/CMakeLists.txt

cmake_minimum_required(VERSION 3.5)
project(myCmakeCpp VERSION 0.1.0 LANGUAGES C CXX)

include(CTest)
enable_testing()

add_subdirectory(operator)  # Add this line

add_executable(myCmakeCpp main.cpp condition.cpp)
target_link_libraries(myCmakeCpp operator_lib)  # Link with the operator_lib library

set(CPACK_PROJECT_NAME ${PROJECT_NAME})
set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
include(CPack)

# Custom clean target
add_custom_target(clean-all
    COMMAND ${CMAKE_COMMAND} --build ${CMAKE_BINARY_DIR} --target clean
    COMMAND ${CMAKE_COMMAND} -E remove_directory ${CMAKE_BINARY_DIR}/*
)
