cmake_minimum_required(VERSION 3.10)
project(vuelto)

# Add the GLFW submodule
add_subdirectory(libs/glfw)

# Set the source files for the library
file(GLOB LIBRARY_SOURCES "src/vuelto/core/*.cpp")

# Create the shared library
add_library(vuelto SHARED ${LIBRARY_SOURCES})

# Set the include directories for the library
target_include_directories(vuelto PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/src)

# Link GLFW to the library
target_link_libraries(vuelto PUBLIC glfw)

# On macOS, link against the OpenGL framework
if(APPLE)
    target_link_libraries(vuelto PUBLIC "-framework OpenGL")
# On Linux, link against the OpenGL library
elseif(UNIX AND NOT APPLE)
    target_link_libraries(vuelto PUBLIC OpenGL)
endif()

# Set the output directory for the library
set_target_properties(vuelto PROPERTIES LIBRARY_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin/vuelto/)

# Add -Wdeprecated-declarations for GCC and Clang
if(CMAKE_COMPILER_IS_GNUCC OR CMAKE_COMPILER_IS_GNUCXX OR CMAKE_C_COMPILER_ID MATCHES "Clang")
    add_compile_options(-Wdeprecated-declarations)
endif()

# Optionally, create an executable for testing
add_executable(test_program test/test.cpp)

# Link the test executable with the library
target_link_libraries(test_program PRIVATE vuelto)
set_target_properties(test_program PROPERTIES RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/bin/test)
