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

# Link GLFW to the library
target_link_libraries(festi PUBLIC glfw)

# On macOS, link against the OpenGL framework
if(APPLE)
    target_link_libraries(festi PUBLIC "-framework OpenGL")
# On Linux, link against the OpenGL library
elseif(UNIX AND NOT APPLE)
    target_link_libraries(festi PUBLIC OpenGL)
# On Windows, link against the OpenGL library
elseif(WIN32)
    target_link_libraries(festi PUBLIC opengl32)
endif()

# Set the output directory for the library
set_target_properties(festi PROPERTIES LIBRARY_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin/festi/)

# Optionally, create an executable for testing
add_executable(test_program test/test.cpp)

# Link the test executable with the library
target_link_libraries(test_program PRIVATE festi)
set_target_properties(test_program PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin/test)
