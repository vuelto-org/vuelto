cmake_minimum_required(VERSION 3.10)
project(festi)

# Set the source files for the library
file(GLOB LIBRARY_SOURCES "src/festi/*.c")

# Create the shared library
add_library(festi SHARED ${LIBRARY_SOURCES})

# Set the include directories for the library
target_include_directories(festi PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/src)

# Find GLFW package
find_package(glfw3 REQUIRED)

# Link GLFW and OpenGL framework to the library
target_link_libraries(festi PUBLIC glfw "-framework OpenGL")

# Set the output directory for the library
set_target_properties(festi PROPERTIES LIBRARY_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/build)

# Optionally, create an executable for testing
add_executable(test_program test/test.c)

# Link the test executable with the library
target_link_libraries(test_program PRIVATE festi)

