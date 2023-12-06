#pragma once

#ifdef __APPLE__
#define GL_SILENCE_DEPRECATION
#include <OpenGL/gl.h>
#elif __linux__
#define GL_SILENCE_DEPRECATION
#include <GL/gl.h>
#elif _WIN32
#define GL_CLAMP_TO_EDGE 0x812f
#elif _WIN64
#define GL_CLAMP_TO_EDGE 0x812f
#else
#include <GL/gl.h>
#endif

#define vueltoForEach(item, array)                                                                                    \
  for (int keep = 1, count = 0, size = sizeof(array) / sizeof *(array); keep && count != size; keep = !keep, count++) \
    for (item = (array) + count; keep; keep = !keep)

#include <GLFW/glfw3.h>
#include <stdbool.h>
#include <stdio.h>
