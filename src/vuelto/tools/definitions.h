#ifdef __APPLE__
#define GL_SILENCE_DEPRECATION
#include <OpenGL/gl.h>
#elif __linux__
#define GL_SILENCE_DEPRECATION
#include <GL/gl.h>
#elif _WIN32
#include <GL/glew.h>
#elif _WIN64
#include <GL/glew.h>
#else
printf("Running on an unknown platform, opengl may not work\n");
#include <GL/gl.h>
#endif

#include <GLFW/glfw3.h>
