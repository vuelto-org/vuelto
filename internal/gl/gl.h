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
