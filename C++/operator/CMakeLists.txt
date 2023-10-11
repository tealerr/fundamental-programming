# C++/operator/CMakeLists.txt

add_library(operator_lib arithmetic.cpp assignment.cpp logical.cpp)

# Make sure to provide the correct path to main_operator.cpp if it's not in the same directory
add_executable(operator_exe main_operator.cpp)
target_link_libraries(operator_exe operator_lib)
