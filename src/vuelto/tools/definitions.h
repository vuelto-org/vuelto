#ifdef __APPLE__
#define GL_SILENCE_DEPRECATION
#include <OpenGL/gl.h>
#elif __linux__
#define GL_SILENCE_DEPRECATION
#include <GL/gl.h>
#elif _WIN32
#elif _WIN64
#else
#include <GL/gl.h>
#endif

#include <GLFW/glfw3.h>
