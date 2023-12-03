cmake_minimum_required(VERSION 3.10)
project(vuelto)

# Add the GLFW submodule
add_subdirectory(libs/glfw)

# Set the source files for the library
file(GLOB LIBRARY_SOURCES
    "src/vuelto/core/*.cpp"
    "src/vuelto/event/*.cpp"
    "src/vuelto/shaders/*.cpp"
    "src/vuelto/vendor/GLAD/src/*.c"
)

file(GLOB LIBRARY_HEADERS
    "src/Vuelto.hpp"
    "src/vuelto/core/*.hpp"
    "src/vuelto/event/*.hpp"
    "src/vuelto/tools/*.hpp"
    "src/vuelto/shaders/*.hpp"
    "src/vuelto/vendor/GLFW/*.h"
    "src/vuelto/vendor/GLAD/include/glad/*.h"
)

# Create the shared library
add_library(vuelto SHARED ${LIBRARY_SOURCES} ${LIBRARY_HEADERS})

set_property(TARGET vuelto PROPERTY CXX_STANDARD 11)

# Set the include directories for the library
target_include_directories(vuelto PUBLIC
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/src>
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/src/vuelto/vendor/GLAD/include>
    $<INSTALL_INTERFACE:include>
)

# Copy header files with folder structure under src
foreach(HEADER ${LIBRARY_HEADERS})
    file(RELATIVE_PATH HEADER_RELATIVE ${CMAKE_CURRENT_SOURCE_DIR}/src ${HEADER})
    get_filename_component(HEADER_NAME ${HEADER} NAME)
    get_filename_component(HEADER_PATH ${HEADER_RELATIVE} DIRECTORY)

    file(MAKE_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/include/vuelto/${HEADER_PATH})
    file(COPY ${HEADER} DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}/include/vuelto/${HEADER_PATH})
endforeach()

# Link GLFW to the library
target_link_libraries(vuelto PUBLIC glfw)

# Add GLAD as a dependency
# target_link_libraries(vuelto PUBLIC glad)

# On macOS, link against the OpenGL framework
if(APPLE)
    target_link_libraries(vuelto PUBLIC "-framework OpenGL")
# On Linux, link against the OpenGL library
elseif(UNIX AND NOT APPLE)
    find_package(OpenGL REQUIRED)
    target_link_libraries(vuelto PUBLIC GL)
elseif(WIN32)
    target_link_libraries(vuelto PUBLIC opengl32)
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
