#ifdef __APPLE__

// #define __gl_h_
#define GL_DO_NOT_WARN_IF_MULTI_GL_VERSION_HEADERS_INCLUDED
// #include <OpenGL/gl.h>  // For macOS
#elif defined _WIN32 || defined _WIN64

#else

#include <GL/glew.h>  // For other platforms

#endif

#define GLFW_INCLUDE_NONE
#include "../vendor/GLFW/glfw3.h"
